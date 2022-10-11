package db

import (
	"context"
	"database/sql"
	"excercise/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomTransfer(t *testing.T, account1, account2 Accounts) Transfers {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomInt(0, 2000),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, account1.ID, transfer.FromAccountID)
	require.Equal(t, account2.ID, transfer.ToAccountID)
	require.NotEqual(t, account2.ID, account1.ID)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.NotZero(t, transfer.ID)
	require.NotEmpty(t, transfer.CreatedAt)
	return transfer
}

func TestQueries_CreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createRandomTransfer(t, account1, account2)
}

func TestQueries_GetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	entry := createRandomTransfer(t, account1, account2)
	actualEntry, err := testQueries.GetTransfer(context.Background(), entry.ID)
	require.NotEmpty(t, actualEntry)
	require.NoError(t, err)
	require.Equal(t, entry.ID, actualEntry.ID)
	require.Equal(t, entry.ToAccountID, actualEntry.ToAccountID)
	require.Equal(t, entry.FromAccountID, actualEntry.FromAccountID)
	require.Equal(t, entry.Amount, actualEntry.Amount)
	require.Equal(t, entry.CreatedAt, actualEntry.CreatedAt)
}

func TestQueries_DeleteTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transfer := createRandomTransfer(t, account1, account2)
	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	transfer2, err2 := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.Error(t, err2, sql.ErrNoRows.Error())
	require.Empty(t, transfer2)
}

func TestQueries_UpdateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transfer := createRandomTransfer(t, account1, account2)
	arg := UpdateTransferParams{
		ID:     transfer.ID,
		Amount: util.RandomInt(0, 2000),
	}
	actualTransfer, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, actualTransfer)
	require.Equal(t, transfer.ID, arg.ID)
	require.Equal(t, transfer.ID, actualTransfer.ID)
	require.Equal(t, transfer.CreatedAt, actualTransfer.CreatedAt)
	require.Equal(t, transfer.ToAccountID, actualTransfer.ToAccountID)
	require.Equal(t, transfer.FromAccountID, actualTransfer.FromAccountID)
	require.NotEqual(t, transfer.Amount, actualTransfer.Amount)
}

func TestQueries_ListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomTransfer(t, account1, account2)
	}
	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         5,
		Offset:        5,
	}
	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)
	require.Len(t, transfers, 5)
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
