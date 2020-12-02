package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var ErrUnableToFind = fmt.Errorf("unable to fing target summation")

func readInput(file string) ([]int, error) {
	r, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func main() {
	nums, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

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
