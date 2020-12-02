package main

import (
	"advent-of.code.2020/pkg"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PasswordPolicyValidator func(int, int, rune, string) bool

func ValidateCharPositionPolicy(pos1, pos2 int, character rune, password string) bool {
	sum := 0
	if password[pos1-1] == byte(character) {
		sum++
	}
	if password[pos2-1] == byte(character) {
		sum++
	}

	return sum == 1
}

func ValidateCountCharPolicy(minOccur, maxOccur int, character rune, password string) bool {
	c := strings.Count(password, string(character))
	return c >= minOccur && c <= maxOccur
}

func GetNumberOfValidPasswords(lines []string, validatePasswordFunc PasswordPolicyValidator) int {
	nValidPasswords := 0

	for _, line := range lines {
		splitted := strings.Split(line, ": ")
		policy := splitted[0]
		password := splitted[1]

		policySplitted := strings.Split(policy[:len(policy)-2], "-")

		n1, err := strconv.Atoi(policySplitted[0])
		if err != nil {
			panic(err)
		}
		n2, err := strconv.Atoi(policySplitted[1])
		if err != nil {
			panic(err)
		}

		character := policy[len(policy)-1]
		if validatePasswordFunc(n1, n2, rune(character), password) {
			nValidPasswords++
		}
	}
	return nValidPasswords
}

func main() {
	file, err := os.Open("02/input.txt")
	if err != nil {
		panic(err)
	}

	lines, err := pkg.ReadLines(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}

	fmt.Println(GetNumberOfValidPasswords(lines, ValidateCountCharPolicy))
	fmt.Println(GetNumberOfValidPasswords(lines, ValidateCharPositionPolicy))

}
