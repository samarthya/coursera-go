package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func logError(err error) {
	log.Fatal(err)
}

// Write a program which reads information from a file and represents it in a slice of structs.
// Assume that there is a text file which contains a series of names. Each line of the text file
// has a first name and a last name, in that order, separated by a single space on the line.
//
// Your program will define a name struct which has two fields, fname for the first name,
// and lname for the last name. Each field will be a string of size 20 (characters).
// Your program should prompt the user for the name of the text file.
// Your program will successively read each line of the text file and create a struct which
// contains the first and last names found in the file. Each struct created will be added to a slice,
// and after all lines have been read from the file, your program will have a slice containing one struct for
// each line in the file. After reading all lines from the file, your program should iterate through your
// slice of structs and print the first and last names found in each struct.

// Submit your source code for the program, “read.go”.

// Week4Assignment2 function
func Week4Assignment2() {
	type Name struct {
		fname, lname string
	}

	// String are UTF-8 by default
	names := make([]Name, 0, 1)
	scanner := bufio.NewReader(os.Stdin)

	cwd, err := os.Getwd()
	fmt.Printf(" CWD: %s\n", cwd)

	fmt.Printf(" Please enter the file to read: ")
	fileName, err := scanner.ReadString('\n')

	if err != nil {
		logError(err)
		return
	}

	if len(strings.Trim(fileName, "\n")) == 0 {
		fileName = "/Users/ss670121/sourcebox/learn/go/workspace/prj1/src/week4/test.txt"
	}

	fileName = strings.Trim(fileName, "\n")

	fmt.Printf(" Reading file %s", fileName)

	if err != nil {
		logError(err)
	}

	// Open file
	handle, err := os.Open(fileName)

	if err != nil {
		logError(err)
		defer handle.Close()
	}

	rd := bufio.NewScanner(handle)
	rd.Split(bufio.ScanLines)

	for rd.Scan() {

		nms := strings.Split(rd.Text(), " ")

		var name Name
		if len(nms) == 2 {
			rlname := []rune(nms[0])
			name.fname = string(rlname[0:20])
			rrname := []rune(nms[1])
			name.lname = string(rrname[0:20])
		} else if len(nms) > 0 {
			rlname := []rune(nms[0])
			name.fname = string(rlname[0:20])
		}
		names = append(names, name)
	}

	// Reading done.

	handle.Close()

	// Print tje slices
	for v, i := range names {
		fmt.Printf("\n Value at %d\n \t[\n", v)
		fmt.Println(i.fname)
		fmt.Println(i.lname)
		fmt.Println("\n\t ]")
	}
}

func main() {
	Week4Assignment2()
}
