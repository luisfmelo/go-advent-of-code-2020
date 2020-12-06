package main

import (
	"advent-of.code.2020/pkg"
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Seat struct {
	row, column int
}

func (s Seat) ID() int {
	return s.row*8 + s.column
}

func DiscoverMySeatID(lines []string) int {
	seatIDs := []int{}
	for _, seatCode := range lines {
		seat := GetSeat(seatCode)
		seatIDs = append(seatIDs, seat.ID())
	}

	// order seat IDs
	sort.Ints(seatIDs)

	// find missing seat IDs that have both neighbours (+1 and -1)
	missingSeatIDs := []int{}
	lastValue := seatIDs[0]
	for _, seatID :=  range seatIDs[1: len(seatIDs) - 1] {
		if seatID == lastValue + 2 {
			missingSeatIDs = append(missingSeatIDs, lastValue + 1)
		}
		lastValue = seatID
	}
	if len(missingSeatIDs) != 1 {
		panic("somthing is bad here")
	}

	return missingSeatIDs[0]
}

func GetSeat(seatCode string) Seat {
	var seat Seat

	// discover row
	minRow := 0
	maxRow := 127
	for _, c := range seatCode[:7] {
		if c == 'F' {
			maxRow = maxRow - (maxRow-minRow)/2 - 1
		} else if c == 'B' {
			minRow = (maxRow-minRow)/2 + minRow + 1
		}
	}
	if minRow != maxRow {
		panic("some error occurred while calculating rows")
	}
	seat.row = minRow

	// discover column
	minCol := 0
	maxCol := 7
	for _, c := range seatCode[7:] {
		if c == 'L' {
			maxCol =  maxCol - (maxCol-minCol)/2 - 1
		}
		if c == 'R' {
			minCol =  (maxCol-minCol)/2 + minCol + 1
		}
	}
	if minCol != maxCol {
		panic("some error occurred while calculating columns")
	}
	seat.column = minCol

	return seat
}

func GetHighestSeatID(lines []string) int {
	highestSeatID := 0
	for _, seatCode := range lines {
		seat := GetSeat(seatCode)
		if seatID := seat.ID(); seatID > highestSeatID {
			highestSeatID = seatID
		}
	}
	return highestSeatID
}

func main() {
	file, err := os.Open("05/input.txt")
	if err != nil {
		panic(err)
	}

	lines, err := pkg.ReadLines(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}

	fmt.Println(GetHighestSeatID(lines))
	fmt.Println(DiscoverMySeatID(lines))
}
