package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Next     string
	Previous string
}

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	conf := Config{
		Next:     "https://pokeapi.co/api/v2//location-area",
		Previous: "https://pokeapi.co/api/v2//location-area",
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
	}
}
