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

func (n nodeType) string() string {
	switch n {
	case nodeTypeFlipFlop:
		return "FlipFlop"
	case nodeTypeConjunction:
		return "Conjunction"
	}

	return ""
}

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

func (n nodeState) string() string {
	switch n {
	case nodeStateOff:
		return "Off"
	case nodeStateOn:
		return "On"
	}

	return ""
}

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

type nodeTree struct {
	name     string
	n        *node
	depth    int
	children []*nodeTree
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

	// current := &nodeTree{
	// 	name:  "kl",
	// 	depth: 1,
	// 	n:     nodeMap["kl"],
	// }
	// root := current

	// populateTree(current, nodeMap, 2)
	// treePrinter(root)

	run(nodeMap)
}

func populateTree(current *nodeTree, nodeMap map[string]*node, depth int) {
	if depth > 4 {
		return
	}

	for childName := range nodeMap[current.name].memory {
		n := nodeMap[childName]
		childNode := &nodeTree{
			name:  childName,
			depth: depth,
			n:     n,
		}

		current.children = append(current.children, childNode)

		populateTree(childNode, nodeMap, depth+1)
	}
}

func treePrinter(current *nodeTree) {
	fmt.Println(current.name)
	nodePrinter(current.n)
	fmt.Println(current.depth)
	fmt.Println("Children")
	for _, c := range current.children {
		treePrinter(c)
	}
}

func nodePrinter(n *node) {
	fmt.Println(n.nType.string())
	fmt.Println(n.state.string())
	fmt.Println(n.memory)
}

func run(nodeMap map[string]*node) {
	tracking := make(map[string][]int)
	for i := 1; i < 10000; i++ {
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
						// determine the periods of each of the high singals to kl
						if output == "kl" && nextSignalType == signalHigh {
							tracking[s.to] = append(tracking[s.to], i)
						}

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

	var periods []int
	for k := range tracking {
		periods = append(periods, tracking[k][1]-tracking[k][0])
	}

	fmt.Println(util.LowestCommonMultiple(periods[0], periods[1], periods[2:]...))
}
