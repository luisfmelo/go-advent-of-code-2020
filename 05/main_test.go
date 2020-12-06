package main

import (
	"advent-of.code.2020/pkg"
	"strings"
	"testing"
)

func TestSeat_ID(t *testing.T) {
	var expectedSeatID, actualSeatID int
	var seat Seat

	seat = Seat{1,1}
	expectedSeatID = 9
	actualSeatID = seat.ID()
	if expectedSeatID !=actualSeatID{
		t.Errorf("Got %v; Want %v", actualSeatID, expectedSeatID)
	}

	seat = Seat{10,2}
	expectedSeatID = 82
	actualSeatID = seat.ID()
	if expectedSeatID !=actualSeatID{
		t.Errorf("Got %v; Want %v", actualSeatID, expectedSeatID)
	}
}

func TestGetSeat(t *testing.T) {
	var seatCode string
	var expectedResult, actualResult Seat

	seatCode = "BFFFBBFRRR"
	expectedResult = Seat{70, 7}
	actualResult = GetSeat(seatCode)
	if expectedResult.row != actualResult.row || expectedResult.column != actualResult.column {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	seatCode = "FFFBBBFRRR"
	expectedResult = Seat{14, 7}
	actualResult = GetSeat(seatCode)
	if expectedResult.row != actualResult.row || expectedResult.column != actualResult.column {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	seatCode = "BBFFBBFRLL"
	expectedResult = Seat{102, 4}
	actualResult = GetSeat(seatCode)
	if expectedResult.row != actualResult.row || expectedResult.column != actualResult.column {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	seatCode = "FBFBBFFRLR"
	expectedResult = Seat{44, 5}
	actualResult = GetSeat(seatCode)
	if expectedResult.row != actualResult.row || expectedResult.column != actualResult.column {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}

func TestGetHighestSeatID(t *testing.T) {
	var expectedResult, actualResult int
	in := "BFFFBBFRRR\nFFFBBBFRRR\nBBFFBBFRLL\nFBFBBFFRLR"
	values, err := pkg.ReadLines(strings.NewReader(in))
	if err != nil {
		panic(err)
	}

	// test allowing north pole credentials
	expectedResult = 102*8 + 4
	actualResult = GetHighestSeatID(values)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}
