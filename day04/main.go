package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mikezange/advent-of-code/common"
)

type passport map[string]string

func main() {
	input, _ := common.ReadWholeFile("input.txt")

	records := strings.Split(string(input), "\n\n")
	for i, record := range records {
		records[i] = strings.ReplaceAll(record, "\n", " ")
	}

	passports := parseRecordsToPassports(records, " ", ":")

	var validPassports int
	var validPassports2 int
	for _, passport := range passports {
		if passport.isValid() {
			validPassports++
		}

		if passport.isValidPart2() {
			validPassports2++
		}
	}

	fmt.Println(validPassports)
	fmt.Println(validPassports2)
}

func (p passport) isValid() bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, field := range requiredFields {
		if _, ok := p[field]; !ok {
			return false
		}
	}
	return true
}

func (p passport) isValidPart2() bool {
	if !p.isValid() {
		return false
	}

	delete(p, "cid")
	rules := map[string]*regexp.Regexp{
		"byr": regexp.MustCompile(`(19[2-9]\d)|(200[0-2])`),
		"iyr": regexp.MustCompile(`201\d|2020`),
		"eyr": regexp.MustCompile(`202\d|2030`),
		"hgt": regexp.MustCompile(`(1[5-8]\d|19[0-3])cm|(59|6\d|7[0-6])in`),
		"hcl": regexp.MustCompile(`#[\da-f]{6}`),
		"ecl": regexp.MustCompile(`amb|blu|brn|gry|grn|hzl|oth`),
		"pid": regexp.MustCompile(`\d{9}`),
	}

	for key, value := range p {
		if key == "pid" && len(value) > 9 {
			return false
		}

		if !rules[key].MatchString(value) {
			return false
		}

	}

	return true
}

func parseRecordsToPassports(records []string, fieldDelim string, attributeDelim string) []passport {
	var passports []passport

	for _, record := range records {
		m := make(passport)
		for _, pair := range strings.Split(record, fieldDelim) {
			z := strings.Split(pair, attributeDelim)
			m[z[0]] = strings.TrimSpace(z[1])
		}

		passports = append(passports, m)
	}

	return passports
}
