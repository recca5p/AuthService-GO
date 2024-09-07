package db

import (
	"AuthService_GO/utils"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func CreateRole(t *testing.T) Role {
	arg := CreateRoleParams{
		Name:      utils.RandomString(8),
		CreatedAt: time.Now(),
	}

	role, err := testQueries.CreateRole(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, role)
	assert.Equal(t, arg.Name, role.Name)
	assert.WithinDuration(t, arg.CreatedAt, time.Now(), time.Second)

	return role
}

func TestCreateRole(t *testing.T) {
	CreateRole(t)
}

func TestGetRole(t *testing.T) {
	role := CreateRole(t)

	actualRole, err := testQueries.GetRole(context.Background(), role.Id)
	require.NoError(t, err)
	require.NotEmpty(t, actualRole)
	assert.Equal(t, role.Id, actualRole.Id)
	assert.Equal(t, role.Name, actualRole.Name)
}

func TestListRoles(t *testing.T) {
	for i := 0; i < 5; i++ {
		CreateRole(t)
	}

	arg := ListRolesParams{
		Limit:  10,
		Offset: 0,
	}

	roles, err := testQueries.ListRoles(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, roles)
	assert.GreaterOrEqual(t, len(roles), 5)
}

func TestUpdateRole(t *testing.T) {
	role := CreateRole(t)

	arg := UpdateRoleParams{
		Id:   role.Id,
		Name: utils.RandomString(8),
	}

	updatedRole, err := testQueries.UpdateRole(context.Background(), arg)
	require.NoError(t, err)
	assert.Equal(t, arg.Name, updatedRole.Name)
}

func TestDeleteRole(t *testing.T) {
	role := CreateRole(t)

	err := testQueries.DeleteRole(context.Background(), role.Id)
	require.NoError(t, err)

	role, err = testQueries.GetRole(context.Background(), role.Id)
	assert.Error(t, err)
	assert.Empty(t, role)
}
