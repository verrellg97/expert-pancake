// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) error
	GetUser(ctx context.Context, id string) (AccountUser, error)
	GetUserAddress(ctx context.Context, userID string) (AccountUserAddress, error)
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (string, error)
	GetUserInfo(ctx context.Context, arg GetUserInfoParams) (AccountUserInfo, error)
	GetUserPassword(ctx context.Context, userID string) (AccountUserPassword, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (AccountUser, error)
	UpsertUser(ctx context.Context, arg UpsertUserParams) (AccountUser, error)
	UpsertUserAddresses(ctx context.Context, arg UpsertUserAddressesParams) (AccountUserAddress, error)
	UpsertUserInfo(ctx context.Context, arg UpsertUserInfoParams) error
	UpsertUserPassword(ctx context.Context, arg UpsertUserPasswordParams) error
}

var _ Querier = (*Queries)(nil)
