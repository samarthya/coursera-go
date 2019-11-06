package week3a

import (
	"fmt"
)

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

func sortArray(array []int) []int {
	element := 0

	fmt.Println(" ---- Sort begins ---")
	for j := range array {
		for k := range array {
			if j != k {
				fmt.Printf(" Compare %d and %d\n", array[j], array[k])
				if array[j] > array[k] {
					fmt.Println(" Swap.")
					element = array[j]
					array[j] = array[k]
					array[k] = element
				}
			}
		}
	}
	fmt.Println(" ---- Sort ends ---")
	return array
}

// Adds the element in the array.
func addTheElement(array []int, element int, capacity int) []int {
	added := false

	/**
	 * Range works with the length.
	 * Issue is the default 3 elements are all populated with 0's
	 * So length and capacity both are 3 even though no element
	 * has been added.
	 */
	fmt.Println(" ---------- loop begins ----------")
	for j := range array {

		fmt.Printf(" Looking for [%d] in array\n", j)
		fmt.Printf(" Found %d at %d capacity %d\n", array[j], j, capacity)

		/**
		 * Check against capacity is required to see which element is the last, as for
		 * the initial state 0's prepopulated are also considered in capacity for cap()
		 */
		if j < capacity {
			if array[j] > element {
				fmt.Printf(" Skipped %d at [%d] > %d\n", array[j], j, element)
				// Move to the next element in the array.
				continue
			}
			if capacity < len(array) {
				fmt.Printf(" Found %d at %d < %d\n", array[j], j, element)
				// Adding it to the index -> Capacity (last element.)
				array[capacity] = element
				capacity++
				fmt.Printf(" Added %d\n It looks like \n", element)
				fmt.Println(array)
				array = sortArray(array)
				fmt.Println(" Sorted the array.")
				fmt.Println(array)

				// When you start with 3 it makes it complicated to rely on cap()
				added = true
				break
			} else /*if capacity == len(array) */ {
				/**
				 * The last element can be less than the new element to be added.
				 */
				array = append(array, element)
				fmt.Println(array)
				array = sortArray(array)
				fmt.Printf(" Edge element added, new length %d\n", len(array))
				added = true
				capacity++
				fmt.Println(array)
				break
			}
		} else {
			break
		}
	}

	if !added {
		fmt.Println(" Out of for.. cases when the element to be added is greater than the greatest in array.")
		fmt.Printf(" locating postion....")
		if capacity < len(array) {
			array[capacity] = element
			fmt.Printf(" Added %d at %d\n", element, capacity)
			capacity++
		} else if capacity == len(array) {
			array = append(array, element)
			capacity++
		}
	}
	fmt.Println(" ---------- loop ends ----------")
	return array
}

/**
 * Adds the element to the splice.
 */
func insertTheElement(array []int, element int, capacity int) []int {

	// If capacity is 0 it is the first element
	if capacity == 0 {
		fmt.Println(" First element!!")
		array[0] = element
	} else {
		/**
		 * Starting to add elements. The existing element in the array are
		 * identified by capacity.
		 */
		fmt.Println(" adding element.......")
		array = addTheElement(array, element, capacity)
	}

	return array
}

// Week3Assignment1 Coursera
func Week3Assignment1() {
	// The make function allocates a zeroed array and returns a slice that refers to that array:
	mySlice := make([]int, 3)
	var num int = 0
	capacity := 0

	fmt.Printf(" Length and Capacity: %d - %d\n", len(mySlice), cap(mySlice))
	for {
		fmt.Printf(" Please enter the num: ")
		fmt.Scanf("%d", &num)

		if num == 0 {
			break
		}

		mySlice = insertTheElement(mySlice, num, capacity)

		/**
		 * Till the element's added are 3, rely on manual.
		 */
		if capacity < 3 {
			capacity++
		} else {
			capacity = len(mySlice)
		}

		fmt.Println(mySlice)
	}

}
