package store

import (
	"fmt"
	"sync"
)

var GlobalStore = make(map[string]int)

type store struct {
	db map[string]int
	t  *transactionStack

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

func (st *store) Get(key string, t *transactionStack) (int, error) {
	// Get
	activeTransaction := t.Peek()
	if activeTransaction == nil {
		if val, ok := GlobalStore[key]; ok {
			return val, nil
		}
		return 0, fmt.Errorf("key %s not set in global store", key)
	} else {
		if val, ok := activeTransaction.store[key]; ok {
			return val, nil
		}
		return 0, fmt.Errorf("key %s not set in active transaction store", key)
	}
}

func (st *store) Set(key string, val int, t *transactionStack) error {
	//
	activeTransaction := t.Peek()
	if activeTransaction == nil {
		GlobalStore[key] = val
	} else {
		activeTransaction.store[key] = val
	}

	return nil
}

func (st *store) Delete(key string, t *transactionStack) error {
	activeTransaction := t.Peek()
	if activeTransaction == nil {
		if _, ok := GlobalStore[key]; ok {
			delete(GlobalStore, key)
		} else {
			return fmt.Errorf("unable to delete key %s as not currently in global store", key)
		}
	} else {
		if _, ok := activeTransaction.store[key]; ok {
			delete(activeTransaction.store, key)
		} else {
			return fmt.Errorf("unable to delete key %s as not currently in transaction store", key)
		}
	}
	return nil
}
