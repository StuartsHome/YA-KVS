package handler

import (
	"testing"

	"github.com/StuartsHome/YA-KVS/store"
)

func TestHandleGet(t *testing.T) {
	// Given
	store := store.NewStore()

	// When
	key := "Stuart"
	val := 100
	err := store.Put(key, val)
	if err != nil {
		t.Error(err)
	}

	// Then
	got, err := store.Get(key)
	if err != nil {
		t.Error(err)
	}

	if got != val {
		t.Error(err)
	}
}

func BenchmarkHandleGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Given
		store := store.NewStore()

		// When
		key := "Stuart"
		val := 100
		err := store.Put(key, val)
		if err != nil {
			b.Error(err)
		}

		// Then
		got, err := store.Get(key)
		if err != nil {
			b.Error(err)
		}

		if got != val {
			b.Error(err)
		}
	}
}
