package main

import (
	"advent-of.code.2020/pkg"
	"strings"
	"testing"
)


func TestCountYesAnswersFromGroups(t *testing.T) {
	var expectedResult, actualResult int
	in := "abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb"
	values, err := pkg.ReadByDelimiter(strings.NewReader(in), "\n\n")
	if err != nil {
		panic(err)
	}

	expectedResult = 11
	actualResult = CountYesAnswersFromGroups(values)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}

func TestCountUnanimityYesAnswersInGroups(t *testing.T) {
	var expectedResult, actualResult int
	in := "abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb"
	values, err := pkg.ReadByDelimiter(strings.NewReader(in), "\n\n")
	if err != nil {
		panic(err)
	}

	expectedResult = 6
	actualResult = CountUnanimityYesAnswersInGroups(values)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}
