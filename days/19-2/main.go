package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2023/util"
)

var ruleListRegex = regexp.MustCompile(`([a-z]+){(.+)}`)
var ruleRegex = regexp.MustCompile(`([xmas])(>|<)([0-9]+):([a-zA-Z]+)`)

type category int

const (
	categoryNone category = iota
	categoryX
	categoryM
	categoryA
	categoryS
)

type comparison int

const (
	comparisonNone comparison = iota
	comparisonGreater
	comparisonLess
)

type rule struct {
	cat     category
	comp    comparison
	value   int
	mapping string
}

type bound struct {
	min int
	max int
}

type bounds struct {
	x bound
	m bound
	a bound
	s bound
}

func main() {
	sections, err := util.ReadStrings(19, false, "\n\n")
	if err != nil {
		panic(err)
	}

	rulesMap := make(map[string][]rule)
	for _, ruleRow := range strings.Split(sections[0], "\n") {
		ruleListParts := ruleListRegex.FindStringSubmatch(ruleRow)
		for _, r := range strings.Split(ruleListParts[2], ",") {
			if strings.Contains(r, ":") {
				ruleParts := ruleRegex.FindStringSubmatch(r)

				cat := categoryNone
				switch ruleParts[1][0] {
				case 'x':
					cat = categoryX
				case 'm':
					cat = categoryM
				case 'a':
					cat = categoryA
				case 's':
					cat = categoryS
				}

				comp := comparisonNone
				switch ruleParts[2][0] {
				case '>':
					comp = comparisonGreater
				case '<':
					comp = comparisonLess
				}

				value, err := strconv.Atoi(ruleParts[3])
				if err != nil {
					panic(err)
				}

				rulesMap[ruleListParts[1]] = append(rulesMap[ruleListParts[1]], rule{
					cat:     cat,
					comp:    comp,
					value:   value,
					mapping: ruleParts[4],
				})
			} else {
				rulesMap[ruleListParts[1]] = append(rulesMap[ruleListParts[1]], rule{
					mapping: r,
				})
			}
		}

		fmt.Println(ruleRow)
		fmt.Printf("%+v\n", rulesMap[ruleListParts[1]])
	}

	var completedBounds []bounds
	b := bounds{
		x: bound{
			min: 1,
			max: 4000,
		},
		m: bound{
			min: 1,
			max: 4000,
		},
		a: bound{
			min: 1,
			max: 4000,
		},
		s: bound{
			min: 1,
			max: 4000,
		},
	}
	r := rulesMap["in"]
	traverseRule(b, r, 0, rulesMap, &completedBounds)

	total := 0
	for _, completed := range completedBounds {
		fmt.Printf("%+v\n", completed)

		row := (completed.x.max - completed.x.min + 1) *
			(completed.m.max - completed.m.min + 1) *
			(completed.a.max - completed.a.min + 1) *
			(completed.s.max - completed.s.min + 1)
		total += row
	}

	fmt.Println(total)
}

func traverseRule(b bounds, r []rule, rulePos int, rulesMap map[string][]rule, completedBounds *[]bounds) {
	if r[rulePos].cat == categoryNone {
		switch r[rulePos].mapping {
		case "A":
			*completedBounds = append(*completedBounds, b)
			return
		case "R":
			return
		}

		traverseRule(b, rulesMap[r[rulePos].mapping], 0, rulesMap, completedBounds)
	} else {
		alternativeBound := b
		switch r[rulePos].cat {
		case categoryX:
			if r[rulePos].comp == comparisonGreater {
				if b.x.min < r[rulePos].value+1 {
					b.x.min = r[rulePos].value + 1
				}
				if alternativeBound.x.max > r[rulePos].value {
					alternativeBound.x.max = r[rulePos].value
				}
			} else {
				if b.x.max > r[rulePos].value-1 {
					b.x.max = r[rulePos].value - 1
				}
				if alternativeBound.x.min < r[rulePos].value {
					alternativeBound.x.min = r[rulePos].value
				}
			}
		case categoryM:
			if r[rulePos].comp == comparisonGreater {
				if b.m.min < r[rulePos].value+1 {
					b.m.min = r[rulePos].value + 1
				}
				if alternativeBound.m.max > r[rulePos].value {
					alternativeBound.m.max = r[rulePos].value
				}
			} else {
				if b.m.max > r[rulePos].value-1 {
					b.m.max = r[rulePos].value - 1
				}
				if alternativeBound.m.min < r[rulePos].value {
					alternativeBound.m.min = r[rulePos].value
				}
			}
		case categoryA:
			if r[rulePos].comp == comparisonGreater {
				if b.a.min < r[rulePos].value+1 {
					b.a.min = r[rulePos].value + 1
				}
				if alternativeBound.a.max > r[rulePos].value {
					alternativeBound.a.max = r[rulePos].value
				}
			} else {
				if b.m.max > r[rulePos].value-1 {
					b.a.max = r[rulePos].value - 1
				}
				if alternativeBound.a.min < r[rulePos].value {
					alternativeBound.a.min = r[rulePos].value
				}
			}
		case categoryS:
			if r[rulePos].comp == comparisonGreater {
				if b.s.min < r[rulePos].value+1 {
					b.s.min = r[rulePos].value + 1
				}
				if alternativeBound.s.max > r[rulePos].value {
					alternativeBound.s.max = r[rulePos].value
				}
			} else {
				if b.s.max > r[rulePos].value-1 {
					b.s.max = r[rulePos].value - 1
				}
				if alternativeBound.s.min < r[rulePos].value {
					alternativeBound.s.min = r[rulePos].value
				}
			}
		}

		if rulePos < len(r)-1 {
			traverseRule(alternativeBound, r, rulePos+1, rulesMap, completedBounds)
		}

		switch r[rulePos].mapping {
		case "A":
			*completedBounds = append(*completedBounds, b)
			return
		case "R":
			return
		}

		traverseRule(b, rulesMap[r[rulePos].mapping], 0, rulesMap, completedBounds)
	}
}
