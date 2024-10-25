package main

import (
	"time"

	pokeapi "github.com/Tomasz3pis/pokedex/internal"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient:    pokeClient,
		prevLocationsURL: nil,
		nextLocationsURL: nil,
	}
	startRepl(cfg)
}
