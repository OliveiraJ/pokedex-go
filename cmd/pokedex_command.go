package cmd

import (
	"fmt"

	"codeberg.org/OliveiraJ/pokedexcli/internal/domain"
)

const (
	PokedexName        = "pokedex"
	PokedexDescription = "Show the name of all pokemons you caught."
)

func PokedexCommand() domain.CliCommand {
	return domain.CliCommand{
		Name:        PokedexName,
		Description: PokedexDescription,
		Callback:    commandPokedex,
	}
}

func commandPokedex(config *domain.Config, args []string) error {
	if len(config.Pokedex.Data) == 0 {
		fmt.Println("Your pokedex is empty")
		return nil
	}

	for name, _ := range config.Pokedex.Data {
		fmt.Printf("%s\n", name)
	}

	return nil
}
