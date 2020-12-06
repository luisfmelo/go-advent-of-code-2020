package main

import (
	"advent-of.code.2020/pkg"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ValidatorFunc func(string) bool

func BirthYearValidator(byr string) bool {
	birthYear, err := strconv.Atoi(byr)
	if err != nil {
		return false
	}
	return birthYear >= 1920 && birthYear <= 2002
}

func IssueYearValidator(iyr string) bool {
	issueYear, err := strconv.Atoi(iyr)
	if err != nil {
		return false
	}
	return issueYear >= 2010 && issueYear <= 2020
}

func ExpirationYearValidator(eyr string) bool {
	expirationYear, err := strconv.Atoi(eyr)
	if err != nil {
		return false
	}
	return expirationYear >= 2020 && expirationYear <= 2030
}

func HeightValidator(hgt string) bool {
	v, err := strconv.Atoi(hgt[:len(hgt)-2])
	if err != nil {
		return false
	}
	measure := hgt[len(hgt)-2:]
	switch measure {
	case "cm":
		return v >= 150 && v <= 193
	case "in":
		return v >= 59 && v <= 76
	default:
		return false
	}
}

func HairColorValidator(hcl string) bool {
	for _, c := range hcl[1:] {
		if !(c >= '0' && c <= '9' || c >= 'a' && c <= 'f') {
			return false
		}
	}
	return hcl[0] == '#'
}

func EyeColorValidator(ecl string) bool {
	eyeColors := map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}
	_, eyeColorIsValid := eyeColors[ecl]
	return eyeColorIsValid
}

func PassportIDValidator(pid string) bool {
	if _, err := strconv.Atoi(pid); err != nil {
		return false
	}
	return len(pid) == 9
}

var mapRequiredFieldsToValidator = map[string]ValidatorFunc{
	"byr": BirthYearValidator,
	"iyr": IssueYearValidator,
	"eyr": ExpirationYearValidator,
	"hgt": HeightValidator,
	"hcl": HairColorValidator,
	"ecl": EyeColorValidator,
	"pid": PassportIDValidator,
}

func ValidatePassport(passport string) bool {
	passport = strings.Replace(passport, "\n", " ", -1)
	// validate passport
	fieldsRequiredCounter := 0
	for _, fields := range strings.Split(passport, " ") {
		keyValue := strings.Split(fields, ":")
		if len(keyValue) != 2 || keyValue[1] == "" {
			continue
		}

		if vFunc, isRequired := mapRequiredFieldsToValidator[keyValue[0]]; isRequired && vFunc(keyValue[1]) {
			fieldsRequiredCounter++
		}
	}

	return fieldsRequiredCounter == 7
}

func CountValidPassports(passports []string) int {
	var validPassports = 0
	for _, passport := range passports {
		if isValid := ValidatePassport(passport); isValid {
			validPassports++
		}
	}
	return validPassports
}

func main() {
	file, err := os.Open("04/input.txt")
	if err != nil {
		panic(err)
	}

	lines, err := pkg.ReadByDelimiter(bufio.NewReader(file), "\n\n")
	if err != nil {
		panic(err)
	}

	fmt.Println(CountValidPassports(lines))
}
