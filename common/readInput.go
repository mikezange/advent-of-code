package common

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
)

func readInput(path string) []string {
	file, _ := os.Open(path)
	defer file.Close()

	s := bufio.NewScanner(file)

	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines
}

func ReadStrings(file string) []string {
	return readInput(file)
}

func ReadInts(file string) []int {
	var ints []int
	for _, str := range readInput(file) {
		x, _ := strconv.Atoi(str)
		ints = append(ints, x)
	}
	return ints
}

func ReadWholeFile(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}
