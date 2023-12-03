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

				if y-1 >= 0 {
					if isGear(grid[y-1][x]) {
						currentGears[fmt.Sprintf("%d,%d", y-1, x)] = struct{}{}
					}

					if x-1 >= 0 && isGear(grid[y-1][x-1]) {
						currentGears[fmt.Sprintf("%d,%d", y-1, x-1)] = struct{}{}
					}

					if x+1 < len(grid[0]) && isGear(grid[y-1][x+1]) {
						currentGears[fmt.Sprintf("%d,%d", y-1, x+1)] = struct{}{}
					}
				}

				if y+1 < len(grid) {
					if isGear(grid[y+1][x]) {
						currentGears[fmt.Sprintf("%d,%d", y+1, x)] = struct{}{}
					}

					if x-1 >= 0 && isGear(grid[y+1][x-1]) {
						currentGears[fmt.Sprintf("%d,%d", y+1, x-1)] = struct{}{}
					}

					if x+1 < len(grid[0]) && isGear(grid[y+1][x+1]) {
						currentGears[fmt.Sprintf("%d,%d", y+1, x+1)] = struct{}{}
					}

				}

				if x-1 >= 0 && isGear(grid[y][x-1]) {
					currentGears[fmt.Sprintf("%d,%d", y, x-1)] = struct{}{}
				}

				if x+1 < len(grid[0]) && isGear(grid[y][x+1]) {
					currentGears[fmt.Sprintf("%d,%d", y, x+1)] = struct{}{}
				}
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
