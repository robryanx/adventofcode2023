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

				if !currentNumberAdjacent && y-1 >= 0 {
					if isSymbol(grid[y-1][x]) {
						currentNumberAdjacent = true
					}

					if !currentNumberAdjacent && x-1 >= 0 && isSymbol(grid[y-1][x-1]) {
						currentNumberAdjacent = true
					}

					if !currentNumberAdjacent && x+1 < len(grid[0]) && isSymbol(grid[y-1][x+1]) {
						currentNumberAdjacent = true
					}
				}

				if !currentNumberAdjacent && y+1 < len(grid) {
					if isSymbol(grid[y+1][x]) {
						currentNumberAdjacent = true
					}

					if !currentNumberAdjacent && x-1 >= 0 && isSymbol(grid[y+1][x-1]) {
						currentNumberAdjacent = true
					}

					if !currentNumberAdjacent && x+1 < len(grid[0]) && isSymbol(grid[y+1][x+1]) {
						currentNumberAdjacent = true
					}

				}

				if !currentNumberAdjacent && x-1 >= 0 && isSymbol(grid[y][x-1]) {
					currentNumberAdjacent = true
				}

				if !currentNumberAdjacent && x+1 < len(grid[0]) && isSymbol(grid[y][x+1]) {
					currentNumberAdjacent = true
				}
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
