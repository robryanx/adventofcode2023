package main

import (
	"fmt"
	"github.com/robryanx/adventofcode2023/util"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var multipleSpace = regexp.MustCompile(`\s+`)

func main() {
	cards, err := util.ReadStrings(4, false, "\n")
	if err != nil {
		panic(err)
	}

	total := 0
	for _, card := range cards {
		cardParts := strings.Split(card, "|")
		winningParts := strings.Split(cardParts[0], ":")

		var winningNumbers []int
		for _, numberRaw := range multipleSpace.Split(strings.Trim(winningParts[1], " "), -1) {
			number, err := strconv.Atoi(numberRaw)
			if err != nil {
				panic(err)
			}

			winningNumbers = append(winningNumbers, number)
		}

		count := 0
		for _, numberRaw := range multipleSpace.Split(strings.Trim(cardParts[1], " "), -1) {
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