package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, params []string) error {
	parameterErrorsMessage := "Please add the area as a parameter like that: 'explore area-name'"
	if len(params) < 1 {
		return errors.New("No area name provided. " + parameterErrorsMessage)
	}
	if len(params) > 1 {
		return errors.New("More than one area name provided. " + parameterErrorsMessage)
	}
	areaName := params[0]
	fmt.Printf("Exploring %s...\n", areaName)

	location, err := cfg.pokeapiClient.GetLocation(areaName)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
