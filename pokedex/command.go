package pokedex

import (
	"fmt"
	"os"
)

func commandHelp(cfg *Config, _ string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("map: Display the next 20 location areas")
	fmt.Println("mapb: Display the previous 20 location areas")
	return nil
}

func commandExit(cfg *Config, _ string) error {
	
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(cfg *Config, _ string) error {

	if cfg.NextURL != nil {
		return MakeAreaRequest(*cfg.NextURL, cfg)
	}

	return MakeAreaRequest(
		"https://pokeapi.co/api/v2/location-area",
		cfg,
	)
}

func commandMapb(cfg *Config, _ string) error {

	if cfg.PreviousURL == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	return MakeAreaRequest(*cfg.PreviousURL, cfg)
}

func commandExplore(cfg *Config, arg string) error {
	
	if arg == "" {
		return fmt.Errorf("you must provide a location area")
	}

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", arg)
	fmt.Printf("Exploring %s...\nFound Pokemon:\n", arg)
	return MakeRequestLocation(cfg, url)
}

func commandCatch(cfg *Config, arg string) error {

	if arg == "" {
		return fmt.Errorf("you must provide the pokemon name")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", arg)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", arg)
	return PokemonCatchRequest(cfg, url)
}

func printPokemonDetails(pokemon PokemonDetails) error {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}
	return nil
}

func commandInspect(cfg *Config, arg string) error {

	if arg == "" {
		return fmt.Errorf("you must provide the pokemon name")
	}

	PokemonDetails, ok := cfg.MapPokemon[arg]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	return printPokemonDetails(PokemonDetails)
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
		"explore": {
			Name:        "explore",
			Description: "Display the name off all pokemon found on especific paramerater",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "This command is used to try to catch a pokemon",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Show Details About the pokemon",
			Callback:    commandInspect,
		},
	}
}
