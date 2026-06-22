package pokedex

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type LocationAreaResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, string) error
}

type CacheInterface interface {
	Add(key string, data []byte)
	Get(key string) ([]byte, bool)
}

type Config struct {
	NextURL       *string        `json:"next"`
	PreviousURL   *string        `json:"previous"`
	Results       []LocationArea `json:"results"`
	CacheStorage  CacheInterface
}
