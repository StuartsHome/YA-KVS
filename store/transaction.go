package store

import "fmt"

type Transaction interface {
	CreateNewTransaction() error
}

type transaction struct {
	// Each transaction has its own local store.
	store map[string]int // TODO: should be an instance of DB, not map.
	next  *transaction
}

// Maintains a list of active/suspended transactions.
type transactionStack struct {
	top  *transaction
	size int
}

// Push a new transaction.
func (ts *transactionStack) PushTransaction() {
	// Push a new transaction, this is the current active transaction.
	temp := transaction{store: make(map[string]int)}
	temp.next = ts.top
	ts.top = &temp
	ts.size++
}

// Pop.
func (ts *transactionStack) PopTransaction() {
	if ts.top == nil {
		fmt.Printf("error: no active transactions.\n")
	} else {
		node := &transaction{}
		ts.top = ts.top.next
		node.next = nil
		ts.size--
	}
}

// Peek.
func (ts *transactionStack) Peek() *transaction {
	return ts.top
}

// Commit.
func (ts *transactionStack) Commit() {
	// Fetch active transaction from stack.
	activeTransaction := ts.Peek()
	if activeTransaction != nil {
		for key, val := range activeTransaction.store {
			// INSERT: db store.
			if activeTransaction.next != nil {
				// Update the parent transaction.
				activeTransaction.next.store[key] = val
			}
		}
	} else {
		fmt.Printf("nothing data to commit.")
	}
}

// RollbackTransaction clears all keys SET within a transaction.
func (ts *transactionStack) RollbackTransaction() {
	if ts.top == nil {
		fmt.Printf("error: no active transaction/\n")
	} else {
		for key := range ts.top.store {
			delete(ts.top.store, key)
		}
		ts.top = ts.top.next
	}
}

func NewTransaction() *transactionStack {
	return &transactionStack{}
}

func (t *transaction) CreateNewTransaction() error {
	return nil
}
