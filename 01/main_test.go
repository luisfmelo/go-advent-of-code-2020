package main

import (
	"advent-of.code.2020/pkg"
	"strings"
	"testing"
)

func Test2GetValuesSumIs2020(t *testing.T) {
	in := "1721\n979\n366\n299\n675\n1456"
	values, err := pkg.ReadInts(strings.NewReader(in))
	if err != nil {
		panic(err)
	}

	expectedResult1 := 1721
	expectedResult2 := 299

	actualResult1, actualResult2 := Get2ValuesSumIs2020(values)
	if actualResult1*actualResult2 != expectedResult1*expectedResult2 {
		t.Errorf(" Got %d, %d;\nWant %d, %d", actualResult1, actualResult2, expectedResult1, expectedResult2)
	}
}

func Test3GetValuesSumIs2020(t *testing.T) {
	in := "1721\n979\n366\n299\n675\n1456"
	values, err := pkg.ReadInts(strings.NewReader(in))
	if err != nil {
		panic(err)
	}

	expectedResult1 := 979
	expectedResult2 := 366
	expectedResult3 := 675

	actualResult1, actualResult2, actualResult3 := Get3ValuesSumIs2020(values)
	if actualResult1*actualResult2*actualResult3 != expectedResult1*expectedResult2*expectedResult3 {
		t.Errorf(" Got %d, %d, %d;\nWant %d, %d, %d",
			actualResult1, actualResult2, actualResult3, expectedResult1, expectedResult2, expectedResult3)
	}
}
