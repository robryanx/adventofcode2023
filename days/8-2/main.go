package main

import (
	"fmt"
	"regexp"

	"github.com/robryanx/adventofcode2023/util"
)

var nodeParse = regexp.MustCompile(`([A-Z0-9]{3}) = \(([A-Z0-9]{3}), ([A-Z0-9]{3})\)`)

func main() {
	rows, err := util.ReadStrings(8, false, "\n")
	if err != nil {
		panic(err)
	}

	var currentNodes []string
	var zPos []int
	graph := make(map[string][2]string)

	for _, node := range rows[2:] {
		nodeParts := nodeParse.FindStringSubmatch(node)

		graph[nodeParts[1]] = [2]string{
			nodeParts[2],
			nodeParts[3],
		}

		if nodeParts[1][2] == 'A' {
			currentNodes = append(currentNodes, nodeParts[1])
			zPos = append(zPos, 0)
		}
	}

	for j := 0; j < len(currentNodes); j++ {
		count := 0
	loop:
		for {
			for pos, dir := range rows[0] {
				match := 0
				if dir == 'R' {
					match = 1
				}

				currentNodes[j] = graph[currentNodes[j]][match]
				if currentNodes[j][2] == 'Z' {
					zPos[j] = (len(rows[0]) * count) + pos + 1
					break loop
				}
			}
			count++
		}
	}

	fmt.Println(lowestCommonMultiple(zPos[0], zPos[1], zPos[2:]...))
}

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lowestCommonMultiple(a, b int, integers ...int) int {
	result := a * b / greatestCommonDivisor(a, b)

	for i := 0; i < len(integers); i++ {
		result = lowestCommonMultiple(result, integers[i])
	}

	return result
}
