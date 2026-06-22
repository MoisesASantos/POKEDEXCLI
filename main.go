package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
	"github.com/MoisesASantos/POKEDEXCLI/repl"
	"github.com/MoisesASantos/POKEDEXCLI/internal"
)

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type config struct {
	Count       int            `json:"count"`
	NextURL     *string        `json:"next"`
	PreviousURL *string        `json:"previous"`
	Results     []LocationArea `json:"results"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type CacheInterface interface {
	Add(key_intro string, data []byte)
	Get(key string) ([]byte, bool)
}

var CacheStorage CacheInterface


func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("map: Display the next 20 location areas")
	fmt.Println("mapb: Display the previous 20 location areas")
	return nil
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func makeRequest(url string, cfg *config) error {

	if bytesGuardados, ok := CacheStorage.Get(url); ok {
		err := json.Unmarshal(bytesGuardados, cfg)
		if err != nil {
			return err
		}
		imprimirResultados(cfg)
		return nil
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("HTTP error code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	CacheStorage.Add(url, body)
	err = json.Unmarshal(body, cfg)
	if err != nil {
		return err
	}

	imprimirResultados(cfg)
	return nil
}

func imprimirResultados(cfg *config) {
	for _, area := range cfg.Results {
		fmt.Println(area.Name)
	}
}


func commandMap(cfg *config) error {
	
	if cfg.NextURL != nil {
		return makeRequest(*cfg.NextURL, cfg)
	}

	return makeRequest(
		"https://pokeapi.co/api/v2/location-area",
		cfg,
	)
}

func commandMapb(cfg *config) error {
	if cfg.PreviousURL == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	return makeRequest(*cfg.PreviousURL, cfg)
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const baseTime = 300000 * time.Millisecond
	CacheStorage = pokecache.NewCache(baseTime)
	cfg := &config{}

	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 location areas",
			callback:    commandMapb,
		},
	}

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		resultList := repl.CleanInput(strings.ToLower(line))

		if len(resultList) == 0 {
			continue
		}

		command, exists := commands[resultList[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.callback(cfg); err != nil {
			fmt.Println(err)
		}
	}
}
