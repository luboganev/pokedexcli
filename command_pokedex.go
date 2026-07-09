package main

import (
	"fmt"
)

func commandPokedex(cfg *config, params []string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You haven't caught any pokemons yet!")
	} else {
		fmt.Println("Your Pokedex:")
		for _, pokemon := range cfg.caughtPokemon {
			fmt.Printf(" -%s\n", pokemon.Name)
		}
	}
	return nil
}
