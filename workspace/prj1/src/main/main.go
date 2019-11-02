package main

import (
	"fmt"
)

// week1_assignment_1 is the week one assignment.
func week1Assignment1() {
	y := 100.230

	// z := 300.20
	x := "Hello, world!"

	fmt.Println(" -------------- ")
	fmt.Printf("%s : x: %f\n", x, y)
	fmt.Println(" -------------- ")
}

/**
 * week2Assignment1
 * Write a program which prompts the user to enter a floating point number and prints
 * the integer which is a truncated version of the floating point number that was entered.
 * Truncation is the process of removing the digits to the right of the decimal place.
 * Submit your source code for the program, “trunc.go”.
 */
func week2Assignment1() {
	var number = 0.0
	fmt.Println(" ****** Week 2 Assignment 1 ********** ")
	fmt.Printf(" Please enter a float number (FLOAT32) : ")
	fmt.Scan(&number)
	fmt.Printf(" Number when converted into Int32 is: %d\n", int32(number))
	fmt.Println(" End.")
}

func main() {
	// week1Assignment1()
	week2Assignment1()
}
