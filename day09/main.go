package main

import (
	"fmt"
	"sort"

	"github.com/mikezange/advent-of-code/common"
)

func main() {
	nums := common.ReadInts("input.txt")

	invalid := part1(nums, 25)
	weakness := part2(nums, invalid)

	fmt.Printf("Part 1: %d\n", invalid)
	fmt.Printf("Part 2: %d\n", weakness)
}

func part1(input []int, preamble int) int {
loop:
	for index := preamble; index < len(input); index++ {
		for i := 0; i <= preamble; i++ {
			for j := i + 1; j <= preamble; j++ {
				if input[index-i]+input[index-j] == input[index] {
					continue loop
				}
			}
		}

		return input[index]
	}

	return -1
}

func part2(input []int, search int) int {
	for i := 0; i < len(input); i++ {
		for j := i; j < len(input); j++ {
			list := input[i:j]
			sum := 0
			for i := 0; i < len(list); i++ {
				sum += list[i]
			}

			if sum == search {
				sort.Ints(list)
				return list[0] + list[len(list)-1]
			}
		}
	}

	return -1
}
