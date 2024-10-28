package main

import (
	"time"

	pokeapi "github.com/Tomasz3pis/pokedex/internal"
	"github.com/Tomasz3pis/pokedex/internal/pokedex"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient:    pokeClient,
		prevLocationsURL: nil,
		nextLocationsURL: nil,
		pokedex:          pokedex.NewPokedex(),
	}
	startRepl(cfg)
}
