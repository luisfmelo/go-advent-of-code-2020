package main

import (
	"advent-of.code.2020/pkg"
	"strings"
	"testing"
)

func TestValidatePassport(t *testing.T) {
	var passport string
	var expectedResult, actualResult bool

	// test notmal input
	passport = "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm"
	expectedResult = true
	actualResult = ValidatePassport(passport, false)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	// test north pole credential without allowing that
	passport = "hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm"
	expectedResult = false
	actualResult = ValidatePassport(passport, false)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	// test north pole credential but allowing that
	passport = "hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm"
	expectedResult = true
	actualResult = ValidatePassport(passport, true)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	// test missing fields
	passport = "hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in"
	expectedResult = false
	actualResult = ValidatePassport(passport, true)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}

func TestCountValidPassports(t *testing.T) {
	var expectedResult, actualResult int
	in := "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\nbyr:1937 iyr:2017 cid:147 hgt:183cm\n\niyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\nhcl:#cfa07d byr:1929\n\nhcl:#ae17e1 iyr:2013\neyr:2024\necl:brn pid:760753108 byr:1931\nhgt:179cm\n\nhcl:#cfa07d eyr:2025 pid:166559648\niyr:2011 ecl:brn hgt:59in"
	values, err := pkg.ReadLines(strings.NewReader(in))
	if err != nil {
		panic(err)
	}

	// test allowing north pole credentials
	expectedResult = 2
	actualResult = CountValidPassports(values, true)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	// test not allowing north pole credentials
	expectedResult = 1
	actualResult = CountValidPassports(values, false)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}
