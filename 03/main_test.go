package main

import (
	"advent-of.code.2020/pkg"
	"strings"
	"testing"
)

func TestPositionMove(t *testing.T) {
	var cmd Command
	var expectedPosition *Position
	p := &Position{0, 0}

	cmd = Command{right: 0, left: 0, up: 0, down: 0}
	expectedPosition = &Position{0, 0}
	p.move(cmd)
	if p.x != expectedPosition.x || p.y != expectedPosition.y {
		t.Errorf("Got %v; Want %v", p, expectedPosition)
	}

	cmd = Command{right: 1, left: 0, up: 0, down: 0}
	expectedPosition = &Position{1, 0}
	p.move(cmd)
	if p.x != expectedPosition.x || p.y != expectedPosition.y {
		t.Errorf("Got %v; Want %v", p, expectedPosition)
	}

	cmd = Command{right: 0, left: 1, up: 0, down: 0}
	expectedPosition = &Position{0, 0}
	p.move(cmd)
	if p.x != expectedPosition.x || p.y != expectedPosition.y {
		t.Errorf("Got %v; Want %v", p, expectedPosition)
	}

	cmd = Command{right: 0, left: 0, up: 0, down: 1}
	expectedPosition = &Position{0, 1}
	p.move(cmd)
	if p.x != expectedPosition.x || p.y != expectedPosition.y {
		t.Errorf("Got %v; Want %v", p, expectedPosition)
	}

	cmd = Command{right: 0, left: 0, up: 1, down: 0}
	expectedPosition = &Position{0, 0}
	p.move(cmd)
	if p.x != expectedPosition.x || p.y != expectedPosition.y {
		t.Errorf("Got %v; Want %v", p, expectedPosition)
	}

	cmd = Command{right: 1, left: 1, up: 1, down: 1}
	expectedPosition = &Position{0, 0}
	p.move(cmd)
	if p.x != expectedPosition.x || p.y != expectedPosition.y {
		t.Errorf("Got %v; Want %v", p, expectedPosition)
	}

	cmd = Command{right: 1, left: 0, up: 0, down: 1}
	expectedPosition = &Position{1, 1}
	p.move(cmd)
	if p.x != expectedPosition.x || p.y != expectedPosition.y {
		t.Errorf("Got %v; Want %v", p, expectedPosition)
	}
}

func TestCountTreesInPath(t *testing.T) {
	var expectedResult, actualResult int
	in := "..##.......\n#...#...#..\n.#....#..#.\n..#.#...#.#\n.#...##..#.\n..#.##.....\n.#.#.#....#\n.#........#\n#.##...#...\n#...##....#\n.#..#...#.#"
	values, err := pkg.ReadLines(strings.NewReader(in))
	if err != nil {
		panic(err)
	}

	expectedResult = 2
	actualResult = CountTreesInPath(values, Command{right: 1, down: 1})
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	expectedResult = 7
	actualResult = CountTreesInPath(values, Command{right: 3, down: 1})
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	expectedResult = 3
	actualResult = CountTreesInPath(values, Command{right: 5, down: 1})
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	expectedResult = 4
	actualResult = CountTreesInPath(values, Command{right: 7, down: 1})
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	expectedResult = 2
	actualResult = CountTreesInPath(values, Command{right: 1, down: 2})
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}
