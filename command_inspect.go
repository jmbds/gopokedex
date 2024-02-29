package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, caught := cfg.caughtPokemon[name]

	if !caught {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, pokemonStat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", pokemonStat.Stat.Name, pokemonStat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("  -%s\n", pokemonType.Type.Name)
	}
	return nil
}
