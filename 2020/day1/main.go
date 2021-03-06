package main

import (
	"fmt"
	"github.com/h8nsre/advent-of-code/utils"
	"log"
)

const target = 2020

func twoSum(data []int, target int) int {
	lookup := make(map[int]bool, len(data))

	for _, v := range data {
		if lookup[v] {
			return v * (target - v)
		}
		lookup[target-v] = true
	}
	return 0
}

func threeSum(data []int, target int) int {
	lookup := make(map[int]bool, len(data))

	for _, v := range data {
		lookup[v] = true
	}

	for i, x := range data {
		for _, y := range data[i+1:] {
			z := target - x - y
			if lookup[z] {
				return x * y * z
			}
		}
	}
	return 0
}

func main() {
	ints, err := utils.GetIntsNL("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(twoSum(ints, target), threeSum(ints, target))

}
