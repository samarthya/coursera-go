package week2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


// LETTERI const 'i'
const LETTERI string = "i"

// RUNEI const 'i'
const RUNEI rune = 'i'

// LETTERA const 'a'
const LETTERA string = "a"

// RUNEA const 'a'
const RUNEA rune = 'a'

// LETTERN const 'n'
const LETTERN string = "n"

// RUNEN const 'n'
const RUNEN rune = 'n'

/**
 * week2Assignment1
 * Write a program which prompts the user to enter a floating point number and prints
 * the integer which is a truncated version of the floating point number that was entered.
 * Truncation is the process of removing the digits to the right of the decimal place.
 * Submit your source code for the program, “trunc.go”.
 */
 func Week2Assignment2() {
	var number = 0.0
	fmt.Println(" ****** Week 2 Assignment 1 ********** ")
	fmt.Printf(" Please enter a float number (FLOAT32) : ")
	fmt.Scan(&number)
	fmt.Printf(" Number when converted into Int32 is: %d\n", int32(number))
	fmt.Println(" End.")
}




/**
 * Write a program which prompts the user to enter a string.
 * The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
 * The program should print “Found!” if the entered string starts with the character ‘i’,
 * ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!”
 * otherwise.
 * The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.
 */

// CheckInStringv1 allows you to addess the following.
func CheckInStringv1() {
	var searchStr string

	fmt.Printf(" CheckinString Function\n Please enter your string: ")
	fmt.Scanf("%s", &searchStr)
	fmt.Printf(" String read from the console: %s\n", searchStr)
	fmt.Println(" Wait while i look for 'i', 'a' & 'n'")

	var loweredStr = strings.ToLower(searchStr)
	if strings.HasPrefix(loweredStr, LETTERI) && strings.HasSuffix(loweredStr, LETTERN) && strings.ContainsAny(loweredStr, LETTERA) {
		fmt.Println("Found!")
		return
	}

	fmt.Println("Not Found!")
}

// CheckInStringv2 allowes to use rune.
func CheckInStringv2() {

	var searchStr string
	var err error

	var reader = bufio.NewReader(os.Stdin)

	fmt.Printf(" CheckinString Function\n Please enter your string: ")
	searchStr, err = reader.ReadString('\n')

	if err == nil {
		var loweredStr = strings.ToLower(strings.TrimSuffix(searchStr, "\n"))
		fmt.Printf(" Wait while i look for 'i', 'a' & 'n' in %s\n", loweredStr)

		var lastElement = len(loweredStr)

		if (strings.IndexRune(loweredStr, RUNEI) == 0) && rune(loweredStr[lastElement-1]) == RUNEN && strings.ContainsRune(loweredStr, RUNEA) {
			fmt.Println("Found!")
			return
		}
	}

	fmt.Println("Not Found!")
}
