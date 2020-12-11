package main

import (
	"advent-of.code.2020/pkg"
	"strings"
	"testing"
)

type TestCase struct {
	input              string
	getNewSeatStateFun NewSeatStateRule
	expectedOutput     int
}

func TestGetNewSeatStatePart1(t *testing.T) {
	var matrix [][]rune
	var actualResult, expectedResult rune

	matrix = [][]rune{
		{'L', '#', 'L', 'L', 'L', 'L'},
		{'#', '#', 'L', 'L', 'L', '#'},
		{'#', '#', 'L', 'L', 'L', '#'},
		{'#', '#', 'L', 'L', 'L', '#'},
		{'#', '#', 'L', 'L', 'L', 'L'},
		{'L', '#', 'L', 'L', 'L', 'L'},
	}

	expectedResult = 'L'
	actualResult = GetNewSeatStatePart1(0, 0, matrix)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	expectedResult = '#'
	actualResult = GetNewSeatStatePart1(1, 3, matrix)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	expectedResult = '#'
	actualResult = GetNewSeatStatePart1(5, 5, matrix)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	expectedResult = 'L'
	actualResult = GetNewSeatStatePart1(0, 5, matrix)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	expectedResult = 'L'
	actualResult = GetNewSeatStatePart1(5, 0, matrix)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	matrix = [][]rune{
		{'#', '.', '#', '#', '.', '#', '#', '.', '#', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '.', '#', '#'},
		{'#', '.', '#', '.', '#', '.', '.', '#', '.', '.'},
		{'#', '#', '#', '#', '.', '#', '#', '.', '#', '#'},
		{'#', '.', '#', '#', '.', '#', '#', '.', '#', '#'},
		{'#', '.', '#', '#', '#', '#', '#', '.', '#', '#'},
		{'.', '.', '#', '.', '#', '.', '.', '.', '.', '.'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '.', '#', '#', '#', '#', '#', '#', '.', '#'},
		{'#', '.', '#', '#', '#', '#', '#', '.', '#', '#'},
	}

	expectedResult = 'L'
	actualResult = GetNewSeatStatePart1(0, 2, matrix)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	expectedResult = '#'
	actualResult = GetNewSeatStatePart1(0, 0, matrix)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	expectedResult = '#'
	actualResult = GetNewSeatStatePart1(9, 0, matrix)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	expectedResult = 'L'
	actualResult = GetNewSeatStatePart1(1, 1, matrix)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	expectedResult = 'L'
	actualResult = GetNewSeatStatePart1(1, 4, matrix)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}

func TestGetNewSeatStatePart2(t *testing.T) {
	var matrix [][]rune
	var actualResult, expectedResult rune

	matrix = [][]rune{
		{'#', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', '#'},
		{'#', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', '.', 'L', '.', '.', 'L', '.', '.'},
		{'L', 'L', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
		{'.', '.', 'L', '.', 'L', '.', '.', '.', '.', '.'},
		{'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', '#'},
		{'#', '.', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L'},
		{'#', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', '#'},
	}

	expectedResult = 'L'
	actualResult = GetNewSeatStatePart2(9, 2, matrix)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}

func TestCountOccupiedSeats(t *testing.T) {
	var matrix [][]rune
	var actualResult, expectedResult int

	matrix = [][]rune{
		{'#', '.', '#', 'L', '.', 'L', '#', '.', '#', '#'},
		{'#', 'L', 'L', 'L', '#', 'L', 'L', '.', 'L', '#'},
		{'L', '.', '#', '.', 'L', '.', '.', '#', '.', '.'},
		{'#', 'L', '#', '#', '.', '#', '#', '.', 'L', '#'},
		{'#', '.', '#', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'#', '.', '#', 'L', '#', 'L', '#', '.', '#', '#'},
		{'.', '.', 'L', '.', 'L', '.', '.', '.', '.', '.'},
		{'#', 'L', '#', 'L', '#', '#', 'L', '#', 'L', '#'},
		{'#', '.', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L'},
		{'#', '.', '#', 'L', '#', 'L', '#', '.', '#', '#'},
	}

	expectedResult = 37
	actualResult = CountOccupiedSeats(matrix)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	matrix = [][]rune{
		{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'L', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', '.', 'L', '.', '.', 'L', '.', '.'},
		{'L', 'L', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
		{'.', '.', 'L', '.', 'L', '.', '.', '.', '.', '.'},
		{'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L'},
		{'L', '.', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L'},
		{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
	}

	expectedResult = 0
	actualResult = CountOccupiedSeats(matrix)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	matrix = [][]rune{
		{'#', '.', '#', '#', '.', '#', '#', '.', '#', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '.', '#', '#'},
		{'#', '.', '#', '.', '#', '.', '.', '#', '.', '.'},
		{'#', '#', '#', '#', '.', '#', '#', '.', '#', '#'},
		{'#', '.', '#', '#', '.', '#', '#', '.', '#', '#'},
		{'#', '.', '#', '#', '#', '#', '#', '.', '#', '#'},
		{'.', '.', '#', '.', '#', '.', '.', '.', '.', '.'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '.', '#', '#', '#', '#', '#', '#', '.', '#'},
		{'#', '.', '#', '#', '#', '#', '#', '.', '#', '#'},
	}

	expectedResult = 71
	actualResult = CountOccupiedSeats(matrix)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}

func TestOccupiedSeatsOnSight(t *testing.T) {
	var matrix [][]rune
	var actualResult, expectedResult int

	matrix = [][]rune{
		{'.', '.', '.', '.', '.', '.', '.', '#', '.'},
		{'.', '.', '.', '#', '.', '.', '.', '.', '.'},
		{'.', '#', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '#', 'L', '.', '.', '.', '.', '#'},
		{'.', '.', '.', '.', '#', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '#', '.', '.', '.', '.', '.'},
	}

	expectedResult = 8
	actualResult = OccupiedSeatsOnSight(4, 3, matrix)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	matrix = [][]rune{
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', 'L', '.', 'L', '.', '#', '.', '#', '.', '#', '.', '#', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}

	expectedResult = 0
	actualResult = OccupiedSeatsOnSight(1, 1, matrix)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	matrix = [][]rune{
		{'.', '#', '#', '.', '#', '#', '.'},
		{'#', '.', '#', '.', '#', '.', '#'},
		{'#', '#', '.', '.', '.', '#', '#'},
		{'.', '.', '.', 'L', '.', '.', '.'},
		{'#', '#', '.', '.', '.', '#', '#'},
		{'#', '.', '#', '.', '#', '.', '#'},
		{'.', '#', '#', '.', '#', '#', '.'},
	}

	expectedResult = 0
	actualResult = OccupiedSeatsOnSight(3, 3, matrix)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	matrix = [][]rune{
		{'#', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', '#'},
		{'#', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', '.', 'L', '.', '.', 'L', '.', '.'},
		{'L', 'L', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
		{'.', '.', 'L', '.', 'L', '.', '.', '.', '.', '.'},
		{'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', '#'},
		{'#', '.', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L'},
		{'#', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', '#'},
	}

	expectedResult = 0
	actualResult = OccupiedSeatsOnSight(0, 3, matrix)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}

func TestNumberOfOccupiedSeatsAfterStabilization(t *testing.T) {
	testCases := []TestCase{
		{
			input:              "L.LL.LL.LL\nLLLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLLL\nL.LLLLLL.L\nL.LLLLL.LL",
			getNewSeatStateFun: GetNewSeatStatePart1,
			expectedOutput:     37,
		},
		{
			input:              "L.LL.LL.LL\nLLLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLLL\nL.LLLLLL.L\nL.LLLLL.LL",
			getNewSeatStateFun: GetNewSeatStatePart2,
			expectedOutput:     26,
		},
	}

	for _, testCase := range testCases {
		matrix, err := pkg.ReadMatrix(strings.NewReader(testCase.input))
		if err != nil {
			panic(err)
		}

		actualResult := NumberOfOccupiedSeatsAfterStabilization(matrix, testCase.getNewSeatStateFun)
		if actualResult != testCase.expectedOutput {
			t.Errorf("Got %v; Want %v", actualResult, testCase.expectedOutput)
		}
	}
}
