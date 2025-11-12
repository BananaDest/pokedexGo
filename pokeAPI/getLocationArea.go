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

func GetLocationAreas(url string) (LocationArea, error) {
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
