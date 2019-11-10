package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
 * Let us assume the following formula for displacement s as a function of time t,
 * acceleration a, initial velocity vo, and initial displacement so.

 * s = ½ a t2 + vot + so
 *
 * Write a program which first prompts the user to enter values for acceleration,
 * initial velocity, and initial displacement. Then the program should prompt the
 * user to enter a value for time and the program should compute the displacement
 * after the entered time.
 *
 * You will need to define and use a function called GenDisplaceFn() which takes
 * three float64 arguments, acceleration a, initial velocity vo, and initial displacement so.
 * GenDisplaceFn() should return a function which computes displacement as a function of time,
 * assuming the given values acceleration, initial velocity, and initial displacement.
 * The function returned by GenDisplaceFn() should take one float64 argument t, representing time,
 * and return one float64 argument which is the displacement travelled after time t.
 *
 * For example, let’s say that I want to assume the following values for acceleration,
 * initial velocity, and initial displacement: a = 10, vo = 2, so = 1. I can use the
 * following statement to call GenDisplaceFn() to generate a function fn which will
 * compute displacement as a function of time.
 *
 * fn := GenDisplaceFn(10, 2, 1)
 *
 * Then I can use the following statement to print the displacement after 3 seconds.
 *
 * fmt.Println(fn(3))
 *
 * And I can use the following statement to print the displacement after 5 seconds.
 *
 * fmt.Println(fn(5))
 *
 * Submit your Go program source code.
 */

func readValue(message string) []float64 {
	myNumbers := make([]float64, 0, 3)

	if message == "" {
		fmt.Printf(" Please enter all the numbers with a space in between and print Enter to finish")
	} else {
		fmt.Printf(" %s", message)
	}

	scanner := bufio.NewReader(os.Stdin)

	lineOfInput, err := scanner.ReadString('\n')

	if err != nil {
		panic(err)
		return nil
	}

	lineOfInput = strings.Trim(lineOfInput, "\n")

	for i, v := range strings.Split(lineOfInput, " ") {
		if i < 4 {
			x, err := strconv.ParseFloat(v, 64)
			if err == nil {
				myNumbers = append(myNumbers, x)
				continue
			}
		}
		break
	}

	return myNumbers
}

// You will need to define and use a function called GenDisplaceFn() which takes
// three float64 arguments, acceleration a, initial velocity vo, and initial displacement so.
// GenDisplaceFn() should return a function which computes displacement as a function of time,
// assuming the given values acceleration, initial velocity, and initial displacement.
// The function returned by GenDisplaceFn() should take one float64 argument t, representing time,
// and return one float64 argument which is the displacement travelled after time t.
// GenDisplaceFn Calculates displacement
func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	// s =½ a t2 + vot + so
	fn := func(t float64) float64 {
		var delay time.Duration = time.Duration(t)
		time.Sleep(delay * time.Second)
		return ((0.50 * (a * math.Pow(t, 2))) + (vo * t) + so)
	}
	return fn
}

func main() {
	var params []float64
	var tm float64 = 0
	params = readValue(" Please enter the value for Acceleration, Velocity & Initial Displacement and press enter to calculate: ")

	if len(params) < 3 {
		fmt.Println(" Error calculating the displacement.")
		return
	}

	var displacement = GenDisplaceFn(params[0], params[1], params[2])
	fmt.Println(" Calculating displacement.....")
	fmt.Print(" Please enter the time: ")
	fmt.Scanf("%f", &tm)
	fmt.Println(displacement(tm))
}
