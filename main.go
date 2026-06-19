package main

import (
	"fmt"
	"github.com/MoisesASantos/POKEDEXCLI/repl"
	"bufio"
	"os"
	"strings"
	"json"
	"net/http"
	"encoding/json"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n\n")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}


func commandMap() error {

}

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
		description: "It displays the names of 20 location areas in the Pokemon world",
		callback:    commandMap,
	},
}

func main() {
	
	scanner := bufio.NewScanner(os.Stdin)

	

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		lower_string := strings.ToLower(line)
		result_list := repl.CleanInput(lower_string)
		command, exist := commands[result_list[0]]
		if !exist {
			fmt.Printf("Unknown command\n")
			os.Exit(1)
		}

		err := command.callback()
		if err != nil {
			fmt.Println("Issue to execute the command\n")
		}
	}
}
