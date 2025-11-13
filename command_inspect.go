package main

import (
	"errors"
	"fmt"
)

func CommandInspect(config *Config) error {
	pokemonName, ok := config.Parameters["name"]
	if !ok {
		return errors.New("no parameters sent")
	}
	pokemon, ok := config.Pokedex[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weigth: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("-%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("- %v\n", t.Type.Name)
	}

	return nil
}
