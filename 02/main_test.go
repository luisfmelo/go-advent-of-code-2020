package main

import (
	"advent-of.code.2020/pkg"
	"strings"
	"testing"
)

type Policy struct {
	n1   int
	n2   int
	char rune
}

type TestCase struct {
	policy   Policy
	password string
	isValid  bool
}

func TestValidateCountCharPolicy(t *testing.T) {
	testCases := []TestCase{
		{Policy{1, 3, 'a'}, "abcde", true},
		{Policy{1, 3, 'b'}, "cdefg", false},
		{Policy{2, 9, 'c'}, "ccccccccc", true},
	}

	for index, testCase := range testCases {
		result := ValidateCountCharPolicy(testCase.policy.n1, testCase.policy.n2, testCase.policy.char, testCase.password)
		if result != testCase.isValid {
			t.Errorf("Test case #%v: Got %v; Want %v", index+1, result, testCase.isValid)
		}
	}
}

func TestValidateCharPositionPolicy(t *testing.T) {
	testCases := []TestCase{
		{Policy{1, 3, 'a'}, "abcde", true},
		{Policy{1, 3, 'b'}, "cdefg", false},
		{Policy{2, 9, 'c'}, "ccccccccc", false},
	}

	for index, testCase := range testCases {
		result := ValidateCharPositionPolicy(testCase.policy.n1, testCase.policy.n2, testCase.policy.char, testCase.password)
		if result != testCase.isValid {
			t.Errorf("Test case #%v: Got %v; Want %v", index+1, result, testCase.isValid)
		}
	}
}

func TestGetNumberOfValidPasswords(t *testing.T) {
	in := "1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc"
	values, err := pkg.ReadLines(strings.NewReader(in))
	if err != nil {
		panic(err)
	}

	expectedResult := 2
	actualResult := GetNumberOfValidPasswords(values, ValidateCountCharPolicy)
	if actualResult != expectedResult {
		t.Errorf(" Got %v; Want %v", actualResult, expectedResult)
	}

	expectedResult = 1
	actualResult = GetNumberOfValidPasswords(values, ValidateCharPositionPolicy)
	if actualResult != expectedResult {
		t.Errorf(" Got %v; Want %v", actualResult, expectedResult)
	}
}
