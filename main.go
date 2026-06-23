package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/MoisesASantos/POKEDEXCLI/internal"
	"github.com/MoisesASantos/POKEDEXCLI/pokedex"
	"github.com/MoisesASantos/POKEDEXCLI/repl"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	const baseTime = 300000 * time.Millisecond

	cache := pokecache.NewCache(baseTime)
	cfg := pokedex.NewConfig(cache)
	commands := pokedex.GetCommands()

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

		var arg string
		if len(resultList) > 1 {
			arg = resultList[1]
		}

		command, exists := commands[resultList[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.Callback(cfg, arg); err != nil {
			fmt.Println(err)
		}
	}
}
