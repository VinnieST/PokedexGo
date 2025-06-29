package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args []string) error {
	if cfg == nil {
		return fmt.Errorf("configuration is nil")
	}

	fmt.Println("Your Pokedex:")
	if len(cfg.pokedex) == 0 {
		fmt.Println("You have not caught any Pokemon yet.")
		return nil
	}
	for name, pokemon := range cfg.pokedex {
		fmt.Printf(" - %s (ID: %d)\n", name, pokemon.ID)
	}
	fmt.Println("Use 'inspect <pokemon_name>' to see more details about a specific Pokemon.")
	return nil
}
