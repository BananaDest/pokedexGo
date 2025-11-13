package main

import (
	"fmt"
)

func CommandPokedex(config *Config) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.Pokedex {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	return nil
}
