package db

import (
	"context"

	db "github.com/expert-pancake/service/accounting/db/sqlc"
	"github.com/expert-pancake/service/accounting/model"
	uuid "github.com/satori/go.uuid"
)

var defaultCompanyChartOfAccounts = []model.ChartOfAccount{
	{AccountCode: "10300", AccountName: "Deposito Berjangka", AccountGroup: "Deposito"},
	{AccountCode: "10401", AccountName: "Piutang Dagang", AccountGroup: "Piutang Usaha"},
	{AccountCode: "10501", AccountName: "Persediaan Minuman", AccountGroup: "Persediaan"},
	{AccountCode: "10502", AccountName: "Persediaan Rokok", AccountGroup: "Persediaan"},
	{AccountCode: "10601", AccountName: "Uang Muka Pembelian - Barang Dagangan", AccountGroup: "Uang Muka"},
	{AccountCode: "10602", AccountName: "Uang Muka Pembelian - Aktiva Tetap", AccountGroup: "Uang Muka"},
	{AccountCode: "10603", AccountName: "Uang Muka Sewa Tempat", AccountGroup: "Uang Muka"},
	{AccountCode: "10701", AccountName: "Bunga Deposito", AccountGroup: "Pendapatan Yang Masih Harus Diterima"},
	{AccountCode: "10702", AccountName: "Bunga Obligasi", AccountGroup: "Pendapatan Yang Masih Harus Diterima"},
	{AccountCode: "10703", AccountName: "Deviden", AccountGroup: "Pendapatan Yang Masih Harus Diterima"},
	{AccountCode: "10704", AccountName: "Jasa Keagenan", AccountGroup: "Pendapatan Yang Masih Harus Diterima"},
	{AccountCode: "10801", AccountName: "PPN Masukan (Pembelian)", AccountGroup: "Pajak Dibayar Dimuka"},
	{AccountCode: "10802", AccountName: "PPh 21", AccountGroup: "Pajak Dibayar Dimuka"},
	{AccountCode: "10803", AccountName: "PPh 23", AccountGroup: "Pajak Dibayar Dimuka"},
	{AccountCode: "10804", AccountName: "PPh Final Ps. 4 ayat 2", AccountGroup: "Pajak Dibayar Dimuka"},
	{AccountCode: "10805", AccountName: "PPh 25", AccountGroup: "Pajak Dibayar Dimuka"},
	{AccountCode: "10901", AccountName: "Sewa Gedung", AccountGroup: "Biaya Dibayar Dimuka"},
	{AccountCode: "10902", AccountName: "Asuransi", AccountGroup: "Biaya Dibayar Dimuka"},
	{AccountCode: "11001", AccountName: "Penyertaan Saham", AccountGroup: "Investasi Jangka Panjang"},
	{AccountCode: "11002", AccountName: "Obligasi", AccountGroup: "Investasi Jangka Panjang"},
	{AccountCode: "11003", AccountName: "Investasi Properti", AccountGroup: "Investasi Jangka Panjang"},
	{AccountCode: "11101", AccountName: "Tanah", AccountGroup: "Aktiva Tetap"},
	{AccountCode: "11102", AccountName: "Bangunan", AccountGroup: "Aktiva Tetap"},
	{AccountCode: "11103", AccountName: "Kendaraan", AccountGroup: "Aktiva Tetap"},
	{AccountCode: "11104", AccountName: "Peralatan", AccountGroup: "Aktiva Tetap"},
	{AccountCode: "11201", AccountName: "Akumulasi Penyusutan Bangunan", AccountGroup: "Akumulasi Penyusutan Aktiva Tetap"},
	{AccountCode: "11202", AccountName: "Akumulasi Penyusutan Kendaraan", AccountGroup: "Akumulasi Penyusutan Aktiva Tetap"},
	{AccountCode: "11203", AccountName: "Akumulasi Penyusutan Peralatan", AccountGroup: "Akumulasi Penyusutan Aktiva Tetap"},
	{AccountCode: "11301", AccountName: "Hak Patent", AccountGroup: "Aktiva Tidak Berwujud"},
	{AccountCode: "11302", AccountName: "Merk dan Cap Dagang", AccountGroup: "Aktiva Tidak Berwujud"},
	{AccountCode: "11303", AccountName: "Goodwill", AccountGroup: "Aktiva Tidak Berwujud"},
	{AccountCode: "11304", AccountName: "Franchise", AccountGroup: "Aktiva Tidak Berwujud"},
	{AccountCode: "11401", AccountName: "Biaya Pra-operasi", AccountGroup: "Aktiva Lain-lain"},
	{AccountCode: "11402", AccountName: "Piutang Kepada Pihak Istimewa", AccountGroup: "Aktiva Lain-lain"},
	{AccountCode: "11403", AccountName: "Uang Jaminan", AccountGroup: "Aktiva Lain-lain"},
	{AccountCode: "11404", AccountName: "Aktiva Tetap Yang Tidak Digunakan", AccountGroup: "Aktiva Lain-lain"},
	{AccountCode: "11499", AccountName: "Lainnya", AccountGroup: "Aktiva Lain-lain"},
	{AccountCode: "20101", AccountName: "Hutang Dagang - Lokal", AccountGroup: "Hutang Dagang"},
	{AccountCode: "20102", AccountName: "Hutang Dagang - Import", AccountGroup: "Hutang Dagang"},
	{AccountCode: "20201", AccountName: "Uang Muka Pelanggan - Penjualan", AccountGroup: "Uang Muka "},
	{AccountCode: "20202", AccountName: "Uang Muka Pelanggan - Sewa", AccountGroup: "Uang Muka "},
	{AccountCode: "20301", AccountName: "PPN Keluaran - Penjualan", AccountGroup: "Hutang Pajak"},
	{AccountCode: "20302", AccountName: "PPnBM", AccountGroup: "Hutang Pajak"},
	{AccountCode: "20303", AccountName: "PPh 21", AccountGroup: "Hutang Pajak"},
	{AccountCode: "20304", AccountName: "PPh 22", AccountGroup: "Hutang Pajak"},
	{AccountCode: "20305", AccountName: "PPh 23", AccountGroup: "Hutang Pajak"},
	{AccountCode: "20306", AccountName: "PPh Final Ps. 4 ayat 2", AccountGroup: "Hutang Pajak"},
	{AccountCode: "20307", AccountName: "PPh 25", AccountGroup: "Hutang Pajak"},
	{AccountCode: "20308", AccountName: "PPh 29", AccountGroup: "Hutang Pajak"},
	{AccountCode: "20309", AccountName: "PBB", AccountGroup: "Hutang Pajak"},
	{AccountCode: "20310", AccountName: "STP Pajak", AccountGroup: "Hutang Pajak"},
	{AccountCode: "20401", AccountName: "Biaya YMH Dibayar - Bunga Pinjaman", AccountGroup: "Biaya yang harus dibayar"},
	{AccountCode: "20402", AccountName: "Biaya YMH Dibayar - Denda Pinjaman", AccountGroup: "Biaya yang harus dibayar"},
	{AccountCode: "20403", AccountName: "Biaya YMH Dibayar - Telepon", AccountGroup: "Biaya yang harus dibayar"},
	{AccountCode: "20404", AccountName: "Biaya YMH Dibayar - Listrik", AccountGroup: "Biaya yang harus dibayar"},
	{AccountCode: "20405", AccountName: "Biaya YMH Dibayar - Sewa", AccountGroup: "Biaya yang harus dibayar"},
	{AccountCode: "20406", AccountName: "Biaya YMH Dibayar - Gaji dan Upah", AccountGroup: "Biaya yang harus dibayar"},
	{AccountCode: "20407", AccountName: "Biaya YMH Dibayar - Asuransi", AccountGroup: "Biaya yang harus dibayar"},
	{AccountCode: "20501", AccountName: "Hutang Bank", AccountGroup: "Hutang Jangka Panjang"},
	{AccountCode: "20502", AccountName: "Hutang Pihak Ketiga", AccountGroup: "Hutang Jangka Panjang"},
	{AccountCode: "29801", AccountName: "Hutang Kepada Pihak Istimewa", AccountGroup: "Hutang Lain-Lain"},
	{AccountCode: "29802", AccountName: "Uang Jaminan", AccountGroup: "Hutang Lain-Lain"},
	{AccountCode: "29803", AccountName: "Hutang Deviden", AccountGroup: "Hutang Lain-Lain"},
	{AccountCode: "29902", AccountName: "Hutang Jangka Panjang - Lembaga Kredit", AccountGroup: "Hutang Jangka Panjang"},
	{AccountCode: "29903", AccountName: "Hutang Jangka Panjang - Pihak Ketiga", AccountGroup: "Hutang Jangka Panjang"},
	{AccountCode: "30100", AccountName: "MODAL", AccountGroup: "Ekuitas"},
	{AccountCode: "30101", AccountName: "Modal Disetor", AccountGroup: "Ekuitas"},
	{AccountCode: "30102", AccountName: "Tambahan Modal Disetor - Agio Saham", AccountGroup: "Ekuitas"},
	{AccountCode: "30201", AccountName: "Saldo Laba Tahun Lalu", AccountGroup: "Saldo Laba"},
	{AccountCode: "30202", AccountName: "Koreksi Saldo Laba", AccountGroup: "Saldo Laba"},
	{AccountCode: "30203", AccountName: "Saldo Laba Tahun Berjalan", AccountGroup: "Saldo Laba"},
	{AccountCode: "40100", AccountName: "Pendapatan Usaha - Penjualan Barang Dagangan", AccountGroup: "Pendapatan Usaha"},
	{AccountCode: "40200", AccountName: "Pendapatan Usaha - Jasa Keagenan dan Distributor", AccountGroup: "Pendapatan Usaha"},
	{AccountCode: "50100", AccountName: "Harga Pokok Penjualan - Minuman", AccountGroup: "HARGA POKOK PENJUALAN"},
	{AccountCode: "50200", AccountName: "Harga Pokok Penjualan - Rokok", AccountGroup: "HARGA POKOK PENJUALAN"},
	{AccountCode: "50300", AccountName: "Harga Pokok Penjualan - Obat-obatan", AccountGroup: "HARGA POKOK PENJUALAN"},
	{AccountCode: "50400", AccountName: "Harga Pokok Penjualan - Perawatan RT", AccountGroup: "HARGA POKOK PENJUALAN"},
	{AccountCode: "50500", AccountName: "Harga Pokok Penjualan - Kosmetik", AccountGroup: "HARGA POKOK PENJUALAN"},
	{AccountCode: "50600", AccountName: "Harga Pokok Penjualan - Sembako", AccountGroup: "HARGA POKOK PENJUALAN"},
	{AccountCode: "59900", AccountName: "Harga Pokok Penjualan - Lainnya", AccountGroup: "HARGA POKOK PENJUALAN"},
	{AccountCode: "60100", AccountName: "B. Pemasaran - Gaji dan Upah", AccountGroup: "BEBAN USAHA"},
	{AccountCode: "60200", AccountName: "B. Pemasaran - Perjalanan Dinas", AccountGroup: "BEBAN USAHA"},
	{AccountCode: "60300", AccountName: "B. Pemasaran - Iklan dan Promosi", AccountGroup: "BEBAN USAHA"},
	{AccountCode: "60400", AccountName: "B. Pemasaran - Pemakaian Perlengkapan", AccountGroup: "BEBAN USAHA"},
	{AccountCode: "60500", AccountName: "B. Pemasaran - Pos, Surat, Meterai", AccountGroup: "BEBAN USAHA"},
	{AccountCode: "60600", AccountName: "B. Pemasaran - Cetakan,Majalah dan Surat Kabar", AccountGroup: "BEBAN USAHA"},
	{AccountCode: "60700", AccountName: "B. Pemasaran - Jasa Profesi", AccountGroup: "BEBAN USAHA"},
	{AccountCode: "60800", AccountName: "B. Pemasaran - Perjamuan", AccountGroup: "BEBAN USAHA"},
	{AccountCode: "60900", AccountName: "B. Pemasaran - Ijin dan Keanggotaan", AccountGroup: "BEBAN USAHA"},
	{AccountCode: "61000", AccountName: "B. Pemasaran - Sewa", AccountGroup: "BEBAN USAHA"},
	{AccountCode: "61100", AccountName: "B. Pemasaran - Telepon", AccountGroup: "BEBAN USAHA"},
	{AccountCode: "61200", AccountName: "B. Pemasaran - Listrik", AccountGroup: "BEBAN USAHA"},
	{AccountCode: "61300", AccountName: "B. Pemasaran - Asuransi", AccountGroup: "BEBAN USAHA"},
	{AccountCode: "61401", AccountName: "B. Penyusutan - Bangunan", AccountGroup: "BIAYA PENYUSUTAN"},
	{AccountCode: "61402", AccountName: "B. Penyusutan - Kendaraan", AccountGroup: "BIAYA PENYUSUTAN"},
	{AccountCode: "61403", AccountName: "B. Penyusutan - Perlt Kantor", AccountGroup: "BIAYA PENYUSUTAN"},
	{AccountCode: "61404", AccountName: "B. Penyusutan - Perlt Toko", AccountGroup: "BIAYA PENYUSUTAN"},
	{AccountCode: "61500", AccountName: "B. Pemeliharaan & Perbaikan", AccountGroup: "BIAYA PENYUSUTAN"},
	{AccountCode: "65000", AccountName: "By. Adm & Umum - Gaji & Upah", AccountGroup: "BIAYA ADM & UMUM"},
	{AccountCode: "65100", AccountName: "By. Adm & Umum - Perjalanan Dinas", AccountGroup: "BIAYA ADM & UMUM"},
	{AccountCode: "65200", AccountName: "By. Adm & Umum - Iklan dan Promosi", AccountGroup: "BIAYA ADM & UMUM"},
	{AccountCode: "65300", AccountName: "By. Adm & Umum - Pemakaian Perlengkapan", AccountGroup: "BIAYA ADM & UMUM"},
	{AccountCode: "65400", AccountName: "By. Adm & Umum - Pos,Surat, Paket", AccountGroup: "BIAYA ADM & UMUM"},
	{AccountCode: "65500", AccountName: "By. Adm & Umum - Cetakan,Majalah dan Srt Kbar", AccountGroup: "BIAYA ADM & UMUM"},
	{AccountCode: "65650", AccountName: "By. Adm & Umum - Jasa Profesi", AccountGroup: "BIAYA ADM & UMUM"},
	{AccountCode: "65700", AccountName: "By. Adm & Umum - Perjamuan", AccountGroup: "BIAYA ADM & UMUM"},
	{AccountCode: "65800", AccountName: "By. Adm & Umum - Ijin dan Keanggotaan", AccountGroup: "BIAYA ADM & UMUM"},
	{AccountCode: "65900", AccountName: "By. Adm & Umum - Sewa", AccountGroup: "BIAYA ADM & UMUM"},
	{AccountCode: "66000", AccountName: "By. Adm & Umum - Telepon", AccountGroup: "BIAYA ADM & UMUM"},
	{AccountCode: "66100", AccountName: "By. Adm & Umum - Listrik", AccountGroup: "BIAYA ADM & UMUM"},
	{AccountCode: "66200", AccountName: "By. Adm & Umum - Asuransi", AccountGroup: "BIAYA ADM & UMUM"},
	{AccountCode: "70100", AccountName: "Penghasilan Bunga Deposito", AccountGroup: "PENGHASILAN LAIN-LAIN"},
	{AccountCode: "70200", AccountName: "Penghasilan Bunga Obligasi", AccountGroup: "PENGHASILAN LAIN-LAIN"},
	{AccountCode: "70300", AccountName: "Penghasilan Deviden", AccountGroup: "PENGHASILAN LAIN-LAIN"},
	{AccountCode: "70400", AccountName: "Penghasilan Bunga Jasa Giro", AccountGroup: "PENGHASILAN LAIN-LAIN"},
	{AccountCode: "70500", AccountName: "Laba Penjualan Aktiva Tetap", AccountGroup: "PENGHASILAN LAIN-LAIN"},
	{AccountCode: "70600", AccountName: "Penghasilan Sewa", AccountGroup: "PENGHASILAN LAIN-LAIN"},
	{AccountCode: "70700", AccountName: "Laba Selisih Kurs", AccountGroup: "PENGHASILAN LAIN-LAIN"},
	{AccountCode: "79900", AccountName: "Penghasilan Lainnya", AccountGroup: "PENGHASILAN LAIN-LAIN"},
	{AccountCode: "80100", AccountName: "Beban Pajak Jasa Giro", AccountGroup: "BEBAN LAIN-LAIN"},
	{AccountCode: "80200", AccountName: "Beban Administrasi Jasa Giro", AccountGroup: "BEBAN LAIN-LAIN"},
	{AccountCode: "80300", AccountName: "Rugi Penjualan Aktiva Tetap", AccountGroup: "BEBAN LAIN-LAIN"},
	{AccountCode: "80400", AccountName: "Rugi Selisih Kurs", AccountGroup: "BEBAN LAIN-LAIN"},
	{AccountCode: "89800", AccountName: "Beban Lainnya", AccountGroup: "BEBAN LAIN-LAIN"},
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
				BranchID:     arg.BranchId,
				AccountCode:  d.AccountCode,
				AccountName:  d.AccountName,
				AccountGroup: d.AccountGroup,
			})
			if err != nil {
				return err
			}
		}

		return err
	})

	return err
}
