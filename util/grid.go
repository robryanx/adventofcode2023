package util

import (
	"fmt"
	"strconv"
	"strings"
)

func PrintGrid(grid [][]byte) {
	for y := 0; y < len(grid); y++ {
		fmt.Printf("%s\n", string(grid[y]))
	}
}

func PrintStrGrid(grid [][]string) {
	for y := 0; y < len(grid); y++ {
		fmt.Printf("%s\n", strings.Join(grid[y], ""))
	}
}

func PrintGridWithPath(grid [][]byte, path []NodePos) {
	grid = CopyGrid(grid)
	for _, p := range path {
		switch p.direction {
		case North:
			grid[p.Y][p.X] = '^'
		case South:
			grid[p.Y][p.X] = 'v'
		case East:
			grid[p.Y][p.X] = '>'
		case West:
			grid[p.Y][p.X] = '<'
		}
	}

	PrintGrid(grid)
}

func PrintGridWithCosts(grid [][]byte, path []NodePos) {
	strGrid := CopyStrGrid(grid)
	for _, p := range path {
		fmt.Println(p.cost)
		strGrid[p.Y][p.X] = fmt.Sprintf("|%s|", strconv.Itoa(p.cost))
	}

	PrintStrGrid(strGrid)
}

func CopyGrid(grid [][]byte) [][]byte {
	newGrid := make([][]byte, 0, len(grid))
	for y := 0; y < len(grid); y++ {
		newRow := append([]byte{}, grid[y]...)
		newGrid = append(newGrid, newRow)
	}

	return newGrid
}

func CopyStrGrid(grid [][]byte) [][]string {
	newGrid := make([][]string, 0, len(grid))
	for y := 0; y < len(grid); y++ {
		newRow := []string{}
		for x := 0; x < len(grid[0]); x++ {
			newRow = append(newRow, string(grid[y][x]))
		}
		newGrid = append(newGrid, newRow)
	}

	return newGrid
}

func CompareGrids(gridA [][]byte, gridB [][]byte) bool {
	if len(gridA) != len(gridB) || len(gridA[0]) != len(gridB[0]) {
		return false
	}

	for y := 0; y < len(gridA); y++ {
		for x := 0; x < len(gridB); x++ {
			if gridA[y][x] != gridB[y][x] {
				return false
			}
		}
	}

	return true
}

func AdjacentMatch(grid [][]byte, y, x int, incDiagonal bool, cb func(char byte, y, x int) bool) {
	if y-1 >= 0 {
		earlyExit := cb(grid[y-1][x], y-1, x)
		if earlyExit {
			return
		}

		if incDiagonal {
			if x-1 >= 0 {
				earlyExit := cb(grid[y-1][x-1], y-1, x-1)
				if earlyExit {
					return
				}
			}

			if x+1 < len(grid[0]) {
				earlyExit := cb(grid[y-1][x+1], y-1, x+1)
				if earlyExit {
					return
				}
			}
		}
	}

	if y+1 < len(grid) {
		earlyExit := cb(grid[y+1][x], y+1, x)
		if earlyExit {
			return
		}

		if incDiagonal {
			if x-1 >= 0 {
				earlyExit := cb(grid[y+1][x-1], y+1, x-1)
				if earlyExit {
					return
				}
			}

			if x+1 < len(grid[0]) {
				earlyExit := cb(grid[y+1][x+1], y+1, x+1)
				if earlyExit {
					return
				}
			}
		}
	}

	if x-1 >= 0 {
		earlyExit := cb(grid[y][x-1], y, x-1)
		if earlyExit {
			return
		}
	}

	if x+1 < len(grid[0]) {
		_ = cb(grid[y][x+1], y, x+1)
	}
}
