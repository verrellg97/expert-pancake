package db

import (
	"context"

	"github.com/calvinkmts/expert-pancake/engine/sql"
	db "github.com/expert-pancake/service/account/db/sqlc"
	"github.com/expert-pancake/service/account/model"
	"github.com/expert-pancake/service/account/util"
	uuid "github.com/satori/go.uuid"
)

type CreateNewUserTrxParams struct {
	ImageUrl         string
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
	ImageUrl    string
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

		userRes, err := q.UpsertUser(ctx, db.UpsertUserParams{
			ID:          id,
			ImageUrl:    arg.ImageUrl,
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
			Key:    model.UserSecurityQuestionKey,
			Value:  arg.SecurityQuestion,
		})
		if err != nil {
			return err
		}

		err = q.UpsertUserInfo(ctx, db.UpsertUserInfoParams{
			UserID: id,
			Key:    model.UserSecurityAnswerKey,
			Value:  arg.SecurityAnswer,
		})
		if err != nil {
			return err
		}

		result.Id = id
		result.ImageUrl = userRes.ImageUrl
		result.FullName = userRes.Fullname
		result.Nickname = userRes.Nickname
		result.Email = userRes.Email.String
		result.PhoneNumber = userRes.PhoneNumber

		return err
	})

	return result, err
}
