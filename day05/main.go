package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/mikezange/advent-of-code/common"
)

var replacer = strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")

func main() {
	entries := common.ReadStrings("input.txt")

	var ids []int
	for _, entry := range entries {
		ids = append(ids, getId(entry))
	}
	sort.Ints(ids)

	for i := 0; i < len(ids); i++ {
		if ids[i+1] != ids[i]+1 {
			fmt.Println(ids[i] + 1)
			break
		}
	}
}

func getId(entry string) int {
	binStr := replacer.Replace(entry)
	id, _ := strconv.ParseInt(binStr, 2, 0)
	return int(id)
}
