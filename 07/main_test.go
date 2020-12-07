package main

import (
	"advent-of.code.2020/pkg"
	"strings"
	"testing"
)


func TestCountHowManyBagsCanContainCertainBag(t *testing.T) {
	var expectedResult, actualResult int
	var in string

	in = "light red bags contain 1 bright white bag, 2 muted yellow bags.\ndark orange bags contain 3 bright white bags, 4 muted yellow bags.\nbright white bags contain 1 shiny gold bag.\nmuted yellow bags contain 2 shiny gold bags, 9 faded blue bags.\nshiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.\ndark olive bags contain 3 faded blue bags, 4 dotted black bags.\nvibrant plum bags contain 5 faded blue bags, 6 dotted black bags.\nfaded blue bags contain no other bags.\ndotted black bags contain no other bags."
	values, err := pkg.ReadLines(strings.NewReader(in))
	if err != nil {
		panic(err)
	}

	expectedResult = 4
	actualResult = CountHowManyBagsCanContainCertainBag(values, "shiny gold")
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}

func TestCountHowManyBagsAreRequiredInsideCertainBag(t *testing.T) {
	var expectedResult, actualResult int
	var in string

	in = "light red bags contain 1 bright white bag, 2 muted yellow bags.\ndark orange bags contain 3 bright white bags, 4 muted yellow bags.\nbright white bags contain 1 shiny gold bag.\nmuted yellow bags contain 2 shiny gold bags, 9 faded blue bags.\nshiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.\ndark olive bags contain 3 faded blue bags, 4 dotted black bags.\nvibrant plum bags contain 5 faded blue bags, 6 dotted black bags.\nfaded blue bags contain no other bags.\ndotted black bags contain no other bags."
	values, err := pkg.ReadLines(strings.NewReader(in))
	if err != nil {
		panic(err)
	}

	expectedResult = 32
	actualResult = CountHowManyBagsAreRequiredInsideCertainBag(values, "shiny gold")
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	in = "shiny gold bags contain 2 dark red bags.\ndark red bags contain 2 dark orange bags.\ndark orange bags contain 2 dark yellow bags.\ndark yellow bags contain 2 dark green bags.\ndark green bags contain 2 dark blue bags.\ndark blue bags contain 2 dark violet bags.\ndark violet bags contain no other bags."
	values, err = pkg.ReadLines(strings.NewReader(in))
	if err != nil {
		panic(err)
	}

	expectedResult = 126
	actualResult = CountHowManyBagsAreRequiredInsideCertainBag(values, "shiny gold")
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}
