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

	var rowSkips []int
	var colSkips []int

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
			rowSkips = append(rowSkips, y)
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
			colSkips = append(colSkips, x)
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
				var startX int
				var endX int
				if posA.x > posB.x {
					startX = posB.x
					endX = posA.x
				} else {
					startX = posA.x
					endX = posB.x
				}

				var startY int
				var endY int
				if posA.y > posB.y {
					startY = posB.y
					endY = posA.y
				} else {
					startY = posA.y
					endY = posB.y
				}

				extraX := 0
				for _, colX := range colSkips {
					if colX >= startX && colX <= endX {
						extraX++
					}
				}

				extraY := 0
				for _, rowY := range rowSkips {
					if rowY >= startY && rowY <= endY {
						extraY++
					}
				}

				total += abs(posA.y-posB.y) + (extraX * 999999) + abs(posA.x-posB.x) + (extraY * 999999)
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
