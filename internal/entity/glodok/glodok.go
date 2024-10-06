package glodok

import "time"

// import "gopkg.in/guregu/null.v3/zero"

// "time"

// "gopkg.in/guregu/null.v3/zero"

type GetAdmin struct {
	AdminID   string `db:"admin_id" json:"admin_id"`
	AdminPass string `db:"admin_pass" json:"admin_pass"`
}

type TableDestinasi struct {
	DestinasiID        string    `db:"destinasi_id" json:"destinasi_id"`
	JenisDestinasiID   string    `db:"jenisdestinasi_id" json:"jenisdestinasi_id"`
	JenisDestinasiKat  string    `db:"jenisdestinasi_kat" json:"jenisdestinasi_kat"`
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
	DestinasiHalal     string    `db:"destinasi_labelhalalyn" json:"destinasi_labelhalalyn"`
	DestinasiOtentik   string    `db:"destinasi_otentikyn" json:"destinasi_otentikyn"`
	DestinasiAktif     string    `db:"destinasi_aktifyn" json:"destinasi_aktifyn"`
}

type TableTipeTransportasi struct {
	TipeTransportasiID   string `db:"tipetransportasi_id" json:"tipetransportasi_id"`
	TipeTransportasiName string `db:"tipetransportasi_name" json:"tipetransportasi_name"`
}

type TableRuteTransportasi struct {
	RuteID                   string `db:"rute_id" json:"rute_id"`
	TipeTransportasiID       string `db:"tipetransportasi_id" json:"tipetransportasi_id"`
	TipeTransportasiName     string `db:"tipetransportasi_name" json:"tipetransportasi_name"`
	TujuanID                 string `db:"tujuan_id" json:"tujuan_id"`
	TujuanAwal               string `db:"tujuan_awal" json:"tujuan_awal"`
	TujuanAkhir              string `db:"tujuan_akhir" json:"tujuan_akhir"`
	PemberhentianID          string `db:"pemberhentian_id" json:"pemberhentian_id"`
	PemberhentianNama        string `db:"pemberhentian_name" json:"pemberhentian_name"`
	PemberhentianPerbaikanYN string `db:"pemberhentian_perbaikanyn" json:"pemberhentian_perbaikanyn"`
	RuteNoBus                string `db:"rute_no" json:"rute_no"`
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

type TableJenisDestinasi struct {
	JenisDestinasiID  string `db:"jenisdestinasi_id" json:"jenisdestinasi_id"`
	JenisDestinasiKat string `db:"jenisdestinasi_kat" json:"jenisdestinasi_kat"`
}

type TableSejarahBeranda struct {
	SejarahBerandaIsi string `db:"sejarahberanda_isi" json:"sejarahberanda_isi"`
}

type TableFotoBeranda struct {
	FotoBerandaID     string `db:"fotoberanda_id" json:"fotoberanda_id"`
	FotoBerandaGambar []byte `db:"fotoberanda_gambar" json:"fotoberanda_gambar"`
	FotoBerandaURL    string `json:"fotoberanda_gambar_url"`
}

type TableVideoBeranda struct {
	VideoBerandaID   string `db:"videoberanda_id" json:"videoberanda_id"`
	VideoBerandaLink string `db:"videoberanda_link" json:"videoberanda_link"`
}

type TableTujuan struct {
	TujuanID             string `db:"tujuan_id" json:"tujuan_id"`
	TujuanAwal           string `db:"tujuan_awal" json:"tujuan_awal"`
	TujuanAkhir          string `db:"tujuan_akhir" json:"tujuan_akhir"`
	TipeTransportasiID   string `db:"tipetransportasi_id" json:"tipetransportasi_id"`
	TipeTransportasiName string `db:"tipetransportasi_name" json:"tipetransportasi_name"`
}

type TablePemberhentian struct {
	PemberhentianID          string `db:"pemberhentian_id" json:"pemberhentian_id"`
	PemberhentianNama        string `db:"pemberhentian_name" json:"pemberhentian_name"`
	TipeTransportasiID       string `db:"tipetransportasi_id" json:"tipetransportasi_id"`
	TipeTransportasiName     string `db:"tipetransportasi_name" json:"tipetransportasi_name"`
	PemberhentianPerbaikanYN string `db:"pemberhentian_perbaikanyn" json:"pemberhentian_perbaikanyn"`
}

type TableMaps struct {
	MapsLink string `db:"maps_link" json:"maps_link"`
}
