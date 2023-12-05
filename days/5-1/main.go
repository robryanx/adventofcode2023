package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2023/util"
)

var seedsRegex = regexp.MustCompile(`^seeds: (.+)$`)

type partMapping struct {
	sourceStart      int
	destinationStart int
	length           int
}

func main() {
	mappings, err := util.ReadStrings(5, false, "\n\n")
	if err != nil {
		panic(err)
	}

	var seeds []int
	seedParts := seedsRegex.FindStringSubmatch(mappings[0])
	for _, seedRaw := range strings.Split(seedParts[1], " ") {
		seed, err := strconv.Atoi(seedRaw)
		if err != nil {
			panic(err)
		}

		seeds = append(seeds, seed)
	}

	var maps [][]partMapping
	for _, mapping := range mappings[1:] {
		mappingLines := strings.Split(mapping, "\n")

		currentPartMapping := []partMapping{}
		for _, mappingLine := range mappingLines[1:] {
			mappingLineNumbersRaw := strings.Split(mappingLine, " ")
			destinationStart, err := strconv.Atoi(mappingLineNumbersRaw[0])
			if err != nil {
				panic(err)
			}
			sourceStart, err := strconv.Atoi(mappingLineNumbersRaw[1])
			if err != nil {
				panic(err)
			}
			length, err := strconv.Atoi(mappingLineNumbersRaw[2])
			if err != nil {
				panic(err)
			}

			currentPartMapping = append(currentPartMapping, partMapping{
				sourceStart:      sourceStart,
				destinationStart: destinationStart,
				length:           length,
			})
		}

		maps = append(maps, currentPartMapping)
	}

	lowestLocation := 10000000000000
	for _, seed := range seeds {
		passValue := seed
		for _, currentMap := range maps {
			for _, values := range currentMap {
				if passValue >= values.sourceStart && passValue <= values.sourceStart+values.length {
					passValue = values.destinationStart + (passValue - values.sourceStart)
					break
				}
			}
		}

		if lowestLocation > passValue {
			lowestLocation = passValue
		}
	}

	fmt.Println(lowestLocation)
}
