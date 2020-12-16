package main

import (
	"advent-of.code.2020/pkg"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Ticket struct {
	randomNumbers []int
}

func NewTicket(ticketStr string) Ticket {
	ticket := Ticket{}
	numbersInTicket := strings.Split(ticketStr, ",")
	for _, numberInTicketStr := range numbersInTicket {
		ticket.randomNumbers = append(ticket.randomNumbers, pkg.StrToInt(numberInTicketStr))
	}
	return ticket
}

type SemiRule struct {
	minVal int
	maxVal int
}

type Rule struct {
	name      string
	semiRules []SemiRule
}

func (r Rule) Validate(number int) bool {
	atLeast1SemiRuleValid := false
	for _, semiRule := range r.semiRules {
		if number >= semiRule.minVal && number <= semiRule.maxVal {
			atLeast1SemiRuleValid = true
			break
		}
	}
	return atLeast1SemiRuleValid
}

func (r Rule) ValidateMultiple(numbers []int) bool {
	for _, number := range numbers {
		if !r.Validate(number) {
			return false
		}
	}
	return true
}

func ParseRules(ticketRulesInfo []string) []Rule {
	var rules []Rule
	for _, ticketRuleInfo := range ticketRulesInfo {
		splitter := strings.Split(ticketRuleInfo, ": ")
		// rule name
		ruleName := splitter[0]
		// remove type field
		ticketRuleInfo = splitter[1]

		var semiRules []SemiRule
		semiRulesStr := strings.Split(ticketRuleInfo, " or ")
		for _, semiRuleStr := range semiRulesStr {
			numberInRule := strings.Split(semiRuleStr, "-")
			semiRule := SemiRule{minVal: pkg.StrToInt(numberInRule[0]), maxVal: pkg.StrToInt(numberInRule[1])}
			semiRules = append(semiRules, semiRule)
		}
		rules = append(rules, Rule{ruleName, semiRules})
	}

	return rules
}

// IsTicketValid will evaluate ticket validity against custom rules. It will return if the ticket is valid and
// the number that failed validation. In case of valid ticket, that number is -1
func IsTicketValid(ticket Ticket, rules []Rule) (bool, int) {
	for _, numberInTicket := range ticket.randomNumbers {
		atLeast1RuleValid := false
		for _, rule := range rules {
			if rule.Validate(numberInTicket) {
				atLeast1RuleValid = true
				break
			}
		}
		if !atLeast1RuleValid {
			return false, numberInTicket
		}
	}
	return true, -1
}

func GetTicketScanningErrorRate(input []string) int {
	var errorRate int

	ticketRulesInfo := strings.Split(input[0], "\n")
	nearbyTicketsInfo := strings.Split(input[2], "\n")

	// transpose ticket rule info to array of rules
	rules := ParseRules(ticketRulesInfo)

	// evaluate every nearby tickets against rules
	for _, nearbyTicketInfo := range nearbyTicketsInfo[1:] {
		nearbyTicket := NewTicket(nearbyTicketInfo)
		isTicketValid, failedNumber := IsTicketValid(nearbyTicket, rules)
		if !isTicketValid {
			errorRate += failedNumber
		}
	}

	return errorRate
}

func LogicallyAssignTicketIndexToRule(mapTicketNumberIndexToValidRules map[int]map[string]Rule) map[int]Rule {
	mapTicketIndexNumberToRule := map[int]Rule{}

	for len(mapTicketNumberIndexToValidRules) > 0 {
		for ticketNumberIndex, validRules := range mapTicketNumberIndexToValidRules {
			if len(validRules) == 1 {
				for ruleName := range validRules {
					mapTicketIndexNumberToRule[ticketNumberIndex] = validRules[ruleName]
					// delete that rule name for other indexes
					for i, mapRuleNameToRule := range mapTicketNumberIndexToValidRules {
						for rName, _ := range mapRuleNameToRule {
							if rName == ruleName {
								delete(mapTicketNumberIndexToValidRules[i], ruleName)
							}
						}
					}
					delete(mapTicketNumberIndexToValidRules, ticketNumberIndex)
				}
			}
		}
	}

	return mapTicketIndexNumberToRule
}

func MultiplyMyTicketNumbersThatHasDepartureAsRuleName(input []string) int {
	ticketRulesInfo := strings.Split(input[0], "\n")
	myTicketInfo := strings.Split(input[1], "\n")
	nearbyTicketsInfo := strings.Split(input[2], "\n")

	// transpose ticket rule info to array of rules
	rules := ParseRules(ticketRulesInfo)

	// evaluate every nearby tickets against rules and append valid tickets
	var validTickets []Ticket
	for _, nearbyTicketInfo := range nearbyTicketsInfo[1:] {
		nearbyTicket := NewTicket(nearbyTicketInfo)
		isTicketValid, _ := IsTicketValid(nearbyTicket, rules)
		if isTicketValid {
			validTickets = append(validTickets, nearbyTicket)
		}
	}

	// discover rule names that has departure in name and multiply my ticket numbers
	myTicket := NewTicket(myTicketInfo[1])
	validTickets = append(validTickets, myTicket)

	// map index of ticket numbers to bulk of ticket numbers
	var mapIndexToTicketNumbers = map[int][]int{}
	for _, validTicket := range validTickets {
		for index, ticketNumber := range validTicket.randomNumbers {
			if _, exists := mapIndexToTicketNumbers[index]; !exists {
				mapIndexToTicketNumbers[index] = []int{}
			}
			mapIndexToTicketNumbers[index] = append(mapIndexToTicketNumbers[index], ticketNumber)
		}
	}

	var mapTicketNumberIndexToValidRules = map[int]map[string]Rule{}
	for ticketNumberIndex, ticketNumbers := range mapIndexToTicketNumbers {
		mapTicketNumberIndexToValidRules[ticketNumberIndex] = map[string]Rule{}
		for _, rule := range rules {
			if rule.ValidateMultiple(ticketNumbers) {
				mapTicketNumberIndexToValidRules[ticketNumberIndex][rule.name] = rule
			}
		}
	}

	mapTicketNumberIndexToRule := LogicallyAssignTicketIndexToRule(mapTicketNumberIndexToValidRules)

	// multiply those values of my ticket
	myTicketMultipliesWhenDepartureInRuleName := 1
	for index, rule := range mapTicketNumberIndexToRule {
		if !strings.Contains(rule.name, "departure") {
			continue
		}
		myTicketMultipliesWhenDepartureInRuleName *= myTicket.randomNumbers[index]
	}

	return myTicketMultipliesWhenDepartureInRuleName
}

func main() {
	file, err := os.Open("16/input.txt")
	if err != nil {
		panic(err)
	}

	ticketParts, err := pkg.ReadByDelimiter(bufio.NewReader(file), "\n\n")
	if err != nil {
		panic(err)
	}

	start := time.Now()

	fmt.Println(GetTicketScanningErrorRate(ticketParts))
	log.Printf("First part took %s", time.Since(start))

	start = time.Now()
	fmt.Println(MultiplyMyTicketNumbersThatHasDepartureAsRuleName(ticketParts))
	log.Printf("Second part took %s", time.Since(start))
}
