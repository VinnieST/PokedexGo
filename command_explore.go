package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if cfg == nil || len(args) < 1 {
		return errors.New("usage: explore <area-name>")
	}
	areaName := args[0]

	area, err := cfg.pokeapiClient.GetLocation(areaName)
	if err != nil {
		return fmt.Errorf("could not find location area: %v", err)
	}

	fmt.Printf("Exploring %s\n", area.Location.Name)
	fmt.Println("Found Pokemon:")
	for _, p := range area.PokemonEncounters {
		fmt.Printf("- %s\n", p.Pokemon.Name)
	}
	return nil
}
