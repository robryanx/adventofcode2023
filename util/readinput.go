package util

import (
	"os"
	"strconv"
	"strings"
)

type Iterator = func(s string)

func ReadStrings(file string, delim string) ([]string, error) {
	var vals []string

	err := read(file, delim, func(s string) {
		vals = append(vals, s)
	})
	if err != nil {
		return vals, err
	}

	return vals, nil
}

func ReadInts(file string, delim string) ([]int, error) {
	var vals []int

	err := read(file, delim, func(s string) {
		i, _ := strconv.Atoi(s)
		vals = append(vals, i)
	})
	if err != nil {
		return vals, err
	}

	return vals, nil
}

func ReadFloats(file string, delim string) ([]float64, error) {
	var vals []float64

	err := read(file, delim, func(s string) {
		i, _ := strconv.ParseFloat(s, 64)
		vals = append(vals, i)
	})
	if err != nil {
		return vals, err
	}

	return vals, nil
}

func read(file string, delim string, iterator Iterator) error {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	for _, row := range strings.Split(string(bytes), delim) {
		iterator(row)
	}

	return nil
}
