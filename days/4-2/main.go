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

type cardNumbers struct {
	winningNumbers []int
	numbers []int
	copies int
}

func main() {
	cardsRaw, err := util.ReadStrings(4, false, "\n")
	if err != nil {
		panic(err)
	}

	var cards []cardNumbers
	for _, card := range cardsRaw {
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

		var numbers []int
		for _, numberRaw := range multipleSpace.Split(strings.Trim(cardParts[1], " "), -1) {
			number, err := strconv.Atoi(numberRaw)
			if err != nil {
				panic(err)
			}

			numbers = append(numbers, number)
		}

		cards = append(cards, cardNumbers{
			winningNumbers: winningNumbers,
			numbers: numbers,
			copies: 1,
		})
	}

	for i:=0; i<len(cards); i++ {
		count := 0
		for _, number := range cards[i].numbers {
			if slices.Contains(cards[i].winningNumbers, number) {
				count++
			}
		}

		for j:=i+1; j<=(i+count) && j<len(cards); j++ {
			cards[j].copies += cards[i].copies
		}
	}

	total := 0
	for _, card := range cards {
		total += card.copies
	}

	fmt.Println(total)
}