# Pokedex CLI

A terminal-based Pokédex built in Go powered by the [PokéAPI](https://pokeapi.co/).

## Prerequisites

Make sure you have the following installed:

- **Go** (1.22 or later) — [https://go.dev/dl/](https://go.dev/dl/)

No external services or databases required — the app talks directly to the PokéAPI and caches responses in memory.

## Installation

Install the CLI using `go install`:

```bash
go install github.com/abolcerek/pokedexcli@latest
```

## Running the Program

Start the interactive REPL by running:

```bash
pokedexcli
```

You'll be dropped into the Pokédex prompt:

```
Pokedex >
```

Type any command and press Enter. Commands and Pokémon names are case-insensitive.

## Commands

| Command | Arguments | Description |
|---|---|---|
| `help` | — | Display a list of available commands |
| `exit` | — | Close the Pokédex |
| `map` | — | Show the next 20 location areas |
| `mapb` | — | Show the previous 20 location areas |
| `explore` | `<location>` | List all Pokémon found in a location area |
| `catch` | `<pokemon>` | Attempt to catch a Pokémon |
| `inspect` | `<pokemon>` | View stats for a caught Pokémon |
| `pokedex` | — | List all Pokémon you've caught |

## Example Session

```
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
...

Pokedex > explore pastoria-city-area
Exploring pastoria-city-area...
Found Pokemon:
- tentacool
- tentacruel
- shellos
...

Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu was caught!

Pokedex > inspect pikachu
Name: pikachu
Height: 4
Weight: 60
Stats:
 -hp: 35
 -attack: 55
 -defense: 40
 -special-attack: 50
 -special-defense: 50
 -speed: 90
Types:
 - electric

Pokedex > pokedex
Your Pokedex:
 - pikachu

Pokedex > exit
Closing the Pokedex... Goodbye!
```
