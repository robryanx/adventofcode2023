package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2023/util"
)

var multipleSpace = regexp.MustCompile(`\s+`)
var cardSplit = regexp.MustCompile(`^Card\s+([0-9]+): ([0-9\s]+)\s\|\s([0-9\s]+)`)

func main() {
	cards, err := util.ReadStrings(4, false, "\n")
	if err != nil {
		panic(err)
	}

	total := 0
	for _, card := range cards {
		parts := cardSplit.FindStringSubmatch(card)

		winningNumbers := make([]int, 5)
		for _, numberRaw := range multipleSpace.Split(strings.Trim(parts[2], " "), -1) {
			number, err := strconv.Atoi(numberRaw)
			if err != nil {
				panic(err)
			}

			winningNumbers = append(winningNumbers, number)
		}

		count := 0
		for _, numberRaw := range multipleSpace.Split(strings.Trim(parts[3], " "), -1) {
			number, err := strconv.Atoi(numberRaw)
			if err != nil {
				panic(err)
			}

			if slices.Contains(winningNumbers, number) {
				if count == 0 {
					count = 1
				} else {
					count *= 2
				}
			}
		}

		total += count
	}

	fmt.Println(total)
}
