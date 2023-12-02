package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2023/util"
)

var colorMax = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	rows, err := util.ReadStrings(2, false, "\n")
	if err != nil {
		panic(err)
	}

	total := 0

	for _, row := range rows {
		rowParts := strings.Split(row, ":")
		gameParts := strings.Split(rowParts[0], " ")

		gameNum, err := strconv.Atoi(gameParts[1])
		if err != nil {
			panic(err)
		}

		possible := true
		for _, pick := range strings.Split(rowParts[1], ";") {
			for _, part := range strings.Split(pick, ",") {
				partParts := strings.Split(strings.Trim(part, " "), " ")
				colorNum, err := strconv.Atoi(partParts[0])
				if err != nil {
					panic(err)
				}

				if colorMax[partParts[1]] < colorNum {
					possible = false
					break
				}
			}

			if !possible {
				break
			}
		}

		if possible {
			total += gameNum
		}
	}

	fmt.Println(total)
}
