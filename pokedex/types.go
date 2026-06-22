package pokedex

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

type CacheInterface interface {
	Add(key string, data []byte)
	Get(key string) ([]byte, bool)
}

type Config struct {
	Count         int            `json:"count"`
	NextURL       *string        `json:"next"`
	PreviousURL   *string        `json:"previous"`
	Results       []LocationArea `json:"results"`
	CacheStorage  CacheInterface
}
