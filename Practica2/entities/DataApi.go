package entities

type Response struct {
	NameRegion     string        `json:"name"`
	PokemonEntries []PokeEntries `json:"pokemon_entries"`
}

type PokeEntries struct {
	ID      int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
}

type Data struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
