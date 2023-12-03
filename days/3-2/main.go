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
	gears := make(map[string][]int)
	for y := 0; y < len(grid); y++ {
		var currentNumber []byte
		var currentGears = make(map[string]struct{})

		for x := 0; x < len(grid[0]); x++ {
			if isNumber(grid[y][x]) {
				currentNumber = append(currentNumber, grid[y][x])

				util.AdjacentMatch(grid, y, x, true, func(char byte, y, x int) bool {
					if isGear(char) {
						currentGears[fmt.Sprintf("%d,%d", y, x)] = struct{}{}
					}

					return false
				})
			}

			if len(currentNumber) > 0 && (x+1 == len(grid[0]) || !isNumber(grid[y][x+1])) {
				for gearPos := range currentGears {
					num, err := strconv.Atoi(string(currentNumber))
					if err != nil {
						panic(err)
					}

					gears[gearPos] = append(gears[gearPos], num)
				}
				currentNumber = nil
				clear(currentGears)
			}
		}
	}

	for _, gear := range gears {
		if len(gear) > 1 {
			val := 1
			for _, num := range gear {
				val *= num
			}

			total += val
		}
	}

	fmt.Println(total)
}

func isNumber(char byte) bool {
	return char >= '0' && char <= '9'
}

func isGear(char byte) bool {
	return char == '*'
}
