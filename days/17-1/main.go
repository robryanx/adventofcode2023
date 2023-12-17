package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2023/util"
)

type dirPos struct {
	y         int
	x         int
	direction Direction
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

	from := util.NodePos{
		Y: 0,
		X: 0,
	}

	to := util.NodePos{
		Y: len(grid) - 1,
		X: len(grid[0]) - 1,
	}

	path, distance, _ := util.Pathfind(grid, from, to)

	fmt.Println(distance)
	util.PrintGridWithPath(grid, path)
}

func paths(grid [][]byte, current dirPos, visited map[int]struct{}, cost int, bestCost *int) {
	if cost > *bestCost {
		return
	}

	fmt.Println(current)

	if current.x == len(grid[0])-1 && current.y == len(grid)-1 {
		fmt.Println(cost)
		if cost < *bestCost {
			*bestCost = cost
		}

		return
	}

	if current.direction != North {
		for i := 1; i < 4; i++ {
			if current.y-i >= 0 {
				if _, ok := visited[(current.y-i)*100+current.x]; !ok {
					newCost := 0
					for j := 1; j <= i; j++ {
						newCost += int(grid[current.y-j][current.x] - '0')
					}

					newPos := dirPos{
						y:         current.y - i,
						x:         current.x,
						direction: North,
					}

					newVisited := make(map[int]struct{}, len(visited))
					for k, v := range visited {
						newVisited[k] = v
					}
					newVisited[newPos.y*100+newPos.x] = struct{}{}

					paths(grid, newPos, newVisited, cost+newCost, bestCost)
				}
			}
		}
	}

	if current.direction != South {
		for i := 1; i < 4; i++ {
			if current.y+i < len(grid) {
				if _, ok := visited[(current.y-i)*100+current.x]; !ok {
					newCost := 0
					for j := 1; j <= i; j++ {
						newCost += int(grid[current.y+j][current.x] - '0')
					}

					newPos := dirPos{
						y:         current.y + i,
						x:         current.x,
						direction: South,
					}

					newVisited := make(map[int]struct{}, len(visited))
					for k, v := range visited {
						newVisited[k] = v
					}
					newVisited[newPos.y*100+newPos.x] = struct{}{}

					paths(grid, newPos, newVisited, cost+newCost, bestCost)
				}
			}
		}
	}

	if current.direction != West {
		for i := 1; i < 4; i++ {
			if current.x-i > 0 {
				if _, ok := visited[(current.y-i)*100+current.x]; !ok {
					newCost := 0
					for j := 1; j <= i; j++ {
						newCost += int(grid[current.y][current.x-j] - '0')
					}

					newPos := dirPos{
						y:         current.y,
						x:         current.x - i,
						direction: West,
					}

					newVisited := make(map[int]struct{}, len(visited))
					for k, v := range visited {
						newVisited[k] = v
					}
					newVisited[newPos.y*100+newPos.x] = struct{}{}

					paths(grid, newPos, newVisited, cost+newCost, bestCost)
				}
			}
		}
	}

	if current.direction != East {
		for i := 1; i < 4; i++ {
			if current.x+i < len(grid[0]) {
				if _, ok := visited[(current.y-i)*100+current.x]; !ok {
					newCost := 0
					for j := 1; j <= i; j++ {
						newCost += int(grid[current.y][current.x+j] - '0')
					}

					newPos := dirPos{
						y:         current.y,
						x:         current.x + i,
						direction: East,
					}

					newVisited := make(map[int]struct{}, len(visited))
					for k, v := range visited {
						newVisited[k] = v
					}
					newVisited[newPos.y*100+newPos.x] = struct{}{}

					paths(grid, newPos, newVisited, cost+newCost, bestCost)
				}
			}
		}
	}
}
