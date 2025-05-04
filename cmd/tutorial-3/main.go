// Function in Go
package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe("Hello Function")
	result, reminder, err := intDivision(12, 2)
	// Check if there is any error or no
	if err != nil {
		fmt.Printf("%s", err.Error())
	} else {
		fmt.Printf("The result of the division is %v with reminder %v", result, reminder)
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	// Error handling in go
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result = numerator / denominator
	var reminder = numerator % denominator
	return result, reminder, err
}
