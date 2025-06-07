package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/ProbsPropps/pokedexcli/internal/pokeapi"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	locations, err := pokeapi.GetLocations("https://pokeapi.co/api/v2/location-area/")
		if err != nil {
		fmt.Errorf("the api is not working correctly: couldn't gather locations: %v", err)
	}
	locPtr := &locations

	var commands = map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help messages",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays 20 location areas in the Pokemon world",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the previous 20 location areas",
			callback: commandMapb,
		},
	}

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		text := scanner.Text()
		command, ok := commands[text]
		if ok {
			switch command.name {
			case "exit": 
				err := command.callback(locPtr)
				if err != nil {
					fmt.Println("Error encountered when running the command")
				}
			case "help":
				err := command.callback(locPtr)
				if err != nil {
					fmt.Println("Error encountered when running the command")
				}
				for key, val := range commands {
					fmt.Printf("%v: %v\n", key, val.description)
				}
			case "map":
				err := command.callback(locPtr)
				if err != nil {
					fmt.Println("Error encountered when running the command")
				}
				locations, err = pokeapi.GetLocations(*locations.NextURL)
				if err != nil {
					fmt.Println("You've reached the end of the location list")
				} else{ locPtr = &locations }
			case "mapb":
				err := command.callback(locPtr)
				if err != nil {
					fmt.Println("Error encountered when running the command")
				}
				locations, err = pokeapi.GetLocations(*locations.PreviousURL)
				if err != nil {
					fmt.Println("You're at the beginning of the location list")
				} else { locPtr = &locations}
				
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	var test []string
	test = strings.Fields(strings.ToLower(text))
	return test
}

func commandExit(locations *pokeapi.LocationArea) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(locations *pokeapi.LocationArea) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	return nil

}

func commandMap(locations *pokeapi.LocationArea) error {
	data := locations.Results
	for _, name := range(data) {
		fmt.Println(name.Name)
	}
	return nil
}

func commandMapb(locations *pokeapi.LocationArea) error {
	data := locations.Results
	for _, name := range(data) {
		fmt.Println(name.Name)
	}
	return nil
}

type cliCommand struct {
	name 		string
	description string
	callback func(l *pokeapi.LocationArea) error
}

type config struct {
	next		*string
	previous	*string
}



