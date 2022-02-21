package store

import (
	"fmt"
	"sync"

	"github.com/StuartsHome/YA-KVS/model"
)

type store struct {
	db map[string]int
	t  Transaction

	// Single writer.
	writer sync.Mutex
}

func NewStore() *store {
	st := &store{
		t: NewTransaction(),
	}

	if err := st.startUp(); err != nil {
		return nil
	}

	return st
}

func (st *store) startUp() error {
	st.db = make(map[string]int)
	return nil
}

func (st *store) Get(key string) (int, error) {
	// Fetch db.
	db := st.db
	var err model.DBError

	// Check if key in db.
	if val, ok := db[key]; ok {
		return val, nil
	}
	return 0, err
}

func (st *store) Put(key string, val int) error {
	// Fetch db.
	db := st.db

	// Check if key in db.
	if _, ok := db[key]; !ok {
		// If not in db, store it!
		db[key] = val
	}

	// If key in db overwrite.
	// TODO: ok for now.
	db[key] = val

	// Save.
	st.db = db
	return nil
}

func (st *store) Delete(key string) error {
	// Fetch db.
	db := st.db
	var err model.DBError

	// Check if key in db.
	if _, ok := db[key]; !ok {
		// If not in db, return an error.
		err.Message = fmt.Sprintf("key %s not found in db", key)
		return err
	}

	delete(db, key)
	return nil
}
