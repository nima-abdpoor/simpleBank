package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStore_TransferTx(t *testing.T) {
	store := newStore(testDB)
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	amount := int64(10)
	arg := TransferTxParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        int64(amount),
	}
	errs := make(chan error)
	results := make(chan TransferTxResult)
	for i := 0; i < 10; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), arg)
			errs <- err
			results <- result
		}()
	}

	n := 5
	existed := make(map[int]bool)

	for i := 0; i < n; i++ {
		err := <-errs
		result := <-results
		require.NoError(t, err)

		require.NotEmpty(t, result)

		//check the transfer
		require.Equal(t, arg.Amount, result.Transfer.Amount)
		require.Equal(t, account1.ID, result.Transfer.FromAccountID)
		require.Equal(t, account2.ID, result.Transfer.ToAccountID)
		require.NotZero(t, result.Transfer.ID)
		require.NotZero(t, result.Transfer.CreatedAt)

		actualTransfer, err2 := store.q.GetTransfer(context.Background(), result.Transfer.ID)
		require.NoError(t, err2)
		require.NotEmpty(t, actualTransfer)
		require.Equal(t, account1.ID, actualTransfer.FromAccountID)
		require.Equal(t, account2.ID, actualTransfer.ToAccountID)

		//check the entry
		require.NotEmpty(t, result.FromEntry)
		require.NotEmpty(t, result.ToEntry)
		require.Equal(t, account1.ID, result.FromEntry.AccountID)
		require.Equal(t, account2.ID, result.ToEntry.AccountID)
		require.Equal(t, account2.ID, result.ToEntry.AccountID)
		require.Equal(t, -arg.Amount, result.FromEntry.Amount)
		require.Equal(t, arg.Amount, result.ToEntry.Amount)
		require.NotZero(t, result.ToEntry.CreatedAt)
		require.NotZero(t, result.FromEntry.CreatedAt)

		actualEntry1, err3 := store.q.GetEntry(context.Background(), result.FromEntry.ID)
		require.NoError(t, err3)
		require.Equal(t, result.FromEntry.ID, actualEntry1.ID)

		actualEntry2, err4 := store.q.GetEntry(context.Background(), result.ToEntry.ID)
		require.NoError(t, err4)
		require.Equal(t, result.ToEntry.ID, actualEntry2.ID)

		//TODO: check account balance
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, account1.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, account2.ID, toAccount.ID)

		//check Account's balance
		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0)

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	//check the final Balances
	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.Equal(t, account1.Balance-int64(n)*amount, updatedAccount1.Balance)

	updatedAccount2, err2 := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err2)
	require.Equal(t, account2.Balance+int64(n)*amount, updatedAccount2.Balance)

}
