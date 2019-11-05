package utils

import (
	"fmt"
)

func readInteger(message string) int {
	var choice int = 0
	if message != "" {
		fmt.Printf("%s", message)
	}
	n, err := fmt.Scanf("%d", &choice)

	if err != nil {
		panic(err)
		return 0
	}

	return choice
}
