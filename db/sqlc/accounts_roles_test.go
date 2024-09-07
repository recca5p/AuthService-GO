package db

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
	"time"
)

// Helper function to create account roles
func CreateAccountRole(t *testing.T) AccountsRole {
	account := CreateAccount(t)
	role := CreateRole(t)

	arg := CreateAccountRoleParams{
		AccountId: account.Id,
		RoleId:    role.Id,
		CreatedAt: time.Now(),
	}

	accountRole, err := testQueries.CreateAccountRole(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accountRole)
	assert.Equal(t, arg.AccountId, accountRole.AccountId)
	assert.Equal(t, arg.RoleId, accountRole.RoleId)

	return accountRole
}

func TestCreateAccountRole(t *testing.T) {
	CreateAccountRole(t)
}

func TestGetAccountRole(t *testing.T) {
	accountRole := CreateAccountRole(t)

	actualAccountRole, err := testQueries.GetAccountRole(context.Background(), accountRole.Id)
	require.NoError(t, err)
	require.NotEmpty(t, actualAccountRole)
	assert.Equal(t, accountRole.Id, actualAccountRole.Id)
	assert.Equal(t, accountRole.AccountId, actualAccountRole.AccountId)
	assert.Equal(t, accountRole.RoleId, actualAccountRole.RoleId)
}

func TestListAccountRoles(t *testing.T) {
	accountRoles := []AccountsRole{}
	for i := 0; i < 10; i++ {
		accountRole := CreateAccountRole(t)
		accountRoles = append(accountRoles, accountRole)
	}

	arg := ListAccountRolesParams{
		AccountId: accountRoles[0].AccountId,
		Limit:     10,
		Offset:    0,
	}

	actualAccountRoles, err := testQueries.ListAccountRoles(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, actualAccountRoles)

	accountRoleFound := false
	for _, ar := range actualAccountRoles {
		if ar.AccountId == arg.AccountId {
			accountRoleFound = true
			break
		}
	}
	assert.True(t, accountRoleFound, "Expected to find an account role with AccountId %s in the result", arg.AccountId)
}

func TestDeleteAccountRole(t *testing.T) {
	accountRole := CreateAccountRole(t)

	err := testQueries.DeleteAccountRole(context.Background(), accountRole.Id)
	require.NoError(t, err)

	accountRole, err = testQueries.GetAccountRole(context.Background(), accountRole.Id)
	assert.Error(t, err)
	assert.Empty(t, accountRole)
}

func TestAccountRoleDeadlock(t *testing.T) {
	// Create an account role
	accountRole := CreateAccountRole(t)

	// Run multiple goroutines to test for deadlock
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := testQueries.GetAccountRoleForUpdate(context.Background(), accountRole.Id)
			require.NoError(t, err)
		}()
	}

	wg.Wait()
}
