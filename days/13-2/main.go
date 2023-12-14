package main

import (
	"fmt"
	"strings"

	"github.com/robryanx/adventofcode2023/util"
)

func main() {
	rowCollection, err := util.ReadStrings(13, false, "\n\n")
	if err != nil {
		panic(err)
	}

	total := 0
	for _, rows := range rowCollection {
		var grid [][]byte
		for _, row := range strings.Split(rows, "\n") {
			grid = append(grid, []byte(row))
		}

		// find vertical symmertry
		for x := 0; x < len(grid[0])-1; x++ {
			difference := 0
			for y := 0; y < len(grid); y++ {
				if grid[y][x] != grid[y][x+1] {
					difference++
					if difference > 1 {
						break
					}
				}
			}

			if difference <= 1 {
				if x > 0 && x+1 < len(grid[0]) {
					xCheck := 0
					xPreCheck := x
					xPostCheck := len(grid[0]) - x - 2
					if xPreCheck > xPostCheck {
						xCheck = xPostCheck
					} else {
						xCheck = xPreCheck
					}

				checkX:
					for i := 1; i <= xCheck; i++ {
						for y := 0; y < len(grid); y++ {
							if grid[y][x-i] != grid[y][x+1+i] {
								difference++
								if difference > 1 {
									break checkX
								}
							}
						}
					}
				}

				if difference == 1 {
					total += x + 1
				}
			}
		}

		// find horizontal symmertry
		for y := 0; y < len(grid)-1; y++ {
			difference := 0
			for x := 0; x < len(grid[0]); x++ {
				if grid[y][x] != grid[y+1][x] {
					difference++
					if difference > 1 {
						break
					}
				}
			}

			if difference <= 1 {
				if y > 0 && y+1 < len(grid) {
					yCheck := 0
					yPreCheck := y
					yPostCheck := len(grid) - y - 2
					if yPreCheck > yPostCheck {
						yCheck = yPostCheck
					} else {
						yCheck = yPreCheck
					}

				checkY:
					for i := 1; i <= yCheck; i++ {
						for x := 0; x < len(grid[0]); x++ {
							if grid[y-i][x] != grid[y+1+i][x] {
								difference++
								if difference > 1 {
									break checkY
								}
							}
						}
					}
				}

				if difference == 1 {
					total += 100 * (y + 1)
				}
			}
		}
	}

	fmt.Println(total)
}
