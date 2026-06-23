package pokedex


import (
	"fmt"
	"math/rand"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func checkIfPokemonWasCaught(cfg *Config, Pokemon PokemonDetails) {

	chance := 100 - (Pokemon.BaseExperience / 4)
	roll := rand.Intn(100)

	if roll < chance {
		fmt.Printf("%s escaped!\n", Pokemon.Name)
	} else {
		cfg.MapPokemon[Pokemon.Name] = Pokemon
		fmt.Printf("%s was caught!\n", Pokemon.Name)
		fmt.Printf("You may now inspect it with the inspect command.\n")
	}
}

func PokemonCatchRequest(cfg *Config, url string) error {
	bytesSaved, ok := cfg.CacheStorage.Get(url)
	var Pokemon PokemonDetails 

	if ok {
		if err := json.Unmarshal(bytesSaved, &Pokemon); err != nil {
			return err
		}
		checkIfPokemonWasCaught(cfg, Pokemon)
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


	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	cfg.CacheStorage.Add(url, body)

	if err := json.Unmarshal(body, &Pokemon); err != nil {
		return err
	}

	checkIfPokemonWasCaught(cfg, Pokemon)
	return nil
}
