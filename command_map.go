package main

import (
	"fmt"

	pokeapi "github.com/BananaDest/pokedexGo/pokeAPI"
)

func CommandMap(config *Config) error {
	if config.Next == "" {
		fmt.Println("You are on last page")
		return nil
	}
	locationAreas, err := pokeapi.GetLocationAreas(config.Next)
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
