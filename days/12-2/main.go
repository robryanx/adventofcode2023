package main

import (
	"fmt"
	"slices"
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

		enlargedPattern := rowParts[0]
		enlargedGroup := rowParts[1]
		for i := 0; i < 4; i++ {
			enlargedPattern += "?" + rowParts[0]
			enlargedGroup += "," + rowParts[1]
		}

		var groups []int
		for _, numberRaw := range strings.Split(enlargedGroup, ",") {
			num, err := strconv.Atoi(numberRaw)
			if err != nil {
				panic(err)
			}

			groups = append(groups, num)
		}
		slices.Reverse(groups)

		posLookup := make(map[string]int)
		count := traverse([]byte(enlargedPattern), groups, &posLookup)

		total += count
	}

	fmt.Println(total)
}

func traverse(pattern []byte, remainingGroups []int, posLookup *map[string]int) int {
	if len(remainingGroups) == 0 {
		found := false
		for _, char := range pattern {
			if char == '#' {
				found = true
				break
			}
		}

		if !found {
			return 1
		}

		return 0
	}

	sum := 0
	for i := 0; i < len(remainingGroups); i++ {
		sum += remainingGroups[i]
	}

	groupMinStart := 0
	for i, char := range pattern {
		if char != '.' {
			groupMinStart = i
			break
		}
	}

	groupMaxStart := len(pattern) - sum - len(remainingGroups) + 1
	for i := 0; i < groupMaxStart; i++ {
		if pattern[i] == '#' {
			groupMaxStart = i
			break
		}
	}

	patternsCount := 0
	group := remainingGroups[len(remainingGroups)-1]
	for i := groupMinStart; i <= groupMaxStart; i++ {
		possible := true
		for j := i; j < i+group; j++ {
			if pattern[j] == '.' {
				possible = false
				break
			}
		}

		if !possible {
			continue
		}

		endOfPattern := i+group >= len(pattern)
		if !endOfPattern && pattern[i+group] == '#' {
			continue
		}

		newPattern := []byte{}
		if i+group+1 < len(pattern) {
			newPattern = append([]byte{}, pattern[i+group+1:]...)
		}

		newGroups := remainingGroups[:len(remainingGroups)-1]

		key := fmt.Sprintf("%s-%d", newPattern, len(newGroups))
		count, ok := (*posLookup)[key]
		if !ok {
			count = traverse(newPattern, newGroups, posLookup)
			(*posLookup)[key] = count
		}

		patternsCount += count
	}

	return patternsCount
}
