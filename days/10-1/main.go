package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2023/util"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type dirPos struct {
	y         int
	x         int
	direction Direction
}

type cycle struct {
	history []dirPos
}

func main() {
	rows, err := util.ReadStrings(10, false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	for _, row := range rows {
		grid = append(grid, []byte(row))
	}

	// find the starting position
	var currentY int
	var currentX int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 'S' {
				currentY = y
				currentX = x
			}
		}
	}

	// check which directions are valid
	var validDirections []cycle
	if currentY-1 > 0 {
		if grid[currentY-1][currentX] == '|' || grid[currentY-1][currentX] == '7' || grid[currentY-1][currentX] == 'F' {
			validDirections = append(validDirections, cycle{
				history: []dirPos{
					{
						y:         currentY - 1,
						x:         currentX,
						direction: North,
					},
				},
			})
		}
	}

	if currentY+1 < len(grid) {
		if grid[currentY+1][currentX] == '|' || grid[currentY+1][currentX] == 'L' || grid[currentY+1][currentX] == 'J' {
			validDirections = append(validDirections, cycle{
				history: []dirPos{
					{
						y:         currentY + 1,
						x:         currentX,
						direction: South,
					},
				},
			})
		}
	}

	if currentX-1 > 0 {
		if grid[currentY][currentX-1] == '-' || grid[currentY][currentX-1] == 'L' || grid[currentY][currentX-1] == 'F' {
			validDirections = append(validDirections, cycle{
				history: []dirPos{
					{
						y:         currentY,
						x:         currentX - 1,
						direction: West,
					},
				},
			})
		}
	}

	if currentX+1 < len(grid[currentY]) {
		if grid[currentY][currentX+1] == '-' || grid[currentY][currentX+1] == '7' || grid[currentY][currentX+1] == 'J' {
			validDirections = append(validDirections, cycle{
				history: []dirPos{
					{
						y:         currentY,
						x:         currentX + 1,
						direction: East,
					},
				},
			})
		}
	}

	for i := 0; i < len(validDirections); i++ {
		for {
			last := validDirections[i].history[len(validDirections[i].history)-1]
			next := next(grid, last)

			if next.x == currentX && next.y == currentY {
				break
			}

			validDirections[i].history = append(validDirections[i].history, next)
		}
	}

	posDistance := make(map[int]int)
	for _, path := range validDirections {
		for i, pos := range path.history {
			distance := i + 1
			cur, ok := posDistance[pos.y*1000+pos.x]

			if !ok || cur > distance {
				posDistance[pos.y*1000+pos.x] = distance
			}
		}
	}

	max := 0
	for _, distance := range posDistance {
		if distance > max {
			max = distance
		}
	}

	fmt.Println(max)
}

func next(grid [][]byte, pos dirPos) dirPos {
	switch pos.direction {
	case East:
		switch grid[pos.y][pos.x] {
		case '-':
			return dirPos{
				y:         pos.y,
				x:         pos.x + 1,
				direction: East,
			}
		case '7':
			return dirPos{
				y:         pos.y + 1,
				x:         pos.x,
				direction: South,
			}
		case 'J':
			return dirPos{
				y:         pos.y - 1,
				x:         pos.x,
				direction: North,
			}
		}
	case West:
		switch grid[pos.y][pos.x] {
		case '-':
			return dirPos{
				y:         pos.y,
				x:         pos.x - 1,
				direction: West,
			}
		case 'F':
			return dirPos{
				y:         pos.y + 1,
				x:         pos.x,
				direction: South,
			}
		case 'L':
			return dirPos{
				y:         pos.y - 1,
				x:         pos.x,
				direction: North,
			}
		}
	case North:
		switch grid[pos.y][pos.x] {
		case '|':
			return dirPos{
				y:         pos.y - 1,
				x:         pos.x,
				direction: North,
			}
		case 'F':
			return dirPos{
				y:         pos.y,
				x:         pos.x + 1,
				direction: East,
			}
		case '7':
			return dirPos{
				y:         pos.y,
				x:         pos.x - 1,
				direction: West,
			}
		}
	case South:
		switch grid[pos.y][pos.x] {
		case '|':
			return dirPos{
				y:         pos.y + 1,
				x:         pos.x,
				direction: South,
			}
		case 'L':
			return dirPos{
				y:         pos.y,
				x:         pos.x + 1,
				direction: East,
			}
		case 'J':
			return dirPos{
				y:         pos.y,
				x:         pos.x - 1,
				direction: West,
			}
		}
	}

	return dirPos{}
}
