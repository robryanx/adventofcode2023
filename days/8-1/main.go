package main

import (
	"fmt"
	"github.com/robryanx/adventofcode2023/util"
	"regexp"
)

var nodeParse = regexp.MustCompile(`([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)`)

func main() {
	rows, err := util.ReadStrings(8, false, "\n")
	if err != nil {
		panic(err)
	}

	graph := make(map[string][2]string)

	for _, node := range rows[2:] {
		nodeParts := nodeParse.FindStringSubmatch(node)

		graph[nodeParts[1]] = [2]string{
			nodeParts[2],
			nodeParts[3],
		}
	}

	done := false
	current := "AAA"
	steps := 1

	for {
		for _, dir := range rows[0] {
			match := 0
			if dir == 'R' {
				match = 1
			}
			current = graph[current][match]

			if current == "ZZZ" {
				done = true
				break
			}
			steps++
		}

		if done {
			break
		}
	}

	fmt.Println(steps)
}
