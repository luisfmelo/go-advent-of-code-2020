package main

import (
	"advent-of.code.2020/pkg"
	"strings"
	"testing"
)

func TestBirthYearValidator(t *testing.T) {
	var byr string
	var expectedResult, actualResult bool

	byr = "2002"
	expectedResult = true
	actualResult = BirthYearValidator(byr)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	byr = "2003"
	expectedResult = false
	actualResult = BirthYearValidator(byr)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	byr = "a2000"
	expectedResult = false
	actualResult = BirthYearValidator(byr)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}

func TestIssueYearValidator(t *testing.T) {
	var iyr string
	var expectedResult, actualResult bool

	iyr = "2010"
	expectedResult = true
	actualResult = IssueYearValidator(iyr)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	iyr = "2003"
	expectedResult = false
	actualResult = IssueYearValidator(iyr)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	iyr = "a2003"
	expectedResult = false
	actualResult = IssueYearValidator(iyr)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}

func TestExpirationYearValidator(t *testing.T) {
	var eyr string
	var expectedResult, actualResult bool

	eyr = "2020"
	expectedResult = true
	actualResult = ExpirationYearValidator(eyr)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	eyr = "2000"
	expectedResult = false
	actualResult = ExpirationYearValidator(eyr)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	eyr = "a2000"
	expectedResult = false
	actualResult = ExpirationYearValidator(eyr)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}

func TestHeightValidator(t *testing.T) {
	var hgt string
	var expectedResult, actualResult bool

	hgt = "150cm"
	expectedResult = true
	actualResult = HeightValidator(hgt)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	hgt = "10cm"
	expectedResult = false
	actualResult = HeightValidator(hgt)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	hgt = "60in"
	expectedResult = true
	actualResult = HeightValidator(hgt)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	hgt = "200in"
	expectedResult = false
	actualResult = HeightValidator(hgt)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	hgt = "a200in"
	expectedResult = false
	actualResult = HeightValidator(hgt)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	hgt = "2dm"
	expectedResult = false
	actualResult = HeightValidator(hgt)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}

func TestHairColorValidator(t *testing.T) {
	var hcl string
	var expectedResult, actualResult bool

	hcl = "#123abc"
	expectedResult = true
	actualResult = HairColorValidator(hcl)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	hcl = "#123abz"
	expectedResult = false
	actualResult = HairColorValidator(hcl)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	hcl = "123abc"
	expectedResult = false
	actualResult = HairColorValidator(hcl)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}

func TestEyeColorValidator(t *testing.T) {
	var ecl string
	var expectedResult, actualResult bool

	ecl = "brn"
	expectedResult = true
	actualResult = EyeColorValidator(ecl)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	ecl = "wat"
	expectedResult = false
	actualResult = EyeColorValidator(ecl)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}

func TestPassportIDValidator(t *testing.T) {
	var pid string
	var expectedResult, actualResult bool

	pid = "000000001"
	expectedResult = true
	actualResult = PassportIDValidator(pid)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	pid = "0123456789"
	expectedResult = false
	actualResult = PassportIDValidator(pid)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	pid = "0123456789"
	expectedResult = false
	actualResult = PassportIDValidator(pid)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	pid = "a123456789"
	expectedResult = false
	actualResult = PassportIDValidator(pid)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}

func TestValidatePassport(t *testing.T) {
	var passport string
	var expectedResult, actualResult bool

	// test normal input
	passport = "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm"
	expectedResult = true
	actualResult = ValidatePassport(passport)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	// test missing fields
	passport = "hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in"
	expectedResult = false
	actualResult = ValidatePassport(passport)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	// test invalid passports
	passport = "eyr:1972 cid:100 hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926"
	expectedResult = false
	actualResult = ValidatePassport(passport)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	passport = "iyr:2019 hcl:#602927 eyr:1967 hgt:170cm ecl:grn pid:012533040 byr:1946"
	expectedResult = false
	actualResult = ValidatePassport(passport)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	passport = "hcl:dab227 iyr:2012 ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277"
	expectedResult = false
	actualResult = ValidatePassport(passport)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	passport = "hgt:59cm ecl:zzz eyr:2038 hcl:74454a iyr:2023 pid:3556412378 byr:2007"
	expectedResult = false
	actualResult = ValidatePassport(passport)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	// test valid passports
	passport = "pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980 hcl:#623a2f"
	expectedResult = true
	actualResult = ValidatePassport(passport)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	passport = "eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm"
	expectedResult = true
	actualResult = ValidatePassport(passport)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	passport = "hcl:#888785 hgt:164cm byr:2001 iyr:2015 cid:88 pid:545766238 ecl:hzl eyr:2022"
	expectedResult = true
	actualResult = ValidatePassport(passport)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}

	passport = "iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719"
	expectedResult = true
	actualResult = ValidatePassport(passport)
	if expectedResult != actualResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}

func TestCountValidPassports(t *testing.T) {
	var expectedResult, actualResult int
	in := "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\nbyr:1937 iyr:2017 cid:147 hgt:183cm\n\niyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\nhcl:#cfa07d byr:1929\n\nhcl:#ae17e1 iyr:2013\neyr:2024\necl:brn pid:760753108 byr:1931\nhgt:179cm\n\nhcl:#cfa07d eyr:2025 pid:166559648\niyr:2011 ecl:brn hgt:59in"
	values, err := pkg.ReadByDelimiter(strings.NewReader(in), "\n\n")
	if err != nil {
		panic(err)
	}

	// test allowing north pole credentials
	expectedResult = 2
	actualResult = CountValidPassports(values)
	if actualResult != expectedResult {
		t.Errorf("Got %v; Want %v", actualResult, expectedResult)
	}
}
