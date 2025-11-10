package main

import (
	"fmt"
	pokeapi "github.com/BananaDest/pokedexGo/pokeAPI"
)

func CommandMapB(config *Config) error {
	if config.Previous == "" {
		fmt.Println("You are on first page")
		return nil
	}
	locationAreas, err := pokeapi.GetLocationAreas(config.Previous)
	if err != nil {
		return err
	}
	for _, locationArea := range locationAreas.Results {
		fmt.Println(locationArea.Name)
	}
	config.Next = locationAreas.Next
	config.Previous = locationAreas.Previous
	return nil
}
