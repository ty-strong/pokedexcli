package cache

import (
	"bytes"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	c := NewCache(interval)

	key := "testkey"
	val := []byte("testvalue")

	c.Add(key, val)

	foundVal, ok := c.Get(key)
	if !ok {
		t.Fatalf("expected to find key %s, but it was not found", key)
	}
	if !bytes.Equal(foundVal, val) {
		t.Errorf("expected value %s, but got %s", val, foundVal)
	}
}

func TestGetNonExistent(t *testing.T) {
	const interval = 5 * time.Second
	c := NewCache(interval)

	_, ok := c.Get("nonexistentkey")
	if ok {
		t.Error("expected not to find key 'nonexistentkey', but it was found")
	}
}

func TestReapLoop(t *testing.T) {
	const reapInterval = 50 * time.Millisecond
	const waitTime = reapInterval + 50*time.Millisecond
	c := NewCache(reapInterval)

	key := "reapkey"
	c.Add(key, []byte("reapvalue"))

	// Ensure the key is present before reaping
	_, ok := c.Get(key)
	if !ok {
		t.Fatalf("key %s should be present before reaping", key)
	}

	time.Sleep(waitTime)

	_, ok = c.Get(key)
	if ok {
		t.Errorf("key %s should have been reaped, but was found", key)
	}
}
