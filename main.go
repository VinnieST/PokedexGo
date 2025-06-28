package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Creating a new scanner for our Pokedex input
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		if scanner.Scan() {
			line := scanner.Text()
			cleanLine := cleanInput(line)

			if len(cleanLine) == 0 {
				fmt.Println("No command entered.")
				continue
			}

			fmt.Printf("Your command was: %s\n", cleanLine[0])
		} else {
			// Handle EOF or error
			fmt.Println("\nExiting...")
			break
		}
	}
}
