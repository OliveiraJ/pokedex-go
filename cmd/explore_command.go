package cmd

import (
	"encoding/json"
	"fmt"

	"codeberg.org/OliveiraJ/pokedexcli/internal/domain"
)

const (
	ExploreName        = "explore"
	ExploreDescription = "List all available Pokemon in the current location"
)

type ExploreCommandGetResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel       int `json:"min_level"`
				PokemonDetails any `json:"pokemon_details"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func ExploreCommand() domain.CliCommand {
	return domain.CliCommand{
		Name:        ExploreName,
		Description: ExploreDescription,
		Callback:    commandExplore,
	}
}

func commandExplore(config *domain.Config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected 1 argument, got %d", len(args))
	}
	locationName := args[0]

	var locationAreas ExploreCommandGetResponse
	body, err := config.PokeAPIClient.Get("https://pokeapi.co/api/v2/location-area/" + locationName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		return err
	}

	for _, m := range locationAreas.PokemonEncounters {
		fmt.Println(m.Pokemon.Name)
	}

	return nil
}
