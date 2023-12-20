package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/robryanx/adventofcode2023/util"
)

var nodeRegex = regexp.MustCompile(`([%&b]{1})([a-z]+) -> (.+)`)

type nodeType int

const (
	nodeTypeBroadcast nodeType = iota
	nodeTypeFlipFlop
	nodeTypeConjunction
)

type signalType int

const (
	signalNone signalType = iota
	signalLow
	signalHigh
)

type nodeState int

const (
	nodeStateOff nodeState = iota
	nodeStateOn
)

type node struct {
	nType   nodeType
	state   nodeState
	memory  map[string]signalType
	outputs []string
}

type signal struct {
	sType signalType
	from  string
	to    string
}

var nodeMap map[string]*node
var inputsMap map[string][]string

func main() {
	nodeLines, err := util.ReadStrings(20, false, "\n")
	if err != nil {
		panic(err)
	}

	nodeMap = make(map[string]*node)
	inputsMap = make(map[string][]string)

	for _, nodeLine := range nodeLines {
		nodeParts := nodeRegex.FindStringSubmatch(nodeLine)

		nType := nodeTypeBroadcast
		switch nodeParts[1] {
		case "%":
			nType = nodeTypeFlipFlop
		case "&":
			nType = nodeTypeConjunction
		}

		memory := make(map[string]signalType)
		outputs := strings.Split(nodeParts[3], ", ")

		for _, input := range inputsMap[nodeParts[2]] {
			memory[input] = signalLow
		}

		for _, output := range outputs {
			if _, ok := nodeMap[output]; ok {
				nodeMap[output].memory[nodeParts[2]] = signalLow
			} else {
				inputsMap[output] = append(inputsMap[output], nodeParts[2])
			}
		}

		n := node{
			nType:   nType,
			state:   nodeStateOff,
			memory:  memory,
			outputs: outputs,
		}

		nodeMap[nodeParts[2]] = &n
	}

	sentLow := 0
	sentHigh := 0

	for i := 0; i < 1000; i++ {
		signals := []signal{
			{
				sType: signalLow,
				from:  "",
				to:    "roadcaster",
			},
		}

		for {
			nextSignals := []signal{}
			for _, s := range signals {
				if s.sType == signalLow {
					sentLow++
				} else {
					sentHigh++
				}

				to, ok := nodeMap[s.to]

				// if something doesn't exist in the node map it will be the output
				if !ok {
					continue
				}

				nextSignalType := signalNone
				switch to.nType {
				case nodeTypeFlipFlop:
					if s.sType == signalLow {
						if to.state == nodeStateOff {
							to.state = nodeStateOn
							nextSignalType = signalHigh
						} else {
							to.state = nodeStateOff
							nextSignalType = signalLow
						}
					}
				case nodeTypeConjunction:
					to.memory[s.from] = s.sType
					allHigh := true
					for _, mem := range to.memory {
						if mem != signalHigh {
							allHigh = false
							break
						}
					}

					if allHigh {
						nextSignalType = signalLow
					} else {
						nextSignalType = signalHigh
					}
				case nodeTypeBroadcast:
					nextSignalType = s.sType
				}

				if nextSignalType != signalNone {
					for _, output := range to.outputs {
						nextSignals = append(nextSignals, signal{
							sType: nextSignalType,
							from:  s.to,
							to:    output,
						})
					}
				}
			}

			if len(nextSignals) == 0 {
				break
			}

			signals = nextSignals
		}
	}

	fmt.Println(sentLow * sentHigh)
}
