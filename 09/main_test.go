package main

import (
	"advent-of.code.2020/pkg"
	"strings"
	"testing"
)

func TestCheck2NumbersSumFoundInArray(t *testing.T) {
	var expectedResult, actualResult bool
	var arr []int
	var sumToFind int

	arr = []int{35,20,15,25,47}
	sumToFind = 40
	expectedResult = true
	actualResult = Check2NumbersSumFoundInArray(arr, sumToFind)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	arr = []int{95, 102, 117, 150, 182}
	sumToFind = 127
	expectedResult = false
	actualResult = Check2NumbersSumFoundInArray(arr, sumToFind)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}

func TestFindXMASEncryptionWeakness(t *testing.T) {
	var expectedResult, actualResult int
	var in string

	in = "35\n20\n15\n25\n47\n40\n62\n55\n65\n95\n102\n117\n150\n182\n127\n219\n299\n277\n309\n576"
	values, err := pkg.ReadInts(strings.NewReader(in))
	if err != nil {
		panic(err)
	}

	expectedResult = 62
	actualResult = FindXMASEncryptionWeakness(values, 5)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}

func TestFirstFailingNumberInXMAS(t *testing.T) {
	var expectedResult, actualResult int
	var in string

	in = "35\n20\n15\n25\n47\n40\n62\n55\n65\n95\n102\n117\n150\n182\n127\n219\n299\n277\n309\n576"
	values, err := pkg.ReadInts(strings.NewReader(in))
	if err != nil {
		panic(err)
	}

	expectedResult = 127
	actualResult = FirstFailingNumberInXMAS(values, 5)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}
