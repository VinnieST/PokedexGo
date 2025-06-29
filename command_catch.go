package main

import (
	"errors"
	"fmt"
	"math/rand"
	"pokedexcli/internal/pokeapi"
	"time"
)

func commandCatch(cfg *config, args []string) error {
	if cfg == nil || len(args) < 1 {
		return errors.New("usage: catch <pokemon-name>")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("could not find pokemon: %v", err)
	}

	if pokemon.ID == 0 {
		return fmt.Errorf("pokemon '%s' not found", pokemonName)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	baseExp := pokemon.BaseExperience
	catchChance := 100 - baseExp
	if catchChance < 10 {
		catchChance = 10
	}

	rand.Seed(time.Now().UnixNano())
	roll := rand.Intn(100) + 1 // 1-100

	if roll <= catchChance {

		if cfg.pokedex == nil {
			cfg.pokedex = make(map[string]pokeapi.Pokemon)
		}

		if _, exists := cfg.pokedex[pokemon.Name]; !exists {
			cfg.pokedex[pokemon.Name] = pokemon
			fmt.Printf("You caught %s! (ID: %d)\n", pokemon.Name, pokemon.ID)
		} else {
			fmt.Printf("You caught %s! (ID: %d) (Already in your Pokedex)\n", pokemon.Name, pokemon.ID)
		}
	} else {
		fmt.Printf("You failed to catch %s. Better luck next time!\n", pokemon.Name)
	}
	return nil
}
