package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2023/util"
)

type pos struct {
	y int
	x int
}

var directionLookup = map[byte]byte{
	'0': 'R',
	'1': 'D',
	'2': 'L',
	'3': 'U',
}

type line struct {
	start pos
	end   pos
}

func main() {
	rows, err := util.ReadStrings(18, true, "\n")
	if err != nil {
		panic(err)
	}

	currentPos := pos{}

	yLookup := make(map[int][]line)

	for _, row := range rows {
		rowParts := strings.Split(row, " ")

		distance64, err := strconv.ParseInt(rowParts[2][2:7], 16, 64)
		if err != nil {
			panic(err)
		}

		distance := int(distance64)
		switch directionLookup[rowParts[2][7:][0]] {
		case 'R':
			l := line{
				start: currentPos,
				end: pos{
					y: currentPos.y,
					x: currentPos.x + distance,
				},
			}
			yLookup[currentPos.y] = append(yLookup[currentPos.y], l)
			currentPos.x += distance
		case 'L':
			l := line{
				start: currentPos,
				end: pos{
					y: currentPos.y,
					x: currentPos.x - distance,
				},
			}
			yLookup[currentPos.y] = append(yLookup[currentPos.y], l)
			currentPos.x -= distance
		case 'U':
			l := line{
				start: currentPos,
				end: pos{
					y: currentPos.y - distance,
					x: currentPos.x,
				},
			}
			yLookup[currentPos.y] = append(yLookup[currentPos.y], l)
			currentPos.y -= distance
		case 'D':
			l := line{
				start: currentPos,
				end: pos{
					y: currentPos.y + distance,
					x: currentPos.x,
				},
			}
			yLookup[currentPos.y] = append(yLookup[currentPos.y], l)
			currentPos.y += distance
		}
	}

	fmt.Println(currentPos)
	fmt.Println(yLookup)

	// yOffset := b.minY
	// yLength := b.maxY - b.minY + 1

	// xOffset := b.minX
	// xLength := b.maxX - b.minX + 1

	// grid := make([][]byte, 0, yLength)
	// for y := 0; y < yLength; y++ {
	// 	row := make([]byte, 0, xLength)
	// 	for x := 0; x < xLength; x++ {
	// 		row = append(row, '.')
	// 	}
	// 	grid = append(grid, row)
	// }

	// for _, row := range rows {
	// 	rowParts := strings.Split(row, " ")

	// 	distance, err := strconv.Atoi(rowParts[1])
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	switch byte(rowParts[0][0]) {
	// 	case 'R':
	// 		for i := 0; i < distance; i++ {
	// 			grid[-yOffset+currentPos.y][-xOffset+currentPos.x+i] = '#'
	// 		}

	// 		currentPos.x += distance
	// 	case 'L':
	// 		for i := 0; i < distance; i++ {
	// 			grid[-yOffset+currentPos.y][-xOffset+currentPos.x-i] = '#'
	// 		}

	// 		currentPos.x -= distance
	// 	case 'U':
	// 		for i := 0; i < distance; i++ {
	// 			grid[-yOffset+currentPos.y-i][-xOffset+currentPos.x] = '#'
	// 		}

	// 		currentPos.y -= distance
	// 	case 'D':
	// 		for i := 0; i < distance; i++ {
	// 			grid[-yOffset+currentPos.y+i][-xOffset+currentPos.x] = '#'
	// 		}

	// 		currentPos.y += distance
	// 	}
	// }

	// total := 0
	// for y := 0; y < len(grid); y++ {
	// 	fmt.Println(y)
	// 	rowFill := 0
	// 	inside := false
	// 	x := 0
	// 	for {
	// 		if grid[y][x] == '#' {
	// 			rowFill++

	// 			placed := 0
	// 			isLine := false
	// 			for j := 1; j < len(grid[0])-x; j++ {
	// 				if grid[y][x+j] == '#' {
	// 					rowFill++
	// 					placed++
	// 					isLine = true
	// 				} else {
	// 					break
	// 				}
	// 			}

	// 			// check the directions of the line
	// 			preUp := false
	// 			preDown := false
	// 			postUp := false
	// 			postDown := false

	// 			if isLine {
	// 				fmt.Println(y, x)
	// 				if y > 0 && grid[y-1][x] == '#' {
	// 					preUp = true
	// 				} else if y < len(grid)-1 && grid[y+1][x] == '#' {
	// 					preDown = true
	// 				}
	// 				if y > 0 && grid[y-1][x+placed] == '#' {
	// 					postUp = true
	// 				} else if y < len(grid)-1 && grid[y+1][x+placed] == '#' {
	// 					postDown = true
	// 				}
	// 				fmt.Println(preUp, preDown, postUp, postDown, y+1, x+placed)

	// 				if (preUp && postDown) || (postUp && preDown) {
	// 					inside = !inside
	// 				}
	// 			} else {
	// 				inside = !inside
	// 			}

	// 			x += placed
	// 		} else if inside {
	// 			rowFill++
	// 			grid[y][x] = 'x'
	// 		}

	// 		x++
	// 		if x == len(grid[0]) {
	// 			break
	// 		}
	// 	}

	// 	total += rowFill
	// }

	// fmt.Println(total)

	// util.PrintGrid(grid)
}
