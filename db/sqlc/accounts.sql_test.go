package db

import (
	"context"
	"database/sql"
	"excercise/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomAccount(t *testing.T) Accounts {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(10),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestQueries_CreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestQueries_GetAccount(t *testing.T) {
	account := createRandomAccount(t)
	actualAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, actualAccount)
	require.Equal(t, account.ID, actualAccount.ID)
	require.Equal(t, account.Owner, actualAccount.Owner)
	require.Equal(t, account.Balance, actualAccount.Balance)
	require.Equal(t, account.Currency, actualAccount.Currency)
}

func TestQueries_UpdateAccount(t *testing.T) {
	account := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: account.Balance,
	}

	actualAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, actualAccount)
	require.Equal(t, account.ID, actualAccount.ID)
	require.Equal(t, arg.Balance, actualAccount.Balance)
	require.Equal(t, account.Currency, actualAccount.Currency)
	require.Equal(t, account.Owner, actualAccount.Owner)
}

func TestQueries_DeleteAccount(t *testing.T) {
	account := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	actualAccount, err2 := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err2)
	require.EqualError(t, err2, sql.ErrNoRows.Error())
	require.Empty(t, actualAccount)
}

func TestQueries_ListAccounts(t *testing.T) {

	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
