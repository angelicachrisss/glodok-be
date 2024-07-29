package glodok

// import "gopkg.in/guregu/null.v3/zero"

// "time"

// "gopkg.in/guregu/null.v3/zero"

// type GetKaryawan struct {
// 	KaryawanID   string      `db:"karyawanID" json:"karyawanID"`
// 	NamaKaryawan string      `db:"namaKaryawan" json:"namaKaryawan"`
// 	NoTelp       int         `db:"noTelp" json:"noTelp"`
// 	Keterangan   zero.String `db:"keterangan" json:"keterangan"`
// }

// type InsertKaryawan struct {
// 	Insertkaryawan GetKaryawan `json:"data"`
// }

// type GetAdmin struct {
// 	AdminID   string `db:"admin_id" json:"admin_id"`
// 	AdminPass string `db:"admin_pass" json: "admin_pass"`
// }

type GetAdmin struct {
	AdminID   string `db:"admin_id" json:"admin_id"`
	AdminNama string `db:"admin_name" json:"admin_name"`
	AdminPass string `db:"admin_pass" json:"admin_pass"`
}

type TableDestinasiIc struct {
	DestinasiID     string  `db:"destinasi_id" json:"destinasi_id"`
	DestinasiName   string  `db:"destinasi_name" json:"destinasi_name"`
	DestinasiDesc   string  `db:"destinasi_desc" json:"destinasi_desc"`
	DestinasiAlamat string  `db:"destinasi_alamat" json:"destinasi_alamat"`
	DestinasiGambar string  `db:"destinasi_gambar" json:"destinasi_gambar"`
	DestinasiLang   float64 `db:"destinasi_lang" json:"destinasi_lang"`
	DestinasiLong   float64 `db:"destinasi_long" json:"destinasi_long"`
	DestinasiHBuka string `db:"destinasi_hbuka" json:"destinasi_hbuka"`
	DestinasiHTutup string `db:"destinasi_htutup" json:"destinasi_htutup"`
}
