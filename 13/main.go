package main

import (
	"advent-of.code.2020/pkg"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func GetEarliestTimestampThatAllBusIDsDepartInSubsequentTimestamps(input []string) uint64 {
	busIDsStr := strings.Split(input[1], ",")

	var n, a []uint64
	for index, id := range busIDsStr {
		if id == "x" {
			continue
		}

		busID, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}

		n = append(n, uint64(busID))
		a = append(a, uint64(busID-index))
	}

	return ChineseRemainder(n, a)
}

func GetEarliestBusIDTimesWaitTime(input []string) uint64 {
	timestampStart, err := strconv.Atoi(input[0])
	if err != nil {
		panic(err)
	}
	busIDs := strings.Split(input[1], ",")

	bestWaitTime := ^uint64(0) >> 1
	selectedBusID := -1
	for _, id := range busIDs {
		if id == "x" {
			continue
		}
		busID, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}

		busWaitTime := uint64(busID - timestampStart%busID)
		if busWaitTime < bestWaitTime {
			bestWaitTime = busWaitTime
			selectedBusID = busID
		}

	}
	return uint64(selectedBusID) * bestWaitTime
}

func main() {
	file, err := os.Open("13/input.txt")
	if err != nil {
		panic(err)
	}

	lines, err := pkg.ReadLines(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}

	fmt.Println(GetEarliestBusIDTimesWaitTime(lines))
	fmt.Println(GetEarliestTimestampThatAllBusIDsDepartInSubsequentTimestamps(lines))
}

func ModularMultiplicativeInverse(a, m uint64) uint64 {
	a = a % m
	var x uint64
	for x = 1; x < m; x++ {
		if (a*x)%m == 1 {
			return x
		}
	}
	return 1
}

func ChineseRemainder(n, a []uint64) uint64 {
	var sum uint64
	var pairs [][]uint64
	var prod uint64 = 1

	for index, elem := range n {
		prod *= elem
		pairs = append(pairs, []uint64{elem, a[index]})
	}

	for _, pair := range pairs {
		p := uint64(math.Floor(float64(prod) / float64(pair[0])))
		sum += pair[1] * ModularMultiplicativeInverse(p, pair[0]) * p
	}

	return sum % prod

}
