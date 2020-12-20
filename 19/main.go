package main

import (
	"advent-of.code.2020/pkg"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

type Rule string

func (r Rule) IsTail() bool {
	return r[0] == '"'
}
func (r Rule) Tail() string {
	return string(r[1])
}
func (r Rule) GetPossibilities() []string {
	return strings.Split(string(r), " | ")
}

type Memoize map[int]string

func GetRegexForRule(ruleID int, mapRuleIDToRule map[int]Rule, mem Memoize) string {
	if re, exists := mem[ruleID]; exists {
		return re
	}
	if mapRuleIDToRule[ruleID].IsTail() {
		mem[ruleID] = mapRuleIDToRule[ruleID].Tail()
		return mem[ruleID]
	}
	re := ""
	for _, s := range mapRuleIDToRule[ruleID].GetPossibilities() {
		re += "|"
		for _, e := range strings.Fields(s) {
			re += GetRegexForRule(pkg.StrToInt(e), mapRuleIDToRule, mem)
		}
	}
	mem[ruleID] = "(?:" + re[1:] + ")"
	return mem[ruleID]
}

// activateLoop will transform this:
// 8: 42
// 11: 42 31
// into this
// 8: 42 | 42 8
// 11: 42 31 | 42 11 31
func CountMessagesThatMatchRule0(rules, messages []string, activateLoop bool) int {
	mapRuleIDToRule := map[int]Rule{}
	for _, rule := range rules {
		splitter := strings.Split(rule, ": ")
		mapRuleIDToRule[pkg.StrToInt(splitter[0])] = Rule(splitter[1])
	}

	memoize := Memoize{}
	if activateLoop {
		memoize8 := fmt.Sprintf("(%s)+", GetRegexForRule(42, mapRuleIDToRule, memoize))
		memoize11 := ""
		for i := 1; i <= 10; i++ {
			memoize11 += fmt.Sprintf("|%s{%d}%s{%d}",
				GetRegexForRule(42, mapRuleIDToRule, memoize), i,
				GetRegexForRule(31, mapRuleIDToRule, memoize), i)
		}
		memoize11 = `(?:` + memoize11[1:] + `)`
		memoize = Memoize{8: memoize8, 11: memoize11}
	}

	re := regexp.MustCompile(fmt.Sprintf("(?m)^%s$", GetRegexForRule(0, mapRuleIDToRule, memoize)))
	return len(re.FindAllString(strings.Join(messages, "\n"), -1))
}

func main() {
	file, err := os.Open("19/input.txt")
	if err != nil {
		panic(err)
	}

	input, err := pkg.ReadByDelimiter(bufio.NewReader(file), "\n\n")
	if err != nil {
		panic(err)
	}

	start := time.Now()
	rules := strings.Split(input[0], "\n")
	messages := strings.Split(input[1], "\n")
	log.Printf("1st Part result: %v\n", CountMessagesThatMatchRule0(rules, messages, false))
	log.Printf("1st Part took: %s", time.Since(start))

	start = time.Now()
	messages = append(messages, "8: 42 | 42 8")
	messages = append(messages, "11: 42 31 | 42 11 31")
	log.Printf("2nd Part result: %v\n", CountMessagesThatMatchRule0(rules, messages, true))
	log.Printf("2nd Part took: %s", time.Since(start))
}
