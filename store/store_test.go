package store

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type StoreTestSuite struct {
	suite.Suite

	transactionStack *transactionStack
}

func (ts *StoreTestSuite) SetupTest() {
	ts.transactionStack = NewTransaction()
}

func TestStoreTestSuite(t *testing.T) {
	suite.Run(t, new(StoreTestSuite))
}

func (ts *StoreTestSuite) TestGet() {
	// Given
	store := NewStore()

	// When
	key := "Stuart"
	val := 100
	err := store.Set(key, val, ts.transactionStack)
	ts.Require().NoError(err)

	// Then
	got, err := store.Get(key, ts.transactionStack)
	ts.Require().NoError(err)

	expected := 100
	ts.Assert().Equal(expected, got)
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Given
		store := NewStore()
		transaction := NewTransaction()

		// When
		key := "Stuart"
		val := 100
		_ = store.Set(key, val, transaction)

		// Then
		got, _ := store.Get(key, transaction)

		expected := 100
		if got != expected {
			b.Errorf("expected %d, got %d", expected, got)
		}
	}
}

func (ts *StoreTestSuite) TestDelete_Success() {
	// Given
	store := NewStore()
	key := "Stuart"
	val := 100
	err := store.Set(key, val, ts.transactionStack)
	ts.Require().NoError(err)

	// When
	// When we delete key from current transaction.
	err = store.Delete(key, ts.transactionStack)
	ts.Require().NoError(err)

	// Then
	// Then key should no longer be in the current transaction stack.
	got, err := store.Get(key, ts.transactionStack)
	ts.Assert().Equal(0, got)
	ts.Assert().EqualError(err, "key Stuart not set in global store")
}

func (ts *StoreTestSuite) TestDelete_Error() {
	// Given
	store := NewStore()
	key := "Stuart"

	// When
	err := store.Delete(key, ts.transactionStack)
	ts.Assert().EqualError(err, "unable to delete key Stuart as not currently in global store")
}
