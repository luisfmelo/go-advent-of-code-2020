package main

import (
	"advent-of.code.2020/pkg"
	"strings"
	"testing"
)

type TestCase struct {
	input          string
	playingTurn    int
	expectedOutput int
}

func TestGetNumberInPlayingTurn(t *testing.T) {
	testCases := []TestCase{
		{input: "0,3,6", playingTurn: 2020, expectedOutput: 436},
		{input: "1,3,2", playingTurn: 2020, expectedOutput: 1},
		{input: "2,1,3", playingTurn: 2020, expectedOutput: 10},
		{input: "1,2,3", playingTurn: 2020, expectedOutput: 27},
		{input: "2,3,1", playingTurn: 2020, expectedOutput: 78},
		{input: "3,2,1", playingTurn: 2020, expectedOutput: 438},
		{input: "3,1,2", playingTurn: 2020, expectedOutput: 1836},

		{input: "0,3,6", playingTurn: 30000000, expectedOutput: 175594},
		{input: "1,3,2", playingTurn: 30000000, expectedOutput: 2578},
		{input: "2,1,3", playingTurn: 30000000, expectedOutput: 3544142},
		{input: "1,2,3", playingTurn: 30000000, expectedOutput: 261214},
		{input: "2,3,1", playingTurn: 30000000, expectedOutput: 6895259},
		{input: "3,2,1", playingTurn: 30000000, expectedOutput: 18},
		{input: "3,1,2", playingTurn: 30000000, expectedOutput: 362},
	}

	for id, testCase := range testCases {
		input, err := pkg.ReadIntsByDelimiter(strings.NewReader(testCase.input), ",")
		if err != nil {
			panic(err)
		}

		actualResult := GetNumberInPlayingTurn(input, testCase.playingTurn)
		if actualResult != testCase.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, testCase.expectedOutput)
		}
	}
}
