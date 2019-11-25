package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

/**
Write a program which allows the user to create a set of animals and to get information about those animals.

Each animal has a name and can be either a cow, bird, or snake.
With each command, the user can either create a new animal of one of the three types,
or the user can request information about an animal that he/she has already created.
Each animal has a unique name, defined by the user.
Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake.
The following table contains the three types of animals and their associated data.

Animal	Food eaten	Locomtion method	Spoken sound
cow		grass		walk				moo
bird	worms		fly					peep
snake	mice		slither				hsss


// Your program should call the appropriate method when the user issues a query command.
**/

// Define an interface type called Animal which describes the methods of an animal.
// Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
// The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method
// should print the animal’s spoken sound.

// ANIMAL interface ANIMAL that should be implemented by the concrete instances.
type ANIMAL interface {
	Speak()
	Eat()
	Move()
}

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

// Define three types Cow, Bird, and Snake.
// For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface.
// When the user creates an animal, create an object of the appropriate type.

// Cow defines the animal
type Cow struct {
	locomotion LOCOMOTION
	speaks     SPEAKS
	food       EATS
}

// Speak for Cow
func (t Cow) Speak() {
	fmt.Println(t.speaks)
}

// Move for Cow
func (t Cow) Move() {
	fmt.Println(t.locomotion)
}

// Eat for Cow
func (t Cow) Eat() {
	fmt.Println(t.food)
}

// Snake defines the animal
type Snake struct {
	locomotion LOCOMOTION
	speaks     SPEAKS
	food       EATS
}

// Speak for Cow
func (t Snake) Speak() {
	fmt.Println(t.speaks)
}

// Move for Cow
func (t Snake) Move() {
	fmt.Println(t.locomotion)
}

// Eat for Cow
func (t Snake) Eat() {
	fmt.Println(t.food)
}

// Bird defines the animal
type Bird struct {
	locomotion LOCOMOTION
	speaks     SPEAKS
	food       EATS
}

// Speak for Cow
func (t Bird) Speak() {
	fmt.Println(t.speaks)
}

// Move for Cow
func (t Bird) Move() {
	fmt.Println(t.locomotion)
}

// Eat for Cow
func (t Bird) Eat() {
	fmt.Println(t.food)
}

const usageMessage string = " Usage:\n newanimal <name> <type>.\n" +
	" NAME: Is arbitary name you want to name your animal.\n TYPE: It can only be cow, snake or bird.\n" +
	" or\n" +
	" query <name> <command>\n" +
	" Command cal be move, eat or speak"

// Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
// Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line.
// Your program should continue in this loop forever.
// Every command from the user must be either a “newanimal” command or a “query” command.

// Each “newanimal” command must be a single line containing three strings.
// The first string is “newanimal”. The second string is an arbitrary string which will be the name of the new animal.
// The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
// Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

// Each “query” command must be a single line containing 3 strings.
// The first string is “query”. The second string is the name of the animal.
// The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
// Your program should process each query command by printing out the requested data.
func readInput() {
	scanner := bufio.NewReader(os.Stdin)
	// Empty map.
	animals := map[string]ANIMAL{}
	var commands []string

	cow := Cow {
		locomotion: walk,
		speaks:     moo,
		food:       grass,
	}

	snake := Snake {
		locomotion: slither,
		speaks:     hisss,
		food:       mice,
	}

	bird := Bird{
		locomotion: fly,
		speaks:     peep,
		food:       worms,
	}

	// Loop starts
	for {
		fmt.Print("\n> ")

		// Read the input.
		fromUser, err := scanner.ReadString('\n')

		// If nil exit.
		if err != nil {
			panic(err)
		}

		fromUser = strings.Trim(fromUser, "\n")

		for _, v := range strings.Split(fromUser, " ") {
			commands = append(commands, strings.Trim(strings.ToLower(v), " "))
			if contains(commands, "exit") {
				fmt.Println(" Exiting....")
				return
			}
		}

		// Execute command
		if len(commands) > 2 {
			switch commands[0] {
			case "newanimal":
				// Add to the animals array.
				var newAnimal bool = true
				switch commands[2] {
				case "cow":
					animals[commands[1]] = cow
				case "snake":
					animals[commands[1]] = snake
				case "bird":
					animals[commands[1]] = bird
				default:
					fmt.Printf("%s", usageMessage)
					newAnimal = false
				}

				if newAnimal {
					newAnimal = true
					fmt.Println(" Created it!!")
				}

			case "query":

				if len(animals) > 0 {
					//Get the animal
					var aIA ANIMAL = animals[commands[1]]
					switch commands[2] {
					case "speak":
						aIA.Speak()
					case "move":
						aIA.Move()
					case "eat":
						aIA.Eat()
					default:
						fmt.Printf("%s", usageMessage)
					}
				} else {
						fmt.Println(" Need to add an element before you query!")
				}
			default:
				fmt.Printf("%s", usageMessage)
			}
		} else {
			fmt.Printf("%s", usageMessage)
		}

		commands = nil
		fmt.Printf("\n Total animals: %d", len(animals))

	}
}

func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}

func main() {
	readInput()
}
