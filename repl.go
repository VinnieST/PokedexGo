package main

import "strings"

// splits the given text into a slice of substring words seperated by whitespace
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
