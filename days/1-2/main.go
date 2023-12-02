package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2023/util"
)

func main() {
	strings, err := util.ReadStrings(1, false, "\n")
	if err != nil {
		panic(err)
	}

	sum := 0

	for _, str := range strings {
		var numbers []int
		chars := []rune(str)
		for i := 0; i < len(chars); i++ {
			if chars[i] >= '0' && chars[i] <= '9' {
				numbers = append(numbers, int(chars[i]-'0'))
			} else if wordNum := isWord(chars, i); wordNum != -1 {
				numbers = append(numbers, wordNum)
			}
		}

		sum += numbers[0]*10 + numbers[len(numbers)-1:][0]
	}

	fmt.Println(sum)
}

func isWord(chars []rune, pos int) int {
	for num, word := range words {
		if len(word)+pos > len(chars) {
			continue
		}

		match := true
		for i := 0; i < len(word); i++ {
			if chars[pos+i] != word[i] {
				match = false
				break
			}
		}

		if match {
			return num + 1
		}
	}

	return -1
}

var words = [][]rune{
	{'o', 'n', 'e'},
	{'t', 'w', 'o'},
	{'t', 'h', 'r', 'e', 'e'},
	{'f', 'o', 'u', 'r'},
	{'f', 'i', 'v', 'e'},
	{'s', 'i', 'x'},
	{'s', 'e', 'v', 'e', 'n'},
	{'e', 'i', 'g', 'h', 't'},
	{'n', 'i', 'n', 'e'},
}
