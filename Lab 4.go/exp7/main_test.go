package main

import (
	"errors"
	"math"
	"testing"
)

func TestDivideValidInput(t *testing.T) {
	result, err := divide(10, 2)
	if err != nil {
		t.Fatalf("divide returned an unexpected error: %v", err)
	}
	if result != 5 {
		t.Fatalf("divide(10, 2) = %v, want 5", result)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := divide(10, 0)
	if err == nil {
		t.Fatal("divide should return an error for a zero divisor")
	}
}

func TestSquareRootValidInput(t *testing.T) {
	result, err := squareRoot(25)
	if err != nil {
		t.Fatalf("squareRoot returned an unexpected error: %v", err)
	}
	if result != 5 {
		t.Fatalf("squareRoot(25) = %v, want 5", result)
	}
}

func TestSquareRootNegativeInput(t *testing.T) {
	_, err := squareRoot(-1)
	if err == nil {
		t.Fatal("squareRoot should return an error for a negative input")
	}
}

func TestValidateInputValidInput(t *testing.T) {
	result, err := validateInput("  hello  ")
	if err != nil {
		t.Fatalf("validateInput returned an unexpected error: %v", err)
	}
	if result != "hello" {
		t.Fatalf("validateInput returned %q, want %q", result, "hello")
	}
}

func TestValidateInputEmptyInput(t *testing.T) {
	_, err := validateInput("   ")
	if !errors.Is(err, errEmptyInput) {
		t.Fatalf("validateInput error = %v, want errEmptyInput", err)
	}
}

func TestSquareRootReturnsFiniteResult(t *testing.T) {
	result, err := squareRoot(4)
	if err != nil || math.IsNaN(result) {
		t.Fatalf("squareRoot(4) returned result=%v, error=%v", result, err)
	}
}
