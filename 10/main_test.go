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

func TestFindNumberOfDifferentArrangements(t *testing.T) {

	testCases := []TestCase{
		{
			input:          "16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4",
			expectedOutput: 8,
		},
		{
			input:          "28\n33\n18\n42\n31\n14\n46\n20\n48\n47\n24\n23\n49\n45\n19\n38\n39\n11\n1\n32\n25\n35\n8\n17\n7\n9\n4\n2\n34\n10\n3",
			expectedOutput: 19208,
		},
	}

	for _, testCase := range testCases {
		values, err := pkg.ReadInts(strings.NewReader(testCase.input))
		if err != nil {
			panic(err)
		}

		actualResult := FindNumberOfDifferentArrangements(values)
		if actualResult != testCase.expectedOutput {
			t.Errorf("Got %v; Want %v", actualResult, testCase.expectedOutput)
		}
	}
}

func TestFindDevicesChainN1JoltDiffTimesN3DiffJoltDifferences(t *testing.T) {
	testCases := []TestCase{
		{
			input:          "16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4",
			expectedOutput: 35,
		},
		{
			input:          "28\n33\n18\n42\n31\n14\n46\n20\n48\n47\n24\n23\n49\n45\n19\n38\n39\n11\n1\n32\n25\n35\n8\n17\n7\n9\n4\n2\n34\n10\n3",
			expectedOutput: 220,
		},
	}

	for _, testCase := range testCases {
		values, err := pkg.ReadInts(strings.NewReader(testCase.input))
		if err != nil {
			panic(err)
		}

		actualResult := FindDevicesChainN1JoltDiffTimesN3DiffJoltDifferences(values)
		if actualResult != testCase.expectedOutput {
			t.Errorf("Got %v; Want %v", actualResult, testCase.expectedOutput)
		}
	}
}
