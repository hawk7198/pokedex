package pokeapi

type Pokemon struct {
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

type PokemonStat struct {
	BaseStat int  `json:"base_stat"`
	Stat     Stat `json:"stat"`
}

type PokemonType struct {
	Type Type `json:"type"`
}

type Stat struct {
	Name string `json:"name"`
}

type Type struct {
	Name string `json:"name"`
}
