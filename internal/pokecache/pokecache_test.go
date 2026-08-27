package pokecache

import (
	"bytes"
	"testing"
	"time"
)

func TestCacheAddAndGet(t *testing.T) {
	c := NewCache(time.Minute)
	want := []byte("value")
	c.Add("key", want)

	got, ok := c.Get("key")
	if !ok {
		t.Fatal("expected cache entry to exist")
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("expected %q, got %q", want, got)
	}
}

func TestCacheReapLoop(t *testing.T) {
	const interval = 10 * time.Millisecond
	c := NewCache(interval)
	c.Add("key", []byte("value"))

	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		if _, ok := c.Get("key"); !ok {
			return
		}
		time.Sleep(interval)
	}

	t.Fatal("expected expired cache entry to be removed")
}
