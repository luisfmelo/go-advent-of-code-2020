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

func TestDecoderVersion1(t *testing.T) {
	testCases := []TestCase{
		{
			input:          "mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X\nmem[8] = 11\nmem[7] = 101\nmem[8] = 0",
			expectedOutput: 165,
		},
		{
			input:          "mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X\nmem[8] = 11\nmem[7] = 101\nmem[8] = 0\nmask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX0\nmem[1] = 7\nmem[7] = 15\n",
			expectedOutput: 84,
		},
	}

	for _, testCase := range testCases {
		lines, err := pkg.ReadLines(strings.NewReader(testCase.input))
		if err != nil {
			panic(err)
		}

		actualResult := DecoderVersion1(lines)
		if actualResult != testCase.expectedOutput {
			t.Errorf("Got %v; Want %v", actualResult, testCase.expectedOutput)
		}
	}
}
func TestDecoderVersion2(t *testing.T) {
	testCases := []TestCase{
		{
			input:          "mask = 000000000000000000000000000000X1001X\nmem[42] = 100\nmask = 00000000000000000000000000000000X0XX\nmem[26] = 1",
			expectedOutput: 208,
		},
	}

	for _, testCase := range testCases {
		lines, err := pkg.ReadLines(strings.NewReader(testCase.input))
		if err != nil {
			panic(err)
		}

		actualResult := DecoderVersion2(lines)
		if actualResult != testCase.expectedOutput {
			t.Errorf("Got %v; Want %v", actualResult, testCase.expectedOutput)
		}
	}
}
