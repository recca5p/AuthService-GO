package db

import (
	"AuthService_GO/utils"
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreatePermission(t *testing.T) Permission {
	arg := CreatePermissionParams{
		Description: sql.NullString{String: utils.RandomString(32), Valid: true},
		Name:        utils.RandomString(16),
	}

	permission, err := testQueries.CreatePermission(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, permission)
	assert.Equal(t, arg.Description, permission.Description)

	return permission
}

func TestCreatePermission(t *testing.T) {
	CreatePermission(t)
}

func TestGetPermission(t *testing.T) {
	permission := CreatePermission(t)

	actualPermission, err := testQueries.GetPermission(context.Background(), permission.Id)
	require.NoError(t, err)
	require.NotEmpty(t, actualPermission)
	assert.Equal(t, permission.Id, actualPermission.Id)
	assert.Equal(t, permission.Description, actualPermission.Description)
}

func TestListPermissions(t *testing.T) {
	for i := 0; i < 5; i++ {
		CreatePermission(t)
	}

	arg := ListPermissionsParams{
		Limit:  10,
		Offset: 0,
	}

	permissions, err := testQueries.ListPermissions(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, permissions)
	assert.GreaterOrEqual(t, len(permissions), 5)
}

func TestUpdatePermission(t *testing.T) {
	permission := CreatePermission(t)

	arg := UpdatePermissionParams{
		Id:          permission.Id,
		Description: sql.NullString{String: utils.RandomString(32), Valid: true},
	}

	updatedPermission, err := testQueries.UpdatePermission(context.Background(), arg)
	require.NoError(t, err)
	assert.Equal(t, arg.Description, updatedPermission.Description)
}

func TestDeletePermission(t *testing.T) {
	permission := CreatePermission(t)

	err := testQueries.DeletePermission(context.Background(), permission.Id)
	require.NoError(t, err)

	permission, err = testQueries.GetPermission(context.Background(), permission.Id)
	assert.Error(t, err)
	assert.Empty(t, permission)
}
