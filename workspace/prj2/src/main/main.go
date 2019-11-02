package main

import (
	"fmt"

	trunc "samarthya.me/trunc"
)

func main() {
	var x float32 = 0
	fmt.Println(" The program begins.")
	fmt.Printf(" Please enter a float Number: ")
	trunc.ScanAFloat(&x)
	trunc.DisplayInt(&x)
	fmt.Println(" THE END.")
}
