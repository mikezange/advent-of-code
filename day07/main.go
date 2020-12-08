package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mikezange/advent-of-code/common"
)

type rules map[string][]bag
type bag struct {
	contains string
	count    int
}

func main() {
	input := common.ReadStrings("input.txt")
	rules := rules{}

	for _, line := range input {
		var contains []bag
		parts := strings.Split(line, " bags contain ")

		if !strings.HasSuffix(line, "contain no other bags.") {
			bags := strings.Split(parts[1], ", ")
			for _, b := range bags {
				b = b[:strings.LastIndex(b, "bag")]
				b = strings.TrimSpace(b)
				sepIndex := strings.Index(b, " ")
				num, _ := strconv.Atoi(b[:sepIndex])
				name := b[sepIndex+1:]

				contains = append(contains, bag{
					contains: name,
					count:    num,
				})
			}
		}

		rules[parts[0]] = contains
	}

	var count int
	for name := range rules {
		if hasGold(rules, name) {
			count++
		}
	}

	fmt.Printf("Part 1: %d\n", count)
	fmt.Printf("Part 2: %d\n", calcFor(rules, "shiny gold"))
}

func hasGold(rules rules, name string) bool {
	for _, bag := range rules[name] {
		if bag.contains == "shiny gold" {
			return true
		}

		if hasGold(rules, bag.contains) {
			return true
		}
	}

	return false
}

func calcFor(rules rules, name string) int {
	count := 0
	for _, bag := range rules[name] {
		count += bag.count * (calcFor(rules, bag.contains) + 1)
	}

	return count
}
