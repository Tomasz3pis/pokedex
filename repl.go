package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokeapi "github.com/Tomasz3pis/pokedex/internal"
	"github.com/Tomasz3pis/pokedex/internal/pokedex"
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
			var params string
			if len(words) > 1 {
				params = words[1]
			}
			err := command.callback(cfg, params)
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
	pokedex          *pokedex.Pokedex
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
		"explore": {
			name:        "explore <area_name>",
			description: "List pokemons in provided area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Add pokemon to users pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Print information about pokemon in pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all pokemons in your pokedex",
			callback:    commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
