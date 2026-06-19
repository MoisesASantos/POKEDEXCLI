package main

import (
	"fmt"
	"github.com/MoisesASantos/POKEDEXCLI/repl"
	"bufio"
	"os"
	"strings"
)

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
		fmt.Printf("Your command was: %s\n", result_list[0])
	}
}
