package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2023/util"
	"golang.org/x/exp/maps"
)

type hand struct {
	typeValue    int
	orderedCards []int
	value        int
}

func main() {
	handRows, err := util.ReadStrings(7, false, "\n")
	if err != nil {
		panic(err)
	}

	var hands []hand
	for _, handRow := range handRows {
		handParts := strings.Split(handRow, " ")

		value, err := strconv.Atoi(handParts[1])
		if err != nil {
			panic(err)
		}

		currentHand := hand{
			value: value,
		}

		handMap := make(map[int]int)
		jokers := 0
		for _, card := range handParts[0] {
			cardValue := normaliseValue(card)
			if cardValue == 1 {
				jokers++
			} else {
				handMap[cardValue]++
			}

			currentHand.orderedCards = append(currentHand.orderedCards, cardValue)
		}

		currentHand.typeValue = handType(maps.Values(handMap), jokers)

		hands = append(hands, currentHand)
	}

	slices.SortFunc(hands, func(a hand, b hand) int {
		if a.typeValue == b.typeValue {
			for i := 0; i < 5; i++ {
				if a.orderedCards[i] > b.orderedCards[i] {
					return 1
				} else if a.orderedCards[i] < b.orderedCards[i] {
					return -1
				}
			}

			return 0
		}

		return a.typeValue - b.typeValue
	})

	total := 0
	for i, finalHand := range hands {
		total += finalHand.value * (i + 1)
	}

	fmt.Println(total)
}

func handType(valueCounts []int, jokers int) int {
	slices.Sort(valueCounts)
	slices.Reverse(valueCounts)

	if len(valueCounts) == 0 {
		valueCounts = append(valueCounts, 5)
	} else {
		valueCounts[0] += jokers
	}

	if len(valueCounts) == 1 {
		return 6
	} else if len(valueCounts) == 2 && valueCounts[0] == 4 {
		return 5
	} else if len(valueCounts) == 2 && valueCounts[0] == 3 {
		return 4
	} else if len(valueCounts) == 3 && valueCounts[0] == 3 {
		return 3
	} else if len(valueCounts) == 3 && valueCounts[0] == 2 {
		return 2
	} else if len(valueCounts) == 4 && valueCounts[0] == 2 {
		return 1
	}

	return 0
}

func normaliseValue(card rune) int {
	if card >= '2' && card <= '9' {
		return int(card - '0')
	} else if card == 'T' {
		return 10
	} else if card == 'Q' {
		return 12
	} else if card == 'K' {
		return 13
	} else if card == 'J' {
		return 1
	}

	return 14
}
