package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/abolcerek/pokedexcli/internal"
	"time"
)

type config struct {
	Next string
	Previous string
	cache *pokecache.Cache
	
}
type cliCommand struct {
	name string
	description string
	callback func(*config) error
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
	Cache := pokecache.NewCache(time.Second * 5)
	Config := config{
		cache: Cache,
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		res := cleanInput(text)
		value, ok := mapper[res[0]]
		if ok == false {
			fmt.Print("Unknown command\n")
		} else {
			value.callback(&Config)
		}
	}
}


