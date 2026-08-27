package pokeapi_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"codeberg.org/OliveiraJ/pokedexcli/internal/pokeapi"
	"codeberg.org/OliveiraJ/pokedexcli/internal/pokecache"
)

func TestClientCachesResponsesByURL(t *testing.T) {
	var requestCount atomic.Int32
	want := []byte(`{"results":[{"name":"canalave-city-area"}]}`)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requestCount.Add(1)
		_, _ = w.Write(want)
	}))
	defer server.Close()

	cache := pokecache.NewCache(time.Minute)
	client := pokeapi.NewClient(cache)

	for i := 0; i < 2; i++ {
		got, err := client.Get(server.URL)
		if err != nil {
			t.Fatalf("request %d failed: %v", i+1, err)
		}
		if !bytes.Equal(got, want) {
			t.Fatalf("request %d: expected %q, got %q", i+1, want, got)
		}
	}

	if got := requestCount.Load(); got != 1 {
		t.Fatalf("expected 1 HTTP request, got %d", got)
	}
}
