package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2023/util"
)

func main() {
	rows, err := util.ReadStrings(14, false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	for _, row := range rows {
		grid = append(grid, []byte(row))
	}

	for y := 1; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 'O' {
				yTest := y
				for yTest > 0 && grid[yTest-1][x] == '.' {
					yTest--
				}

				if yTest != y {
					grid[yTest][x] = 'O'
					grid[y][x] = '.'
				}
			}
		}
	}

	total := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 'O' {
				total += len(grid) - y
			}
		}
	}

	fmt.Println(total)
}
