package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(5, 5)
	if result != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", result, 10)
	}
}

func TestMinus(t *testing.T) {
	result := minus(5, 5)
	if result != 0 {
		t.Errorf("Minus was incorrect, got: %d, want: %d.", result, 0)
	}
}

func TestMultiply(t *testing.T) {
	result := multiply(5, 5)
	if result != 25 {
		t.Errorf("Multiply was incorrect, got: %d, want: %d.", result, 25)
	}
}

func TestDivide(t *testing.T) {
	result := divide(5, 5)
	if result != 1 {
		t.Errorf("Divide was incorrect, got: %d, want: %d.", result, 1)
	}
}
