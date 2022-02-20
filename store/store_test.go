package store

import "testing"

func TestGet(t *testing.T) {
	// Given
	store := NewStore()

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

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Given
		store := NewStore()

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

func TestDelete_Success(t *testing.T) {
	// Given
	store := NewStore()
	key := "Stuart"
	val := 100
	err := store.Put(key, val)
	if err != nil {
		t.Error(err)
	}

	// When
	err = store.Delete(key)
	if err != nil {
		t.Error(err)
	}

	// Then
	_, err = store.Get(key)
	if err == nil {
		t.Error("this throw an error")
	}
}

func TestDelete_Error(t *testing.T) {
	// Given
	store := NewStore()
	key := "Stuart"

	// When
	err := store.Delete(key)
	if err == nil {
		t.Error("this throw an error")
	}
}
