package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

var errEmptyInput = errors.New("input cannot be empty")

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("cannot divide %.2f by zero", dividend)
	}
	return dividend / divisor, nil
}

func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("cannot calculate the square root of negative number %.2f", number)
	}
	return math.Sqrt(number), nil
}

func validateInput(input string) (string, error) {
	trimmedInput := strings.TrimSpace(input)
	if trimmedInput == "" {
		return "", errEmptyInput
	}
	return trimmedInput, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division result:", result)
	}

	root, err := squareRoot(25)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Square root:", root)
	}

	input, err := validateInput("Lab 4")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Validated input:", input)
	}
}
