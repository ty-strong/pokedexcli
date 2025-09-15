package main

import (
	"fmt"
)

func commandExit(cfg *Config, args ...string) (bool, error) {
	fmt.Println("Closing the Pokedex... Goodbye!")
	return true, nil
}
