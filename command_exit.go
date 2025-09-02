package main

import (
	"fmt"
)

func commandExit(cfg *Config) (bool, error) {
	fmt.Println("Closing the Pokedex... Goodbye!")
	return true, nil
}
