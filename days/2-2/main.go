package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2023/util"
)

func main() {
	rows, err := util.ReadStrings(2, false, "\n")
	if err != nil {
		panic(err)
	}

	total := 0
	for _, row := range rows {
		rowParts := strings.Split(row, ":")

		colorMin := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, pick := range strings.Split(rowParts[1], ";") {
			for _, part := range strings.Split(pick, ",") {
				partParts := strings.Split(strings.Trim(part, " "), " ")
				colorNum, err := strconv.Atoi(partParts[0])
				if err != nil {
					panic(err)
				}

				if colorMin[partParts[1]] < colorNum {
					colorMin[partParts[1]] = colorNum
				}
			}
		}

		total += colorMin["red"] * colorMin["green"] * colorMin["blue"]
	}

	fmt.Println(total)
}
