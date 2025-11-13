package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokeapi "github.com/BananaDest/pokedexGo/pokeAPI"
)

type Config struct {
	Next       string
	Previous   string
	Parameters map[string]string
	Pokedex    map[string]pokeapi.Pokemon
}

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	conf := Config{
		Next:       "https://pokeapi.co/api/v2",
		Previous:   "https://pokeapi.co/api/v2",
		Parameters: make(map[string]string),
		Pokedex:    make(map[string]pokeapi.Pokemon),
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := scanner.Text()
		cleaned := CleanInput(input)
		if len(cleaned) == 0 {
			continue
		}
		command, ok := GetCommands()[cleaned[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if command.name == "explore" {
			if len(cleaned) > 2 {
				fmt.Println("Missing Argument: Location area name")
				continue
			}
			conf.Parameters["locationArea"] = cleaned[1]
		}
		if command.name == "catch" || command.name == "inspect" {
			if len(cleaned) > 2 {
				fmt.Println("Missing Argument: pokemon name")
				continue
			}
			conf.Parameters["name"] = cleaned[1]
		}
		err := command.callback(&conf)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		continue
	}
}

func CleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    CommandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    CommandHelp,
		},
		"map": {
			name:        "map",
			description: "Return Next 20 locations",
			callback:    CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Return Previous 20 locations",
			callback:    CommandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Lists all pokemon in location",
			callback:    CommandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    CommandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon",
			callback:    CommandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Returns pokedex",
			callback:    CommandPokedex,
		},
	}
}
