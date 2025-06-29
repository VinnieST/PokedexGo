package main

import (
	"pokedexcli/internal/pokeapi"
	"time"
)

type config struct {
	pokeapiClient    *pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          map[string]pokeapi.Pokemon
}

func main() {
	cfg := &config{
		pokeapiClient:    pokeapi.NewClient(5*time.Second, 5*time.Second),
		nextLocationsURL: nil,
		prevLocationsURL: nil,
	}
	startRepl(cfg)
}
