package main

import (
	"advent-of.code.2020/pkg"
	"strings"
	"testing"
)

type TestCase struct {
	input          string
	dimensions int
	expectedOutput int
}

func TestCountActiveCubes(t *testing.T) {
	testCases := []TestCase{
		{input: ".#.\n..#\n###", dimensions: 3, expectedOutput: 112},
		{input: ".#.\n..#\n###", dimensions: 4, expectedOutput: 848},
	}

	for id, testCase := range testCases {
		input, err := pkg.ReadMatrix(strings.NewReader(testCase.input))
		if err != nil {
			panic(err)
		}

		actualResult := CountActiveCubes(input, 6, testCase.dimensions, false)
		if actualResult != testCase.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, testCase.expectedOutput)
		}
	}
}