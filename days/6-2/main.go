package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/robryanx/adventofcode2023/util"
)

var multipleSpace = regexp.MustCompile(`\s+`)

func main() {
	parts, err := util.ReadStrings(6, false, "\n")
	if err != nil {
		panic(err)
	}

	timesStr := ""
	times := multipleSpace.Split(parts[0], -1)
	for _, time := range times[1:] {
		timesStr += time
	}

	time, err := strconv.Atoi(timesStr)
	if err != nil {
		panic(err)
	}

	distancesStr := ""
	distances := multipleSpace.Split(parts[1], -1)
	for _, distance := range distances[1:] {
		distancesStr += distance
	}

	distance, err := strconv.Atoi(distancesStr)
	if err != nil {
		panic(err)
	}

	ways := 0
	for hold := 1; hold < time-1; hold++ {
		traveled := (time - hold) * hold
		if traveled > distance {
			ways++
		}
	}

	fmt.Println(ways)
}
