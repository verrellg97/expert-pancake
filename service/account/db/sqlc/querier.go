// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) error
	UpsertUserInfo(ctx context.Context, arg UpsertUserInfoParams) error
	UpsertUserPassword(ctx context.Context, arg UpsertUserPasswordParams) error
}

var _ Querier = (*Queries)(nil)
