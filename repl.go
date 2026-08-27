package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"codeberg.org/OliveiraJ/pokedexcli/cmd"
	"codeberg.org/OliveiraJ/pokedexcli/internal/domain"
	"codeberg.org/OliveiraJ/pokedexcli/internal/pokeapi"
	"codeberg.org/OliveiraJ/pokedexcli/internal/pokecache"
)

func main() {
	commands := cmd.GetCommands()
	scanner := bufio.NewScanner(os.Stdin)
	cache := pokecache.NewCache(5 * time.Minute)
	config := &domain.Config{
		Next:          "https://pokeapi.co/api/v2/location-area/",
		Previous:      "",
		PokeAPIClient: pokeapi.NewClient(cache),
		Pokedex:       domain.Pokedex{Data: map[string]domain.Pokemon{}},
	}

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()

		inputText := cleanInput(scanner.Text())
		if len(inputText) == 0 {
			continue
		}

		command, exists := commands[inputText[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := command.Callback(config, inputText[1:])
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	words := []string{}
	for word := range strings.FieldsSeq(text) {
		words = append(words, strings.ToLower(word))
	}
	return words
}
