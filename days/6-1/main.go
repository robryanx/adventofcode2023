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

	times := multipleSpace.Split(parts[0], -1)
	distances := multipleSpace.Split(parts[1], -1)

	total := 1
	for i := 1; i < len(times); i++ {
		ways := 0
		time, err := strconv.Atoi(times[i])
		if err != nil {
			panic(err)
		}

		distance, err := strconv.Atoi(distances[i])
		if err != nil {
			panic(err)
		}

		for hold := 1; hold < time-1; hold++ {
			traveled := (time - hold) * hold
			if traveled > distance {
				ways++
			}
		}

		total *= ways
	}

	fmt.Println(total)
}
