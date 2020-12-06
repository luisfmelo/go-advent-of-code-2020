package main

import (
	"advent-of.code.2020/pkg"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var mapRequiredFields = map[string]bool{
	"byr": true,
	"iyr": true,
	"eyr": true,
	"hgt": true,
	"hcl": true,
	"ecl": true,
	"pid": true,
}

func ValidatePassport(passport string, acceptNorthPoleCredentials bool) bool{
	// validate passport
	fieldsRequiredCounter := 0
	for _, fields := range strings.Split(passport, " ") {
		keyValue := strings.Split(fields, ":")
		if len(keyValue) != 2 || keyValue[1] == "" {
			continue
		}

		if _, isRequired := mapRequiredFields[keyValue[0]]; isRequired {
			fieldsRequiredCounter++
		}
	}

	return fieldsRequiredCounter == 7
}

func CountValidPassports(lines []string, acceptNorthPoleCredentials bool) int {
	var passport []string
	var validPassports = 0
	for index, line := range lines {
		if line == "" || index == len(lines) - 1{
			if isValid := ValidatePassport(strings.Join(passport, " "), acceptNorthPoleCredentials); isValid{
				validPassports++
			}

			// clear old passport
			passport = []string{}
			continue
		}
		passport = append(passport, line)
	}
	return validPassports
}

func main() {
	file, err := os.Open("04/input.txt")
	if err != nil {
		panic(err)
	}

	lines, err := pkg.ReadLines(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}

	fmt.Println(CountValidPassports(lines, true))
}
