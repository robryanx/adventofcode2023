package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2023/util"
)

func main() {
	rows, err := util.ReadStrings(9, false, "\n")
	if err != nil {
		panic(err)
	}

	total := 0
	for _, row := range rows {
		var numbers [][]int
		var firstRow []int
		for _, numberRaw := range strings.Split(row, " ") {
			number, err := strconv.Atoi(numberRaw)
			if err != nil {
				panic(err)
			}

			firstRow = append(firstRow, number)
		}

		numbers = append(numbers, firstRow)
		currentRow := 0

		for {
			nonZeroAdded := false

			nextRow := make([]int, 0, len(firstRow))
			for i := 0; i < len(numbers[currentRow])-1; i++ {
				nextValue := numbers[currentRow][i+1] - numbers[currentRow][i]
				if nextValue != 0 {
					nonZeroAdded = true
				}

				nextRow = append(nextRow, nextValue)
			}

			numbers = append(numbers, nextRow)
			currentRow++

			if !nonZeroAdded {
				break
			}
		}

		numbers[currentRow] = append([]int{0}, numbers[currentRow]...)

		for i := currentRow - 1; i >= 0; i-- {
			newValue := numbers[i][0]
			newValue -= numbers[i+1][0]

			numbers[i] = append([]int{newValue}, numbers[i]...)
		}

		total += numbers[0][0]
	}

	fmt.Println(total)
}
