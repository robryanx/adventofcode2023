package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2023/util"
)

func main() {
	rows, err := util.ReadStrings(12, false, "\n")
	if err != nil {
		panic(err)
	}

	total := 0
	for _, row := range rows {
		rowParts := strings.Split(row, " ")
		var groups []int
		for _, numberRaw := range strings.Split(rowParts[1], ",") {
			num, err := strconv.Atoi(numberRaw)
			if err != nil {
				panic(err)
			}

			groups = append(groups, num)
		}

		completed := 0
		traverse([]byte(rowParts[0]), 0, groups, []byte(rowParts[0]), &completed)

		total += completed
	}

	fmt.Println(total)
}

func traverse(pattern []byte, currentPos int, remainingGroups []int, currentPattern []byte, completed *int) {
	for currentPos < len(pattern) && pattern[currentPos] == '.' {
		currentPos++
	}

	if currentPos >= len(pattern) && len(remainingGroups) == 0 {
		*completed++

		return
	} else if currentPos >= len(pattern) {
		return
	}

	// check what lengths is it possible to consume here
	lengthCount := 0
	for currentPos+lengthCount < len(pattern) && pattern[currentPos+lengthCount] != '.' {
		lengthCount++
	}

	if (currentPos == 0 || currentPattern[currentPos-1] != '#') &&
		len(remainingGroups) != 0 {
		group := remainingGroups[0]

		if group <= lengthCount {
			newGroups := append([]int{}, remainingGroups[1:]...)
			newPattern := append([]byte{}, currentPattern...)
			for j := 0; j < group; j++ {
				newPattern[currentPos+j] = '#'
			}

			traverse(pattern, currentPos+group, newGroups, newPattern, completed)
		}
	}

	if pattern[currentPos] == '?' {
		newPattern := append([]byte{}, currentPattern...)
		newPattern[currentPos] = '.'

		traverse(pattern, currentPos+1, remainingGroups, newPattern, completed)
	}

	return
}
