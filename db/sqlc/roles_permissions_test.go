package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
	"time"
)

// Helper function to create random UUIDs for testing
func CreateRandomUUID(t *testing.T) uuid.UUID {
	id, err := uuid.NewRandom()
	require.NoError(t, err)
	return id
}

// Helper function to create role permissions
func CreateRolePermission(t *testing.T) RolesPermission {
	role := CreateRole(t)
	permission := CreatePermission(t)

	arg := CreateRolePermissionParams{
		RoleId:       role.Id,
		PermissionId: permission.Id,
		CreatedAt:    time.Now(),
	}

	rolePermission, err := testQueries.CreateRolePermission(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, rolePermission)
	assert.Equal(t, arg.RoleId, rolePermission.RoleId)
	assert.Equal(t, arg.PermissionId, rolePermission.PermissionId)

	return rolePermission
}

func TestCreateRolePermission(t *testing.T) {
	CreateRolePermission(t)
}

func TestGetRolePermission(t *testing.T) {
	rolePermission := CreateRolePermission(t)

	actualRolePermission, err := testQueries.GetRolePermission(context.Background(), rolePermission.Id)
	require.NoError(t, err)
	require.NotEmpty(t, actualRolePermission)
	assert.Equal(t, rolePermission.Id, actualRolePermission.Id)
	assert.Equal(t, rolePermission.RoleId, actualRolePermission.RoleId)
	assert.Equal(t, rolePermission.PermissionId, actualRolePermission.PermissionId)
}

func TestListRolePermissions(t *testing.T) {
	rolePermissions := []RolesPermission{}
	for i := 0; i < 10; i++ {
		rolePermission := CreateRolePermission(t)
		rolePermissions = append(rolePermissions, rolePermission)
	}

	arg := ListRolePermissionsParams{
		RoleId: rolePermissions[0].RoleId,
		Limit:  10,
		Offset: 0,
	}

	actualRolePermissions, err := testQueries.ListRolePermissions(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, actualRolePermissions)

	rolePermissionFound := false
	for _, rp := range actualRolePermissions {
		if rp.RoleId == arg.RoleId {
			rolePermissionFound = true
			break
		}
	}
	assert.True(t, rolePermissionFound, "Expected to find a role permission with RoleId %s in the result", arg.RoleId)
}

func TestDeleteRolePermission(t *testing.T) {
	rolePermission := CreateRolePermission(t)

	err := testQueries.DeleteRolePermission(context.Background(), rolePermission.Id)
	require.NoError(t, err)

	rolePermission, err = testQueries.GetRolePermission(context.Background(), rolePermission.Id)
	assert.Error(t, err)
	assert.Empty(t, rolePermission)
}

func TestRolePermissionDeadlock(t *testing.T) {
	// Create a role permission
	rolePermission := CreateRolePermission(t)

	// Run multiple goroutines to test for deadlock
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := testQueries.GetRolePermissionForUpdate(context.Background(), rolePermission.Id)
			require.NoError(t, err)
		}()
	}

	wg.Wait()
}
