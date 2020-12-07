package main

import (
	"advent-of.code.2020/pkg"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type InnerBags struct {
	name   string
	number int
}

type Rules map[string][]InnerBags

var innerBagRe = regexp.MustCompile(`(\d+)\s(\w+\s\w+)\sbag(s)?`)

func GetRules(ruleLines []string) Rules {
	mapRules := Rules{}
	for _, ruleLine := range ruleLines {
		ruleLine = ruleLine[:len(ruleLine)-1] // remove the dot

		var innerBags []InnerBags
		splitted := strings.Split(ruleLine, " bags contain ")
		bagName := splitted[0]
		if splitted[1] == "no other bags" {
			innerBags = []InnerBags{}
		} else {
			for _, insideBag := range strings.Split(splitted[1], ", ") {
				number, err := strconv.Atoi(innerBagRe.ReplaceAllString(insideBag, `$1`))
				if err != nil {
					panic(err)
				}
				innerBags = append(innerBags,
					InnerBags{name: innerBagRe.ReplaceAllString(insideBag, `$2`), number: number})
			}
		}
		mapRules[bagName] = innerBags
	}

	return mapRules
}

func GetUniqueOuterBagsThatCanContainCertainBag(rules Rules, searchingBag string, outerBags map[string]bool) map[string]bool {
	for bagName, innerBags := range rules {
		var bagsToSearch []string

		// check each bag for its inside bags
		for _, innerBag := range innerBags {
			if innerBag.name == searchingBag {
				outerBags[bagName] = true
				bagsToSearch = append(bagsToSearch, bagName)
				break
			}
		}
		for _, bagToSearch := range bagsToSearch {
			outerBags = GetUniqueOuterBagsThatCanContainCertainBag(rules, bagToSearch, outerBags)
		}
	}

	return outerBags
}

func CountHowManyBagsCanContainCertainBag(lines []string, searchingBag string) int {
	uniqueOuterBags := GetUniqueOuterBagsThatCanContainCertainBag(GetRules(lines), searchingBag, map[string]bool{})
	return len(uniqueOuterBags)
}

func CountHowManyBagsAreRequiredInsideCertainBag(lines []string, searchingBag string) int {
	var counter int
	rules := GetRules(lines)

	for _, innerBag := range rules[searchingBag] {
		counter += innerBag.number + innerBag.number*CountHowManyBagsAreRequiredInsideCertainBag(lines, innerBag.name)
	}

	return counter
}

func main() {
	file, err := os.Open("07/input.txt")
	if err != nil {
		panic(err)
	}

	lines, err := pkg.ReadLines(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}

	fmt.Println(CountHowManyBagsCanContainCertainBag(lines, "shiny gold"))
	fmt.Println(CountHowManyBagsAreRequiredInsideCertainBag(lines, "shiny gold"))
}
