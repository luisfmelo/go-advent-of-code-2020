package main

import (
	"reflect"
	"testing"
)

func TestGetRegexForRule(t *testing.T) {
	type TestCase struct {
		mapRuleIDToRawRule map[int]Rule
		ruleID             int
		expectedOutput     string
	}

	testCases := []TestCase{
		{mapRuleIDToRawRule: map[int]Rule{1: "\"a\""}, ruleID: 1, expectedOutput: "a"},
		{mapRuleIDToRawRule: map[int]Rule{1: "2", 2: "\"a\""}, ruleID: 1, expectedOutput: "(?:a)"},
		{mapRuleIDToRawRule: map[int]Rule{1: "2 3", 2: "\"a\"", 3: "\"b\""}, ruleID: 1, expectedOutput: "(?:ab)"},
		{mapRuleIDToRawRule: map[int]Rule{1: "2 2 | 3 3", 2: "\"a\"", 3: "\"b\""}, ruleID: 1, expectedOutput: "(?:aa|bb)"},
		{mapRuleIDToRawRule: map[int]Rule{0: "3 1 2", 1: "3 3 | 4 4", 2: "3 4 | 4 3", 3: "\"a\"", 4: "\"b\""}, ruleID: 0, expectedOutput: "(?:a(?:aa|bb)(?:ab|ba))"},
	}

	for id, testCase := range testCases {
		actualResult := GetRegexForRule(testCase.ruleID, testCase.mapRuleIDToRawRule, map[int]string{})
		if !reflect.DeepEqual(actualResult, testCase.expectedOutput) {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, testCase.expectedOutput)
		}
	}
}

func TestCountMessagesThatMatchRule0(t *testing.T) {
	type TestCase struct {
		rules          []string
		messages       []string
		activateLoop   bool
		expectedOutput int
	}
	testCases := []TestCase{
		{
			rules:          []string{"0: 4 1 5", "1: 2 3 | 3 2", "2: 4 4 | 5 5", "3: 4 5 | 5 4", "4: \"a\"", "5: \"b\""},
			messages:       []string{"ababbb", "bababa", "abbbab", "aaabbb", "aaaabbb"},
			activateLoop: false,
			expectedOutput: 2,
		},
		{
			rules: []string{
				"42: 9 14 | 10 1",
				"9: 14 27 | 1 26",
				"10: 23 14 | 28 1",
				"1: \"a\"",
				"11: 42 31",
				"5: 1 14 | 15 1",
				"19: 14 1 | 14 14",
				"12: 24 14 | 19 1",
				"16: 15 1 | 14 14",
				"31: 14 17 | 1 13",
				"6: 14 14 | 1 14",
				"2: 1 24 | 14 4",
				"0: 8 11",
				"13: 14 3 | 1 12",
				"15: 1 | 14",
				"17: 14 2 | 1 7",
				"23: 25 1 | 22 14",
				"28: 16 1",
				"4: 1 1",
				"20: 14 14 | 1 15",
				"3: 5 14 | 16 1",
				"27: 1 6 | 14 18",
				"14: \"b\"",
				"21: 14 1 | 1 14",
				"25: 1 1 | 1 14",
				"22: 14 14",
				"8: 42",
				"26: 14 22 | 1 20",
				"18: 15 15",
				"7: 14 5 | 1 21",
				"24: 14 1",
			},
			messages: []string{
				"abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa",
				"bbabbbbaabaabba",
				"babbbbaabbbbbabbbbbbaabaaabaaa",
				"aaabbbbbbaaaabaababaabababbabaaabbababababaaa",
				"bbbbbbbaaaabbbbaaabbabaaa",
				"bbbababbbbaaaaaaaabbababaaababaabab",
				"ababaaaaaabaaab",
				"ababaaaaabbbaba",
				"baabbaaaabbaaaababbaababb",
				"abbbbabbbbaaaababbbbbbaaaababb",
				"aaaaabbaabaaaaababaa",
				"aaaabbaaaabbaaa",
				"aaaabbaabbaaaaaaabbbabbbaaabbaabaaa",
				"babaaabbbaaabaababbaabababaaab",
				"aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba",
			},
			activateLoop: false,
			expectedOutput: 3,
		},
		{
			rules: []string{
				"42: 9 14 | 10 1",
				"9: 14 27 | 1 26",
				"10: 23 14 | 28 1",
				"1: \"a\"",
				"11: 42 31",
				"5: 1 14 | 15 1",
				"19: 14 1 | 14 14",
				"12: 24 14 | 19 1",
				"16: 15 1 | 14 14",
				"31: 14 17 | 1 13",
				"6: 14 14 | 1 14",
				"2: 1 24 | 14 4",
				"0: 8 11",
				"13: 14 3 | 1 12",
				"15: 1 | 14",
				"17: 14 2 | 1 7",
				"23: 25 1 | 22 14",
				"28: 16 1",
				"4: 1 1",
				"20: 14 14 | 1 15",
				"3: 5 14 | 16 1",
				"27: 1 6 | 14 18",
				"14: \"b\"",
				"21: 14 1 | 1 14",
				"25: 1 1 | 1 14",
				"22: 14 14",
				"8: 42",
				"26: 14 22 | 1 20",
				"18: 15 15",
				"7: 14 5 | 1 21",
				"24: 14 1",
			},
			messages: []string{
				"abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa",
				"bbabbbbaabaabba",
				"babbbbaabbbbbabbbbbbaabaaabaaa",
				"aaabbbbbbaaaabaababaabababbabaaabbababababaaa",
				"bbbbbbbaaaabbbbaaabbabaaa",
				"bbbababbbbaaaaaaaabbababaaababaabab",
				"ababaaaaaabaaab",
				"ababaaaaabbbaba",
				"baabbaaaabbaaaababbaababb",
				"abbbbabbbbaaaababbbbbbaaaababb",
				"aaaaabbaabaaaaababaa",
				"aaaabbaaaabbaaa",
				"aaaabbaabbaaaaaaabbbabbbaaabbaabaaa",
				"babaaabbbaaabaababbaabababaaab",
				"aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba",
			},
			activateLoop: true,
			expectedOutput: 12,
		},
	}

	for id, testCase := range testCases {
		actualResult := CountMessagesThatMatchRule0(testCase.rules, testCase.messages, testCase.activateLoop)
		if actualResult != testCase.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, testCase.expectedOutput)
		}
	}
}
