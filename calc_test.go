package main

import "testing"

func TestAddition(t *testing.T) {
	tables := []struct {
		firstOperand   int
		operator       string
		secondOperand  int
		expectedResult int
	}{
		{48, "+", 21, 69},
		{-6, "+", 56, 50},
		{-34, "+", -21, -55},
	}

	for _, table := range tables {
		actualResult := PerformCalculation(
			table.firstOperand, table.secondOperand, table.operator,
		)
		if actualResult != table.expectedResult {
			t.Errorf("Test failed: %d + %d expected %d, got %d",
				table.firstOperand,
				table.secondOperand,
				table.expectedResult,
				actualResult,
			)
		}
	}
}

func TestSubstraction(t *testing.T) {
	tables := []struct {
		firstOperand   int
		operator       string
		secondOperand  int
		expectedResult int
	}{
		{48, "-", 21, 27},
		{-6, "-", 56, -62},
		{-34, "-", -35, 1},
	}

	for _, table := range tables {
		actualResult := PerformCalculation(
			table.firstOperand, table.secondOperand, table.operator,
		)
		if actualResult != table.expectedResult {
			t.Errorf("Test failed: %d - %d expected %d, got %d",
				table.firstOperand,
				table.secondOperand,
				table.expectedResult,
				actualResult,
			)
		}
	}
}

func TestMultiplication(t *testing.T) {
	tables := []struct {
		firstOperand   int
		operator       string
		secondOperand  int
		expectedResult int
	}{
		{5, "*", 2, 10},
		{-6, "*", 7, -42},
		{-34, "*", -2, 68},
	}

	for _, table := range tables {
		actualResult := PerformCalculation(
			table.firstOperand, table.secondOperand, table.operator,
		)
		if actualResult != table.expectedResult {
			t.Errorf("Test failed: %d * %d expected %d, got %d",
				table.firstOperand,
				table.secondOperand,
				table.expectedResult,
				actualResult,
			)
		}
	}
}

func TestDivision(t *testing.T) {
	tables := []struct {
		firstOperand   int
		operator       string
		secondOperand  int
		expectedResult int
	}{
		{5, "/", 2, 2},
		{-6, "/", 2, -3},
		{-100, "/", -2, 50},
	}

	for _, table := range tables {
		actualResult := PerformCalculation(
			table.firstOperand, table.secondOperand, table.operator,
		)
		if actualResult != table.expectedResult {
			t.Errorf("Test failed: %d / %d expected %d, got %d",
				table.firstOperand,
				table.secondOperand,
				table.expectedResult,
				actualResult,
			)
		}
	}
}
