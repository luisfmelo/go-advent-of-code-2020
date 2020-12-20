package main

import (
	"advent-of.code.2020/pkg"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var spaceRegex = regexp.MustCompile(`\s+`)

type ParenthesisPair struct {
	Deep         int
	OpeningIndex int
	ClosingIndex int
}

type mathOperation rune

const (
	sum mathOperation = '+'
	mul               = '*'
)

type SimpleExpressionCalc func (expression string) int

func SimpleExpressionCalculatorSumHasPrecedent(expression string) int {
	// clean expression
	expression = strings.Trim(spaceRegex.ReplaceAllString(expression, " "), " ")

	// do the work
	parameters := strings.Split(expression, " ")
	var onlyMultiplicationsParameters = parameters
	for index, parameter := range parameters {
		if parameter == "+" {
			v1 := pkg.StrToInt(parameters[index-1])
			v2 := pkg.StrToInt(parameters[index+1])
			onlyMultiplicationsParameters[index-1] = ""
			onlyMultiplicationsParameters[index] = ""
			onlyMultiplicationsParameters[index+1] = strconv.Itoa(v1 + v2)
		}
	}
	return SimpleExpressionCalculatorEqualPrecedence(strings.Join(onlyMultiplicationsParameters, " "))
}

func SimpleExpressionCalculatorEqualPrecedence(expression string) int {
	// clean expression
	expression = strings.Trim(spaceRegex.ReplaceAllString(expression, " "), " ")

	// do the work
	result := 0
	parameters := strings.Split(expression, " ")
	var nexOperation = sum
	for _, parameter := range parameters {
		switch parameter {
		case "":
			continue
		case "+":
			nexOperation = sum
		case "*":
			nexOperation = mul
		default:
			value := pkg.StrToInt(parameter)
			switch nexOperation {
			case sum:
				result += value
			case mul:
				result *= value
			default:
				panic("operation not found")
			}
		}
	}
	return result
}

func EvaluateComplexExpression(expression string, simpleExpressionCalc SimpleExpressionCalc) int {
	var deep, maxDeep int
	var parenthesisPairs = map[int]ParenthesisPair{}
	var mapDeepToParenthesisPairs = map[int][]ParenthesisPair{}

	for index, char := range expression {
		switch char {
		case '(':
			// update deep
			deep++
			if deep > maxDeep {
				maxDeep = deep
			}

			// create parenthesis pair
			parenthesisPairs[deep] = ParenthesisPair{deep, index, -1}
		case ')':
			// update parenthesis pairs
			parenthesisPair := parenthesisPairs[deep]
			parenthesisPair.ClosingIndex = index
			parenthesisPairs[deep] = parenthesisPair

			mapDeepToParenthesisPairs[deep] = append(mapDeepToParenthesisPairs[deep], parenthesisPair)
			deep--
		}
	}

	// iterate from max deep to zero to make simple operations and remove the parenthesis
	for d := maxDeep; d > 0; d-- {
		parenthesisPairs := mapDeepToParenthesisPairs[d]
		for _, parenthesisPair := range parenthesisPairs {
			opIndex := parenthesisPair.OpeningIndex
			clIndex := parenthesisPair.ClosingIndex
			// evaluate simple expression
			operationResult := simpleExpressionCalc(expression[opIndex+1 : clIndex])
			// replace parenthesis in full expression with the right amount of parenthesis
			resultStr := strconv.Itoa(operationResult)
			var leadingSpaces, trailingSpaces string
			var numberOfSpacesMissing = float64(clIndex - opIndex + 1.0 - len(resultStr))
			for i := 0.0; i < math.Ceil(numberOfSpacesMissing/2); i++ {
				leadingSpaces += " "
			}
			for i := 0.0; i < math.Floor(numberOfSpacesMissing/2); i++ {
				trailingSpaces += " "
			}
			expression = fmt.Sprintf("%s%s%s%s%s", expression[:opIndex], leadingSpaces, resultStr, trailingSpaces, expression[clIndex+1:])
		}
	}

	return simpleExpressionCalc(expression)
}

func SumExpressionsResults(expressions []string, sumHasPrecedence bool) int {
	var simpleExpressionCalc SimpleExpressionCalc
	if sumHasPrecedence {
		simpleExpressionCalc = SimpleExpressionCalculatorSumHasPrecedent
	} else {
		simpleExpressionCalc = SimpleExpressionCalculatorEqualPrecedence
	}

	totalSum := 0
	for _, expression := range expressions {
		totalSum += EvaluateComplexExpression(expression, simpleExpressionCalc)
	}
	return totalSum
}

func main() {
	file, err := os.Open("18/input.txt")
	if err != nil {
		panic(err)
	}

	expressions, err := pkg.ReadLines(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}

	start := time.Now()
	log.Printf("1st Part result: %v\n", SumExpressionsResults(expressions, false))
	log.Printf("1st Part took: %s", time.Since(start))

	start = time.Now()
	log.Printf("2nd Part result: %v\n", SumExpressionsResults(expressions, true))
	log.Printf("2nd Part took: %s", time.Since(start))
}
