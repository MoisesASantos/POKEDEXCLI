package pokedex

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func MakeAreaRequest(url string, cfg *Config) error {
	if bytesGuardados, ok := cfg.CacheStorage.Get(url); ok {
		if err := json.Unmarshal(bytesGuardados, cfg); err != nil {
			return err
		}

		printAreaResults(cfg)
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

	cfg.CacheStorage.Add(url, body)

	if err := json.Unmarshal(body, cfg); err != nil {
		return err
	}

	printAreaResults(cfg)
	return nil
}

func printAreaResults(cfg *Config) {
	for _, area := range cfg.Results {
		fmt.Println(area.Name)
	}
}







