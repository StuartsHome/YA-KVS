package store

import (
	"testing"
)

func TestNewTransaction(t *testing.T) {
	// Given
	store := NewStore()
	key := "Stuart"
	val := 100
	store.Put(key, val)

	got, err := store.Get(key)
	if err != nil {
		t.Log(err)
	}

	if got != val {
		t.Log("this should be set")
	}

	store.t.CreateNewTransaction()

	// When

	// Then
}
