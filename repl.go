package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ty-strong/pokedexcli/internal/pokeapi"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	initialURL := "https://pokeapi.co/api/v2/location-area"
	config := &Config{
		PokeapiClient:       pokeClient,
		NextLocationAreaURL: &initialURL,
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		shouldExit, err := command.callback(config, args...)
		if err != nil {
			fmt.Println(err)
		}
		if shouldExit {
			break
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) (bool, error)
}

type Config struct {
	PokeapiClient           pokeapi.Client
	NextLocationAreaURL     *string
	PreviousLocationAreaURL *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Lists the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Lists the names of the pokemon in an area",
			callback:    commandExplore,
		},
	}
}
