package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/mikezange/advent-of-code/common"
)

type passwordRule struct {
	min      int
	max      int
	char     byte
	password string
}

func main() {
	lines := common.ReadStrings("input.txt")
	rules := parseLinesToRules(lines)

	var validCount int
	for _, rule := range rules {
		count := bytes.Count([]byte(rule.password), []byte{rule.char})
		if count >= rule.min && count <= rule.max {
			validCount++
		}
	}

	var validCount2 int
	for _, rule := range rules {
		var tmp int
		if rule.password[rule.min-1] == rule.char {
			tmp++
		}

		if rule.password[rule.max-1] == rule.char {
			tmp++
		}

		if tmp == 1 {
			validCount2++
		}
	}

	fmt.Println(validCount)
	fmt.Println(validCount2)
}

func parseLinesToRules(lines []string) []passwordRule {
	var rules []passwordRule
	for _, line := range lines {
		parts := strings.Split(line, " ")
		bounds := strings.Split(parts[0], "-")
		min, _ := strconv.Atoi(bounds[0])
		max, _ := strconv.Atoi(bounds[1])

		rule := passwordRule{
			min:      min,
			max:      max,
			char:     strings.TrimSuffix(parts[1], ":")[0],
			password: parts[2],
		}

		rules = append(rules, rule)
	}

	return rules
}
