package main

import (
	"advent-of.code.2020/pkg"
	"bufio"
	"fmt"
	"os"
)

func Get3ValuesSumIs2020(values []int) (int, int, int) {
	for index1, v1 := range values {
		for index2, v2 := range values[index1:] {
			for _, v3 := range values[index2:] {
				if v1+v2+v3 == 2020 {
					return v1, v2, v3
				}
			}
		}
	}
	return 0, 0, 0
}

func Get2ValuesSumIs2020(values []int) (int, int) {
	for index, v1 := range values {
		for _, v2 := range values[index:] {
			if v1+v2 == 2020 {
				return v1, v2
			}
		}
	}
	return 0, 0
}

func main() {
	file, err := os.Open("01/input.txt")
	if err != nil {
		panic(err)
	}

	values, err := pkg.ReadInts(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}

	n1, n2 := Get2ValuesSumIs2020(values)
	fmt.Println(n1 * n2)

	n1, n2, n3 := Get3ValuesSumIs2020(values)
	fmt.Println(n1 * n2 * n3)
}
