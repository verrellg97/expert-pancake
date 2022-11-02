package db

import (
	"context"
	"github.com/calvinkmts/expert-pancake/engine/sql"
	db "github.com/expert-pancake/service/account/db/sqlc"
	"github.com/expert-pancake/service/account/model"
	"reflect"
)

type UpdateUserTrxParams struct {
	AccountId   string
	FullName    string
	Nickname    string
	Email       string
	PhoneNumber string
	Location    model.Location
}

type UpdateUserTrxResult struct {
	AccountId   string
	FullName    string
	Nickname    string
	Email       string
	PhoneNumber string
	Location    model.Location
}

func (trx *Trx) UpdateUserTrx(ctx context.Context, arg UpdateUserTrxParams) (UpdateUserTrxResult, error) {
	var result UpdateUserTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		err = q.UpdateUser(ctx, db.UpdateUserParams{
			ID:          arg.AccountId,
			Fullname:    arg.FullName,
			Nickname:    arg.Nickname,
			Email:       sql.StringToNullableString(arg.Email),
			PhoneNumber: arg.PhoneNumber,
		})
		if err != nil {
			return err
		}

		if !reflect.ValueOf(arg.Location).IsZero() {
			err = q.UpsertUserAddresses(ctx, db.UpsertUserAddressesParams{
				UserID:      arg.AccountId,
				Country:     "INDONESIA",
				Province:    arg.Location.Province,
				Regency:     arg.Location.Regency,
				District:    arg.Location.District,
				FullAddress: arg.Location.FullAddress,
			})
			if err != nil {
				return err
			}
		}

		result.AccountId = arg.AccountId
		result.FullName = arg.FullName
		result.Nickname = arg.Nickname
		result.Email = arg.Email
		result.PhoneNumber = arg.PhoneNumber
		result.Location = arg.Location

		return err
	})

	return result, err
}
