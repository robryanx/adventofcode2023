package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2023/util"
)

type dirPos struct {
	y    int
	x    int
	cost int
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
	Unknown
)

func main() {
	rows, err := util.ReadStrings(17, true, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	for _, row := range rows {
		grid = append(grid, []byte(row))
	}

	graph := util.NewGraph()

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			for _, path := range paths(grid, y, x) {
				graph.AddEdge(fmt.Sprintf("%d-%d", y, x), fmt.Sprintf("%d-%d", path.y, path.x), path.cost)
			}
		}
	}

	score, path := graph.GetPath(fmt.Sprintf("%d-%d", 0, 0), fmt.Sprintf("%d-%d", len(grid)-1, len(grid[0])-1))
	fmt.Println(score)
	fmt.Println(path)
}

func paths(grid [][]byte, y, x int) []dirPos {
	newPaths := make([]dirPos, 0, 12)

	for i := 1; i < 4; i++ {
		if y-i >= 0 {
			newCost := 0
			for j := 1; j <= i; j++ {
				newCost += int(grid[y-j][x] - '0')
			}

			newPaths = append(newPaths, dirPos{
				y:    y - i,
				x:    x,
				cost: newCost,
			})
		}
	}

	for i := 1; i < 4; i++ {
		if y+i < len(grid) {
			newCost := 0
			for j := 1; j <= i; j++ {
				newCost += int(grid[y+j][x] - '0')
			}

			newPaths = append(newPaths, dirPos{
				y:    y + i,
				x:    x,
				cost: newCost,
			})
		}
	}

	for i := 1; i < 4; i++ {
		if x-i >= 0 {
			newCost := 0
			for j := 1; j <= i; j++ {
				newCost += int(grid[y][x-j] - '0')
			}

			newPaths = append(newPaths, dirPos{
				y:    y,
				x:    x - i,
				cost: newCost,
			})
		}
	}

	for i := 1; i < 4; i++ {
		if x+i < len(grid[0]) {
			newCost := 0
			for j := 1; j <= i; j++ {
				newCost += int(grid[y][x+j] - '0')
			}

			newPaths = append(newPaths, dirPos{
				y:    y,
				x:    x + i,
				cost: newCost,
			})
		}
	}

	return newPaths
}
