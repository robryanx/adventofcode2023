package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/robryanx/adventofcode2023/util"
)

type pos struct {
	y int
	x int
}

type total struct {
	start int
	end   int
	value int
}

var totalLookup = []total{
	// {
	// 	start: 0,
	// 	end:   62,
	// 	value: 12743875833,
	// },
	// {
	// 	start: 126,
	// 	end:   188,
	// 	value: 599106836376,
	// },
	// {
	// 	start: 189,
	// 	end:   251,
	// 	value: 70376010709,
	// },
	// {
	// 	start: 252,
	// 	end:   314,
	// 	value: 58453489475,
	// },
	// {
	// 	start: 315,
	// 	end:   377,
	// 	value: 72059982054,
	// },
	// {
	// 	start: 378,
	// 	end:   440,
	// 	value: 14275601673,
	// },
	// {
	// 	start: 504,
	// 	end:   566,
	// 	value: 188568636,
	// },
	// {
	// 	start: 630,
	// 	end:   692,
	// 	value: 7108169198,
	// },
	// {
	// 	start: 693,
	// 	end:   755,
	// 	value: 16997635383,
	// },
	// {
	// 	start: 756,
	// 	end:   818,
	// 	value: 3045080831,
	// },
	// {
	// 	start: 882,
	// 	end:   944,
	// 	value: 6657808502,
	// },
	// {
	// 	start: 945,
	// 	end:   1007,
	// 	value: 85664464597,
	// },
}

func main() {
	sample := false
	rows, err := util.ReadStrings(12, sample, "\n")
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	total := 0
	completed := 0

	if sample {
		for i := 0; i < 6; i++ {
			wg.Add(1)
			i := i
			go func() {
				defer wg.Done()
				total += runRows(rows[i:i+1], i)
			}()
		}
	} else {
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			i := i
			go func() {
				defer wg.Done()

				found := false
				for _, rangeTotal := range totalLookup {
					if i >= rangeTotal.start && i <= rangeTotal.end {
						found = true
					}
				}

				if !found {
					startTime := time.Now()
					fmt.Printf("%d: Started\n", i)
					collect := runRows(rows[i:i+1], i)
					fmt.Printf("%d: Value: %d - %s\n", i, collect, time.Since(startTime))

					total += collect
				}

				completed++
				fmt.Printf("%d/1000\n", completed)
			}()
		}
	}

	wg.Wait()

	for _, rangeTotal := range totalLookup {
		total += rangeTotal.value
	}

	fmt.Printf("Grand total: %d\n", total)
}

func runRows(rows []string, start int) int {
	total := 0
	for i, row := range rows {
		rowParts := strings.Split(row, " ")

		enlargedPattern := rowParts[0]
		enlargedGroup := rowParts[1]
		for i := 0; i < 4; i++ {
			enlargedPattern += "?" + rowParts[0]
			enlargedGroup += "," + rowParts[1]
		}

		fmt.Printf("run: %d, %d/%d - %s %s\n", start/len(rows), i+start, i+start+len(rows), enlargedPattern, enlargedGroup)

		var groups []int
		for _, numberRaw := range strings.Split(enlargedGroup, ",") {
			num, err := strconv.Atoi(numberRaw)
			if err != nil {
				panic(err)
			}

			groups = append(groups, num)
		}

		completed := 0
		maybeBroken := 0
		pattern := []byte(enlargedPattern)
		for _, pos := range pattern {
			if pos != '.' {
				maybeBroken++
			}
		}
		traverse(pattern, 0, groups, -1, maybeBroken, &completed)

		total += completed
	}

	return total
}

func traverse(pattern []byte, currentPos int, remainingGroups []int, lastBrokenPos, maybeBroken int, completed *int) {
	for currentPos < len(pattern) && pattern[currentPos] == '.' {
		currentPos++
	}

	if currentPos >= len(pattern) && len(remainingGroups) == 0 {
		*completed++

		return
	} else if currentPos >= len(pattern) {
		return
	}

	sum := 0
	for i := 0; i < len(remainingGroups); i++ {
		sum += remainingGroups[i]
	}
	if sum > maybeBroken {
		return
	}

	// check what lengths is it possible to consume here
	lengthCount := 0
	for currentPos+lengthCount < len(pattern) && pattern[currentPos+lengthCount] != '.' {
		lengthCount++
	}

	if (currentPos == 0 || lastBrokenPos != currentPos-1) &&
		len(remainingGroups) != 0 {
		group := remainingGroups[0]

		if group <= lengthCount {
			newGroups := append([]int{}, remainingGroups[1:]...)

			traverse(pattern, currentPos+group, newGroups, currentPos+group-1, maybeBroken-group, completed)
		}
	}

	if pattern[currentPos] == '?' {
		traverse(pattern, currentPos+1, remainingGroups, lastBrokenPos, maybeBroken-1, completed)
	}
}
