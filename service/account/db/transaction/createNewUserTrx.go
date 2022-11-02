package db

import (
	"context"
	"github.com/calvinkmts/expert-pancake/engine/sql"
	db "github.com/expert-pancake/service/account/db/sqlc"
	"github.com/expert-pancake/service/account/util"
	uuid "github.com/satori/go.uuid"
)

type CreateNewUserTrxParams struct {
	FullName         string
	Nickname         string
	Email            string
	PhoneNumber      string
	Password         string
	SecurityQuestion string
	SecurityAnswer   string
}

type CreateNewUserTrxResult struct {
	Id          string
	FullName    string
	Nickname    string
	Email       string
	PhoneNumber string
}

func (trx *Trx) CreateNewUserTrx(ctx context.Context, arg CreateNewUserTrxParams) (CreateNewUserTrxResult, error) {
	var result CreateNewUserTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		id := uuid.NewV4().String()

		err = q.CreateUser(ctx, db.CreateUserParams{
			ID:          id,
			Fullname:    arg.FullName,
			Nickname:    arg.Nickname,
			Email:       sql.StringToNullableString(arg.Email),
			PhoneNumber: arg.PhoneNumber,
		})
		if err != nil {
			return err
		}

		hashedPassword, _ := util.HashPassword(arg.Password)

		err = q.UpsertUserPassword(ctx, db.UpsertUserPasswordParams{
			UserID:   id,
			Password: hashedPassword,
		})
		if err != nil {
			return err
		}

		err = q.UpsertUserInfo(ctx, db.UpsertUserInfoParams{
			UserID: id,
			Key:    "security_questions",
			Value:  arg.SecurityQuestion,
		})
		if err != nil {
			return err
		}

		err = q.UpsertUserInfo(ctx, db.UpsertUserInfoParams{
			UserID: id,
			Key:    "security_answers",
			Value:  arg.SecurityAnswer,
		})
		if err != nil {
			return err
		}

		result.Id = id
		result.FullName = arg.FullName
		result.Nickname = arg.Nickname
		result.Email = arg.Email
		result.PhoneNumber = arg.PhoneNumber

		return err
	})

	return result, err
}
