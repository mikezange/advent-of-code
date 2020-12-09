package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mikezange/advent-of-code/common"
)

type instruction struct {
	seen   bool
	opcode string
	count  int
}

func main() {
	lines := common.ReadStrings("input.txt")
	var bootCode []instruction

	for _, op := range lines {
		parts := strings.Split(op, " ")
		count, _ := strconv.Atoi(strings.TrimSpace(parts[1]))

		bootCode = append(bootCode, instruction{
			seen:   false,
			opcode: parts[0],
			count:  count,
		})
	}

	//run the broken code
	acc, _ := runCode(bootCode)
	fmt.Printf("code failed: acc: %d\n", acc)

	//attempt to fix the code
	for i := 0; i < len(bootCode); i++ {
		orig := bootCode[i]
		switch bootCode[i].opcode {
		case "nop":
			bootCode[i] = instruction{opcode: "jmp", count: bootCode[i].count}
		case "jmp":
			bootCode[i] = instruction{opcode: "nop", count: bootCode[i].count}
		default:
			continue
		}

		acc, err := runCode(bootCode)
		if err != nil {
			bootCode[i] = orig
			continue
		}

		fmt.Printf("code finished: acc: %d\n", acc)
		break
	}
}

func runCode(code []instruction) (int, error) {
	acc := 0
	curOp := 0
	seen := map[int]bool{}

	for !seen[curOp] {
		if curOp >= len(code) {
			return acc, nil
		}

		seen[curOp] = true
		op := code[curOp]
		switch op.opcode {
		case "nop":
			curOp++
		case "acc":
			acc += op.count
			curOp++
		case "jmp":
			curOp += op.count
			continue
		}
	}

	return acc, fmt.Errorf("code failed")
}
