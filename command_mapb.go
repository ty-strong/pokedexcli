package main

import (
	"errors"
	"fmt"
)

func commandMapb(cfg *Config) (bool, error) {
	if cfg.PreviousLocationAreaURL == nil {
		return false, errors.New("you're on the first page")
	}

	locationsResp, err := cfg.PokeapiClient.ListLocations(cfg.PreviousLocationAreaURL)
	if err != nil {
		return false, err
	}

	cfg.NextLocationAreaURL = locationsResp.Next
	cfg.PreviousLocationAreaURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	return false, nil
}
