package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, params []string) error {
	parameterErrorsMessage := "Please add the pokemon name as a parameter like that: 'catch pokemon-name'"
	if len(params) < 1 {
		return errors.New("No pokemon name provided. " + parameterErrorsMessage)
	}
	if len(params) > 1 {
		return errors.New("More than one pokemon name provided. " + parameterErrorsMessage)
	}
	pokemonName := params[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	catchSuccessRate := float64(rand.Intn(pokemon.BaseExperience)) / float64(pokemon.BaseExperience)
	catchThreshold := float64(0.75)
	fmt.Printf("Catching attempt -> %.2f\n", catchSuccessRate)
	if catchSuccessRate > catchThreshold {
		fmt.Println("Got them! Congratulations!")
		cfg.caughtPokemon[pokemonName] = pokemon
	} else {
		fmt.Printf("%s got away! Good luck next time!\n", pokemonName)
	}

	return nil
}
