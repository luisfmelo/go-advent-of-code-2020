package main

import (
	"advent-of.code.2020/pkg"
	"bufio"
	"fmt"
	"os"
	"sort"
)

func FindNumberOfDifferentArrangements(devicesJolt []int) int {
	devicesJolt = append(devicesJolt, 0)
	sort.Ints(devicesJolt)

	terminationJolt := devicesJolt[len(devicesJolt)-1] + 3
	devicesJolt = append(devicesJolt, terminationJolt)

	mapJolts := map[int]int{}
	for _, jolt := range devicesJolt {
		mapJolts[jolt] = 0
	}

	mapJolts[0] = 1
	for _, n := range devicesJolt {
		for i:= 1; i <=3; i++{
			if v, exists := mapJolts[n-i]; exists {
				mapJolts[n] += v
			}
		}
	}

	return mapJolts[terminationJolt]
}

func FindDevicesChainN1JoltDiffTimesN3DiffJoltDifferences(devicesJolt []int) int {
	sort.Ints(devicesJolt)

	var currentJolt, n1JoltDiff, n3JoltDiff int
	for _, deviceJolt := range devicesJolt {
		if deviceJolt-currentJolt == 1 {
			n1JoltDiff++
		} else if deviceJolt-currentJolt == 3 {
			n3JoltDiff++
		}
		currentJolt = deviceJolt
	}

	// sum my device built in joltage: +3 that the highest rated
	n3JoltDiff++

	return n1JoltDiff * n3JoltDiff
}

func main() {
	defer pkg.Elapsed("AdventO of Code 2020 - Day 10")()
	file, err := os.Open("10/input.txt")
	if err != nil {
		panic(err)
	}

	ints, err := pkg.ReadInts(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}

	fmt.Println(FindDevicesChainN1JoltDiffTimesN3DiffJoltDifferences(ints))
	fmt.Println(FindNumberOfDifferentArrangements(ints))
}
