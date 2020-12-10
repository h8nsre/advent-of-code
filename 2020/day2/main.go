package main

import (
	"errors"
	"fmt"
	"github.com/h8nsre/advent-of-code/utils"
	"log"
	"strconv"
	"strings"
)

type Range struct {
	Max int
	Min int
}

func NewRange(r string) (*Range, error) {
	splitRanges := strings.Split(r, "-")

	if len(splitRanges) != 2 {
		return nil, fmt.Errorf("range split of length %d", len(splitRanges))
	}

	min, err := strconv.Atoi(splitRanges[0])

	if err != nil {
		return nil, err
	}

	max, err := strconv.Atoi(splitRanges[1])

	if err != nil {
		return nil, err
	}

	newRange := &Range{
		Max: max,
		Min: min,
	}

	return newRange, nil
}

func valid(s string) (bool, bool, error) {
	countValid1 := 0

	split := strings.Split(s, " ")

	if len(split) != 3 {
		return false, false, errors.New("can't split input")
	}

	newRange, err := NewRange(split[0])

	if err != nil {
		return false, false, err
	}

	if len(split[1]) != 2 {
		return false, false, errors.New("too many letter bytes")
	}

	letter := split[1][0]

	letterRune := rune(letter)

	pass := split[2]

	if newRange.Min < 1 || newRange.Max > len(pass) {
		return false, false, fmt.Errorf("range %v does not fit password length %d", newRange, len(pass))
	}

	for _, c := range pass {
		if c == letterRune {
			countValid1++
		}
	}

	valid1 := countValid1 >= newRange.Min && countValid1 <= newRange.Max

	valid2 := (rune(pass[newRange.Min-1]) != letterRune && rune(pass[newRange.Max-1]) == letterRune) || (rune(pass[newRange.Min-1]) == letterRune && rune(pass[newRange.Max-1]) != letterRune)

	return valid1, valid2, nil
}

func main() {
	var count1, count2 int
	err := utils.ScanLine("input.txt", func(s string) error {
		v1, v2, err := valid(s)
		if err != nil {
			return err
		}
		if v1 {
			count1++
		}
		if v2 {
			count2++
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count1, count2)
}
