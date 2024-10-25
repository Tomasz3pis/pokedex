package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokeapi "github.com/Tomasz3pis/pokedex/internal"
)

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Pokedex! Choose one of avaliable commands:")
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}
type Command struct {
	name        string
	description string
	callback    func(*config) error
}

func commandHelp(cfg *config) error {
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

func commandExit(cfg *config) error {
	fmt.Println("Bye bye!")
	os.Exit(0)
	return nil
}

func commandMapf(cfg *config) error {
	//TODO get locations, set next and prev, print locations
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

func commandMapb(cfg *config) error {
	//TODO check prev, get locations, set next and prev, list locations
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

func getCommands() map[string]Command {
	return map[string]Command{

		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display next 20 areas neme",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 areas name",
			callback:    commandMapb,
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
