package cmd

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"

	"codeberg.org/OliveiraJ/pokedexcli/internal/domain"
)

const (
	CatchName        = "catch"
	CatchDescription = "Catch a pokemon"
)

func CatchCommand() domain.CliCommand {
	return domain.CliCommand{
		Name:        CatchName,
		Description: CatchDescription,
		Callback:    commandCatch,
	}
}

func commandCatch(config *domain.Config, args []string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])

	pokemon, err := config.PokeAPIClient.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", args[0]))
	if err != nil {
		return err
	}

	var caughtPokemon domain.Pokemon
	if err := json.Unmarshal([]byte(pokemon), &caughtPokemon); err != nil {
		return err
	}

	if !calculateCatchChance(caughtPokemon.BaseExperience) {
		fmt.Println("Failed to catch the Pokemon!")
		return nil
	}

	config.Pokedex.Data[caughtPokemon.Name] = caughtPokemon

	fmt.Printf("Caught a %s!\n", caughtPokemon.Name)

	return nil
}

func calculateCatchChance(baseExperience int) bool {
	chance := 1 / (1 + math.Pow10((baseExperience-50)/200))
	inf := 0.1 + rand.Float64()*0.9

	return chance > inf
}
