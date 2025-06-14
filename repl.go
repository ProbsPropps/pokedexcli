package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/ProbsPropps/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient 	pokeapi.Client
	nextURL 		*string
	previousURL 	*string
	caughtPokemon 	map[string]pokeapi.ShallowPokemon
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	pokeCaught := make(map[string]pokeapi.ShallowPokemon)
	cfg.caughtPokemon = pokeCaught
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		
		if len(words) == 0 {
			continue
		}
		
		commandName := words[0]

		command, exists := getCommands()[commandName]

		if exists {
			if command.needsMultiple {
				if len(words) == 1 {
					fmt.Println("Missing argument for command")
					continue
				}
				err := command.callback(cfg, words[1])
				if err != nil {
					fmt.Println(err)
				}
				continue
			} 
			command.callback(cfg, "")
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}


func cleanInput(text string) []string {
	output := strings.ToLower(text)
	return strings.Fields(output)
}

type cliCommand struct {
	name 			string
	description 	string
	callback 		func(cfg *config, str string) error
	needsMultiple	bool
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
			needsMultiple: false,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
			needsMultiple: false,
		},
		"map": {
			name: "map",
			description: "Get the next 20 locations",
			callback: commandMapf,
			needsMultiple: false,
		},
		"mapb": {
			name: "mapb",
			description: "Get the previous 20 locations",
			callback: commandMapb,
			needsMultiple: false,
		},
		"explore": {
			name: "explore",
			description: "Explore a region to find the natural Pokemon",
			callback: commandExplore,
			needsMultiple: true,
		},
		"catch": {
			name: "catch",
			description: "Try to catch a pokemon!",
			callback: commandCatch,
			needsMultiple: true,
		},
		"inspect": {
			name: "inspect",
			description: "Look at a Pokemon's Pokedex record",
			callback: commandInspect,
			needsMultiple: true,
		},
		"pokedex": {
			name: "pokedex",
			description: "See what you have in your Pokedex",
			callback: commandPokedex,
			needsMultiple: false,
		},
	}
}
