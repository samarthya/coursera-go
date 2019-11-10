package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers.
The program should print the integers out on one line, in sorted order, from least to greatest.
Use your favorite search tool to find a description of how the bubble sort algorithm works.

A recurring operation in the bubble sort algorithm is the Swap operation which swaps
the position of two adjacent elements in the slice. You should write a Swap()
function which performs this operation. Your Swap() function should take two arguments,
a slice of integers and an index value i which indicates a position in the slice.
The Swap() function should return nothing, but it should swap the contents of the
slice in position i with the contents in position i+1.

Submit your Go program source code.
*/

func swap(swapSlice []int, index int) {
	// fmt.Println(" Swap Called")
	x := swapSlice[index]
	swapSlice[index] = swapSlice[index+1]
	swapSlice[index+1] = x
}

// As part of this program, you should write a function called BubbleSort() which takes a slice
// of integers as an argument and returns nothing. The BubbleSort() function should modify
// the slice so that the elements are in sorted order.

// BubbleSort Sort the array
func BubbleSort(mySequence []int) {

	fmt.Printf("\n Will run till %d\n", len(mySequence))
	for i := 0; i < len(mySequence); i++ {
		for j := 0; j < (len(mySequence) - (i + 1)); j++ {
			// fmt.Printf("\n Comparing %d and %d", mySequence[j], mySequence[j+1])
			if mySequence[j] > mySequence[j+1] {
				swap(mySequence, j)
			}
		}
	}
}

func main() {
	myNumbers := make([]int, 0, 1)

	var lineOfInput string

	fmt.Println(" Bubble Sort")
	fmt.Println(" Please enter all the numbers with a space in between and print Enter to finish")
	scanner := bufio.NewReader(os.Stdin)

	lineOfInput, err := scanner.ReadString('\n')

	if err != nil {
		panic(err)
		return
	}

	lineOfInput = strings.Trim(lineOfInput, "\n")
	fmt.Println(" Splitting and Sorting under process.")
	numbers := strings.Split(lineOfInput, " ")

	for i, v := range numbers {
		if i < 10 {

			fmt.Printf(" -- %s", v)
			x, err := strconv.Atoi(numbers[i])

			if err == nil {
				myNumbers = append(myNumbers, x)
			}

		}
	}

	// NUmbers read.
	BubbleSort(myNumbers)
	fmt.Println(myNumbers)
}
