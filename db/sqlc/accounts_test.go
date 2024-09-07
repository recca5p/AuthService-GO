package db

import (
	"AuthService_GO/utils"
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func CreateAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Username:     utils.RandomUser(),
		PasswordHash: utils.RandomUser(),
		CreatedAt:    time.Now(),
		UpdatedAt:    sql.NullTime{Time: time.Now(), Valid: true},
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	assert.Equal(t, arg.Username, account.Username)
	assert.Equal(t, arg.PasswordHash, account.PasswordHash)
	assert.WithinDuration(t, arg.CreatedAt, account.CreatedAt, time.Second)
	assert.True(t, account.UpdatedAt.Valid)
	assert.WithinDuration(t, arg.UpdatedAt.Time, account.UpdatedAt.Time, time.Second)

	return account
}

func TestCreateAccount(t *testing.T) {
	CreateAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := CreateAccount(t)

	actualAccount, err := testQueries.GetAccount(context.Background(), account.Id)
	require.NoError(t, err)
	require.NotEmpty(t, actualAccount)
	assert.Equal(t, account.Id, actualAccount.Id)
	assert.Equal(t, account.Username, actualAccount.Username)
	assert.Equal(t, account.PasswordHash, actualAccount.PasswordHash)
	if account.UpdatedAt.Valid {
		assert.WithinDuration(t, account.UpdatedAt.Time, actualAccount.UpdatedAt.Time, time.Second)
	} else {
		assert.False(t, actualAccount.UpdatedAt.Valid)
	}
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 5; i++ {
		CreateAccount(t)
	}

	arg := ListAccountsParams{
		Username: "",
		Limit:    10,
		Offset:   0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	assert.GreaterOrEqual(t, len(accounts), 1)

	arg = ListAccountsParams{
		Username: "", // Use empty string to ensure it handles the case correctly
		Limit:    10,
		Offset:   0,
	}

	accounts, err = testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	arg = ListAccountsParams{
		Username: "", // Empty string should still retrieve all accounts
		Limit:    10,
		Offset:   0,
	}

	accounts, err = testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	assert.GreaterOrEqual(t, len(accounts), 5)
}

func TestUpdateAccount(t *testing.T) {
	account := CreateAccount(t)

	arg := UpdateAccountParams{
		Id:           account.Id,
		PasswordHash: "newpasswordhash",
		UpdatedAt:    sql.NullTime{Time: time.Now(), Valid: true},
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	assert.Equal(t, arg.PasswordHash, updatedAccount.PasswordHash)
	assert.NotEqual(t, account.UpdatedAt, updatedAccount.UpdatedAt)
}

func TestDeleteAccount(t *testing.T) {
	account := CreateAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.Id)
	require.NoError(t, err)

	account, err = testQueries.GetAccount(context.Background(), account.Id)
	assert.Error(t, err)
	assert.Empty(t, account)
}
