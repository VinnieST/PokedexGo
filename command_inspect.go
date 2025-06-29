package main

import (
	"errors"
	"fmt"
	"strings"
)

func commandInspect(cfg *config, args []string) error {
	if cfg == nil || len(args) < 1 {
		return errors.New("invalid arguments")
	}
	name := strings.ToLower(args[0])

	pokemon, ok := cfg.pokedex[name]
	if !ok {
		fmt.Println("You have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}
	return nil
}
