// Package pokeapi get locationArea
package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	pokecache "github.com/BananaDest/pokedexGo/internal/pokecache"
)

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
type LocationAreaByName struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetLocationAreas(url string) (LocationArea, error) {
	url += "/location-area"
	cache := pokecache.NewCache(10000)
	entry, exists := cache.Get(url)
	if exists {
		locationArea := LocationArea{}
		err := json.Unmarshal(entry, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}
	if res.StatusCode > 299 {
		return LocationArea{}, fmt.Errorf("non success status code: %v", res.Status)
	}
	cache.Add(url, body)
	locationArea := LocationArea{}
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	return locationArea, nil
}

func GetLocationArea(name string) (LocationAreaByName, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + name
	cache := pokecache.NewCache(10000)
	entry, exists := cache.Get(url)
	if exists {
		locationArea := LocationAreaByName{}
		err := json.Unmarshal(entry, &locationArea)
		if err != nil {
			return LocationAreaByName{}, err
		}
		return locationArea, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return LocationAreaByName{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaByName{}, err
	}
	if res.StatusCode > 299 {
		return LocationAreaByName{}, fmt.Errorf("non success status code: %v", res.Status)
	}
	cache.Add(url, body)
	locationArea := LocationAreaByName{}
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return LocationAreaByName{}, err
	}
	return locationArea, nil
}
