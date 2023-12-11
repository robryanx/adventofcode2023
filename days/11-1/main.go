package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2023/util"
)

type pos struct {
	y int
	x int
}

func main() {
	rows, err := util.ReadStrings(11, false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	for _, row := range rows {
		grid = append(grid, []byte(row))
	}

	// check rows for expansion
	for y := 0; y < len(grid); y++ {
		emptyRow := true
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == '#' {
				emptyRow = false
				break
			}
		}

		if emptyRow {
			row := make([]byte, 0, len(grid[0]))
			for i := 0; i < len(grid[0]); i++ {
				row = append(row, '.')
			}

			grid = append(grid[:y+1], grid[y:]...)
			grid[y] = row
			y++
		}
	}

	// check for cols in expansion
	for x := 0; x < len(grid[0]); x++ {
		emptyCol := true
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == '#' {
				emptyCol = false
				break
			}
		}

		if emptyCol {
			for i := 0; i < len(grid); i++ {
				grid[i] = append(grid[i][:x+1], grid[i][x:]...)
				grid[i][x] = '.'
			}

			x++
		}
	}

	posLookup := make(map[int]pos)

	replacement := 1
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == '#' {
				posLookup[replacement] = pos{
					y: y,
					x: x,
				}
				replacement++
			}
		}
	}

	total := 0
	for numA, posA := range posLookup {
		for numB, posB := range posLookup {
			if numA != numB {
				total += abs(posA.y-posB.y) + abs(posA.x-posB.x)
			}
		}
	}

	fmt.Println(total / 2)
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}
