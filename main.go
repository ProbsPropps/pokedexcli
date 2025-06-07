package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
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
				err := command.callback()
				if err != nil {
					fmt.Println("Error encountered when running the command")
				}
			case "help":
				err := command.callback()
				if err != nil {
					fmt.Println("Error encountered when running the command")
				}
				for key, val := range commands {
					fmt.Printf("%v: %v\n", key, val.description)
				}
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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	return nil

}

func commandMap() error {
	return nil
}

func commandMapb() error {
	return nil
}

type cliCommand struct {
	name string
	description string
	callback func() error
}
