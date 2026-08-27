package cmd

import (
	"encoding/json"
	"io"
	"os"
	"testing"

	"codeberg.org/OliveiraJ/pokedexcli/internal/domain"
)

func TestCommandInspect(t *testing.T) {
	var pikachu domain.Pokemon
	err := json.Unmarshal([]byte(`{
		"name": "pikachu",
		"height": 4,
		"weight": 60,
		"base_experience": 112,
		"stats": [
			{"base_stat": 35, "stat": {"name": "hp"}},
			{"base_stat": 90, "stat": {"name": "speed"}}
		],
		"types": [
			{"slot": 1, "type": {"name": "electric"}}
		]
	}`), &pikachu)
	if err != nil {
		t.Fatalf("failed to prepare pokemon: %v", err)
	}

	config := &domain.Config{
		Pokedex: domain.Pokedex{
			Data: map[string]domain.Pokemon{"pikachu": pikachu},
		},
	}

	output := captureStdout(t, func() {
		err = commandInspect(config, []string{"pikachu"})
	})
	if err != nil {
		t.Fatalf("commandInspect returned an error: %v", err)
	}

	want := "Name: pikachu\n" +
		"Height: 4\n" +
		"Weight: 60\n" +
		"Base Experience: 112\n" +
		"Stats:\n" +
		"  -hp: 35\n" +
		"  -speed: 90\n" +
		"Types:\n" +
		"  - electric\n"
	if output != want {
		t.Fatalf("expected output %q, got %q", want, output)
	}
}

func TestCommandInspectPokemonNotFound(t *testing.T) {
	config := &domain.Config{
		Pokedex: domain.Pokedex{Data: map[string]domain.Pokemon{}},
	}

	var err error
	output := captureStdout(t, func() {
		err = commandInspect(config, []string{"missingno"})
	})
	if err != nil {
		t.Fatalf("commandInspect returned an error: %v", err)
	}

	want := "No pokemon found in your pokedex with name missingno\n"
	if output != want {
		t.Fatalf("expected output %q, got %q", want, output)
	}
}

func TestCommandInspectWithoutArguments(t *testing.T) {
	if err := commandInspect(&domain.Config{}, nil); err != nil {
		t.Fatalf("commandInspect returned an error: %v", err)
	}
}

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	previousStdout := os.Stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create stdout pipe: %v", err)
	}
	os.Stdout = writer

	fn()

	if err := writer.Close(); err != nil {
		t.Fatalf("failed to close stdout writer: %v", err)
	}
	os.Stdout = previousStdout

	output, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("failed to read stdout: %v", err)
	}
	if err := reader.Close(); err != nil {
		t.Fatalf("failed to close stdout reader: %v", err)
	}

	return string(output)
}
