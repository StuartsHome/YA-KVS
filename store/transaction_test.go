package store

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TransactionTestSuite struct {
	suite.Suite

	key              string
	val              int
	transactionStack *transactionStack
}

func (ts *TransactionTestSuite) SetupTest() {
	ts.key = "Stuart"
	ts.val = 100
	ts.transactionStack = NewTransaction()
}

func TestTransactionTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionTestSuite))
}

func (ts *TransactionTestSuite) TestNewTransaction_EmptyUponCreation() {
	// Given
	store := NewStore()
	err := store.Set(ts.key, ts.val, ts.transactionStack)
	ts.Require().NoError(err)

	got, err := store.Get(ts.key, ts.transactionStack)
	ts.Require().NoError(err)
	ts.Assert().Equal(ts.val, got)

	// When
	ts.transactionStack.PushTransaction()
	item := ts.transactionStack.Peek()
	// Then
	ts.Assert().Empty(item.store)
}

func (ts *TransactionTestSuite) TestNewTransaction_OneTransaction_Success() {
	// Given
	store := NewStore()
	err := store.Set(ts.key, ts.val, ts.transactionStack)
	ts.Require().NoError(err)

	got, err := store.Get(ts.key, ts.transactionStack)
	ts.Require().NoError(err)
	ts.Assert().Equal(ts.val, got)

	// When
	ts.transactionStack.PushTransaction()
	item := ts.transactionStack.Peek()
	// Then
	ts.Assert().Empty(item.store)

	// Store a new val in transaction at top of stack.
	newVal := 200
	err = store.Set(ts.key, newVal, ts.transactionStack)
	ts.Require().NoError(err)

	got, err = store.Get(ts.key, ts.transactionStack)
	ts.Require().NoError(err)
	ts.Assert().Equal(newVal, got)

	// Remove transaction.
	ts.transactionStack.PopTransaction()
	got, err = store.Get(ts.key, ts.transactionStack)
	ts.Require().NoError(err)
	ts.Assert().Equal(ts.val, got)

	// No more transactions.
	ts.Assert().Equal(0, ts.transactionStack.size)
	ts.Assert().Nil(ts.transactionStack.top)

}
