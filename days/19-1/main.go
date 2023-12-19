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
var partRegex = regexp.MustCompile(`{x=([0-9]+),m=([0-9]+),a=([0-9]+),s=([0-9]+)}`)

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

func (r rule) matchPart(p part) string {
	if r.cat == categoryNone {
		return r.mapping
	}

	pVal := p.value(r.cat)

	if r.compare(pVal) {
		return r.mapping
	}

	return ""
}

func (r rule) compare(pVal int) bool {
	switch r.comp {
	case comparisonGreater:
		return pVal > r.value
	case comparisonLess:
		return pVal < r.value
	}

	return false
}

type part struct {
	x int
	m int
	a int
	s int
}

func (p part) value(cat category) int {
	switch cat {
	case categoryX:
		return p.x
	case categoryM:
		return p.m
	case categoryA:
		return p.a
	case categoryS:
		return p.s
	}

	return 0
}

func (p part) total() int {
	return p.x + p.m + p.a + p.s
}

func main() {
	sections, err := util.ReadStrings(19, false, "\n\n")
	if err != nil {
		panic(err)
	}

	rulesMap := make(map[string][]rule)

	total := 0
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
	}

	for _, partRow := range strings.Split(sections[1], "\n") {
		partCategories := partRegex.FindStringSubmatch(partRow)

		x, err := strconv.Atoi(partCategories[1])
		if err != nil {
			panic(err)
		}

		m, err := strconv.Atoi(partCategories[2])
		if err != nil {
			panic(err)
		}

		a, err := strconv.Atoi(partCategories[3])
		if err != nil {
			panic(err)
		}

		s, err := strconv.Atoi(partCategories[4])
		if err != nil {
			panic(err)
		}

		p := part{
			x: x,
			m: m,
			a: a,
			s: s,
		}

		accepted := false
		rule := rulesMap["in"]
	complete:
		for {
			for _, subRule := range rule {
				mapping := subRule.matchPart(p)
				if len(mapping) != 0 {
					if mapping == "A" {
						accepted = true
						break complete
					}
					if mapping == "R" {
						break complete
					}

					rule = rulesMap[subRule.mapping]
					break
				}
			}
		}

		if accepted {
			total += p.total()
		}
	}

	fmt.Println(total)
}
