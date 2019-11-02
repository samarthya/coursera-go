package trunc

import (
	"fmt"
)

func displayHelloWorld() string {
	return "Hello World!"
}

/**
 * Write a program which prompts the user to enter a floating point number and prints
 * the integer which is a truncated version of the floating point number that was entered.
 * Truncation is the process of removing the digits to the right of the decimal place.
 * Submit your source code for the program, “trunc.go”.
 */

// ScanAFloat allows you to scan an input FLOAT32 and return.
func ScanAFloat(number *float32) float32 {
	num, err := fmt.Scan(number)

	if err != nil {
		fmt.Printf(" Error occurred.")
	} else {
		fmt.Printf(" The number of items scanned: %d\n", num)
	}
	return (*number)
}

// DisplayInt allows you to display the input float as int.
func DisplayInt(number *float32) {
	var y = int32(*number)
	fmt.Printf(" Value output as Int32: %d\n", y)
}

func test1() {
	var x int32 = 1
	var y int16 = 2

	const (
		f = 2
	)

	fmt.Println(" Project 2 - Main")

	// x = y
	// cannot use y (type int16) as type int32 in assignment
	// Converting a value using the Conversion.

	x = int32(y)

	fmt.Printf(" X : %d\n", x)
}
