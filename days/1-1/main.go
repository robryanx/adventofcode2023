package main

import (
	"fmt"
	"github.com/robryanx/adventofcode2023/util"
)

func main() {
	strings, err := util.ReadStrings("inputs/1.txt", "\n")
	if err != nil {
		panic(err)
	}

	sum := 0

	for _, str := range strings {
		var numbers []int
		chars := []rune(str)
		for i:=0; i<len(chars); i++ {
			if chars[i] >= '0' && chars[i] <= '9' {
				numbers = append(numbers, int(chars[i] - '0'))
			}
		}

		sum += numbers[0] * 10 + numbers[len(numbers)-1:][0]
	}

	fmt.Println(sum)
}
