package main

import (
	"testing"
)

type TestCase struct {
	expression           string
	simpleExpressionCalc SimpleExpressionCalc
	expectedOutput       int
}

func TestEvaluateComplexExpression(t *testing.T) {
	testCases := []TestCase{
		{expression: "1 + 2 * 3 + 4 * 5 + 6", simpleExpressionCalc: SimpleExpressionCalculatorEqualPrecedence, expectedOutput: 71},
		{expression: "1 + (2 * 3) + (4 * (5 + 6))", simpleExpressionCalc: SimpleExpressionCalculatorEqualPrecedence, expectedOutput: 51},
		{expression: "2 * 3 + (4 * 5)", simpleExpressionCalc: SimpleExpressionCalculatorEqualPrecedence, expectedOutput: 26},
		{expression: "5 + (8 * 3 + 9 + 3 * 4 * 3)", simpleExpressionCalc: SimpleExpressionCalculatorEqualPrecedence, expectedOutput: 437},
		{expression: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", simpleExpressionCalc: SimpleExpressionCalculatorEqualPrecedence, expectedOutput: 12240},
		{expression: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", simpleExpressionCalc: SimpleExpressionCalculatorEqualPrecedence, expectedOutput: 13632},

		{expression: "1 + 2 * 3 + 4 * 5 + 6", simpleExpressionCalc: SimpleExpressionCalculatorSumHasPrecedent, expectedOutput: 231},
		{expression: "1 + (2 * 3) + (4 * (5 + 6))", simpleExpressionCalc: SimpleExpressionCalculatorSumHasPrecedent, expectedOutput: 51},
		{expression: "2 * 3 + (4 * 5)", simpleExpressionCalc: SimpleExpressionCalculatorSumHasPrecedent, expectedOutput: 46},
		{expression: "5 + (8 * 3 + 9 + 3 * 4 * 3)", simpleExpressionCalc: SimpleExpressionCalculatorSumHasPrecedent, expectedOutput: 1445},
		{expression: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", simpleExpressionCalc: SimpleExpressionCalculatorSumHasPrecedent, expectedOutput: 669060},
		{expression: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", simpleExpressionCalc: SimpleExpressionCalculatorSumHasPrecedent, expectedOutput: 23340},
	}

	for id, testCase := range testCases {
		actualResult := EvaluateComplexExpression(testCase.expression, testCase.simpleExpressionCalc)
		if actualResult != testCase.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, testCase.expectedOutput)
		}
	}
}
