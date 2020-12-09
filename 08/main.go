package main

import (
	"advent-of.code.2020/pkg"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	op    string
	value int
}

func InterpretOperationLine(line string) Operation {
	splitter := strings.Split(line, " ")
	operation := splitter[0]
	v, err := strconv.Atoi(splitter[1][1:])
	if err != nil {
		panic(err)
	}
	if splitter[1][0] == '-' {
		v *= -1
	}

	return Operation{op: operation, value: v}
}

func RunBootCodeAndFixLoop(lines []string) int {
	var acc int
	var inLoop bool
	for index, line := range lines {
		if line[:3] == "nop" {
			lines[index] = strings.Replace(line, "nop", "jmp", 1)
			acc, inLoop = RunBootCode(lines)
			if !inLoop {
				return acc
			}
			lines[index] = strings.Replace(line, "jmp", "nop", 1)
		} else if line[:3] == "jmp" {
			lines[index] = strings.Replace(line, "jmp", "nop", 1)
			acc, inLoop = RunBootCode(lines)
			if !inLoop {
				return acc
			}
			lines[index] = strings.Replace(line, "nop", "jmp", 1)
		}
	}

	return -1
}

func RunBootCode(lines []string) (int, bool) {
	var inLoop bool
	var mapOperationsExecuted = map[int]bool{}
	var acc int
	for i := 0; i < len(lines); i++ {
		if _, exists := mapOperationsExecuted[i]; exists {
			inLoop = true
			break
		}
		mapOperationsExecuted[i] = true

		operation := InterpretOperationLine(lines[i])
		switch operation.op {
		case "nop":
			continue
		case "acc":
			acc += operation.value
		case "jmp":
			i += operation.value - 1
			if i < -1 || i > len(lines)-1 {
				return -1, true
			}
		}
	}

	return acc, inLoop
}

func main() {
	file, err := os.Open("08/input.txt")
	if err != nil {
		panic(err)
	}

	lines, err := pkg.ReadLines(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}

	fmt.Println(RunBootCode(lines))
	fmt.Println(RunBootCodeAndFixLoop(lines))
}
