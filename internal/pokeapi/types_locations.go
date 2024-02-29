package pokeapi

// RespShallowLocations -
type RespShallowLocations struct {
	Count    int                `json:"count"`
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []NamedApiResource `json:"results"`
}

type Location struct {
	Id                     int                     `json:"id"`
	Name                   string                  `json:"name"`
	GameIndex              int                     `json:"game_index"`
	EncounteredMethodRates []EncounteredMethodRate `json:"encountered_method_rates"`
	Location               NamedApiResource        `json:"location"`
	Names                  []Name                  `json:"names"`
	PokemonEncounters      []PokemonEncounter      `json:"pokemon_encounters"`
}

type EncounteredMethodRate struct {
	EncounteredMethod NamedApiResource          `json:"encountered_method"`
	VersionDetails    []EncounterVersionDetails `json:"version_details"`
}

type EncounterVersionDetails struct {
	Rate    int              `json:"rate"`
	Version NamedApiResource `json:"version"`
}

type Name struct {
	Name     string           `json:"name"`
	Language NamedApiResource `json:"language"`
}

type PokemonEncounter struct {
	Pokemon        NamedApiResource         `json:"pokemon"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

type NamedApiResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type VersionEncounterDetail struct {
	Version          NamedApiResource `json:"version"`
	MaxChance        int              `json:"max_chance"`
	EncounterDetails []Encounter      `json:"encounter_details"`
}

type Encounter struct {
	MinLevel        int                `json:"min_level"`
	MaxLevel        int                `json:"max_level"`
	ConditionValues []NamedApiResource `json:"condition_values"`
	Chance          int                `json:"chance"`
	Method          NamedApiResource   `json:"method"`
}
