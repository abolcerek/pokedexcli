package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/abolcerek/pokedexcli/internal"
	"time"
)

type PokemonStats struct {
	hp int
	attack int
	defense int
	special_attack int
	special_defense int
	speed int
}

type Pokemon struct {
	name string
	height int
	weight int
	Stats PokemonStats
	Types []string
}

type config struct {
	Next string
	Previous string
	cache *pokecache.Cache
	Location string
	Pokemon string
	Pokedex map[string]Pokemon
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
	"explore": {
        name:        "explore",
        description: "Lists all pokemon in location",
        callback:    commandExplore,
    	},
	"catch": {
        name:        "catch",
        description: "Catches a pokemon",
        callback:    commandCatch,
    	},
	"inspect": {
        name:        "inspect",
        description: "Inspects a pokemon",
        callback:    commandInspect,
    	},
	"pokedex": {
        name:        "pokedex",
        description: "Displays pokedex",
        callback:    commandPokedex,
    	},
	}
	scanner := bufio.NewScanner(os.Stdin)
	Cache := pokecache.NewCache(time.Second * 5)
	Pokedex := make(map[string]Pokemon)
	Config := config{
		cache: Cache,
		Pokedex: Pokedex,
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
			if value.name == "explore" {
				if len(res) > 1 {
					Config.Location = res[1]
				} else {
					fmt.Print("Please provide location\n")
					continue
				}
			}
			if value.name == "catch" {
				if len(res) > 1 {
					Config.Pokemon = res[1]
				} else {
					fmt.Print("Please provide pokemon\n")
					continue
				}
			}
			if value.name == "inspect" {
				if len(res) > 1 {
					_, ok := Config.Pokedex[res[1]]
					if ok == false {
						fmt.Print("you have not caught that pokemon\n")
						continue
					}
					Config.Pokemon = res[1]
				} else {
					fmt.Print("Please provide pokemon\n")
					continue
				}
			}
			if value.name == "pokedex" {
				if len(Config.Pokedex) < 1 {
					fmt.Print("you have not caught any pokemon\n")
					continue
				}
			}
			value.callback(&Config)
		}
	}
}


