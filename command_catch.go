package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, name string) error {
	if name == "" {
		return fmt.Errorf("invalid name")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemonResp, err := cfg.pokeapiClient.CatchPokemon(name)
	if err != nil {
		return err
	}
	if rand.IntN(1000) > pokemonResp.BaseExperience {
		fmt.Printf("%s was caught!", name)
		cfg.pokedex[name] = pokemonResp
		return nil
	}
	fmt.Printf("%s escaped!", name)
	return nil
}
