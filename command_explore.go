package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *Config, args ...string) (bool, error) {
	if len(args) != 1 {
		return false, errors.New("you must provide a location name")
	}
	locationAreaName := args[0]

	fmt.Printf("Exploring %s...\n", locationAreaName)
	locationArea, err := cfg.PokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return false, err
	}

	fmt.Println("Found Pokemon:")
	if len(locationArea.PokemonEncounters) == 0 {
		fmt.Println("No pokemon found in this area.")
	}
	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return false, nil
}
