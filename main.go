package main

import (
	"bufio"
	"fmt"
	"os"
)

type config struct {
	Next string
	Previous string
}
type cliCommand struct {
	name string
	description string
	callback func() error
	tracker *config
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
	"map": {
        name:        "map",
        description: "Displays a location map",
        callback:    commandMap,
    	},
	"mapb": {
        name:        "map",
        description: "Displays the previous locations",
        callback:    commandMapb,
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
			fmt.Print("Unknown command\n")
		} else {
			value.callback()
		}
	}
}


