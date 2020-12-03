package main

import (
	"advent-of.code.2020/pkg"
	"bufio"
	"fmt"
	"os"
)

type Command struct {
	right, left, up, down int
}

type Position struct {
	x, y int
}

func (p *Position) move(cmd Command) {
	p.x += cmd.right
	p.x -= cmd.left
	p.y -= cmd.up
	p.y += cmd.down
}

func CountTreesInPath(lines []string, cmd Command) int {
	p := Position{0, 0}

	treesCounter := 0
	for p.y < len(lines) {
		block := lines[p.y][p.x%len(lines[0])]
		if block == '#' {
			treesCounter++
		}
		p.move(cmd)
	}
	return treesCounter
}

func MultiplyTreeCountersInPaths(lines []string, cmds []Command) int {
	result := 1
	for _, cmd := range cmds {
		result *= CountTreesInPath(lines, cmd)
	}
	return result
}

func main() {
	file, err := os.Open("03/input.txt")
	if err != nil {
		panic(err)
	}

	lines, err := pkg.ReadLines(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}

	fmt.Println(CountTreesInPath(lines, Command{right: 3, down: 1}))
	fmt.Println(MultiplyTreeCountersInPaths(lines, []Command{{right: 1, down: 1}, {right: 3, down: 1}, {right: 5, down: 1}, {right: 7, down: 1}, {right: 1, down: 2}}))
}
