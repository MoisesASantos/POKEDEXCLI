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

type StatPokemon struct {
	Name string `json:"name"`
}

type StatsPokemons struct {
	BaseStat int         `json:"base_stat"`
	Stat     StatPokemon `json:"stat"`
}

type PokemonType struct {
	Name string `json:"name"`
}

type TypeSlot struct {
	Type PokemonType `json:"type"`
}

type PokemonDetails struct {
	Name           string          `json:"name"`
	Height         int             `json:"height"`
	Weight         int             `json:"weight"`
	BaseExperience int             `json:"base_experience"`
	Types          []TypeSlot      `json:"types"`
	Stats          []StatsPokemons `json:"stats"`
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
	NextURL       *string
	PreviousURL   *string
	Results       []LocationArea
	CacheStorage  CacheInterface
	MapPokemon    map[string]PokemonDetails
}

func NewConfig(cache CacheInterface) *Config {
	return &Config{
		CacheStorage: cache,
		MapPokemon:   make(map[string]PokemonDetails),
	}
}
