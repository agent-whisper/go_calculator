/*
Ideas for improvement:
- Use float value
- Add modulo, power, and root
- Parse a complete arithmetic string instead
- Able to get function as input, and perform calculation multiple times
using that function
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// EOL is for handling strings on windows
const EOL = '\r'

func main() {
	firstOperand := readCliInputAsInt("Enter first number: ")
	operator := readCliInputAsStr("Enter operator (+, -, *, /): ")
	secondOperand := readCliInputAsInt("Enter second number: ")

	CalculateAndPrintResult(firstOperand, secondOperand, operator)
}

// CalculateAndPrintResult prints a formatted string showing the operands, operator, and the result
func CalculateAndPrintResult(firstOperand, secondOperand int, operator string) {
	result := PerformCalculation(firstOperand, secondOperand, operator)
	fmt.Printf("%d %s %d = %d\n", firstOperand, operator, secondOperand, result)
}

// PerformCalculation returns the result of basic arithmetic (+, -, *, /) beetwen two operands
func PerformCalculation(firstOperand, secondOperand int, operator string) int {
	result := 0
	switch operator {
	case "+":
		result = firstOperand + secondOperand
	case "-":
		result = firstOperand - secondOperand
	case "*":
		result = firstOperand * secondOperand
	case "/":
		result = firstOperand / secondOperand
	default:
		fmt.Println("Operator is invalid..")
	}
	return result
}

func readCliInputAsInt(placeholderText string) int {
	cliInput := readCliInputAsStr(placeholderText)
	return convertStrToInt(cliInput)
}

func readCliInputAsStr(placeholderText string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(placeholderText)
	cliInput, _ := reader.ReadString(EOL)

	// Trims "\r\n" runes on windows. Won't have effect on Linux
	cliInput = strings.TrimRight(cliInput, "\r\n")
	return cliInput
}

func convertStrToInt(strToConvert string) int {
	i, err := strconv.Atoi(strToConvert)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	return i
}
