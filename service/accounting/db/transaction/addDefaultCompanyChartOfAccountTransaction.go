package db

import (
	"context"

	db "github.com/expert-pancake/service/accounting/db/sqlc"
	"github.com/expert-pancake/service/accounting/model"
	uuid "github.com/satori/go.uuid"
)

var defaultCompanyChartOfAccounts = []model.ChartOfAccount{
	{CurrencyCode: "IDR", ReportType: "NERACA", AccountType: "RECEIVABLE", AccountGroup: "PIUTANG USAHA", AccountCode: "1.2.1.0001", AccountName: "PIUTANG PELANGGAN"},
	{CurrencyCode: "IDR", ReportType: "NERACA", AccountType: "INVENTORY", AccountGroup: "PERSEDIAAN", AccountCode: "1.3.1.0001", AccountName: "PERSEDIAAN USAHA"},
	{CurrencyCode: "IDR", ReportType: "NERACA", AccountType: "CURRENT ASSET", AccountGroup: "ASET LANCAR", AccountCode: "1.4.1.0001", AccountName: "POS SEMENTARA"},
	{CurrencyCode: "IDR", ReportType: "NERACA", AccountType: "NON CURRENT ASSET", AccountGroup: "ASET TETAP", AccountCode: "1.5.1.0001", AccountName: "ASET TETAP"},
	{CurrencyCode: "IDR", ReportType: "NERACA", AccountType: "PAYABLE", AccountGroup: "HUTANG USAHA", AccountCode: "1.6.1.0001", AccountName: "HUTANG PEMASOK"},
	{CurrencyCode: "IDR", ReportType: "NERACA", AccountType: "CURRENT LIABILITY", AccountGroup: "KEWAJIBAN LANCAR", AccountCode: "1.7.1.0001", AccountName: "HUTANG LAIN"},
	{CurrencyCode: "IDR", ReportType: "NERACA", AccountType: "NON CURRENT LIABILITY", AccountGroup: "KEWAJIBAN JANGKA PANJANG", AccountCode: "1.8.1.0001", AccountName: "HUTANG LAIN JANGKA PANJANG"},
	{CurrencyCode: "IDR", ReportType: "NERACA", AccountType: "EQUITY", AccountGroup: "MODAL", AccountCode: "1.9.1.0001", AccountName: "MODAL"},
	{CurrencyCode: "IDR", ReportType: "NERACA", AccountType: "EQUITY", AccountGroup: "MODAL", AccountCode: "1.9.1.0002", AccountName: "PERUBAHAN MODAL"},
	{CurrencyCode: "IDR", ReportType: "NERACA", AccountType: "CURRENT YEAR EARNING", AccountGroup: "CURRENT YEAR EARNING", AccountCode: "1.10.1.0001", AccountName: "LABA RUGI BERJALAN"},
	{CurrencyCode: "IDR", ReportType: "LABA RUGI", AccountType: "MAIN REVENUE", AccountGroup: "PENDAPATAN USAHA", AccountCode: "2.1.1.0001", AccountName: "PENDAPATAN UTAMA"},
	{CurrencyCode: "IDR", ReportType: "LABA RUGI", AccountType: "OTHER REVENUE", AccountGroup: "PENDAPATAN USAHA", AccountCode: "2.2.1.0001", AccountName: "PENDAPATAN LAIN"},
	{CurrencyCode: "IDR", ReportType: "LABA RUGI", AccountType: "MAIN COST OF REVENUE", AccountGroup: "BEBAN POKOK USAHA", AccountCode: "2.3.1.0001", AccountName: "BEBAN UTAMA"},
	{CurrencyCode: "IDR", ReportType: "LABA RUGI", AccountType: "OTHER COST OF REVENUE", AccountGroup: "BEBAN POKOK USAHA", AccountCode: "2.4.1.0001", AccountName: "BEBAN LAIN"},
	{CurrencyCode: "IDR", ReportType: "LABA RUGI", AccountType: "EXPENSE", AccountGroup: "PENDAPATAN / BIAYA LAIN", AccountCode: "2.5.1.0001", AccountName: "BIAYA 1"},
	{CurrencyCode: "IDR", ReportType: "LABA RUGI", AccountType: "EXPENSE", AccountGroup: "PENDAPATAN / BIAYA LAIN", AccountCode: "2.5.1.0002", AccountName: "BIAYA 2"},
	{CurrencyCode: "IDR", ReportType: "LABA RUGI", AccountType: "INCOME", AccountGroup: "PENDAPATAN / BIAYA LAIN", AccountCode: "2.6.1.0001", AccountName: "PEMASUKAN 1"},
	{CurrencyCode: "IDR", ReportType: "LABA RUGI", AccountType: "INCOME", AccountGroup: "PENDAPATAN / BIAYA LAIN", AccountCode: "2.6.1.0002", AccountName: "PEMASUKAN 2"},
}

type AddDefaultCompanyChartOfAccountTrxParams struct {
	CompanyId string
	BranchId  string
}

func (trx *Trx) AddDefaultCompanyChartOfAccountTransactionTrx(ctx context.Context, arg AddDefaultCompanyChartOfAccountTrxParams) error {

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		for _, d := range defaultCompanyChartOfAccounts {

			_, err = q.InsertCompanyChartOfAccount(ctx, db.InsertCompanyChartOfAccountParams{
				ID:           uuid.NewV4().String(),
				CompanyID:    arg.CompanyId,
				CurrencyCode: d.CurrencyCode,
				ReportType:   d.ReportType,
				AccountType:  d.AccountType,
				AccountGroup: d.AccountGroup,
				AccountCode:  d.AccountCode,
				AccountName:  d.AccountName,
			})
			if err != nil {
				return err
			}
		}

		return err
	})

	return err
}
