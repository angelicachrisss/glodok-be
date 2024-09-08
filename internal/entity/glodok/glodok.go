package glodok

import "time"

// import "gopkg.in/guregu/null.v3/zero"

// "time"

// "gopkg.in/guregu/null.v3/zero"

type GetAdmin struct {
	AdminID   string `db:"admin_id" json:"admin_id"`
	AdminNama string `db:"admin_name" json:"admin_name"`
	AdminPass string `db:"admin_pass" json:"admin_pass"`
}

type TableDestinasi struct {
	DestinasiID     string    `db:"destinasi_id" json:"destinasi_id"`
	DestinasiName   string    `db:"destinasi_name" json:"destinasi_name"`
	DestinasiDesc   string    `db:"destinasi_desc" json:"destinasi_desc"`
	DestinasiAlamat string    `db:"destinasi_alamat" json:"destinasi_alamat"`
	DestinasiGambar []byte    `db:"destinasi_gambar" json:"destinasi_gambar"`
	DestinasiLang   float64   `db:"destinasi_lang" json:"destinasi_lang"`
	DestinasiLong   float64   `db:"destinasi_long" json:"destinasi_long"`
	DestinasiHBuka  string    `db:"destinasi_hbuka" json:"destinasi_hbuka"`
	DestinasiHTutup string    `db:"destinasi_htutup" json:"destinasi_htutup"`
	DestinasiJBuka  time.Time `db:"destinasi_jbuka" json:"destinasi_jbuka"`
	DestinasiJTutup time.Time `db:"destinasi_jtutup" json:"destinasi_jtutup"`
	DestinasiKet    string    `db:"destinasi_kat" json:"destinasi_kat"`
	DestinasiHalal  string    `db:"destinasi_labelhalal" json:"destinasi_labelhalal"`
}

type TableTipeTransportasi struct {
	TipeTransportasiID   string `db:"tipetransportasi_id" json:"tipetransportasi_id"`
	TipeTransportasiName string `db:"tipetransportasi_name" json:"tipetransportasi_name"`
}

type TableRuteTransportasi struct {
	RuteID               string `db:"rute_id" json:"rute_id"`
	TipeTransportasiID   string `db:"tipetransportasi_id" json:"tipetransportasi_id"`
	RuteNoBus            string `db:"rute_no" json:"rute_no"`
	RuteTujuanAwal       string `db:"rute_tujuanawal" json:"rute_tujuanawal"`
	RuteTujuanAkhir      string `db:"rute_tujuanakhir" json:"rute_tujuanakhir"`
	TipeTransportasiName string `db:"tipetransportasi_name" json:"tipetransportasi_name"`
}
