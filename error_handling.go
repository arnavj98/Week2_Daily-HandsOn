package main

import (
	"errors"
	"fmt"
)

// Function to divide two numbers with error handling
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero") // Return error when dividing by zero
	}
	return a / b, nil // Return result with no error
}

func main() {
	// Test case 1: valid division
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}

	// Test case 2: division by zero
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // Will print the error
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}
}
