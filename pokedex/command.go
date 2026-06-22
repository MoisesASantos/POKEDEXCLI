package pokedex

import (
	"fmt"
	"os"
)

func commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("map: Display the next 20 location areas")
	fmt.Println("mapb: Display the previous 20 location areas")
	return nil
}

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(cfg *Config) error {
	if cfg.NextURL != nil {
		return MakeRequest(*cfg.NextURL, cfg)
	}

	return MakeRequest(
		"https://pokeapi.co/api/v2/location-area",
		cfg,
	)
}

func commandMapb(cfg *Config) error {
	if cfg.PreviousURL == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	return MakeRequest(*cfg.PreviousURL, cfg)
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Display a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			Name:        "map",
			Description: "Display the next 20 location areas",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Display the previous 20 location areas",
			Callback:    commandMapb,
		},
	}
}
