package main

import (
	"errors"
	"fmt"
	"math/rand/v2"

	pokeapi "github.com/BananaDest/pokedexGo/pokeAPI"
)

func CommandCatch(config *Config) error {
	pokemonName, ok := config.Parameters["name"]
	if !ok {
		return errors.New("no parameters sent")
	}
	pokemon, err := pokeapi.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	catchRate := rand.IntN(3)
	if catchRate == 0 {
		fmt.Printf("catched!\n")
		config.Pokedex[pokemon.Name] = pokemon
		return nil
	}
	fmt.Println("It escaped!")
	return nil
}
