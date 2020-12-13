package main

import (
	"advent-of.code.2020/pkg"
	"strings"
	"testing"
)

type TestCase struct {
	input          string
	expectedOutput uint64
}

func TestGetEarliestBusIDTimesWaitTime(t *testing.T) {

	testCases := []TestCase{
		{
			input:          "939\n7,13,x,x,59,x,31,19",
			expectedOutput: 295,
		},
	}

	for _, testCase := range testCases {
		lines, err := pkg.ReadLines(strings.NewReader(testCase.input))
		if err != nil {
			panic(err)
		}

		actualResult := GetEarliestBusIDTimesWaitTime(lines)
		if actualResult != testCase.expectedOutput {
			t.Errorf("Got %v; Want %v", actualResult, testCase.expectedOutput)
		}
	}
}

func TestGetEarliestTimestampThatAllBusIDsDepartInSubsequentTimestamps(t *testing.T) {

	testCases := []TestCase{
		{input: "\n17,x,13,19", expectedOutput: 3417},
		{input: "\n67,7,59,61", expectedOutput: 754018},
		{input: "\n67,x,7,59,61", expectedOutput: 779210},
		{input: "\n7,13,x,x,59,x,31,19", expectedOutput: 1068781},
		{input: "\n67,7,x,59,61", expectedOutput: 1261476},
		{input: "\n1789,37,47,1889", expectedOutput: 1202161486},
	}

	for _, testCase := range testCases {
		lines, err := pkg.ReadLines(strings.NewReader(testCase.input))
		if err != nil {
			panic(err)
		}

		actualResult := GetEarliestTimestampThatAllBusIDsDepartInSubsequentTimestamps(lines)
		if actualResult != testCase.expectedOutput {
			t.Errorf("Got %v; Want %v", actualResult, testCase.expectedOutput)
		}
	}
}
