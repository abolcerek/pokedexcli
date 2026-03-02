package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type results struct {
	Name string `json:"name"`
	Url string `json:"url"`
}
type Response struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Prevous string `json:"previous"`
	Res []results `json:"results"`
}

var Next string
var Previous string


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
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\nmap: Displays 20 location areas\nmapb: Displays previous 20 location areas\n")
	return nil
}
func commandMap() error {
	url := "https://pokeapi.co/api/v2/location-area"
	if Next != "" {
		url = Next
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	res := Response{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}
	for _, location := range res.Res {
		fmt.Printf("%s\n", location.Name)
	}
	Next = res.Next
	Previous = res.Prevous
	return nil
}
func commandMapb() error {
	url := "https://pokeapi.co/api/v2/location-area"
	if Previous != "" {
		url = Previous
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	res := Response{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}
	if Previous == res.Prevous {
		fmt.Print("you're on the first page\n")
		Previous = ""
		Next = ""
		return nil
	}
	for _, location := range res.Res {
		fmt.Printf("%s\n", location.Name)
	}
	Next = res.Next
	Previous = res.Prevous
	return nil
}
