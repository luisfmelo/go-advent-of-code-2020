package main

import (
	"advent-of.code.2020/pkg"
	"strings"
	"testing"
)

func TestRunBootCodeAndFixLoop(t *testing.T) {
	var expectedResult, actualResult int
	var in string

	in = "nop +0\nacc +1\njmp +4\nacc +3\njmp -3\nacc -99\nacc +1\njmp -4\nacc +6"
	values, err := pkg.ReadLines(strings.NewReader(in))
	if err != nil {
		panic(err)
	}

	expectedResult = 8
	actualResult = RunBootCodeAndFixLoop(values)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}

func TestRunBootCode(t *testing.T) {
	var expectedResult, actualResult int
	var in string

	in = "nop +0\nacc +1\njmp +4\nacc +3\njmp -3\nacc -99\nacc +1\njmp -4\nacc +6"
	values, err := pkg.ReadLines(strings.NewReader(in))
	if err != nil {
		panic(err)
	}

	expectedResult = 5
	actualResult, _ = RunBootCode(values)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}
