package main

import (
	"fmt"
)

func commandExplore(cfg *config, area string) error {
	if area == "" {
		return fmt.Errorf("Must enter valid area")
	}
	fmt.Printf("Exploring %s...\n", area)
	locationsResp, err := cfg.pokeapiClient.ExploreLocation(area)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, loc := range locationsResp.PokemonEncounters {
		fmt.Printf(" - %s\n", loc.Pokemon.Name)
	}
	return nil

}
