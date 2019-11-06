package week3b

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
	listOfNumbers := make([]int, 3)

	// Original capacity is 0 against the computer cap()
	capacity := 0
	added := false

	fmt.Println(" Begins with")
	fmt.Printf(" LEN and CAP: [%d] : [%d]\n", len(listOfNumbers), cap(listOfNumbers))
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

		fmt.Printf(" LEN and CAP: [%d] : [%d]\n", len(listOfNumbers), capacity)

		if capacity == 0 {
			fmt.Println(" First element!!")
			listOfNumbers[0] = number
			capacity++
			fmt.Println(listOfNumbers)
			continue
		}

		/**
		 * Starting to add elements. The existing element in the array are
		 * identified by capacity.
		 */
		fmt.Println(" adding element.......")
		//////////////////////////////////////////////////////////

		/**
		 * Range works with the length.
		 * Issue is the default 3 elements are all populated with 0's
		 * So length and capacity both are 3 even though no element
		 * has been added.
		 */
		fmt.Printf(" Starting to look for position to add %d  in array\n", number)

		for j := range listOfNumbers {
			fmt.Printf(" Element: [%d] at index: %d Total element- %d\n", listOfNumbers[j], j, capacity)

			/**
			 * Check against capacity is required to see which element is the last, as for
			 * the initial state 0's prepopulated are also considered in capacity for cap()
			 */
			if j < capacity {
				fmt.Printf(" index -> %d\n", j)

				if listOfNumbers[j] > number {
					fmt.Printf(" Skipped [%d] > [%d]\n", listOfNumbers[j], number)

					// Move to the next element in the array.
					continue
				}

				// Check for element which is lesser than number is found so now we need to place it
				// 1. Is there sufficient space in array.
				//     - capacity < length means we can squeeze in elements.
				// Capacity usually points to the last element.
				// E.g. Capacity 1 means only one element and empty at index 1.
				if capacity < len(listOfNumbers) {
					fmt.Printf(" Found [%d] < %d\n", listOfNumbers[j], number)

					// Adding it to the index -> Capacity (last element.)
					listOfNumbers[capacity] = number

					fmt.Printf(" Added [%d] at %d\n", number, capacity)

					//Increase index.
					capacity++

					listOfNumbers = sortArray(listOfNumbers)
					fmt.Println(listOfNumbers)
					added = true
					break
				} else if capacity == len(listOfNumbers) {
					/**
					 * The last element can be less than the new element to be added.
					 */
					listOfNumbers = append(listOfNumbers, number)

					// After adding sort it
					// Can be optimised.
					listOfNumbers = sortArray(listOfNumbers)
					added = true
					capacity++
					fmt.Println(listOfNumbers)
					break
				}
			} else {
				break
			}
		} // Loop ends

		// Edge cases.
		if !added {
			fmt.Println(" No postion within the array can be used to add this element.")
			if capacity < len(listOfNumbers) {
				listOfNumbers[capacity] = number
				fmt.Printf(" Added [%d] at %d\n", number, capacity)
			} else if capacity == len(listOfNumbers) {
				// No slot available at the new one.
				listOfNumbers = append(listOfNumbers, number)
			}
			capacity++
			fmt.Println(listOfNumbers)
		}

		added = false
	}
}
