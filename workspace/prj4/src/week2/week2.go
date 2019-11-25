package main

import "fmt"

func secondFunction() {
	fmt.Println(" Second Function")
}

func thirdFunction() {
	fmt.Println(" Third Function")
}
func main() {
	go secondFunction()
	fmt.Print(" Main function\n")
	go thirdFunction()
}