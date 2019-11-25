package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

/**
 * Write a program which allows the user to get information about a predefined set of animals.
 * Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak.
 * The user can issue a request to find out one of three things about an animal:
 * 1) the food that it eats,
 * 2) its method of locomotion, and
 * 3) the sound it makes when it speaks.
 *
 * The following table contains the three animals and their associated data which should be hard-coded into your program.
 * Animal	Food eaten	Locomotion method	Spoken sound
 * cow	    grass	    walk				moo
 * bird	    worms		fly					peep
 * snake	mice		slither				hsss

 * Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
 * Your program accepts one request at a time from the user, prints out the answer to the request, and
 * prints out a new prompt.
 * Your program should continue in this loop forever.
 * Every request from the user must be a single line containing 2 strings.
 * The first string is the name of an animal, either “cow”, “bird”, or “snake”.
 * The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
 * Your program should process each request by printing out the requested data.
 *
 * You will need a data structure to hold the information about each animal.
 * Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings.
 * Make three methods called Eat(), Move(), and Speak().
 * The receiver type of all of your methods should be your Animal type.
 * The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and
 * the Speak() method should print the animal’s spoken sound.
 *
 * Your program should call the appropriate method when the user makes a request.
 */

// EATS What the animal speaks
type EATS int

const (
	grass EATS = iota
	worms
	mice
)

func (a EATS) String() string {
	names := [...]string{"grass", "worms", "mice"}

	if a < grass || a > mice {
		return "Unknown"
	}

	return names[a]
}

// SPEAKS What the animal speaks
type SPEAKS int

const (
	moo SPEAKS = iota
	peep
	hisss
)

func (s SPEAKS) String() string {
	tell := [...]string{"moo", "peep", "hisss"}

	if s < moo || s > hisss {
		return "Unknown"
	}

	return tell[s]
}

// LOCOMOTION constant.
type LOCOMOTION int

const (
	walk LOCOMOTION = iota
	fly
	slither
)

func (a LOCOMOTION) String() string {
	names := [...]string{"walk", "fly", "slither"}

	if a < walk || a > slither {
		return "Unknown"
	}

	return names[a]
}

//Animal structure
type Animal struct {
	// LOCOMOTION how it walks.
	locomotion LOCOMOTION
	// SPEAKS How it speaks.
	speaks SPEAKS
	// EATS what it eats.
	food EATS
}

// Speaks Functino
func (animal *Animal) Speaks() string {
	return animal.speaks.String()
}

// Eats function
func (animal *Animal) Eats() string {
	return animal.food.String()
}

// Walks function
func (animal *Animal) Walks() string {
	return animal.locomotion.String()
}

// ANIMAL identifies the Animal
type ANIMAL int

const (
	cow ANIMAL = iota
	bird
	snake
)

func (a ANIMAL) String() string {
	names := [...]string{"cow", "bird", "snake"}
	if a < cow || a > snake {
		return "Unknown"
	}
	return names[a]
}

func execCommand(cmd string, animal Animal) {
	// fmt.Println(animal)
	switch cmd {
	case "eat":
		fmt.Println(animal.Eats())
	case "move":
		fmt.Println(animal.Walks())
	case "speak":
		fmt.Println(animal.Speaks())
	default:
		fmt.Println(" Invalid command. You can only enter eat move or speak")
	}
}

func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}

func main() {
	scanner := bufio.NewReader(os.Stdin)

	var animals = map[ANIMAL]Animal{
		cow: Animal{
			walk,
			moo,
			grass,
		},
		snake: Animal{
			slither,
			hisss,
			mice,
		},
		bird: Animal{
			fly,
			peep,
			worms,
		},
	}

	// for i, v := range animals {
	// 	fmt.Printf(" %d - %v ", i, v)
	// }

	var commands []string
	var counts int = 0

	for {
		// Loop starts
		fmt.Print("> ")

		// Read the input.
		fromUser, err := scanner.ReadString('\n')

		// If null exit.
		if err != nil {
			panic(err)
		}

		fromUser = strings.Trim(fromUser, "\n")

		for i, v := range strings.Split(fromUser, " ") {
			if i < 2 {	
				commands = append(commands, strings.ToLower(v))
				counts++
			}
		}

		if contains(commands, "exit") {
			fmt.Println(" Exiting....")
			return
		}

		if counts == 2 {
			// Commands recieved no execute and break
			switch commands[0] {
			case cow.String():
				fmt.Printf(" Cow - %s \n", commands[1])
				execCommand(commands[1], animals[cow])
			case snake.String():
				fmt.Printf(" Snake - %s \n", commands[1])
				execCommand(commands[1], animals[snake])
			case bird.String():
				fmt.Printf(" Bird - %s \n", commands[1])
				execCommand(commands[1], animals[bird])
			default:
				fmt.Printf(" Error: Please specific the correct animal (cow, snake or bird)\n")
			}
		} else {
			fmt.Println(" Error: Enter Valid input please.")
			fmt.Println(" Format: You need to provide two inputs after the prompt. E.g. cow speak")
		}

		// fmt.Printf("\n CAP(%d) LEN(%d)\n", cap(commands), len(commands))
		commands = commands[:0]
		fromUser = ""
		counts = 0
	}
}
