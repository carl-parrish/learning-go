// Write your answer here, and then test your code.
package main

// Uncomment this import section to use required Go packages
 import (
 	"fmt"
 	"strconv"
 	"strings"
 )

// Change these boolean values to control whether you see 
// the expected answer and/or hints.
const showExpectedResult = false;
const showHints = false;

// calculate() returns the the result of adding 2 numbers 
// after parsing them from strings
func calculate(value1 string, value2 string) float64 {
    // Your code goes here.
	var num1, num2 float64
	var err error

	// Convert the first string to a float64
	  num1, err = strconv.ParseFloat(strings.TrimSpace(value1), 64)
	  if (err != nil) {
		fmt.Printf("%T, %v\n", num1, num1)
	 }

	// [Convert the second string to a float64
	 num2, err = strconv.ParseFloat(strings.TrimSpace(value2), 64)
	 if  (err != nil) {
		fmt.Printf("%T %w\n", num2, num2)
	}



	// Calculate and return the result
	result := num1 + num2

	return result
}
