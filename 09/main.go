package main

import (
	"advent-of.code.2020/pkg"
	"bufio"
	"fmt"
	"os"
)

func FindXMASEncryptionWeakness(numbers []int, preamble int) int {
	failingNumber := FirstFailingNumberInXMAS(numbers, preamble)

	for index, number := range numbers {
		if number < failingNumber {
			var contagiousSetSum int
			var contagiousSet []int
			for _, n2 := range numbers[index:] {
				contagiousSet = append(contagiousSet, n2)
				contagiousSetSum += n2
				if contagiousSetSum > failingNumber {
					break
				} else if contagiousSetSum == failingNumber {
					min := contagiousSet[0]
					max := contagiousSet[0]
					for _, contagiousSetNumber := range contagiousSet {
						if contagiousSetNumber < min {
							min = contagiousSetNumber
						} else if contagiousSetNumber > max {
							max = contagiousSetNumber
						}
					}
					return min + max
				}
			}
		}
	}

	return -1
}

func Check2NumbersSumFoundInArray(arr []int, sumToFind int) bool {
	for i1, n1 := range arr {
		for _, n2 := range arr[i1:] {
			if n1+n2 == sumToFind {
				return true
			}
		}
	}
	return false
}

func FirstFailingNumberInXMAS(numbers []int, preamble int) int {
	for index, number := range numbers[preamble:] {
		if !Check2NumbersSumFoundInArray(numbers[index:index+preamble], number) {
			return number
		}
	}

	return -1
}

func main() {
	file, err := os.Open("09/input.txt")
	if err != nil {
		panic(err)
	}

	ints, err := pkg.ReadInts(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}

	fmt.Println(FirstFailingNumberInXMAS(ints, 25))
	fmt.Println(FindXMASEncryptionWeakness(ints, 25))
}
