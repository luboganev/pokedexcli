package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, params []string) error {
	parameterErrorsMessage := "Please add the area as a parameter like that: 'catch pokemon-name'"
	if len(params) < 1 {
		return errors.New("No pokemon name provided." + parameterErrorsMessage)
	}
	if len(params) > 1 {
		return errors.New("More than one pokemon name provided." + parameterErrorsMessage)
	}
	pokemonName := params[0]
	if pokemon, ok := cfg.caughtPokemon[pokemonName]; ok {
		pokemonDescription := fmt.Sprintf("Name: %s\nHeight: %d\nWeight: %d", pokemon.Name, pokemon.Height, pokemon.Weight)

		// Add stats
		pokemonDescription += "\nStats:"
		for _, stat := range pokemon.Stats {
			pokemonDescription += fmt.Sprintf("\n  -%s: %d", stat.Stat.Name, stat.BaseStat)
		}

		// Add types
		pokemonDescription += "\nTypes:"
		for _, t := range pokemon.Types {
			pokemonDescription += fmt.Sprintf("\n  - %s", t.Type.Name)
		}

		fmt.Println(pokemonDescription)
	} else {
		return fmt.Errorf("You don't have '%s'. Did you forget to catch it? Maybe it broke out... Who knows. It's a mystery.", pokemonName)
	}
	return nil
}
