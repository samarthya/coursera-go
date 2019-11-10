package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Assignment1 Function
// Write a program which prompts the user to first enter a name, and then enter an address.
// Your program should create a map and add the name and address to the map using the keys “name”
// and “address”, respectively. Your program should use Marshal() to create a JSON object from
// the map, and then your program should print the JSON object.
func Assignment1() {
	var name string
	var address string

	scanner := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter a name: ")
	name, err := scanner.ReadString('\n')

	fmt.Print("Please enter address: ")
	address, err = scanner.ReadString('\n')

	myMap := map[string]string{
		"name": strings.Trim(name, "\n"), "address": strings.Trim(address, "\n"),
	}

	rep, err := json.Marshal(myMap)

	if err == nil {
		fmt.Println(string(rep))
	}
}

func main() {
	Assignment1()
}
