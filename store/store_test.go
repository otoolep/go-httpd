package store

import (
	"testing"
)

// Test_StoreOpen tests that the store can be opened and closed.
func Test_StoreOpen(t *testing.T) {
	s := New()
	if err := s.Open(); err != nil {
		t.Fatalf("failed to open store: %s", err)
	}
	if err := s.Close(); err != nil {
		t.Fatalf("failed to close store: %s", err)
	}
}

func Test_StoreOpenSingleNode(t *testing.T) {
	s := New()
	if err := s.Open(); err != nil {
		t.Fatalf("failed to open store: %s", err)
	}

	if err := s.Set("foo", "bar"); err != nil {
		t.Fatalf("failed to set key: %s", err.Error())
	}

	value, err := s.Get("foo")
	if err != nil {
		t.Fatalf("failed to get key: %s", err.Error())
	}
	if value != "bar" {
		t.Fatalf("key has wrong value: %s", value)
	}

	if err := s.Delete("foo"); err != nil {
		t.Fatalf("failed to delete key: %s", err.Error())
	}

	value, err = s.Get("foo")
	if err != nil {
		t.Fatalf("failed to get key: %s", err.Error())
	}
	if value != "" {
		t.Fatalf("key has wrong value: %s", value)
	}
}
