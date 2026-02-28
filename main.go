package main

import (
	"bufio"
	"fmt"
	"os"
)
type cliCommand struct {
	name string
	description string
	callback func() error
}

func main() {
	mapper := map[string]cliCommand{
    "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    	},
	"help": {
        name:        "help",
        description: "Displays a help message",
        callback:    commandHelp,
    	},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		res := cleanInput(text)
		value, ok := mapper[res[0]]
		if ok == false {
			fmt.Print("Unknown command")
		} else {
			value.callback()
		}
	}
}


