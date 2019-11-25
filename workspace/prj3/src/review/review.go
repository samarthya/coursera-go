package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//
// GenDisplaceFn  returns a function of time to calculate the displacement
//
func GenDisplaceFn(acceleration float64, velocity float64, displacement float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*acceleration*t*t + velocity*t + displacement
	}
}

func main() {
	fmt.Println(" Test")
	scanner := bufio.NewScanner(os.Stdin)
	acc := getParameterFromConsole(scanner, "acceleration")
	v := getParameterFromConsole(scanner, "velocity")
	dis := getParameterFromConsole(scanner, "displacement")
	fn := GenDisplaceFn(acc, v, dis)

	fmt.Println("Trying to calculate the displacement as a function of time 3 times")
	for i := 0; i < 3; i++ {
		t := getParameterFromConsole(scanner, "time")
		fmt.Printf("The displacement after time %f is %f\n", t, fn(t))
	}
}

func getParameterFromConsole(scanner *bufio.Scanner, parameter string) float64 {
	fmt.Printf("Please enter %s and press 'return': ", parameter)
	var para float64
	for scanner.Scan() {
		parastr := scanner.Text()
		val, err := strconv.ParseFloat(strings.TrimSpace(parastr), 10)
		if err != nil {
			fmt.Printf("Invalid input, please enter %s and press 'return': ", parameter)
		} else {
			para = val
			break
		}
	}
	return para
}
