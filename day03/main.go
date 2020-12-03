package main

import (
	"fmt"

	"github.com/mikezange/advent-of-code/common"
)

func main() {
	rows := common.ReadStrings("input.txt")

	results := []int{
		run(1, 1, rows),
		run(3, 1, rows),
		run(5, 1, rows),
		run(7, 1, rows),
		run(1, 2, rows),
	}

	part2Result := 1
	for i := 0; i < len(results); i++ {
		part2Result *= results[i]
	}

	fmt.Printf("Part 1: %d\n", results[1])
	fmt.Printf("Part 2: %d\n", part2Result)

}

func run(x, y int, rows []string) (treeCount int) {
	col, row := x, y
	width := len(rows[0])

	for row < len(rows) {
		if col >= width {
			col -= width
		}

		if rows[row][col] == '#' {
			treeCount++
		}

		row += y
		col += x
	}

	return
}
