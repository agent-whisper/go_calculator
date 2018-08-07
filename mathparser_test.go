package main

import (
  "testing"
)

func TestParseOne(t *testing.T) {
  expectedSlice := []interface{}{3, 4, 2, "*", 1, 5, "-", 2, 3, "^", "^", "/", "+"}
  actualStack, err := convertToPostfix("3 + 4 * 2 / ( 1 - 5 ) ^ 2 ^ 3")
  if err != nil {
    t.Errorf("Error when converting string to postfix")
  }
  for i := len(expectedSlice)-1; i >= 0; i-- {
    x, err := actualStack.Pop()
    if err != nil {
      t.Errorf("Error when popping stack")
    }
    if x != expectedSlice[i] {
      t.Errorf("Mismatched at index %d: want(%v), got(%v)", i, expectedSlice, x)
    }
  }
}
