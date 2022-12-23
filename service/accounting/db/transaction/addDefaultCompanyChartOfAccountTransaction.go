package db

import (
	"context"

	db "github.com/expert-pancake/service/accounting/db/sqlc"
	"github.com/expert-pancake/service/accounting/model"
	uuid "github.com/satori/go.uuid"
)

var defaultCompanyChartOfAccountGroups = []model.ChartOfAccountGroup{
	{ReportType: "NERACA", AccountType: "CASH & EQUALS", AccountGroupName: "KAS"},
	{ReportType: "NERACA", AccountType: "CASH & EQUALS", AccountGroupName: "BANK"},
	{ReportType: "NERACA", AccountType: "RECEIVABLE", AccountGroupName: "PIUTANG USAHA"},
	{ReportType: "NERACA", AccountType: "INVENTORY", AccountGroupName: "PERSEDIAAN"},
	{ReportType: "NERACA", AccountType: "CURRENT ASSET", AccountGroupName: "ASET LANCAR"},
	{ReportType: "NERACA", AccountType: "NON CURRENT ASSET", AccountGroupName: "ASET TETAP"},
	{ReportType: "NERACA", AccountType: "PAYABLE", AccountGroupName: "HUTANG USAHA"},
	{ReportType: "NERACA", AccountType: "CURRENT LIABILITY", AccountGroupName: "KEWAJIBAN LANCAR"},
	{ReportType: "NERACA", AccountType: "NON CURRENT LIABILITY", AccountGroupName: "KEWAJIBAN JANGKA PANJANG"},
	{ReportType: "NERACA", AccountType: "EQUITY", AccountGroupName: "MODAL"},
	{ReportType: "NERACA", AccountType: "CURRENT YEAR EARNING", AccountGroupName: "CURRENT YEAR EARNING"},
	{ReportType: "LABA RUGI", AccountType: "MAIN REVENUE", AccountGroupName: "PENDAPATAN USAHA"},
	{ReportType: "LABA RUGI", AccountType: "OTHER REVENUE", AccountGroupName: "PENDAPATAN USAHA"},
	{ReportType: "LABA RUGI", AccountType: "MAIN COST OF REVENUE", AccountGroupName: "BEBAN POKOK USAHA"},
	{ReportType: "LABA RUGI", AccountType: "OTHER COST OF REVENUE", AccountGroupName: "BEBAN POKOK USAHA"},
	{ReportType: "LABA RUGI", AccountType: "EXPENSE", AccountGroupName: "PENDAPATAN / BIAYA LAIN"},
	{ReportType: "LABA RUGI", AccountType: "INCOME", AccountGroupName: "PENDAPATAN / BIAYA LAIN"},
}

var defaultCompanyChartOfAccounts = []model.ChartOfAccount{
	{CurrencyCode: "IDR", AccountType: "RECEIVABLE", AccountGroup: "PIUTANG USAHA", AccountCode: "1.2.1.0001", AccountName: "PIUTANG PELANGGAN"},
	{CurrencyCode: "IDR", AccountType: "INVENTORY", AccountGroup: "PERSEDIAAN", AccountCode: "1.3.1.0001", AccountName: "PERSEDIAAN USAHA"},
	{CurrencyCode: "IDR", AccountType: "CURRENT ASSET", AccountGroup: "ASET LANCAR", AccountCode: "1.4.1.0001", AccountName: "POS SEMENTARA"},
	{CurrencyCode: "IDR", AccountType: "NON CURRENT ASSET", AccountGroup: "ASET TETAP", AccountCode: "1.5.1.0001", AccountName: "ASET TETAP"},
	{CurrencyCode: "IDR", AccountType: "PAYABLE", AccountGroup: "HUTANG USAHA", AccountCode: "1.6.1.0001", AccountName: "HUTANG PEMASOK"},
	{CurrencyCode: "IDR", AccountType: "CURRENT LIABILITY", AccountGroup: "KEWAJIBAN LANCAR", AccountCode: "1.7.1.0001", AccountName: "HUTANG LAIN"},
	{CurrencyCode: "IDR", AccountType: "NON CURRENT LIABILITY", AccountGroup: "KEWAJIBAN JANGKA PANJANG", AccountCode: "1.8.1.0001", AccountName: "HUTANG LAIN JANGKA PANJANG"},
	{CurrencyCode: "IDR", AccountType: "EQUITY", AccountGroup: "MODAL", AccountCode: "1.9.1.0001", AccountName: "MODAL"},
	{CurrencyCode: "IDR", AccountType: "EQUITY", AccountGroup: "MODAL", AccountCode: "1.9.1.0002", AccountName: "PERUBAHAN MODAL"},
	{CurrencyCode: "IDR", AccountType: "CURRENT YEAR EARNING", AccountGroup: "CURRENT YEAR EARNING", AccountCode: "1.10.1.0001", AccountName: "LABA RUGI BERJALAN"},
	{CurrencyCode: "IDR", AccountType: "MAIN REVENUE", AccountGroup: "PENDAPATAN USAHA", AccountCode: "2.1.1.0001", AccountName: "PENDAPATAN UTAMA"},
	{CurrencyCode: "IDR", AccountType: "OTHER REVENUE", AccountGroup: "PENDAPATAN USAHA", AccountCode: "2.2.1.0001", AccountName: "PENDAPATAN LAIN"},
	{CurrencyCode: "IDR", AccountType: "MAIN COST OF REVENUE", AccountGroup: "BEBAN POKOK USAHA", AccountCode: "2.3.1.0001", AccountName: "BEBAN UTAMA"},
	{CurrencyCode: "IDR", AccountType: "OTHER COST OF REVENUE", AccountGroup: "BEBAN POKOK USAHA", AccountCode: "2.4.1.0001", AccountName: "BEBAN LAIN"},
	{CurrencyCode: "IDR", AccountType: "EXPENSE", AccountGroup: "PENDAPATAN / BIAYA LAIN", AccountCode: "2.5.1.0001", AccountName: "BIAYA 1"},
	{CurrencyCode: "IDR", AccountType: "EXPENSE", AccountGroup: "PENDAPATAN / BIAYA LAIN", AccountCode: "2.5.1.0002", AccountName: "BIAYA 2"},
	{CurrencyCode: "IDR", AccountType: "INCOME", AccountGroup: "PENDAPATAN / BIAYA LAIN", AccountCode: "2.6.1.0001", AccountName: "PEMASUKAN 1"},
	{CurrencyCode: "IDR", AccountType: "INCOME", AccountGroup: "PENDAPATAN / BIAYA LAIN", AccountCode: "2.6.1.0002", AccountName: "PEMASUKAN 2"},
}

type AddDefaultCompanyChartOfAccountTrxParams struct {
	CompanyId string
	BranchId  string
}

func (trx *Trx) AddDefaultCompanyChartOfAccountTransactionTrx(ctx context.Context, arg AddDefaultCompanyChartOfAccountTrxParams) error {

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		for _, d := range defaultCompanyChartOfAccountGroups {

			_, err = q.InsertChartOfAccountGroup(ctx, db.InsertChartOfAccountGroupParams{
				ID:               uuid.NewV4().String(),
				CompanyID:        arg.CompanyId,
				ReportType:       d.ReportType,
				AccountType:      d.AccountType,
				AccountGroupName: d.AccountGroupName,
			})
			if err != nil {
				return err
			}
		}

		for _, d := range defaultCompanyChartOfAccounts {

			resultGroup, err := q.GetChartOfAccountGroupByAccTypeAccGroup(context.Background(), db.GetChartOfAccountGroupByAccTypeAccGroupParams{
				CompanyID:        arg.CompanyId,
				AccountType:      d.AccountType,
				AccountGroupName: d.AccountGroup,
			})
			if err != nil {
				return err
			}

			_, err = q.InsertCompanyChartOfAccount(ctx, db.InsertCompanyChartOfAccountParams{
				ID:                    uuid.NewV4().String(),
				CompanyID:             arg.CompanyId,
				CurrencyCode:          d.CurrencyCode,
				ChartOfAccountGroupID: resultGroup.ID,
				AccountCode:           d.AccountCode,
				AccountName:           d.AccountName,
			})
			if err != nil {
				return err
			}
		}

		return err
	})

	return err
}
