package cmd

import (
	"fmt"

	"codeberg.org/OliveiraJ/pokedexcli/internal/domain"
)

const (
	InspectName        = "inspect"
	InspectDescription = "Inspect a pokemon"
)

func InspectCommand() domain.CliCommand {
	return domain.CliCommand{
		Name:        InspectName,
		Description: InspectDescription,
		Callback:    commandInspect,
	}
}

func commandInspect(config *domain.Config, args []string) error {
	if len(args) == 0 {
		return nil
	}
	pokemonName := args[0]

	pokemon, ok := config.Pokedex.Data[pokemonName]
	if !ok {
		fmt.Printf("No pokemon found in your pokedex with name %s\n", pokemonName)
		return nil
	}

	fmt.Printf("Name: %s\n", pokemonName)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Base Experience: %d\n", pokemon.BaseExperience)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokemonType.Type.Name)
	}

	return nil
}
