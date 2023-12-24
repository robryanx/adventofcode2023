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

	// find the starting position
found:
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 'S' {
				currentPos[pos{y, x}] = struct{}{}
				break found
			}
		}
	}

	for i := 0; i < 64; i++ {
		nextPos := make(map[pos]struct{})

		for cPos := range currentPos {
			if cPos.y > 0 && grid[cPos.y-1][cPos.x] != '#' {
				next := pos{cPos.y - 1, cPos.x}
				if _, ok := nextPos[next]; !ok {
					nextPos[next] = struct{}{}
				}
			}
			if cPos.y < len(grid)-1 && grid[cPos.y+1][cPos.x] != '#' {
				next := pos{cPos.y + 1, cPos.x}
				if _, ok := nextPos[next]; !ok {
					nextPos[next] = struct{}{}
				}
			}
			if cPos.x > 0 && grid[cPos.y][cPos.x-1] != '#' {
				next := pos{cPos.y, cPos.x - 1}
				if _, ok := nextPos[next]; !ok {
					nextPos[next] = struct{}{}
				}
			}
			if cPos.x < len(grid)-1 && grid[cPos.y][cPos.x+1] != '#' {
				next := pos{cPos.y, cPos.x + 1}
				if _, ok := nextPos[next]; !ok {
					nextPos[next] = struct{}{}
				}
			}
		}

		currentPos = nextPos
	}

	fmt.Println(len(currentPos))
}
