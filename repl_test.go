package main

import (
    "testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
	input    string
	expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello  ",
			expected: []string{"hello"},
		},
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}
	for _, c := range cases {
	actual := cleanInput(c.input)
	if len(actual) != len(c.expected) {
		t.Errorf("Error: Lengths dont match")
	}
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		if word != expectedWord {
			t.Errorf("Error: Words don't match")
		}
	}
}
}