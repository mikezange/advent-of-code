package main

import (
	"fmt"
	"strings"

	"github.com/mikezange/advent-of-code/common"
)

func main() {
	input, _ := common.ReadWholeFile("input.txt")
	groups := strings.Split(string(input), "\n\n")

	totalTrue := 0
	totalEveryone := 0

	for _, group := range groups {
		people := strings.Split(group, "\n")
		m := map[rune]int{}
		for _, person := range people {
			for _, answer := range person {
				m[answer]++
			}
		}

		totalTrue += len(m)

		for _, count := range m {
			if count == len(people) {
				totalEveryone++
			}
		}
	}

	fmt.Println(totalTrue)
	fmt.Println(totalEveryone)

}
