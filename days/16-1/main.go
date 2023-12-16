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
	ended     bool
}

func main() {
	rows, err := util.ReadStrings(16, false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	for _, row := range rows {
		grid = append(grid, []byte(row))
	}

	beams := []*dirPos{
		{
			y:         0,
			x:         -1,
			direction: East,
		},
	}

	startingPos := map[string]struct{}{}
	visited := map[int]struct{}{}

	for {
		allEnded := true
		for _, beam := range beams {
			if beam.ended {
				continue
			}

			allEnded = false

			switch beam.direction {
			case East:
				if beam.x < len(grid[0])-1 {
					beam.x++
					visited[beam.y*1000+beam.x] = struct{}{}
					switch grid[beam.y][beam.x] {
					case '|':
						beam.direction = South
						if _, ok := startingPos[fmt.Sprintf("%d-%d-%d", beam.y, beam.x, beam.direction)]; ok {
							beam.ended = true
						} else {
							startingPos[fmt.Sprintf("%d-%d-%d", beam.y, beam.x, beam.direction)] = struct{}{}
						}
						if _, ok := startingPos[fmt.Sprintf("%d-%d-%d", beam.y, beam.x, North)]; !ok {
							beams = append(beams, &dirPos{
								y:         beam.y,
								x:         beam.x,
								direction: North,
							})

							startingPos[fmt.Sprintf("%d-%d-%d", beam.y, beam.x, North)] = struct{}{}
						}
					case '\\':
						beam.direction = South
					case '/':
						beam.direction = North
					}
				} else {
					beam.ended = true
				}
			case West:
				if beam.x > 0 {
					beam.x--
					visited[beam.y*1000+beam.x] = struct{}{}
					switch grid[beam.y][beam.x] {
					case '|':
						beam.direction = South
						if _, ok := startingPos[fmt.Sprintf("%d-%d-%d", beam.y, beam.x, beam.direction)]; ok {
							beam.ended = true
						} else {
							startingPos[fmt.Sprintf("%d-%d-%d", beam.y, beam.x, beam.direction)] = struct{}{}
						}
						if _, ok := startingPos[fmt.Sprintf("%d-%d-%d", beam.y, beam.x, North)]; !ok {
							beams = append(beams, &dirPos{
								y:         beam.y,
								x:         beam.x,
								direction: North,
							})

							startingPos[fmt.Sprintf("%d-%d-%d", beam.y, beam.x, North)] = struct{}{}
						}
					case '\\':
						beam.direction = North
					case '/':
						beam.direction = South
					}
				} else {
					beam.ended = true
				}
			case North:
				if beam.y > 0 {
					beam.y--
					visited[beam.y*1000+beam.x] = struct{}{}
					switch grid[beam.y][beam.x] {
					case '-':
						beam.direction = West
						if _, ok := startingPos[fmt.Sprintf("%d-%d-%d", beam.y, beam.x, beam.direction)]; ok {
							beam.ended = true
						} else {
							startingPos[fmt.Sprintf("%d-%d-%d", beam.y, beam.x, beam.direction)] = struct{}{}
						}
						if _, ok := startingPos[fmt.Sprintf("%d-%d-%d", beam.y, beam.x, East)]; !ok {
							beams = append(beams, &dirPos{
								y:         beam.y,
								x:         beam.x,
								direction: East,
							})

							startingPos[fmt.Sprintf("%d-%d-%d", beam.y, beam.x, East)] = struct{}{}
						}
					case '\\':
						beam.direction = West
					case '/':
						beam.direction = East
					}
				} else {
					beam.ended = true
				}
			case South:
				if beam.y < len(grid)-1 {
					beam.y++
					visited[beam.y*1000+beam.x] = struct{}{}
					switch grid[beam.y][beam.x] {
					case '-':
						beam.direction = West
						if _, ok := startingPos[fmt.Sprintf("%d-%d-%d", beam.y, beam.x, beam.direction)]; ok {
							beam.ended = true
						} else {
							startingPos[fmt.Sprintf("%d-%d-%d", beam.y, beam.x, beam.direction)] = struct{}{}
						}
						if _, ok := startingPos[fmt.Sprintf("%d-%d-%d", beam.y, beam.x, East)]; !ok {
							beams = append(beams, &dirPos{
								y:         beam.y,
								x:         beam.x,
								direction: East,
							})

							startingPos[fmt.Sprintf("%d-%d-%d", beam.y, beam.x, East)] = struct{}{}
						}
					case '\\':
						beam.direction = East
					case '/':
						beam.direction = West
					}
				} else {
					beam.ended = true
				}
			}
		}

		if allEnded {
			break
		}
	}

	fmt.Println(len(visited))
}
