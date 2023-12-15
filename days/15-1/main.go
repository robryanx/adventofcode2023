package main

import (
	"fmt"
	"strings"

	"github.com/robryanx/adventofcode2023/util"
)

func main() {
	rows, err := util.ReadStrings(15, false, "\n")
	if err != nil {
		panic(err)
	}

	total := 0
	for _, sequence := range strings.Split(rows[0], ",") {
		hash := 0
		for _, char := range []byte(sequence) {
			hash += int(char)
			hash *= 17
			hash %= 256
		}

		total += hash
	}

	fmt.Println(total)

}
