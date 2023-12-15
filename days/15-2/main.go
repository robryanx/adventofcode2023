package main

import (
	"fmt"
	"strings"

	"github.com/robryanx/adventofcode2023/util"
)

type lens struct {
	name   string
	length int
}

func main() {
	rows, err := util.ReadStrings(15, false, "\n")
	if err != nil {
		panic(err)
	}

	boxes := make(map[int][]lens, 256)

	total := 0
	for _, sequence := range strings.Split(rows[0], ",") {
		if strings.Contains(sequence, "-") {
			code := sequence[:len(sequence)-1]

			h := hash([]byte(code))
			if _, ok := boxes[h]; ok {
				for i := 0; i < len(boxes[h]); i++ {
					if boxes[h][i].name == code {
						boxes[h] = append(boxes[h][:i], boxes[h][i+1:]...)
						break
					}
				}
			}
		} else {
			parts := strings.Split(sequence, "=")

			l := lens{
				name:   parts[0],
				length: int(parts[1][0] - '0'),
			}

			h := hash([]byte(parts[0]))
			if _, ok := boxes[h]; !ok {
				boxes[h] = []lens{l}
			} else {
				found := false
				for i := 0; i < len(boxes[h]); i++ {
					if boxes[h][i].name == parts[0] {
						found = true
						boxes[h][i] = l
						break
					}
				}

				if !found {
					boxes[h] = append(boxes[h], l)
				}
			}
		}
	}

	for boxPos, box := range boxes {
		for i, lens := range box {
			total += (boxPos + 1) * (i + 1) * lens.length
		}
	}

	fmt.Println(total)
}

func hash(sequence []byte) int {
	hash := 0
	for _, char := range []byte(sequence) {
		hash += int(char)
		hash *= 17
		hash %= 256
	}

	return hash
}
