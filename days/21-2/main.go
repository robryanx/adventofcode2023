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
	rows, err := util.ReadStrings(21, false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	for _, row := range rows {
		grid = append(grid, []byte(row))
	}

	currentPos := make(map[pos]struct{})

	gridCopies := 300

	// find the starting position
found:
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 'S' {
				grid[y][x] = '.'
				currentPos[pos{y + (gridCopies/2)*len(grid), x + (gridCopies/2)*len(grid[0])}] = struct{}{}
				break found
			}
		}
	}

	var extendedGrid [][]byte
	for y := 0; y < (len(grid) * gridCopies); y++ {
		extendedRow := append([]byte{}, grid[y%len(grid)]...)
		for i := 0; i < gridCopies; i++ {
			extendedRow = append(extendedRow, grid[y%len(grid)]...)
		}

		extendedGrid = append(extendedGrid, extendedRow)
	}

	//display(extendedGrid, currentPos)

	// currentSize := 12321
	// currentDiff := 9185
	// for i := 3; i < 481844; i++ {
	// 	currentDiff += 6050
	// 	currentSize += currentDiff
	// 	fmt.Printf("iteration: %d, current size: %d\n", 55*i, currentSize)
	// }

	var prevSize int
	var prevDiff int
	compareSize := 65
	for i := 0; i < 1000; i++ {
		nextPos := make(map[pos]struct{})

		for cPos := range currentPos {
			if extendedGrid[cPos.y-1][cPos.x] != '#' {
				next := pos{cPos.y - 1, cPos.x}
				if _, ok := nextPos[next]; !ok {
					nextPos[next] = struct{}{}
				}
			}
			if extendedGrid[cPos.y+1][cPos.x] != '#' {
				next := pos{cPos.y + 1, cPos.x}
				if _, ok := nextPos[next]; !ok {
					nextPos[next] = struct{}{}
				}
			}
			if extendedGrid[cPos.y][cPos.x-1] != '#' {
				next := pos{cPos.y, cPos.x - 1}
				if _, ok := nextPos[next]; !ok {
					nextPos[next] = struct{}{}
				}
			}
			if extendedGrid[cPos.y][cPos.x+1] != '#' {
				next := pos{cPos.y, cPos.x + 1}
				if _, ok := nextPos[next]; !ok {
					nextPos[next] = struct{}{}
				}
			}
		}

		if (i+1)%compareSize == 0 {
			diff := len(nextPos) - prevSize
			fmt.Printf("iteration: %d, current size: %d, diff: %d, diff diff: %d\n", i+1, len(nextPos), diff, diff-prevDiff)
			prevSize = len(nextPos)
			prevDiff = diff
		}

		currentPos = nextPos
	}

}

func display(grid [][]byte, currentPos map[pos]struct{}) {
	printGrid := util.CopyGrid(grid)
	for cPos := range currentPos {
		printGrid[cPos.y][cPos.x] = 'O'
	}
	util.PrintGrid(printGrid)
}
