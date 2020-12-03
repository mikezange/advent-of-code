package main

import (
	"fmt"
	"os"

	"github.com/mikezange/advent-of-code/common"
)

var ErrUnableToFind = fmt.Errorf("unable to find target summation")

func main() {
	nums := common.ReadInts("input.txt")

	multiple, err := findSummation(nums, 2020)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(multiple)

	part2, err := findSummationPart2(nums, 2020)
	if err != nil {
		panic(err)
	}

	fmt.Println(part2)
}

func findSummation(nums []int, target int) (int, error) {
	for _, num := range nums {
		for i := 0; i < len(nums); i++ {
			if num+nums[i] == target {
				return num * nums[i], nil
			}
		}
	}

	return 0, ErrUnableToFind
}

func findSummationPart2(nums []int, target int) (int, error) {
	for _, num := range nums {
		for i := 0; i < len(nums); i++ {
			for j := 0; j < len(nums); j++ {
				if num+nums[i]+nums[j] == target {
					return num * nums[i] * nums[j], nil
				}
			}
		}
	}

	return 0, ErrUnableToFind
}
