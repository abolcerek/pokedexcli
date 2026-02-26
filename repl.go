package main

import (
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
