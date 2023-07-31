package db

import (
	"context"

	"github.com/calvinkmts/expert-pancake/engine/sql"
	db "github.com/expert-pancake/service/account/db/sqlc"
	"github.com/expert-pancake/service/account/model"
)

type UpdateUserTrxParams struct {
	AccountId   string
	ImageUrl    string
	FullName    string
	Nickname    string
	Email       string
	PhoneNumber string
	Location    model.Location
}

type UpdateUserTrxResult struct {
	AccountId   string
	ImageUrl    string
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

		userRes, err := q.UpsertUser(ctx, db.UpsertUserParams{
			ID:          arg.AccountId,
			ImageUrl:    arg.ImageUrl,
			Fullname:    arg.FullName,
			Nickname:    arg.Nickname,
			Email:       sql.StringToNullableString(arg.Email),
			PhoneNumber: arg.PhoneNumber,
		})
		if err != nil {
			return err
		}

		userAddressRes, err := q.UpsertUserAddresses(ctx, db.UpsertUserAddressesParams{
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

		result.AccountId = userRes.ID
		result.FullName = userRes.Fullname
		result.Nickname = userRes.Nickname
		result.Email = userRes.Email.String
		result.PhoneNumber = userRes.PhoneNumber
		result.Location = model.Location{
			Province:    userAddressRes.Province,
			Regency:     userAddressRes.Regency,
			District:    userAddressRes.District,
			FullAddress: userAddressRes.FullAddress,
		}

		return err
	})

	return result, err
}
