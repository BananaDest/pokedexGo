package main

import (
	"errors"
	"fmt"

	pokeapi "github.com/BananaDest/pokedexGo/pokeAPI"
)

func CommandExplore(config *Config) error {
	location, ok := config.Parameters["locationArea"]
	if !ok {
		return errors.New("no parameters sent")
	}
	locationArea, err := pokeapi.GetLocationArea(location)
	if err != nil {
		return err
	}
	fmt.Println("Found pokemon:")
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
