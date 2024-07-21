package main

import (
	"testing"
)

func TestFIFOCache(t *testing.T) {
	c := NewFIFO(2)

	c.Add("a", 1)
	c.Add("b", 2)

	if v, ok := c.Get("a"); !ok || v != 1 {
		t.Fatalf("expected 1, got %v", v)
	}

	c.Add("c", 3)

	if _, ok := c.Get("a"); ok {
		t.Fatalf("expected 'a' to be evicted")
	}

	if v, ok := c.Get("c"); !ok || v != 3 {
		t.Fatalf("expected 3, got %v", v)
	}
}
