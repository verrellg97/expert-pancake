package db

import (
	"context"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	db "github.com/expert-pancake/service/business-relation/db/sqlc"
	"github.com/expert-pancake/service/business-relation/impl/client"
	"github.com/expert-pancake/service/business-relation/model"
	uuid "github.com/satori/go.uuid"
)

type AddContactInvitationTrxParams struct {
	PrimaryContactBookId string
	PrimaryCompanyId     string
	SecondaryCompanyId   string
}

type AddContactInvitationTrxResult struct {
	Message string
}

func (trx *Trx) AddContactInvitationTrx(ctx context.Context, arg AddContactInvitationTrxParams) (AddContactInvitationTrxResult, error) {
	var result AddContactInvitationTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		_, err = q.InsertContactInvitation(context.Background(), db.InsertContactInvitationParams{
			ID:                   uuid.NewV4().String(),
			PrimaryContactBookID: arg.PrimaryContactBookId,
			PrimaryCompanyID:     arg.PrimaryCompanyId,
			SecondaryCompanyID:   arg.SecondaryCompanyId,
		})
		if err != nil {
			return errors.NewServerError(model.AddNewContactInvitationError, err.Error())
		}

		branches, err := client.GetCompanyBranches(
			client.GetCompanyBranchesRequest{
				CompanyId: arg.SecondaryCompanyId,
			})
		if err != nil {
			return err
		}

		contactBook, err := q.GetMyContactBook(context.Background(), arg.PrimaryCompanyId)
		if err != nil {
			return errors.NewServerError(model.AddNewContactInvitationError, err.Error())
		}

		_, err = client.InsertNotification(
			client.InsertNotificationRequest{
				CompanyId: arg.SecondaryCompanyId,
				BranchId:  branches.Result[0].BranchId,
				Title:     "INVITATION REQUEST",
				Content:   contactBook.Name + " wants to connect with you",
				Type:      "INVITATION",
			})
		if err != nil {
			return err
		}

		result.Message = "OK"

		return err
	})

	return result, err
}
