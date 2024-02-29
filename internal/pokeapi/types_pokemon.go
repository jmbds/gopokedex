package pokeapi

type Pokemon struct {
	Id                     int                `json:"id"`
	Name                   string             `json:"name"`
	BaseExperience         int                `json:"base_experience"`
	Height                 int                `json:"height"`
	IsDefault              bool               `json:"is_default"`
	Order                  int                `json:"order"`
	Weight                 int                `json:"weight"`
	Abilities              []PokemonAbility   `json:"abilities"`
	Forms                  []NamedApiResource `json:"forms"`
	GameIndices            []VersionGameIndex `json:"game_indices"`
	HeldItems              []PokemonHeldItem  `json:"held_items"`
	LocationAreaEncounters string             `json:"location_area_encounters"`
	Moves                  []PokemonMove      `json:"moves"`
	PastTypes              []PokemonTypePast  `json:"past_types"`
	Sprites                PokemonSprites     `json:"sprites"`
	Cries                  PokemonCries       `json:"cries"`
	Species                NamedApiResource   `json:"species"`
	Stats                  []PokemonStat      `json:"stats"`
	Types                  []PokemonType      `json:"types"`
}

type PokemonAbility struct {
	IsHidden bool             `json:"is_hidden"`
	Slot     int              `json:"slot"`
	Ability  NamedApiResource `json:"ability"`
}

type VersionGameIndex struct {
	GameIndex int              `json:"game_index"`
	Version   NamedApiResource `json:"version"`
}

type PokemonType struct {
	Slot int              `json:"slot"`
	Type NamedApiResource `json:"type"`
}

type PokemonFormType struct {
	Slot int              `json:"slot"`
	Type NamedApiResource `json:"type"`
}

type PokemonTypePast struct {
	Generation NamedApiResource `json:"generation"`
	Types      []PokemonType    `json:"types"`
}

type PokemonHeldItem struct {
	Item           NamedApiResource         `json:"item"`
	VersionDetails []PokemonHeldItemVersion `json:"version_details"`
}

type PokemonHeldItemVersion struct {
	Version NamedApiResource `json:"version"`
	Rarity  int              `json:"rarity"`
}

type PokemonMove struct {
	Move                NamedApiResource     `json:"move"`
	VersionGroupDetails []PokemonMoveVersion `json:"version_group_details"`
}

type PokemonMoveVersion struct {
	MoveLearnMethod NamedApiResource `json:"move_learn_method"`
	VersionGroup    NamedApiResource `json:"version_group"`
	LevelLearnedAt  int              `json:"level_learned_at"`
}

type PokemonStat struct {
	Stat     NamedApiResource `json:"stat"`
	Effort   int              `json:"effort"`
	BaseStat int              `json:"base_stat"`
}

type PokemonSprites struct {
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
	FrontFemale      string `json:"front_female"`
	FrontShinyFemale string `json:"front_shiny_female"`
	BackDefault      string `json:"back_default"`
	BackShiny        string `json:"back_shiny"`
	BackFemale       string `json:"back_female"`
	BackShinyFemale  string `json:"back_shiny_female"`
}

type PokemonCries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}
