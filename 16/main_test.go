package main

import (
	"advent-of.code.2020/pkg"
	"strings"
	"testing"
)

type TestCase struct {
	input          string
	expectedOutput int
}

func TestGetNumberInPlayingTurn(t *testing.T) {
	testCases := []TestCase{
		{input: "class: 1-3 or 5-7\nrow: 6-11 or 33-44\nseat: 13-40 or 45-50\n\nyour ticket:\n7,1,14\n\nnearby tickets:\n7,3,47\n40,4,50\n55,2,20\n38,6,12", expectedOutput: 71},
	}

	for id, testCase := range testCases {
		input, err := pkg.ReadByDelimiter(strings.NewReader(testCase.input), "\n\n")
		if err != nil {
			panic(err)
		}

		actualResult := GetTicketScanningErrorRate(input)
		if actualResult != testCase.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, testCase.expectedOutput)
		}
	}
}

func TestMultiplyMyTicketNumbersThatHasDepartureAsRuleName(t *testing.T) {
	testCases := []TestCase{
		{input: "class departure: 0-1 or 4-19\nrow: 0-5 or 8-19\nseat departure: 0-13 or 16-19\n\nyour ticket:\n11,12,13\n\nnearby tickets:\n3,9,18\n15,1,5\n5,14,9", expectedOutput: 12 * 13},
	}

	for id, testCase := range testCases {
		input, err := pkg.ReadByDelimiter(strings.NewReader(testCase.input), "\n\n")
		if err != nil {
			panic(err)
		}

		actualResult := MultiplyMyTicketNumbersThatHasDepartureAsRuleName(input)
		if actualResult != testCase.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, testCase.expectedOutput)
		}
	}
}
