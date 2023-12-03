package main

import (
	"fmt"
	"strconv"

	"github.com/robryanx/adventofcode2023/util"
)

func main() {
	var grid [][]byte

	rows, err := util.ReadStrings(3, false, "\n")
	if err != nil {
		panic(err)
	}

	for _, row := range rows {
		grid = append(grid, []byte(row))
	}

	total := 0
	for y := 0; y < len(grid); y++ {
		var currentNumber []byte
		var currentNumberAdjacent bool

		for x := 0; x < len(grid[0]); x++ {
			if isNumber(grid[y][x]) {
				currentNumber = append(currentNumber, grid[y][x])

				util.AdjacentMatch(grid, y, x, true, func(char byte, y, x int) bool {
					if isSymbol(char) {
						currentNumberAdjacent = true
						return true
					}

					return false
				})
			}

			if len(currentNumber) > 0 && (x+1 == len(grid[0]) || !isNumber(grid[y][x+1])) {
				if currentNumberAdjacent {
					num, err := strconv.Atoi(string(currentNumber))
					if err != nil {
						panic(err)
					}

					total += num
					currentNumberAdjacent = false
				}
				currentNumber = nil
			}
		}
	}

	fmt.Println(total)
}

func isNumber(char byte) bool {
	return char >= '0' && char <= '9'
}

func isSymbol(char byte) bool {
	return !isNumber(char) && char != '.'
}
