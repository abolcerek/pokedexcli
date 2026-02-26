package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!.")
	response := cleanInput("Charmander Bulbasaur PIKACHU")
	fmt.Println(response)
	fmt.Println(len(response))

}


