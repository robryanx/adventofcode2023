package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2023/util"
)

type pos struct {
	x, y, z int
}

type brick struct {
	start      pos
	end        pos
	supporting []int
}

func newBrick(startParts []string, endParts []string) brick {
	var startInts []int
	for i := 0; i < 3; i++ {
		val, err := strconv.Atoi(startParts[i])
		if err != nil {
			panic(err)
		}
		startInts = append(startInts, val)
	}

	var endInts []int
	for i := 0; i < 3; i++ {
		val, err := strconv.Atoi(endParts[i])
		if err != nil {
			panic(err)
		}
		endInts = append(endInts, val)
	}

	return brick{
		start: pos{startInts[0], startInts[1], startInts[2]},
		end:   pos{endInts[0], endInts[1], endInts[2]},
	}
}

func intersect(a, b brick) bool {
	intersectX := false

	if a.start.x <= b.start.x && a.end.x >= b.end.x { // brick a fully encoses brick b
		intersectX = true
	} else if a.start.x <= b.start.x && a.end.x >= b.start.x && a.end.x <= b.end.x { // brick a starts before but ends during brick b
		intersectX = true
	} else if a.start.x >= b.start.x && a.start.x <= b.end.x && a.end.x >= b.end.x { // brick a starts during but ends after brick b
		intersectX = true
	} else if a.start.x >= b.start.x && a.end.x <= b.end.x { // test brick is fully enclosed
		intersectX = true
	}

	if !intersectX {
		return false
	}

	if a.start.y <= b.start.y && a.end.y >= b.end.y { // brick a fully encoses brick b
		return true
	} else if a.start.y <= b.start.y && a.end.y >= b.start.y && a.end.y <= b.end.y { // brick a starts before but ends during brick b
		return true
	} else if a.start.y >= b.start.y && a.start.y <= b.end.y && a.end.y >= b.end.y { // brick a starts during but ends after brick b
		return true
	} else if a.start.y >= b.start.y && a.end.y <= b.end.y { // test brick is fully enclosed
		return true
	}

	return false
}

func main() {
	rows, err := util.ReadStrings(22, false, "\n")
	if err != nil {
		panic(err)
	}

	var bricks []brick
	for _, row := range rows {
		start, end, _ := strings.Cut(row, "~")

		bricks = append(bricks, newBrick(
			strings.Split(start, ","),
			strings.Split(end, ","),
		))
	}

	slices.SortFunc(bricks, func(a brick, b brick) int {
		if a.start.z > b.start.z {
			return 1
		} else if a.start.z < b.start.z {
			return -1
		}

		return 0
	})

	supportedBy := make(map[int][]int)
	for i := 0; i < len(bricks); i++ {
		if bricks[i].start.z == 1 {
			continue
		}

		// check if the brick can move down
		for {
			testZ := bricks[i].start.z - 1
			canMove := true
			for j := 0; j < i; j++ {
				if bricks[j].end.z >= testZ {
					if intersect(bricks[i], bricks[j]) {
						canMove = false
						bricks[j].supporting = append(bricks[j].supporting, i)
						supportedBy[i] = append(supportedBy[i], j)
					}
				}
			}

			if !canMove {
				break
			}

			bricks[i].start.z--
			bricks[i].end.z--

			if bricks[i].start.z == 1 {
				break
			}
		}
	}

	total := 0
	for i := 0; i < len(bricks); i++ {
		removed := []int{i}
		closed := []int{}
		count := 0
		cascade(bricks, i, supportedBy, &removed, &closed, &count)

		total += (count - 1)
	}

	fmt.Println(total)
}

func cascade(bricks []brick, b int, supportedBy map[int][]int, removed *[]int, closed *[]int, count *int) {
	if slices.Contains(*closed, b) {
		return
	}

	*count++

	cascadeTo := []int{}
	for _, support := range bricks[b].supporting {
		// check if all of the blocks support has been removed
		allRemoved := true
		for _, by := range supportedBy[support] {
			if !slices.Contains(*removed, by) {
				allRemoved = false
				break
			}
		}

		if allRemoved {
			*removed = append(*removed, support)
			cascadeTo = append(cascadeTo, support)
		}
	}

	*closed = append(*closed, b)
	for _, c := range cascadeTo {
		cascade(bricks, c, supportedBy, removed, closed, count)
	}
}
