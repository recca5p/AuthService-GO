// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateAccountRole(ctx context.Context, arg CreateAccountRoleParams) (AccountsRole, error)
	CreatePermission(ctx context.Context, arg CreatePermissionParams) (Permission, error)
	CreateRole(ctx context.Context, arg CreateRoleParams) (Role, error)
	CreateRolePermission(ctx context.Context, arg CreateRolePermissionParams) (RolesPermission, error)
	DeleteAccount(ctx context.Context, id uuid.UUID) error
	DeleteAccountRole(ctx context.Context, id uuid.UUID) error
	DeletePermission(ctx context.Context, id uuid.UUID) error
	DeleteRole(ctx context.Context, id uuid.UUID) error
	DeleteRolePermission(ctx context.Context, id uuid.UUID) error
	GetAccount(ctx context.Context, id uuid.UUID) (Account, error)
	GetAccountRole(ctx context.Context, id uuid.UUID) (AccountsRole, error)
	GetAccountRoleForUpdate(ctx context.Context, id uuid.UUID) (AccountsRole, error)
	GetPermission(ctx context.Context, id uuid.UUID) (Permission, error)
	GetPermissionForUpdate(ctx context.Context, id uuid.UUID) (Permission, error)
	GetRole(ctx context.Context, id uuid.UUID) (Role, error)
	GetRoleForUpdate(ctx context.Context, id uuid.UUID) (Role, error)
	GetRolePermission(ctx context.Context, id uuid.UUID) (RolesPermission, error)
	GetRolePermissionForUpdate(ctx context.Context, id uuid.UUID) (RolesPermission, error)
	ListAccountRoles(ctx context.Context, arg ListAccountRolesParams) ([]AccountsRole, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListPermissions(ctx context.Context, arg ListPermissionsParams) ([]Permission, error)
	ListRolePermissions(ctx context.Context, arg ListRolePermissionsParams) ([]RolesPermission, error)
	ListRoles(ctx context.Context, arg ListRolesParams) ([]Role, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	UpdatePermission(ctx context.Context, arg UpdatePermissionParams) (Permission, error)
	UpdateRole(ctx context.Context, arg UpdateRoleParams) (Role, error)
}

var _ Querier = (*Queries)(nil)
