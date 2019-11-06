package main

import (
	"fmt"

	"samarthya.me/week1"
	"samarthya.me/week2"
	"samarthya.me/week3a"
	"samarthya.me/week3b"
)

func readChoice(maxOption int) int {
	var choice int = 0
	fmt.Scanf("%d", &choice)

	if choice < maxOption {
		return choice
	}

	fmt.Println(" Invalid choice, please choose again.\n")
	return readChoice(maxOption)
}

func menu() int {

	fmt.Println(" ************************ ")
	fmt.Println(" 1. Week 1 Assignment")
	fmt.Println(" 2. Week 2 Assignment 1")
	fmt.Println(" 3. Week 2 Assignment 2")
	fmt.Println(" 4. Week 3 Assignment 1")
	fmt.Println(" 5. Week 3 Assignment 1")
	fmt.Println("------------------------")
	fmt.Println(" 0. Exit.")
	fmt.Println(" ************************ ")
	fmt.Print(" Please enter your choice: ")

	return readChoice(6)
}

func main() {

	// week1Assignment1()
	// week2Assignment1()

	// utils.CheckInStringv1()

	for {
		x := menu()

		switch x {
		case 1:
			week1.Week1Assignment1()
		case 2:
			week2.Week2Assignment2()
		case 3:
			week2.CheckInStringv2()
		case 4:
			week3a.Week3Assignment1()
		case 5:
			week3b.Week3Assignment1()
		default:
			break
		}

		if x == 0 {
			break
		}
	}
}
