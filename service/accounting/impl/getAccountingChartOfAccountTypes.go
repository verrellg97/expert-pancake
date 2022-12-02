package impl

import (
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/accounting/model"
)

var defaultChartOfAccounts = []string{
	"AKTIVA LAIN-LAIN",
	"AKTIVA TETAP",
	"AKTIVA TIDAK BERWUJUD",
	"AKUMULASI PENYUSUTAN AKTIVA TETAP",
	"BANK",
	"BEBAN LAIN-LAIN",
	"BEBAN USAHA",
	"BIAYA ADM & UMUM",
	"BIAYA DIBAYAR DIMUKA",
	"BIAYA PENYUSUTAN",
	"BIAYA YANG HARUS DIBAYAR",
	"DEPOSITO",
	"EKUITAS",
	"HARGA POKOK PENJUALAN",
	"HUTANG DAGANG",
	"HUTANG JANGKA PANJANG",
	"HUTANG LAIN-LAIN",
	"HUTANG PAJAK",
	"INVESTASI JANGKA PANJANG",
	"KAS",
	"PAJAK DIBAYAR DIMUKA",
	"PENDAPATAN USAHA",
	"PENDAPATAN YANG MASIH HARUS DITERIMA",
	"PENGHASILAN LAIN-LAIN",
	"PERSEDIAAN",
	"PIUTANG USAHA",
	"SALDO LABA",
	"UANG MUKA",
}

func (a accountingService) GetAccountingChartOfAccountTypes(w http.ResponseWriter, r *http.Request) error {

	res := model.GetAccountingChartOfAccountTypesResponse{ChartOfAccountTypes: defaultChartOfAccounts}
	httpHandler.WriteResponse(w, res)

	return nil
}
