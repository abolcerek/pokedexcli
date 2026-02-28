package main

import (
	"fmt"
	"os"
	"strings"
)




func cleanInput(text string)[]string {
	if len(text) == 0 {
		return []string{}
	}
	lowerstring := strings.ToLower(text)
	words := strings.Fields(lowerstring)
	return words
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
	return nil
}
