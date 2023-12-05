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
	source      passRange
	destination passRange
}

type passRange struct {
	start int
	end   int
}

func main() {
	mappings, err := util.ReadStrings(5, false, "\n\n")
	if err != nil {
		panic(err)
	}

	var seeds []passRange
	seedParts := seedsRegex.FindStringSubmatch(mappings[0])
	seedSplit := strings.Split(seedParts[1], " ")

	for i := 0; i < len(seedSplit); i += 2 {
		start, err := strconv.Atoi(seedSplit[i])
		if err != nil {
			panic(err)
		}

		length, err := strconv.Atoi(seedSplit[i+1])
		if err != nil {
			panic(err)
		}

		seeds = append(seeds, passRange{
			start: start,
			end:   start + length - 1,
		})
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
				source: passRange{
					start: sourceStart,
					end:   sourceStart + length - 1,
				},
				destination: passRange{
					start: destinationStart,
					end:   destinationStart + length - 1,
				},
			})
		}

		maps = append(maps, currentPartMapping)
	}

	lowest := 100000000000
	for i := 0; i < len(seeds); i++ {
		sourceRanges := []passRange{
			{
				start: seeds[i].start,
				end:   seeds[i].end,
			},
		}

		var destinationRanges []passRange
		for _, currentMap := range maps {
			var unhandledSourceRanges []passRange
			for _, values := range currentMap {
				for j := 0; j < len(sourceRanges); j++ {
					// seed range is contained within a single source range
					if sourceRanges[j].start >= values.source.start && sourceRanges[j].end <= values.source.end {
						destinationRanges = append(destinationRanges, passRange{
							start: values.destination.start + (sourceRanges[j].start - values.source.start),
							end:   values.destination.start + (sourceRanges[j].end - values.source.start),
						})
					} else if sourceRanges[j].start < values.source.start && sourceRanges[j].end > values.source.end { // bigger than the range
						destinationRanges = append(destinationRanges, passRange{
							start: values.destination.start,
							end:   values.destination.end,
						})

						unhandledSourceRanges = append(unhandledSourceRanges, passRange{
							start: sourceRanges[j].start,
							end:   values.source.start - 1,
						})

						unhandledSourceRanges = append(unhandledSourceRanges, passRange{
							start: values.source.end + 1,
							end:   sourceRanges[j].end,
						})
					} else if sourceRanges[j].start >= values.source.start && sourceRanges[j].start <= values.source.end && sourceRanges[j].end > values.source.end { // starts in the range but doesn't end in it
						destinationRanges = append(destinationRanges, passRange{
							start: values.destination.start + (sourceRanges[j].start - values.source.start),
							end:   values.destination.end,
						})

						unhandledSourceRanges = append(unhandledSourceRanges, passRange{
							start: values.source.end + 1,
							end:   sourceRanges[j].end,
						})
					} else if sourceRanges[j].start < values.source.start && sourceRanges[j].end >= values.source.start && sourceRanges[j].end <= values.source.end { // ends in the range but doesn't start in it
						destinationRanges = append(destinationRanges, passRange{
							start: values.destination.start,
							end:   values.destination.start + (sourceRanges[j].end - values.source.start),
						})

						unhandledSourceRanges = append(unhandledSourceRanges, passRange{
							start: sourceRanges[j].start,
							end:   values.source.start - 1,
						})
					} else {
						unhandledSourceRanges = append(unhandledSourceRanges, passRange{
							start: sourceRanges[j].start,
							end:   sourceRanges[j].end,
						})
					}
				}

				sourceRanges = unhandledSourceRanges
				unhandledSourceRanges = nil
			}

			for _, sourceRange := range sourceRanges {
				destinationRanges = append(destinationRanges, sourceRange)
			}

			sourceRanges = destinationRanges
			destinationRanges = nil
		}

		for _, sourceRange := range sourceRanges {
			if sourceRange.start < lowest && sourceRange.start > 0 {
				lowest = sourceRange.start
			}
		}
	}

	fmt.Println(lowest)
}
