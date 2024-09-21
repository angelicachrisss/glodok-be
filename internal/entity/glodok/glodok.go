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
	DestinasiID        string    `db:"destinasi_id" json:"destinasi_id"`
	DestinasiName      string    `db:"destinasi_name" json:"destinasi_name"`
	DestinasiDesc      string    `db:"destinasi_desc" json:"destinasi_desc"`
	DestinasiAlamat    string    `db:"destinasi_alamat" json:"destinasi_alamat"`
	DestinasiGambar    []byte    `db:"destinasi_gambar" json:"destinasi_gambar"`
	DestinasiGambarURL string    `json:"destinasi_gambar_url"`
	DestinasiLang      float64   `db:"destinasi_lang" json:"destinasi_lang"`
	DestinasiLong      float64   `db:"destinasi_long" json:"destinasi_long"`
	DestinasiHBuka     string    `db:"destinasi_hbuka" json:"destinasi_hbuka"`
	DestinasiHTutup    string    `db:"destinasi_htutup" json:"destinasi_htutup"`
	DestinasiJBuka     time.Time `db:"destinasi_jbuka" json:"destinasi_jbuka"`
	DestinasiJTutup    time.Time `db:"destinasi_jtutup" json:"destinasi_jtutup"`
	DestinasiKet       string    `db:"destinasi_kat" json:"destinasi_kat"`
	DestinasiHalal     string    `db:"destinasi_labelhalal" json:"destinasi_labelhalal"`
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
	RuteTurun1           string `db:"rute_turun1" json:"rute_turun1"`
	RuteTurun2           string `db:"rute_turun2" json:"rute_turun2"`
	RuteFlagPerbaikan1   string `db:"rute_flagperbaikan1" json:"rute_flagperbaikan1"`
	RuteFlagPerbaikan2   string `db:"rute_flagperbaikan2" json:"rute_flagperbaikan2"`
}

type TableReview struct {
	ReviewID     string    `db:"review_id" json:"review_id"`
	ReviewRating int       `db:"review_rating" json:"review_rating"`
	Reviewer     string    `db:"reviewer_name" json:"reviewer_name"`
	ReviewDesc   string    `db:"review_desc" json:"review_desc"`
	ReviewDate   time.Time `db:"review_date" json:"review_date"`
}

type TableBerita struct {
	BeritaID         string    `db:"berita_id" json:"berita_id"`
	DestinasiID      string    `db:"destinasi_id" json:"destinasi_id"`
	DestinasiName    string    `db:"destinasi_name" json:"destinasi_name"`
	BeritaJudul      string    `db:"berita_judul" json:"berita_judul"`
	BeritaDesc       string    `db:"berita_desc" json:"berita_desc"`
	BeritaGambar     []byte    `db:"berita_foto" json:"berita_foto"`
	BeritaGambarURL  string    `json:"berita_foto_url"`
	BeritaDate       time.Time `db:"berita_date_update" json:"berita_date_update"`
	BeritaLinkSumber string    `db:"berita_linksumber" json:"berita_linksumber"`
}
