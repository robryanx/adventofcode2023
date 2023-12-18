package main

import (
	"fmt"
	"slices"
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

	lines := []line{}

	currentPos := pos{}
	for _, row := range rows {
		rowParts := strings.Split(row, " ")

		// distance64, err := strconv.ParseInt(rowParts[2][2:7], 16, 64)
		// if err != nil {
		// 	panic(err)
		// }
		//
		// distance := int(distance64)

		distance, err := strconv.Atoi(rowParts[1])
		if err != nil {
			panic(err)
		}
		// directionLookup[rowParts[2][7:][0]]
		switch byte(rowParts[0][0]) {
		case 'R':
			currentPos.x += distance
		case 'L':
			currentPos.x -= distance
		case 'U':
			lines = append(lines, line{
				start: pos{
					x: currentPos.x,
					y: currentPos.y - distance,
				},
				end: currentPos,
			})
			currentPos.y -= distance
		case 'D':
			lines = append(lines, line{
				start: currentPos,
				end: pos{
					x: currentPos.x,
					y: currentPos.y + distance,
				},
			})
			currentPos.y += distance
		}
	}

	slices.SortFunc(lines, func(a line, b line) int {
		if a.start.y > b.start.y {
			return 1
		} else if a.start.y < b.start.y {
			return -1
		}

		return 0
	})

	total := 0

	var lineSet []line
	current := lines[0].start.y
	nextStart := lines[0].start.y
	nextEnd := 10000000000
	for {
		if current == nextStart {
			maxAdded := -1
			for i := 0; i < len(lines); i++ {
				if lines[i].start.y == nextStart {
					lineSet = append(lineSet, lines[i])
					maxAdded = i
				}
			}

			if maxAdded > -1 {
				lines = lines[maxAdded+1:]
				if len(lines) > 0 {
					nextStart = lines[0].start.y
				} else {
					nextStart = -1
				}
			}

			nextEnd = 10000000000
			for _, l := range lineSet {
				if l.end.y < nextEnd {
					nextEnd = l.end.y
				}
			}
		}

		// sort the line set by x value
		slices.SortFunc(lineSet, func(a line, b line) int {
			if a.start.x > b.start.x {
				return 1
			} else if a.start.x < b.start.x {
				return -1
			}

			return 0
		})

		inside := false
		var prev *line
		across := 0
		for _, l := range lineSet {
			l := l
			if prev == nil {
				prev = &l
				inside = true
			} else if !inside {
				inside = true
			} else {
				across += l.start.x - prev.end.x + 3
				prev = &l
				inside = false
			}
		}

		fmt.Printf("current: %d\n", current)
		fmt.Printf("line set: %+v\n", lineSet)
		fmt.Println(nextEnd)
		fmt.Println(across)
		fmt.Println(across * (nextEnd - current))
		fmt.Println()

		if current == nextEnd {
			var newLineSet []line
			for i := 0; i < len(lineSet); i++ {
				if lineSet[i].end.y != nextEnd {
					newLineSet = append(newLineSet, lineSet[i])
				}
			}

			lineSet = newLineSet
			if len(lineSet) == 0 {
				break
			}

			nextEnd = 10000000000
			for _, l := range lineSet {
				if l.end.y < nextEnd {
					nextEnd = l.end.y
				}
			}
		}

		if nextStart < nextEnd && nextStart != -1 {
			current = nextStart
		} else {
			current = nextEnd
		}
	}

	// fmt.Println(lineSet)
	// fmt.Println(current)
	// fmt.Println(nextStart)
	// fmt.Println(nextEnd)
	fmt.Println(total)

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
