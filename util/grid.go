package util

import "fmt"

func PrintGrid(grid [][]byte) {
	for y := 0; y < len(grid); y++ {
		fmt.Printf("%s\n", string(grid[y]))
	}
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
