package main

import (
	"fmt"
	"math/rand/v2"
	"os"
)

type Command struct {
	name        string
	description string
	callback    func(*config, string) error
}

func commandHelp(cfg *config, params string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit(cfg *config, params string) error {
	fmt.Println("Bye bye!")
	os.Exit(0)
	return nil
}

func commandMapf(cfg *config, params string) error {
	location, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = location.Next
	cfg.prevLocationsURL = location.Previous

	for _, v := range location.Results {
		fmt.Println(v.Name)
	}

	return nil
}

func commandMapb(cfg *config, params string) error {
	if cfg.prevLocationsURL == nil {
		return fmt.Errorf("No previous locations avaliable")
	}
	location, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = location.Next
	cfg.prevLocationsURL = location.Previous

	for _, v := range location.Results {
		fmt.Println(v.Name)
	}
	return nil
}

func commandExplore(cfg *config, params string) error {
	if params == "" {
		fmt.Println("Area name not provided...")
		return nil
	}
	fmt.Printf("Exploring %v...\n", params)
	pokemons, err := cfg.pokeapiClient.ListPokemons(params)
	if err != nil {
		return err
	}
	fmt.Printf("Found Pokemon:\n")
	for _, v := range pokemons.PokemonEncounters {
		fmt.Println(v.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, params string) error {
	if params == "" {
		fmt.Println("Pokemon name missing...")
		return nil
	}
	pokemon, err := cfg.pokeapiClient.GetPokemon(params)

	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", params)
	catchChance := pokemon.BaseExperience / 4
	roll := rand.IntN(100)
	if roll > catchChance {
		fmt.Printf("%v was cought!\n", params)
		cfg.pokedex.AddEntry(params, pokemon)
		fmt.Printf("You may now inspect it with the inspect command.\n")
	} else {
		fmt.Printf("%v escaped!\n", params)
	}

	return nil
}

func commandInspect(cfg *config, params string) error {
	if pok, exists := cfg.pokedex.GetEntry(params); exists {
		fmt.Printf("Name: %v\nHeight: %v\nWeight: %v\nStats:\n", pok.Name, pok.Height, pok.Weight)
		for _, v := range pok.Stats {
			fmt.Printf("  -%v: %v\n", v.Name, v.Value)
		}
		fmt.Println("Types:")
		for _, v := range pok.Types {
			fmt.Printf("  - %v\n", v.Name)
		}
		return nil
	} else {
		fmt.Println("You have not caught that pokemon")
		return nil
	}
}

func commandPokedex(cfg *config, params string) error {
	fmt.Printf("Your Pokedex:\n")
	for _, v := range cfg.pokedex.ListEntries() {
		fmt.Printf("  - %v\n", v)
	}
	return nil
}
