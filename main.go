package main

import (
	"time"

	"github.com/luboganev/pokedexcli/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeapiClient,
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}

	startRepl(cfg)
}
