package main

import (
	"bufio"
	"fmt"
	"os"

	pokeClient "github.com/Tomasz3pis/pokedex/internal"
)

type Command struct {
	name        string
	description string
	callback    func(*pokeClient.Config) error
}

func commandHelp(cfg *pokeClient.Config) error {
	fmt.Println("This is a help command. Not so helpfull yet.")
	return nil
}

func commandExit(cfg *pokeClient.Config) error {
	fmt.Println("Bye bye!")
	os.Exit(0)
	return nil
}

func commandMap(cfg *pokeClient.Config) error {
	return nil
}

func commandMapb(cfg *pokeClient.Config) error {
	return nil
}

func createCommands() map[string]Command {
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
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 areas name",
			callback:    commandMapb,
		},
	}
}

func main() {
	cfg, err := pokeClient.NewConfig()
	if err != nil {
		fmt.Errorf("Faild to load config: %w", err)
	}
	cmds := createCommands()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Pokedex! Choose one of avaliable commands:")
	for _, v := range cmds {
		fmt.Printf("Name: %s\n", v.name)
		fmt.Printf("Description: %s\n\n", v.description)
	}
	fmt.Println("Provide command")
	for scanner.Scan() {
		if cmd, exists := cmds[scanner.Text()]; exists {
			cmd.callback(&cfg)
		}

	}
}
