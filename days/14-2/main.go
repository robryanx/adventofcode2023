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

	period := 0
	var baseGrid [][]byte
	for i := 0; i <= 1000000; i += 1 {
		if i == 100 {
			baseGrid = util.CopyGrid(grid)
		}

		rollNorth(grid)
		rollWest(grid)
		rollSouth(grid)
		rollEast(grid)

		if util.CompareGrids(grid, baseGrid) && i != 100 {
			period = i - 99
			break
		}
	}

	offset := (1000000000 - 100 - period) % period
	for i := 0; i < offset; i++ {
		rollNorth(grid)
		rollWest(grid)
		rollSouth(grid)
		rollEast(grid)
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

func rollNorth(grid [][]byte) {
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
}

func rollWest(grid [][]byte) {
	for x := 1; x < len(grid[0]); x++ {
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == 'O' {
				xTest := x
				for xTest > 0 && grid[y][xTest-1] == '.' {
					xTest--
				}

				if xTest != x {
					grid[y][xTest] = 'O'
					grid[y][x] = '.'
				}
			}
		}
	}
}

func rollSouth(grid [][]byte) {
	for y := len(grid) - 2; y >= 0; y-- {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 'O' {
				yTest := y
				for yTest < len(grid)-1 && grid[yTest+1][x] == '.' {
					yTest++
				}

				if yTest != y {
					grid[yTest][x] = 'O'
					grid[y][x] = '.'
				}
			}
		}
	}
}

func rollEast(grid [][]byte) {
	for x := len(grid) - 2; x >= 0; x-- {
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == 'O' {
				xTest := x
				for xTest < len(grid)-1 && grid[y][xTest+1] == '.' {
					xTest++
				}

				if xTest != x {
					grid[y][xTest] = 'O'
					grid[y][x] = '.'
				}
			}
		}
	}
}
