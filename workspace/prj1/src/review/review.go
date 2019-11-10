package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sortArray(array []int) []int {
	element := 0
	for j := range array {
		for k := range array {
			if j != k {
				if array[j] > array[k] {
					element = array[j]
					array[j] = array[k]
					array[k] = element
				}
			}
		}
	}
	return array
}

func isX(input string) bool {
	if strings.Compare(input, "X") == 0 || strings.Compare(input, "x") == 0 {
		fmt.Println(" Exit pressed. ")
		return true
	}
	return false
}

/**
 * Reads the input from the console for number
 */
func readInput(message string) string {
	scanner := bufio.NewScanner(os.Stdin)
	var valueT string
	fmt.Printf(message)

	for scanner.Scan() {
		input := scanner.Text()
		valueT = input
		break
	}
	return valueT
}

func printArrayMetric(listOfNumbers []int) {
	fmt.Printf(" LEN [%d] : CAP [%d]\n", len(listOfNumbers), cap(listOfNumbers))
}

/**
 * Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
 * The program should be written as a loop. Before entering the loop, the program should create an empty
 * integer slice of size (length) 3. During each pass through the loop, the program prompts the user to
 * enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice,
 * and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any
 * number of integers which the user decides to enter. The program should only quit (exiting the loop)
 * when the user enters the character ‘X’ instead of an integer.
 *
 * Submit your source code for the program, “slice.go”.
 **/

// Week3Assignment1  Week 3 assignment 1
func Week3Assignment1() {

	// Define the original slice
	array := make([]int, 0, 3)

	listOfNumbers := array[0:]

	fmt.Println(" Begins with")
	printArrayMetric(listOfNumbers)
	fmt.Println(listOfNumbers)
	fmt.Println("-------------------------")

	for {
		input := readInput(" Please enter the number: ")

		if isX(input) {
			return
		}

		number, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println(" Did not understand; Please provide input again. To exit press x or X")
			continue
		}

		printArrayMetric(listOfNumbers)

		if len(listOfNumbers) == 0 {
			fmt.Println(" First element!!")
			listOfNumbers = append(listOfNumbers, number)
			fmt.Println(listOfNumbers)
			continue
		}

		/**
		 * Starting to add elements. The existing element in the array are
		 * identified by capacity.
		 */
		fmt.Println(" adding element.......")
		//////////////////////////////////////////////////////////

		listOfNumbers = append(listOfNumbers, number)
		listOfNumbers = sortArray(listOfNumbers)
		fmt.Println(listOfNumbers)
	}
}

/**
 * Executes the week3 assignment.
 */
func main() {
	Week3Assignment1()
}
