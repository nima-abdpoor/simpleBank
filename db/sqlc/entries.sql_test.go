package db

import (
	"context"
	"database/sql"
	"excercise/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomEntry(t *testing.T) Entries {
	account := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomInt(-2000, +2000),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, account.ID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.NotEmpty(t, entry.ID)
	require.NotEmpty(t, entry.CreatedAt)
	return entry
}

func TestQueries_CreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestQueries_DeleteEntry(t *testing.T) {
	entry := createRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	entry2, err2 := testQueries.GetEntry(context.Background(), entry.ID)
	require.Error(t, err2, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}

func TestQueries_GetEntry(t *testing.T) {
	entry := createRandomEntry(t)
	actualEntry, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, actualEntry)
	require.Equal(t, entry.AccountID, actualEntry.AccountID)
	require.Equal(t, entry.ID, actualEntry.ID)
	require.Equal(t, entry.CreatedAt, actualEntry.CreatedAt)
	require.Equal(t, entry.Amount, actualEntry.Amount)
}

func TestQueries_UpdateEntry(t *testing.T) {
	entry := createRandomEntry(t)
	arg := UpdateEntryParams{
		ID:        entry.ID,
		AccountID: entry.AccountID,
		Amount:    util.RandomInt(-2000, +2000),
	}
	actualEntry, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, actualEntry)
	require.Equal(t, entry.AccountID, actualEntry.AccountID)
	require.Equal(t, entry.ID, actualEntry.ID)
	require.Equal(t, entry.CreatedAt, actualEntry.CreatedAt)
	require.Equal(t, arg.Amount, actualEntry.Amount)
	require.Equal(t, arg.AccountID, actualEntry.AccountID)
	require.NotEqual(t, entry.Amount, actualEntry.Amount)
}

func TestQueries_GetAllEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}
	entry := createRandomEntry(t)
	arg := ListEntriesParams{
		AccountID: entry.AccountID,
		Limit:     5,
		Offset:    5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)
	for _, entr := range entries {
		require.NotEmpty(t, entr)
	}
}
