package main

import (
  "errors"
  "fmt"
  "strconv"
  "strings"
  "zilch/stack"
)

const DEBUG = false

func convertToPostfix(inputExp string) (*stack.Stack, error) {
  expSlice := strings.Split(inputExp, " ")
  operator := stack.New()
  output := stack.New()
  for _, token := range expSlice {
    if DEBUG {
      fmt.Println()
      fmt.Printf("> Current token: %s\n", token)
      fmt.Println("> Operator:")
      operator.Print()
      fmt.Println("> Output:")
      output.Print()
    }
    operand, err := strconv.Atoi(token)
    if err == nil {
      output.Push(operand)
      continue
    }
    if !isValidToken(token) {
      errorMsg := fmt.Sprintf("Invalid token detected: %s", token)
      return nil, errors.New(errorMsg)
    }
    if isOperator(token) {
      if DEBUG {
        fmt.Printf("Detected valid token: %s\n", token)
      }
      handleOperatorToken(token, operator, output)
    }
    if token == "(" {
      operator.Push(token)
    }
    if token == ")" {
      err := handleRightBracket(token, operator, output)
      if err != nil {
        return nil, err
      }
    }
  }
  for {
    if !operator.IsEmpty() {
      if operator.Peek() == "(" {
        return nil, errors.New("Missing right bracket")
      }
      top, err := operator.Pop()
      if err != nil {
        return nil, err
      }
      output.Push(top)
    } else {
      break
    }
  }
  return output, nil
}

func isOperator(token string) bool {
  return token == "+" || token == "-" || token == "*" || token == "/" || token == "^"
}

func isValidToken(token string) bool {
  operatorList := [...]string{"+", "-", "*", "/", "(", ")", "^"}
  for _, op := range operatorList {
    if token == op {
      return true
    }
  }
  return false
}

func isHigherPrecedence(op1, op2 string) bool {
  if isValidToken(op1) && isValidToken(op2) {
    op1Rank := getPrecedence(op1)
    op2Rank := getPrecedence(op2)
    return op1Rank > op2Rank
  }
  return false
}

func isEqualPrecedence(op1, op2 string) bool {
  if isValidToken(op1) && isValidToken(op2) {
    op1Rank := getPrecedence(op1)
    op2Rank := getPrecedence(op2)
    return op1Rank == op2Rank
  }
  return false
}

func getPrecedence(op string) int {
  switch {
    case op == "+" || op == "-":
      return 0
    case op == "*" || op == "/":
      return 5
    case op == "^":
      return 7
    case op == "(" || op == ")":
      return 10
    default:
      return -1
  }
}

func handleOperatorToken(token string, operatorStack, outputStack *stack.Stack) {
  for {
    if operatorStack.IsEmpty() {
      break
    }
    op := operatorStack.Peek().(string)
    if (isHigherPrecedence(op, token) || (isEqualPrecedence(op, token) && isLeftAssociative(token))) && op != "(" {
      top, _ := operatorStack.Pop()
      outputStack.Push(top)
    } else {
      break
    }
  }
  operatorStack.Push(token)
}

func isLeftAssociative(token string) bool {
  return token == "+" || token == "-" || token == "*" || token == "/"
}

func handleRightBracket(token string, operatorStack, outputStack *stack.Stack) error {
  mismatched := true
  for {
    if operatorStack.IsEmpty() {
      break
    }
    top, _ := operatorStack.Pop()
    outputStack.Push(top)
    if operatorStack.Peek() == "(" {
      mismatched = false
      _, _ = operatorStack.Pop()
      break
    }
  }
  if mismatched {
    return errors.New("Missing left bracket")
  }
  return nil
}
