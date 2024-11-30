package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2023/util"
)

func main() {
	rows, err := util.ReadStrings(23, true, "\n")
	if err != nil {
		panic(err)
	}

	fmt.Println(rows)
}
