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

	grid[currentY][currentX] = 'F'

	newGrid := make([][]byte, 0, len(grid))
	for y := 0; y < len(grid); y++ {
		var row []byte
		for x := 0; x < len(grid[0]); x++ {
			row = append(row, '.')
		}
		newGrid = append(newGrid, row)
	}

	for i := 0; i < len(validDirections); i++ {
		for {
			last := validDirections[i].history[len(validDirections[i].history)-1]
			next := next(grid, last)

			newGrid[next.y][next.x] = grid[next.y][next.x]

			if next.x == currentX && next.y == currentY {
				break
			}

			validDirections[i].history = append(validDirections[i].history, next)
		}
	}

	expandedGridHorizontal := make([][]byte, 0, len(newGrid))
	for y := 0; y < len(newGrid); y++ {
		var row []byte
		for x := 0; x < len(newGrid[0]); x++ {
			row = append(row, newGrid[y][x])
			if (newGrid[y][x] == 'J' || newGrid[y][x] == '7' || newGrid[y][x] == '|' || newGrid[y][x] == '.') && x+1 < len(grid[0]) && (newGrid[y][x+1] == 'F' || newGrid[y][x+1] == 'L' || newGrid[y][x+1] == '|' || newGrid[y][x+1] == '.') {
				row = append(row, '.')
			} else if newGrid[y][x] == '-' && x+1 < len(grid[0]) && newGrid[y][x+1] == '-' {
				row = append(row, '-')
			} else if x+1 < len(grid[0]) {
				row = append(row, 'X')
			}
		}
		expandedGridHorizontal = append(expandedGridHorizontal, row)
	}

	expandedGridVertical := make([][]byte, 0, len(newGrid)*2)
	for y := 0; y < len(newGrid); y++ {
		var row []byte
		var expandedRow []byte
		for x := 0; x < len(expandedGridHorizontal[0]); x++ {
			row = append(row, expandedGridHorizontal[y][x])
			if (expandedGridHorizontal[y][x] == '-' || expandedGridHorizontal[y][x] == '.' || expandedGridHorizontal[y][x] == 'J' || expandedGridHorizontal[y][x] == 'L' || expandedGridHorizontal[y][x] == 'X') &&
				y+1 < len(newGrid) && (expandedGridHorizontal[y+1][x] == '-' || expandedGridHorizontal[y+1][x] == '.' || expandedGridHorizontal[y+1][x] == 'F' || expandedGridHorizontal[y+1][x] == '7' || expandedGridHorizontal[y+1][x] == 'X') {
				expandedRow = append(expandedRow, '.')
			} else if expandedGridHorizontal[y][x] == '|' && y+1 < len(newGrid) && expandedGridHorizontal[y+1][x] == '|' {
				expandedRow = append(expandedRow, '|')
			} else if y+1 < len(newGrid) {
				expandedRow = append(expandedRow, 'X')
			} else {
				expandedRow = append(expandedRow, expandedGridHorizontal[y][x])
			}
		}
		expandedGridVertical = append(expandedGridVertical, row, expandedRow)
	}

	total := 0
	enclosed := make(map[int]struct{})
	for y := 0; y < len(newGrid); y++ {
		firstNonDot := 10000
		lastNonDot := -1
		for x := 0; x < len(newGrid[0]); x++ {
			if newGrid[y][x] != '.' {
				if x < firstNonDot {
					firstNonDot = x
				}
				if x > lastNonDot {
					lastNonDot = x
				}
			}
		}

		for x := 0; x < len(newGrid[0]); x++ {
			if newGrid[y][x] == '.' && x > firstNonDot && x < lastNonDot {
				if _, ok := enclosed[y*2000+x*2]; !ok {
					visited := make(map[int]struct{})
					completed := false
					navigate(expandedGridVertical, y*2, x*2, &visited, &completed)
					if !completed {
						total++
						for pos := range visited {
							enclosed[pos] = visited[pos]
						}
					}
				} else {
					total++
				}
			}
		}
	}

	fmt.Println(total)
}

func navigate(grid [][]byte, y, x int, visited *map[int]struct{}, completed *bool) {
	if y-1 >= 0 {
		if _, ok := (*visited)[(y-1)*1000+x]; !ok {
			if grid[y-1][x] == '.' {
				(*visited)[(y-1)*1000+x] = struct{}{}
				navigate(grid, y-1, x, visited, completed)
				if *completed {
					return
				}
			}
		}
	} else if grid[y][x] == '.' {
		*completed = true
		return
	}

	if y+1 < len(grid) {
		if _, ok := (*visited)[(y+1)*1000+x]; !ok {
			if grid[y+1][x] == '.' {
				(*visited)[(y+1)*1000+x] = struct{}{}
				navigate(grid, y+1, x, visited, completed)
				if *completed {
					return
				}
			}
		}
	} else if grid[y][x] == '.' {
		*completed = true
		return
	}

	if x-1 >= 0 {
		if _, ok := (*visited)[y*1000+x-1]; !ok {
			if grid[y][x-1] == '.' {
				(*visited)[y*1000+x-1] = struct{}{}
				navigate(grid, y, x-1, visited, completed)
				if *completed {
					return
				}
			}
		}
	} else if grid[y][x] == '.' {
		*completed = true
		return
	}

	if x+1 < len(grid[0]) {
		if _, ok := (*visited)[y*1000+x+1]; !ok {
			if grid[y][x+1] == '.' {
				(*visited)[y*1000+x+1] = struct{}{}
				navigate(grid, y, x+1, visited, completed)
				if *completed {
					return
				}
			}
		}
	} else if grid[y][x] == '.' {
		*completed = true
		return
	}
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
