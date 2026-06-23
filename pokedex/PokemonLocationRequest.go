package pokedex

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func printPokemonLocationResult(data LocationAreaResponse) {
	for _, encounter := range data.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
}

func MakeRequestLocation(cfg *Config, url string) error {
	bytesGuardados, ok := cfg.CacheStorage.Get(url)
	var data LocationAreaResponse

	if ok {
		if err := json.Unmarshal(bytesGuardados, &data); err != nil {
			return err
		}
		printPokemonLocationResult(data)
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

	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	printPokemonLocationResult(data)
	return nil
}
