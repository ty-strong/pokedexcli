package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *Config) (bool, error) {
	if cfg.NextLocationAreaURL == nil {
		return false, errors.New("you've reached the end of the map")
	}

	locationsResp, err := cfg.PokeapiClient.ListLocations(cfg.NextLocationAreaURL)
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
