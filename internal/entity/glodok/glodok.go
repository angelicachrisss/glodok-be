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

