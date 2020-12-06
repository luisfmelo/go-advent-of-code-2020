package main

import (
	"advent-of.code.2020/pkg"
	"bufio"
	"fmt"
	"os"
	"strings"
)


func CountUnanimityYesAnswersInGroups(groupsAnswers []string) int {
	counter := 0
	for _, groupAnswers := range groupsAnswers {
		mapGroupAnswerCounter := map[rune]int{}
		groupMembersAnswers := strings.Split(groupAnswers, "\n")
		for _, groupMemberAnswer := range groupMembersAnswers {
			for _, answer := range groupMemberAnswer {
				mapGroupAnswerCounter[answer]++
			}
		}
		for _, answerCounter := range mapGroupAnswerCounter {
			if answerCounter == len(groupMembersAnswers){
				counter++
			}
		}
	}
	return counter
}

func CountYesAnswersFromGroups(groupsAnswers []string) int {
	counter := 0
	for _, groupAnswers := range groupsAnswers {
		mapGroupAnswers := map[rune]bool{}
		for _, answer := range groupAnswers {
			if answer == '\n' {
				continue
			}
			mapGroupAnswers[answer] = true
		}
		counter += len(mapGroupAnswers)
	}
	return counter
}

func main() {
	file, err := os.Open("06/input.txt")
	if err != nil {
		panic(err)
	}

	lines, err := pkg.ReadByDelimiter(bufio.NewReader(file), "\n\n")
	if err != nil {
		panic(err)
	}

	fmt.Println(CountYesAnswersFromGroups(lines))
	fmt.Println(CountUnanimityYesAnswersInGroups(lines))
}
