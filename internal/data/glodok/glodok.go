package glodok

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"

	jaegerLog "glodok-be/pkg/log"
)

type (
	// Data ...
	Data struct {
		db   *sqlx.DB
		stmt *map[string]*sqlx.Stmt

		tracer opentracing.Tracer
		logger jaegerLog.Factory
	}

	// statement ...
	statement struct {
		key   string
		query string
	}
)

const (
	// query get
	//--admin
	getAdmin  = "GetAdmin"
	qGetAdmin = `SELECT admin_id, admin_pass FROM t_admin`

	getAdminbyID  = "GetAdminByID"
	qGetAdminByID = `SELECT admin_id, admin_pass FROM t_admin WHERE admin_id =?`

	//--destinasi
	getJenisDestinasi  = "GetJenisDestinasi"
	qGetJenisDestinasi = `SELECT jenisdestinasi_id, jenisdestinasi_kat FROM t_jenisdestinasi`

	// getTableDestinasi  = "GetTableDestinasi"
	// qGetTableDestinasi = `SELECT destinasi_id, destinasi_name, destinasi_desc, destinasi_alamat, destinasi_gambar, destinasi_lang, destinasi_long,destinasi_hbuka,destinasi_htutup,destinasi_jbuka,destinasi_jtutup, destinasi_kat, destinasi_labelhalal FROM t_destinasi WHERE destinasi_kat = ? LIMIT ?,?`

	// getCountDestinasi  = "GetCountDestinasi"
	// qGetCountDestinasi = `SELECT COUNT(destinasi_id) AS TotalCount FROM t_destinasi WHERE destinasi_kat = ?`

	fetchLastDestinasiID  = "FetchLastDestinasiID"
	qFetchLastDestinasiID = `SELECT destinasi_id FROM t_destinasi ORDER BY destinasi_id DESC LIMIT 1`

	getImageDestinasi  = "GetImageDestinasi"
	qGetImageDestinasi = `SELECT destinasi_gambar from t_destinasi WHERE destinasi_id = ?`

	// getSearchDestinasi  = "GetSearchDestinasi"
	// qGetSearchDestinasi = `SELECT destinasi_id, destinasi_name, destinasi_desc, destinasi_alamat, destinasi_gambar, destinasi_lang, destinasi_long,destinasi_hbuka,destinasi_htutup,destinasi_jbuka,destinasi_jtutup, destinasi_kat, destinasi_labelhalal FROM t_destinasi WHERE destinasi_kat = ? AND( destinasi_id LIKE ? OR destinasi_name LIKE ? ) LIMIT ?,?`

	// getCountSearchDestinasi  = "GetCountSearchDestinasi"
	// qGetCountSearchDestinasi = `SELECT COUNT(destinasi_id) AS TotalCount FROM t_destinasi WHERE destinasi_kat = ? AND( destinasi_id LIKE ? OR destinasi_name LIKE ? )`

	getTableAllDestinasi  = "GetTableAllDestinasi"
	qGetTableAllDestinasi = `SELECT
    d.destinasi_id,
    d.jenisdestinasi_id,
    j.jenisdestinasi_kat,
    d.destinasi_name,
    d.destinasi_desc,
    d.destinasi_alamat,
    d.destinasi_gambar,
    d.destinasi_lang,
    d.destinasi_long,
    d.destinasi_hbuka,
    d.destinasi_htutup,
    d.destinasi_jbuka,
    d.destinasi_jtutup,
    d.destinasi_labelhalalyn,
    d.destinasi_otentikyn,
    d.destinasi_aktifyn FROM t_destinasi AS d JOIN t_jenisdestinasi AS j
    ON d.jenisdestinasi_id COLLATE utf8mb4_general_ci = j.jenisdestinasi_id COLLATE utf8mb4_general_ci LIMIT ?, ?`

	getCountTableAllDestinasi  = "GetCountTableAllDestinasi"
	qGetCountTableAllDestinasi = "SELECT COUNT(destinasi_id) AS TotalCount FROM t_destinasi"

	getSearchTableAllDestinasi  = "GetSearchTableAllDestinasi"
	qGetSearchTableAllDestinasi = `SELECT
    d.destinasi_id,
    d.jenisdestinasi_id,
    j.jenisdestinasi_kat,
    d.destinasi_name,
    d.destinasi_desc,
    d.destinasi_alamat,
    d.destinasi_gambar,
    d.destinasi_lang,
    d.destinasi_long,
    d.destinasi_hbuka,
    d.destinasi_htutup,
    d.destinasi_jbuka,
    d.destinasi_jtutup,
    d.destinasi_labelhalalyn,
    d.destinasi_otentikyn,
    d.destinasi_aktifyn FROM t_destinasi AS d JOIN t_jenisdestinasi AS j
	ON d.jenisdestinasi_id COLLATE utf8mb4_general_ci = j.jenisdestinasi_id COLLATE utf8mb4_general_ci
	WHERE d.destinasi_id COLLATE utf8mb4_general_ci LIKE ? OR d.destinasi_name COLLATE utf8mb4_general_ci LIKE ?
	LIMIT ?, ?`

	getCountSearchTableAllDestinasi  = "GetCountSearchTableAllDestinasi"
	qGetCountSearchTableAllDestinasi = `SELECT COUNT(destinasi_id) AS TotalCount FROM t_destinasi WHERE destinasi_id LIKE ? OR destinasi_name LIKE ?`

	getTableDestinasiByJenis  = "GetTableDestinasiByJenis"
	qGetTableDestinasiByJenis = `SELECT
    d.destinasi_id,
    d.jenisdestinasi_id,
    j.jenisdestinasi_kat,
    d.destinasi_name,
    d.destinasi_desc,
    d.destinasi_alamat,
    d.destinasi_gambar,
    d.destinasi_lang,
    d.destinasi_long,
    d.destinasi_hbuka,
    d.destinasi_htutup,
    d.destinasi_jbuka,
    d.destinasi_jtutup,
    d.destinasi_labelhalalyn,
    d.destinasi_otentikyn,
    d.destinasi_aktifyn FROM t_destinasi AS d JOIN t_jenisdestinasi AS j
    ON d.jenisdestinasi_id COLLATE utf8mb4_general_ci = j.jenisdestinasi_id COLLATE utf8mb4_general_ci
	WHERE d.jenisdestinasi_id COLLATE utf8mb4_general_ci = ? LIMIT ?, ?`

	getCountTableDestinasiByJenis  = "GetCountTableDestinasiByJenis"
	qGetCountTableDestinasiByJenis = `SELECT COUNT(destinasi_id) AS TotalCount FROM t_destinasi WHERE jenisdestinasi_id = ?`

	getSearchTableDestinasiByJenis  = "GetSearchTableDestinasiByJenis"
	qGetSearchTableDestinasiByJenis = `SELECT
    d.destinasi_id,
    d.jenisdestinasi_id,
    j.jenisdestinasi_kat,
    d.destinasi_name,
    d.destinasi_desc,
    d.destinasi_alamat,
    d.destinasi_gambar,
    d.destinasi_lang,
    d.destinasi_long,
    d.destinasi_hbuka,
    d.destinasi_htutup,
    d.destinasi_jbuka,
    d.destinasi_jtutup,
    d.destinasi_labelhalalyn,
    d.destinasi_otentikyn,
    d.destinasi_aktifyn FROM t_destinasi AS d JOIN t_jenisdestinasi AS j
    ON d.jenisdestinasi_id COLLATE utf8mb4_general_ci = j.jenisdestinasi_id COLLATE utf8mb4_general_ci
	WHERE d.jenisdestinasi_id COLLATE utf8mb4_general_ci = ? AND (d.destinasi_id COLLATE utf8mb4_general_ci LIKE ? OR d.destinasi_name COLLATE utf8mb4_general_ci LIKE ?)
	LIMIT ?, ?`

	getCountSearchTableDestinasiByJenis  = "GetCountSearchTableDestinasiByJenis"
	qGeCounttSearchTableDestinasiByJenis = `SELECT COUNT(destinasi_id) AS TotalCount FROM t_destinasi WHERE jenisdestinasi_id = ? AND (destinasi_id LIKE ? OR destinasi_name LIKE ?)`

	//--tipetransportasi
	fetchLastTipeTransportasiID  = "FetchLastTipeTransportasiID"
	qFetchLastTipeTransportasiID = `SELECT tipetransportasi_id FROM t_tipetransportasi ORDER BY tipetransportasi_id DESC LIMIT 1`

	getTableTipeTransportasi  = "GetTableTipeTransportasi"
	qGetTableTipeTransportasi = `SELECT tipetransportasi_id, tipetransportasi_name FROM t_tipetransportasi LIMIT ?, ?`

	getCountTableTipeTransportasi  = "GetCountTableTipeTransportasi"
	qGetCountTableTipeTransportasi = "SELECT COUNT(tipetransportasi_id) AS TotalCount FROM t_tipetransportasi"

	getSearchTipeTransportasi  = "GetSearchTipeTransportasi"
	qGetSearchTipeTransportasi = `SELECT tipetransportasi_id, tipetransportasi_name FROM t_tipetransportasi WHERE tipetransportasi_id LIKE ? OR tipetransportasi_name LIKE ? LIMIT ?, ?`

	getCountSearchTipeTransportasi  = "GetCountSearchTipeTransportasi"
	qGetCountSearchTipeTransportasi = `SELECT COUNT(tipetransportasi_id) AS TotalCount FROM t_tipetransportasi WHERE tipetransportasi_id LIKE ? OR tipetransportasi_name LIKE ?`

	//--rutetransportasi
	getTipeTransportasi  = "GetTipeTransportasi"
	qGetTipeTransportasi = `SELECT tipetransportasi_id, tipetransportasi_name FROM t_tipetransportasi`

	getTujuanTransportasiDropDown  = "GetTujuanTransportasiDropDown"
	qGetTujuanTransportasiDropDown = `SELECT t_tujuan.tujuan_id, t_tujuan.tujuan_awal, t_tujuan.tujuan_akhir
	FROM t_tujuan
	JOIN t_tipetransportasi ON t_tujuan.tipetransportasi_id = t_tipetransportasi.tipetransportasi_id
	LEFT JOIN t_rutetransportasi ON t_tujuan.tujuan_id = t_rutetransportasi.tujuan_id
	WHERE t_tipetransportasi.tipetransportasi_id COLLATE utf8mb4_unicode_ci = ?
	AND t_rutetransportasi.tujuan_id IS NULL`

	getPemberhentianDropDown  = "GetPemberhentianDropDown"
	qGetPemberhentianDropDown = `SELECT t_pemberhentian.pemberhentian_id, t_pemberhentian.pemberhentian_name, t_pemberhentian.pemberhentian_perbaikanyn
	FROM t_pemberhentian
	JOIN t_tipetransportasi ON t_pemberhentian.tipetransportasi_id = t_tipetransportasi.tipetransportasi_id
	WHERE t_tipetransportasi.tipetransportasi_id = ?`

	fetchLastRuteTransportasiID  = "FetchLastRuteTransportasiID"
	qFetchLastRuteTransportasiID = `SELECT rute_id FROM t_rutetransportasi ORDER BY rute_id DESC LIMIT 1`

	getTableRuteTransportasi  = "GetTableRuteTransportasi"
	qGetTableRuteTransportasi = `SELECT 
    r.rute_id,
    r.tipetransportasi_id,
    tt.tipetransportasi_name,
    r.pemberhentian_id,
    tp.pemberhentian_name,
	tp.pemberhentian_perbaikanyn,
    r.tujuan_id,
    tu.tujuan_awal,
    tu.tujuan_akhir,
    r.rute_no
	FROM 
    t_rutetransportasi r
	JOIN 
    t_tipetransportasi tt ON r.tipetransportasi_id = tt.tipetransportasi_id
	JOIN 
    t_tujuan tu ON r.tujuan_id = tu.tujuan_id
	JOIN 
    t_pemberhentian tp ON r.pemberhentian_id = tp.pemberhentian_id LIMIT ?,?;`

	getCountTableRuteTransportasi  = "GetCountTableRuteTransportasi"
	qGetCountTableRuteTransportasi = "SELECT COUNT(rute_id) AS TotalCount FROM t_rutetransportasi"

	getSearchRuteTransportasi  = "GetSearchRuteTransportasi"
	qGetSearchRuteTransportasi = `SELECT 
    r.rute_id,
    r.tipetransportasi_id,
    tt.tipetransportasi_name,
    r.pemberhentian_id,
    tp.pemberhentian_name,
	tp.pemberhentian_perbaikanyn,
    r.tujuan_id,
    tu.tujuan_awal,
    tu.tujuan_akhir,
    r.rute_no
	FROM 
    t_rutetransportasi r
	JOIN 
    t_tipetransportasi tt ON r.tipetransportasi_id = tt.tipetransportasi_id
	JOIN 
    t_tujuan tu ON r.tujuan_id = tu.tujuan_id
	JOIN 
    t_pemberhentian tp ON r.pemberhentian_id = tp.pemberhentian_id
	WHERE 
    tt.tipetransportasi_name LIKE ? 
    OR tu.tujuan_awal LIKE ? 
    OR tu.tujuan_akhir LIKE ? LIMIT ?,?`

	getCountSearchRuteTransportasi  = "GetCountSearchRuteTransportasi"
	qGetCountSearchRuteTransportasi = `   SELECT COUNT(r.rute_id) AS TotalCount 
	FROM t_rutetransportasi AS r
	JOIN t_tipetransportasi AS t ON r.tipetransportasi_id = t.tipetransportasi_id
	JOIN t_tujuan AS tu ON r.tujuan_id = tu.tujuan_id
	WHERE 
    t.tipetransportasi_name LIKE ? 
    OR tu.tujuan_awal LIKE ?
    OR tu.tujuan_akhir LIKE ?`

	//review
	fetchLastReviewID  = "FetchLastReviewID"
	qFetchLastReviewID = `SELECT review_id FROM t_review ORDER BY review_id DESC LIMIT 1`

	//-get all
	// getTableReview  = "GetTableReview"
	// qGetTableReview = `SELECT review_id, review_rating, review_desc, review_date FROM t_review ORDER BY review_date DESC LIMIT ?,? `

	// getCountReview  = "GetCountReview"
	// qGetCountReview = `SELECT COUNT(review_id) AS TotalCount FROM t_review`

	// getSearchTableReview  = "GetSearchTableReview"
	// qGetSearchTableReview = `SELECT review_id, review_rating, review_desc, review_date
	// FROM t_review
	// WHERE review_id LIKE ?
	// ORDER BY review_id DESC
	// LIMIT ?,?`

	// getCountSearchReview  = "GetCountSearchReview"
	// qGetCountSearchReview = `SELECT COUNT(review_id) AS TotalCount FROM t_review WHERE review_id LIKE ?`

	// //-by rating
	// getTableReviewByRating  = "GetTableReviewByRating"
	// qGetTableReviewByRating = `SELECT review_id, review_rating, review_desc, review_date FROM t_review WHERE review_rating = ? LIMIT ?,?`

	// getCountReviewByRating  = "GetCountReviewByRating"
	// qGetCountReviewByRating = `SELECT COUNT(review_id) AS TotalCount FROM t_review WHERE review_rating = ?`

	// getSearchReviewByRating  = "GetSearchReviewByRating"
	// qGetSearchReviewByRating = `SELECT review_id, review_rating review_desc, review_date
	// FROM t_review
	// WHERE review_rating = ? AND (review_id LIKE ?)
	// ORDER BY review_id DESC
	// LIMIT ?,?`

	// getCountSearchReviewByRating  = "GetCountSearchReviewByRating"
	// qGetCountSearchReviewByRating = `SELECT COUNT(review_id) AS TotalCount
	// FROM t_review WHERE review_rating = ? AND (review_id LIKE ?)`

	getTableReview  = "GetTableReview"
	qGetTableReview = `	SELECT r.review_id, r.destinasi_id, r.user_id, u.user_name, r.review_rating, r.review_desc, r.review_date, r.review_anonyn 
	FROM t_review AS r 
	JOIN t_user AS u ON u.user_id COLLATE utf8mb4_unicode_ci = r.user_id COLLATE utf8mb4_unicode_ci 
	WHERE r.destinasi_id COLLATE utf8mb4_unicode_ci LIKE ?
	AND r.user_id LIKE ?
	ORDER BY r.review_id DESC 
	LIMIT ?,?`

	getCountTableReview  = "GetCountTableReview"
	qGetCountTableReview = `SELECT COUNT(r.review_id) AS total_reviews
	FROM t_review AS r
	JOIN t_user AS u ON u.user_id COLLATE utf8mb4_unicode_ci = r.user_id COLLATE utf8mb4_unicode_ci
	WHERE r.destinasi_id COLLATE utf8mb4_unicode_ci LIKE ?
	AND r.user_id LIKE ?`

	//berita
	getDestinasi  = "GetDestinasi"
	qGetDestinasi = `SELECT destinasi_id, destinasi_name FROM t_destinasi`

	fetchLastBeritaID  = "FetchLastBeritaID"
	qFetchLasBeritaiID = `SELECT berita_id FROM t_berita ORDER BY berita_id DESC LIMIT 1`

	getTableBerita  = "GetTableBerita"
	qGetTableBerita = `SELECT
	b.berita_id,
	b.destinasi_id,
	d.destinasi_name,
	b.berita_judul,
	b.berita_desc,
	b.berita_foto,
	b.berita_date_update,
	b.berita_linksumber FROM t_berita AS b JOIN t_destinasi AS d ON b.destinasi_id = d.destinasi_id LIMIT ?,?`

	getCountTableBerita  = "GetCountTableBerita"
	qGetCountTableBerita = "SELECT COUNT(b.berita_id) as TotalCount FROM t_berita AS b JOIN t_destinasi AS d ON b.destinasi_id = d.destinasi_id"

	getImageBerita  = "GetImageBerita"
	qGetImageBerita = `SELECT berita_foto from t_berita WHERE berita_id = ?`

	getSearchBerita = "GetSearchBerita"
	qSearchBerita   = `SELECT
	b.berita_id,
	b.destinasi_id,
	d.destinasi_name,
	b.berita_judul,
	b.berita_desc,
	b.berita_foto,
	b.berita_date_update,
	b.berita_linksumber FROM t_berita AS b JOIN t_destinasi AS d ON b.destinasi_id = d.destinasi_id WHERE b.berita_id LIKE ? OR d.destinasi_name LIKE ? OR b.berita_judul LIKE ? LIMIT ?,?`

	getCountSearchBerita = "GetCountSearchBerita"
	qCountSearchBerita   = `SELECT COUNT(b.berita_id) as TotalCount FROM t_berita AS b JOIN t_destinasi AS d ON b.destinasi_id = d.destinasi_id WHERE b.berita_id LIKE ? OR d.destinasi_name LIKE ? OR b.berita_judul LIKE ?`

	//jenisdestinasi
	fetchJenisDestinasiID  = "FetchJenisDestinasiID"
	qFetchJenisDestinasiID = `SELECT jenisdestinasi_id FROM t_jenisdestinasi ORDER BY jenisdestinasi_id DESC LIMIT 1`

	getTableJenisDestinasi  = "GetTableJenisDestinasi"
	qGetTableJenisDestinasi = `SELECT jenisdestinasi_id, jenisdestinasi_kat  FROM t_jenisdestinasi WHERE jenisdestinasi_id LIKE ? OR jenisdestinasi_kat LIKE ? LIMIT ?,?`

	getCountTableJenisDestinasi  = `GetCountTableJenisDestinasi`
	qGetCountTableJenisDestinasi = `SELECT COUNT(jenisdestinasi_id) AS TotalCount FROM t_jenisdestinasi WHERE jenisdestinasi_id LIKE ? OR jenisdestinasi_kat LIKE ?`

	//sejarahberanda
	getSejarahBeranda  = "GetSejarahBeranda"
	qGetSejarahBeranda = `SELECT sejarahberanda_isi FROM t_sejarahberanda`

	//maps
	getMaps  = "GetMaps"
	qGetMaps = `SELECT maps_link FROM t_maps`

	//fotoberanda
	fetchFotoBerandaID  = "FetchFotoBerandaID"
	qFetchFotoBerandaID = `SELECT fotoberanda_id FROM t_fotoberanda ORDER BY fotoberanda_id DESC LIMIT 1`

	getTableFotoBeranda  = "GetTableFotoBeranda"
	qGetTableFotoBeranda = `SELECT fotoberanda_id, fotoberanda_gambar FROM t_fotoberanda WHERE fotoberanda_id LIKE ? LIMIT ?,?`

	getCountTableFotoBeranda  = "GetCountTableFotoBeranda"
	qGetCountTableFotoBeranda = `SELECT COUNT(fotoberanda_id) AS TotalCount FROM t_fotoberanda WHERE fotoberanda_id LIKE ?`

	getImageFotoBeranda  = "GetImageFotoBeranda"
	qGetImageFotoBeranda = `SELECT fotoberanda_gambar from t_fotoberanda WHERE fotoberanda_id = ?`

	//videoberanda
	fetchVideoBerandaID  = "FetchVideoBerandaID"
	qFetchVideoBerandaID = `SELECT videoberanda_id FROM t_videoberanda ORDER BY videoberanda_id DESC LIMIT 1`

	getTableVideoBeranda  = "GetTableVideoBeranda"
	qGetTableVideoBeranda = `SELECT videoberanda_id, videoberanda_link FROM t_videoberanda WHERE videoberanda_id LIKE ? LIMIT ?,?`

	getCountTableVideoBeranda  = "GetCountTableVideoBeranda"
	qGetCountTableVideoBeranda = `SELECT COUNT(videoberanda_id) AS TotalCount FROM t_videoberanda WHERE videoberanda_id LIKE ?`

	//tujuan
	fetchLastTujuanTransportasi  = "FetchLastTujuanTransportasi"
	qFetchLastTujuanTransportasi = `SELECT tujuan_id FROM t_tujuan ORDER BY tujuan_id DESC LIMIT 1`

	getTableTujuanTransportasi  = "GetTableTujuanTransportasi"
	qGetTableTujuanTransportasi = `SELECT
    j.tujuan_id,
    j.tipetransportasi_id,
    t.tipetransportasi_name,
    j.tujuan_awal,
	j.tujuan_akhir
	FROM t_tujuan AS j
	JOIN t_tipetransportasi AS t
    ON j.tipetransportasi_id COLLATE utf8mb4_general_ci = t.tipetransportasi_id COLLATE utf8mb4_general_ci
	WHERE j.tujuan_id COLLATE utf8mb4_general_ci LIKE ?
   	OR t.tipetransportasi_name COLLATE utf8mb4_general_ci LIKE ?
   	OR j.tujuan_awal COLLATE utf8mb4_general_ci LIKE ?
	OR j.tujuan_akhir COLLATE utf8mb4_general_ci LIKE ?
	LIMIT ?, ?`

	getCountTableTujuanTransportasi  = "GetCountTableTujuanTransportasi"
	qGetCountTableTujuanTransportasi = `SELECT COUNT(j.tujuan_id) AS TotalCount FROM t_tujuan AS j JOIN t_tipetransportasi AS t ON j.tipetransportasi_id COLLATE utf8mb4_general_ci = t.tipetransportasi_id COLLATE utf8mb4_general_ci
	WHERE j.tujuan_id COLLATE utf8mb4_general_ci LIKE ?
  	OR t.tipetransportasi_name COLLATE utf8mb4_general_ci LIKE ?
   	OR j.tujuan_awal COLLATE utf8mb4_general_ci LIKE ?
	OR j.tujuan_akhir COLLATE utf8mb4_general_ci LIKE ?`

	//pemberhentian
	fetchLastPemberhentianTransportasi  = "FetchLastPemberhentianTransportasi"
	qFetchLastPemberhentianTransportasi = `SELECT pemberhentian_id FROM t_pemberhentian ORDER BY pemberhentian_id DESC LIMIT 1`

	getTablePemberhentianTransportasi  = "GetTablePemberhentianTransportasi"
	qGetTablePemberhentianTransportasi = `SELECT
    p.pemberhentian_id,
    p.tipetransportasi_id,
    t.tipetransportasi_name,
    p.pemberhentian_name,
	p.pemberhentian_perbaikanyn
	FROM t_pemberhentian AS p
	JOIN t_tipetransportasi AS t
    ON p.tipetransportasi_id COLLATE utf8mb4_general_ci = t.tipetransportasi_id COLLATE utf8mb4_general_ci
	WHERE p.pemberhentian_id COLLATE utf8mb4_general_ci LIKE ?
   	OR t.tipetransportasi_name COLLATE utf8mb4_general_ci LIKE ?
   	OR p.pemberhentian_name COLLATE utf8mb4_general_ci LIKE ? LIMIT ?, ?`

	getCountTablePemberhentianTransportasi  = "GetCountTablePemberhentianTransportasi"
	qGetCountTablePemberhentianTransportasi = `SELECT COUNT(p.pemberhentian_id) AS TotalCount FROM t_pemberhentian AS p JOIN t_tipetransportasi AS t ON p.tipetransportasi_id COLLATE utf8mb4_general_ci = t.tipetransportasi_id COLLATE utf8mb4_general_ci
	WHERE p.pemberhentian_id COLLATE utf8mb4_general_ci LIKE ?
   	OR t.tipetransportasi_name COLLATE utf8mb4_general_ci LIKE ?
   	OR p.pemberhentian_name COLLATE utf8mb4_general_ci LIKE ?`

	getTableUser  = "GetTableUser"
	qGetTableUser = `SELECT user_id, user_name, user_pass FROM t_user WHERE user_id LIKE ? OR user_name LIKE ? LIMIT ?,?`

	getCountTableUser  = "GetCountTableUser"
	qGetCountTableUser = `SELECT COUNT(user_id) FROM t_user WHERE user_id LIKE ? OR user_name LIKE ?`

	//------------------------------------------------------------------------
	//query insert
	//--admin
	submitLogin  = "SubmitLogin"
	qSubmitLogin = `SELECT admin_id, admin_pass FROM t_admin WHERE admin_id = ?`

	//--destinasi
	insertDestinasi  = "InsertDestinasi"
	qInsertDestinasi = `INSERT INTO t_destinasi (destinasi_id, jenisdestinasi_id, destinasi_name,destinasi_desc,destinasi_alamat, destinasi_gambar, destinasi_lang, destinasi_long,destinasi_hbuka, destinasi_htutup, destinasi_jbuka, destinasi_jtutup,destinasi_labelhalalyn, destinasi_otentikyn, destinasi_aktifyn)
	values (?,?,?,?,?,?,?,?,?,?,TIME(?),TIME(?),?,?,?)`

	//--tipetransportasi
	insertTipeTransportasi  = "InsertTipeTransportasi"
	qInsertTipeTransportasi = `INSERT INTO t_tipetransportasi (tipetransportasi_id, tipetransportasi_name) VALUES (?,?)`

	//--rutetransportasi
	insertRuteTransportasi  = "InsertRuteTransportasi"
	qInsertRuteTransportasi = `INSERT INTO t_rutetransportasi (rute_id, tipetransportasi_id, pemberhentian_id, tujuan_id, rute_no) VALUES (?,?,?,?,?)`

	//berita
	insertBerita  = "InsertBerita"
	qInsertBerita = `INSERT INTO t_berita (berita_id, destinasi_id, berita_judul, berita_desc, berita_foto, berita_date_update, berita_linksumber) VALUES (?,?,?,?,?,CONVERT_TZ(NOW(), '+00:00', '+07:00'),?)`

	//jenisdestinasi
	insertJenisDestinasi  = "InsertJenisDestinasi"
	qInsertJenisDestinasi = `INSERT INTO t_jenisdestinasi (jenisdestinasi_id, jenisdestinasi_kat) VALUES (?,?)`

	//fotoberanda
	insertFotoBeranda  = "InsertFotoBeranda"
	qInsertFotoBeranda = `INSERT INTO t_fotoberanda (fotoberanda_id, fotoberanda_gambar) VALUES (?,?)`

	//videoberanda
	insertVideoBeranda  = "InsertVideoBeranda"
	qInsertVideoBeranda = `INSERT INTO t_videoberanda (videoberanda_id, videoberanda_link) VALUES (?,?)`

	//tujuan
	insertTujuanTransportasi  = "InsertTujuanTransportasi"
	qInsertTujuanTransportasi = `INSERT INTO t_tujuan (tujuan_id, tipetransportasi_id, tujuan_awal, tujuan_akhir) VALUES (?,?,?,?)`

	//pemberhentian
	insertPemberhentianTransportasi  = "InsertPemberhentianTransportasi"
	qInsertPemberhentianTransportasi = `INSERT INTO t_pemberhentian (pemberhentian_id, tipetransportasi_id, pemberhentian_name, pemberhentian_perbaikanyn) VALUES (?,?,?,?)`

	//------------------------------------------------------------------------
	//query update
	updateAdmin  = "UpdateAdmin"
	qUpdateAdmin = `UPDATE t_admin SET admin_pass =? WHERE admin_id =?`

	updateTipeTransportasi  = "UpdateTipeTransportasi"
	qUpdateTipeTransportasi = `UPDATE t_tipetransportasi SET tipetransportasi_name =?  WHERE tipetransportasi_id =?`

	updateRuteTransportasi  = "UpdateRuteTransportasi"
	qUpdateRuteTransportasi = `UPDATE t_rutetransportasi SET pemberhentian_id = ?, rute_no = ? WHERE rute_id =?`

	updateDestinasi  = "UpdateDestinasi"
	qUpdateDestinasi = `UPDATE t_destinasi SET destinasi_name =?, destinasi_desc =?, destinasi_gambar =?, destinasi_hbuka =?, destinasi_htutup =?, destinasi_jbuka =?, destinasi_jtutup =?, destinasi_labelhalalyn =?,destinasi_otentikyn =? WHERE destinasi_id =?`

	updateStatusDestinasi  = "UpdateStatusDestinasi"
	qUpdateStatusDestinasi = `UPDATE t_destinasi SET destinasi_aktifyn = ? WHERE destinasi_id =?`

	updateBerita  = "UpdateBerita"
	qUpdateBerita = `UPDATE t_berita SET destinasi_id = ?, berita_judul = ?, berita_desc = ?, berita_foto= ?, berita_date_update=CONVERT_TZ(NOW(), '+00:00', '+07:00'), berita_linksumber=? WHERE berita_id=?`

	updateJenisDestinasi  = "UpdateJenisDestinasi"
	qUpdateJenisDestinasi = `UPDATE t_jenisdestinasi SET jenisdestinasi_kat = ?  WHERE jenisdestinasi_id =?`

	updateSejarahBeranda  = "UpdateSejarahBeranda"
	qUpdateSejarahBeranda = `UPDATE t_sejarahberanda SET sejarahberanda_isi = ? WHERE sejarahberanda_id = 1`

	updateMaps  = "UpdateMaps"
	qUpdateMaps = `UPDATE t_maps SET maps_link = ? WHERE maps_link = ?`

	updateTujuan  = "UpdateTujuan"
	qUpdateTujuan = `UPDATE t_tujuan SET tujuan_awal = ?, tujuan_akhir = ? WHERE tujuan_id =?`

	updatePemberhentian  = "UpdatePemberhentian"
	qUpdatePemberhentian = `UPDATE t_pemberhentian SET pemberhentian_name = ? , pemberhentian_perbaikanyn = ? WHERE pemberhentian_id =?`

	//------------------------------------------------------------------------
	//query delete
	deleteAdmin  = "DeleteAdmin"
	qDeleteAdmin = `DELETE FROM t_admin WHERE admin_id =?`

	deleteDestinasi  = "DeleteDestinasi"
	qDeleteDestinasi = `DELETE FROM t_destinasi WHERE destinasi_id =?`

	deleteTipeTransportasi  = "DeleteTipeTransportasi"
	qDeleteTipeTransportasi = `DELETE FROM t_tipetransportasi WHERE tipetransportasi_id =?`

	deleteRuteTransportasi  = "DeleteRuteTransportasi"
	qDeleteRuteTransportasi = `DELETE FROM t_rutetransportasi WHERE rute_id =?`

	deleteReview  = "DeleteReview"
	qDeleteReview = `DELETE FROM t_review WHERE review_id =?`

	deleteBerita  = "DeleteBerita"
	qDeleteBerita = `DELETE FROM t_berita WHERE berita_id =?`

	deleteJenisDestinasi  = "DeleteJenisDestinasi"
	qDeleteJenisDestinasi = `DELETE FROM t_jenisdestinasi WHERE jenisdestinasi_id =?`

	deleteFotoBeranda  = "DeleteFotoBeranda"
	qDeleteFotoBeranda = `DELETE FROM t_fotoberanda WHERE fotoberanda_id =?`

	deleteVideoBeranda  = "DeleteVideoBeranda"
	qDeleteVideoBeranda = `DELETE FROM t_videoberanda WHERE videoberanda_id =?`

	deleteTujuan  = "DeleteTujuan"
	qDeleteTujuan = `DELETE FROM t_tujuan WHERE tujuan_id =?`

	deletePemberhentian  = "DeletePemberhentian"
	qDeletePemberhentian = `DELETE FROM t_pemberhentian WHERE pemberhentian_id =?`

	deleteRuteByPemberhentian  = "DeleteRuteByPemberhentian"
	qDeleteRuteByPemberhentian = `DELETE FROM t_rutetransportasi
	WHERE pemberhentian_id NOT IN (SELECT pemberhentian_id FROM t_pemberhentian);`

	deleteUser  = "DeleteUser"
	qDeleteUser = `DELETE FROM t_user WHERE user_id =?`

	deleteReviewByUser  = "DeleteReviewByUser"
	qDeleteReviewByUser = `DELETE FROM t_review
	WHERE user_id NOT IN (SELECT user_id COLLATE utf8mb4_unicode_ci FROM t_user)`

	//FOR MASYARAKAT
	//destinasi
	getDestinasiByID  = "GetDestinasiByID"
	qGetDestinasiByID = `SELECT d.destinasi_id, d.jenisdestinasi_id, j.jenisdestinasi_kat, d.destinasi_name, d.destinasi_desc, d.destinasi_alamat, d.destinasi_gambar, d.destinasi_lang, d.destinasi_long, d.destinasi_hbuka, d.destinasi_htutup, d.destinasi_jbuka, d.destinasi_jtutup, d.destinasi_labelhalalyn, d.destinasi_otentikyn, d.destinasi_aktifyn 
	FROM t_destinasi AS d 
	JOIN t_jenisdestinasi AS j ON d.jenisdestinasi_id = j.jenisdestinasi_id COLLATE utf8mb4_unicode_ci
	WHERE d.destinasi_id = ?`

	getAllDestinasi  = "GetAllDestinasi"
	qGetAllDestinasi = `SELECT d.destinasi_id, d.jenisdestinasi_id, j.jenisdestinasi_kat, d.destinasi_name, d.destinasi_desc, d.destinasi_alamat, d.destinasi_gambar, d.destinasi_lang, d.destinasi_long, d.destinasi_hbuka, d.destinasi_htutup, d.destinasi_jbuka, d.destinasi_jtutup, d.destinasi_labelhalalyn, d.destinasi_otentikyn, d.destinasi_aktifyn 
	FROM t_destinasi AS d 
	JOIN t_jenisdestinasi AS j ON d.jenisdestinasi_id = j.jenisdestinasi_id COLLATE utf8mb4_unicode_ci
	WHERE d.jenisdestinasi_id = ? AND d.destinasi_aktifyn = "Y" AND (d.destinasi_name LIKE ?)`

	//review
	insertReview  = "InsertReview"
	qInsertReview = `INSERT INTO t_review (review_id, destinasi_id, user_id, review_rating, review_desc, review_date, review_anonyn) VALUES (?,?,?,?,?,CONVERT_TZ(NOW(), '+00:00', '+07:00'),?)`

	getAllReview  = "GetAllReview"
	qGetAllReview = `SELECT r.review_id, r.destinasi_id, r.user_id, u.user_name, r.review_rating, r.review_desc, r.review_date, r.review_anonyn 
	FROM t_review AS r 
	JOIN t_user AS u ON u.user_id COLLATE utf8mb4_unicode_ci = r.user_id COLLATE utf8mb4_unicode_ci 
	WHERE r.destinasi_id COLLATE utf8mb4_unicode_ci = ? 
	AND r.review_rating LIKE ?
	ORDER BY r.review_id DESC 
	LIMIT ?,?`

	getCountAllReview  = `GetCountAllReview`
	qGetCountAllReview = `SELECT COUNT(review_id) AS TotalCount FROM t_review WHERE destinasi_id = ? AND review_rating LIKE ?`

	getAvgReview  = "GetAvgReview"
	qGetAvgReview = `SELECT AVG(review_rating) AS AverageRating FROM t_review WHERE destinasi_id = ?`

	//foto
	getFotoBerandaML  = "GetFotoBerandaML"
	qGetFotoBerandaML = `SELECT fotoberanda_id, fotoberanda_gambar FROM t_fotoberanda`

	//video
	getVideoBerandaML  = "GetVideoBerandaML"
	qGetVideoBerandaML = `SELECT videoberanda_id, videoberanda_link FROM t_videoberanda`

	//transportasi
	getTransportasiML  = "GetTransportasiML"
	qGetTransportasiML = `SELECT 
    r.rute_id,
    r.tipetransportasi_id,
    tt.tipetransportasi_name,
    r.pemberhentian_id,
    tp.pemberhentian_name,
	tp.pemberhentian_perbaikanyn,
    r.tujuan_id,
    tu.tujuan_awal,
    tu.tujuan_akhir,
    r.rute_no
	FROM 
    t_rutetransportasi r
	JOIN 
    t_tipetransportasi tt ON r.tipetransportasi_id = tt.tipetransportasi_id
	JOIN 
    t_tujuan tu ON r.tujuan_id = tu.tujuan_id
	JOIN 
    t_pemberhentian tp ON r.pemberhentian_id = tp.pemberhentian_id WHERE tp.pemberhentian_perbaikanyn LIKE ?`

	getCountTransportasiML  = "GetCountTransportasiML"
	qGetCountTransportasiML = `SELECT 
	COUNT(r.rute_id) AS TotalCount
	FROM 
    t_rutetransportasi r
	JOIN 
    t_tipetransportasi tt ON r.tipetransportasi_id = tt.tipetransportasi_id
	JOIN 
    t_tujuan tu ON r.tujuan_id = tu.tujuan_id
	JOIN 
    t_pemberhentian tp ON r.pemberhentian_id = tp.pemberhentian_id WHERE tp.pemberhentian_perbaikanyn LIKE ?`

	getBeritaML  = "GetBeritaML"
	qGetBeritaML = `SELECT
	b.berita_id,
	b.destinasi_id,
	d.destinasi_name,
	b.berita_judul,
	b.berita_desc,
	b.berita_foto,
	b.berita_date_update,
	b.berita_linksumber FROM t_berita AS b JOIN t_destinasi AS d ON b.destinasi_id = d.destinasi_id WHERE b.berita_judul LIKE ? ORDER BY b.berita_id DESC LIMIT ?,?`

	getCountBeritaML  = "GetCountBeritaML"
	qGetCountBeritaML = `SELECT COUNT(b.berita_id) AS TotalCount FROM t_berita AS b JOIN t_destinasi AS d ON b.destinasi_id = d.destinasi_id WHERE berita_judul LIKE ?`

	// getBeritaML  = "GetBeritaML"
	// qGetBeritaML = `SELECT
	// berita_id,
	// destinasi_id,
	// berita_judul,
	// berita_desc,
	// berita_foto,
	// berita_date_update,
	// berita_linksumber FROM t_berita WHERE berita_judul LIKE ? ORDER BY berita_id DESC LIMIT ?,?`

	// getCountBeritaML  = "GetCountBeritaML"
	// qGetCountBeritaML = `SELECT COUNT(berita_id) AS TotalCount FROM t_berita WHERE berita_judul LIKE ?`

	getBeritaMLByID  = "GetBeritaMLByID"
	qGetBeritaMLByID = `SELECT
	berita_id,
	destinasi_id,
	berita_judul,
	berita_desc,
	berita_foto,
	berita_date_update,
	berita_linksumber FROM t_berita WHERE berita_id=?`

	getJenisDestinasiML  = "GetJenisDestinasiML"
	qGetJenisDestinasiML = `SELECT DISTINCT jd.jenisdestinasi_id, jd.jenisdestinasi_kat
	FROM t_jenisdestinasi jd
	JOIN t_destinasi d ON jd.jenisdestinasi_id = d.jenisdestinasi_id COLLATE utf8mb4_unicode_ci`

	getDestinasiDDML  = "GetDestinasiDDML"
	qGetDestinasiDDML = `SELECT destinasi_id, destinasi_name FROM t_destinasi WHERE destinasi_aktifyn = "Y"`

	//user
	insertUser  = "InsertUser"
	qInsertUser = `INSERT INTO t_user (user_id, user_name, user_pass) VALUES (?,?,?)`

	submitLoginML  = "SubmitLoginML"
	qSubmitLoginML = `SELECT user_id, user_name, user_pass FROM t_user WHERE user_id = ?`

	getUser  = "GetUser"
	qGetUser = `SELECT user_id, user_name, user_pass FROM t_user WHERE user_id =?`

	updateUser  = "UpdateUser"
	qUpdateUser = `UPDATE t_user SET user_name = ?, user_pass = ? WHERE user_id = ?`
)

var (
	readStmt = []statement{
		//---admin
		{getAdmin, qGetAdmin},
		{submitLogin, qSubmitLogin},
		{getAdminbyID, qGetAdminByID},

		//---destinasi
		{fetchLastDestinasiID, qFetchLastDestinasiID},
		{getJenisDestinasi, qGetJenisDestinasi},
		// {getTableDestinasi, qGetTableDestinasi},
		// {getCountDestinasi, qGetCountDestinasi},
		{getImageDestinasi, qGetImageDestinasi},
		// {getSearchDestinasi, qGetSearchDestinasi},
		// {getCountSearchDestinasi, qGetCountSearchDestinasi},
		{getTableAllDestinasi, qGetTableAllDestinasi},
		{getCountTableAllDestinasi, qGetCountTableAllDestinasi},
		{getTableDestinasiByJenis, qGetTableDestinasiByJenis},
		{getCountTableDestinasiByJenis, qGetCountTableDestinasiByJenis},
		{getSearchTableAllDestinasi, qGetSearchTableAllDestinasi},
		{getCountSearchTableAllDestinasi, qGetCountSearchTableAllDestinasi},
		{getSearchTableDestinasiByJenis, qGetSearchTableDestinasiByJenis},
		{getCountSearchTableDestinasiByJenis, qGeCounttSearchTableDestinasiByJenis},

		//---tipetransportasi
		{fetchLastTipeTransportasiID, qFetchLastTipeTransportasiID},
		{getTableTipeTransportasi, qGetTableTipeTransportasi},
		{getCountTableTipeTransportasi, qGetCountTableTipeTransportasi},
		{getSearchTipeTransportasi, qGetSearchTipeTransportasi},
		{getCountSearchTipeTransportasi, qGetCountSearchTipeTransportasi},

		//--rutetransportasi
		{getTipeTransportasi, qGetTipeTransportasi},
		{getTujuanTransportasiDropDown, qGetTujuanTransportasiDropDown},
		{getPemberhentianDropDown, qGetPemberhentianDropDown},
		{fetchLastRuteTransportasiID, qFetchLastRuteTransportasiID},
		{getTableRuteTransportasi, qGetTableRuteTransportasi},
		{getCountTableRuteTransportasi, qGetCountTableRuteTransportasi},
		{getSearchRuteTransportasi, qGetSearchRuteTransportasi},
		{getCountSearchRuteTransportasi, qGetCountSearchRuteTransportasi},

		//review
		{fetchLastReviewID, qFetchLastReviewID},
		// {getTableReview, qGetTableReview},
		// {getCountReview, qGetCountReview},
		// {getSearchTableReview, qGetSearchTableReview},
		// {getCountSearchReview, qGetCountSearchReview},
		// {getTableReviewByRating, qGetTableReviewByRating},
		// {getCountReviewByRating, qGetCountReviewByRating},
		// {getSearchReviewByRating, qGetSearchReviewByRating},
		// {getCountSearchReviewByRating, qGetCountSearchReviewByRating},
		{getTableReview, qGetTableReview},
		{getCountTableReview, qGetCountTableReview},

		//berita
		{getDestinasi, qGetDestinasi},
		{fetchLastBeritaID, qFetchLasBeritaiID},
		{getTableBerita, qGetTableBerita},
		{getCountTableBerita, qGetCountTableBerita},
		{getImageBerita, qGetImageBerita},
		{getSearchBerita, qSearchBerita},
		{getCountSearchBerita, qCountSearchBerita},

		//jenisdestinasi
		{fetchJenisDestinasiID, qFetchJenisDestinasiID},
		{getTableJenisDestinasi, qGetTableJenisDestinasi},
		{getCountTableJenisDestinasi, qGetCountTableJenisDestinasi},

		//sejarahberanda
		{getSejarahBeranda, qGetSejarahBeranda},

		//maps
		{getMaps, qGetMaps},

		//fotoberanda
		{fetchFotoBerandaID, qFetchFotoBerandaID},
		{getTableFotoBeranda, qGetTableFotoBeranda},
		{getCountTableFotoBeranda, qGetCountTableFotoBeranda},
		{getImageFotoBeranda, qGetImageFotoBeranda},

		//videoberanda
		{fetchVideoBerandaID, qFetchVideoBerandaID},
		{getTableVideoBeranda, qGetTableVideoBeranda},
		{getCountTableVideoBeranda, qGetCountTableVideoBeranda},

		//tujuan
		{fetchLastTujuanTransportasi, qFetchLastTujuanTransportasi},
		{getTableTujuanTransportasi, qGetTableTujuanTransportasi},
		{getCountTableTujuanTransportasi, qGetCountTableTujuanTransportasi},

		//pemberhentian
		{fetchLastPemberhentianTransportasi, qFetchLastPemberhentianTransportasi},
		{getTablePemberhentianTransportasi, qGetTablePemberhentianTransportasi},
		{getCountTablePemberhentianTransportasi, qGetCountTablePemberhentianTransportasi},

		//user
		{getTableUser, qGetTableUser},
		{getCountTableUser, qGetCountTableUser},
		{getUser, qGetUser},

		//for masyarakat
		{getDestinasiByID, qGetDestinasiByID},
		{getAllDestinasi, qGetAllDestinasi},
		{getAllReview, qGetAllReview},
		{getAvgReview, qGetAvgReview},
		{getCountAllReview, qGetCountAllReview},
		{getFotoBerandaML, qGetFotoBerandaML},
		{getVideoBerandaML, qGetVideoBerandaML},
		{getTransportasiML, qGetTransportasiML},
		{getCountTransportasiML, qGetCountTransportasiML},
		{getBeritaML, qGetBeritaML},
		{getCountBeritaML, qGetCountBeritaML},
		{getBeritaMLByID, qGetBeritaMLByID},
		{getJenisDestinasiML, qGetJenisDestinasiML},
		{getDestinasiDDML, qGetDestinasiDDML},
		{submitLoginML, qSubmitLoginML},
	}
	insertStmt = []statement{
		{insertDestinasi, qInsertDestinasi},
		{insertTipeTransportasi, qInsertTipeTransportasi},
		{insertRuteTransportasi, qInsertRuteTransportasi},
		{insertReview, qInsertReview},
		{insertBerita, qInsertBerita},
		{insertJenisDestinasi, qInsertJenisDestinasi},
		{insertFotoBeranda, qInsertFotoBeranda},
		{insertVideoBeranda, qInsertVideoBeranda},
		{insertTujuanTransportasi, qInsertTujuanTransportasi},
		{insertPemberhentianTransportasi, qInsertPemberhentianTransportasi},
		{insertUser, qInsertUser},
	}
	updateStmt = []statement{
		{updateAdmin, qUpdateAdmin},
		{updateTipeTransportasi, qUpdateTipeTransportasi},
		{updateRuteTransportasi, qUpdateRuteTransportasi},
		{updateDestinasi, qUpdateDestinasi},
		{updateStatusDestinasi, qUpdateStatusDestinasi},
		{updateBerita, qUpdateBerita},
		{updateJenisDestinasi, qUpdateJenisDestinasi},
		{updateSejarahBeranda, qUpdateSejarahBeranda},
		{updateMaps, qUpdateMaps},
		{updateTujuan, qUpdateTujuan},
		{updatePemberhentian, qUpdatePemberhentian},
		{updateUser, qUpdateUser},
	}
	deleteStmt = []statement{
		{deleteAdmin, qDeleteAdmin},
		{deleteDestinasi, qDeleteDestinasi},
		{deleteTipeTransportasi, qDeleteTipeTransportasi},
		{deleteRuteTransportasi, qDeleteRuteTransportasi},
		{deleteReview, qDeleteReview},
		{deleteBerita, qDeleteBerita},
		{deleteJenisDestinasi, qDeleteJenisDestinasi},
		{deleteFotoBeranda, qDeleteFotoBeranda},
		{deleteVideoBeranda, qDeleteVideoBeranda},
		{deleteTujuan, qDeleteTujuan},
		{deletePemberhentian, qDeletePemberhentian},
		{deleteRuteByPemberhentian, qDeleteRuteByPemberhentian},
		{deleteUser, qDeleteUser},
		{deleteReviewByUser, qDeleteReviewByUser},
	}
)

// New ...
func New(db *sqlx.DB, tracer opentracing.Tracer, logger jaegerLog.Factory) *Data {
	var (
		stmts = make(map[string]*sqlx.Stmt)
	)
	d := &Data{
		db:     db,
		tracer: tracer,
		logger: logger,
		stmt:   &stmts,
	}

	d.InitStmt()
	return d
}

func (d *Data) InitStmt() {
	var (
		err   error
		stmts = make(map[string]*sqlx.Stmt)
	)

	for _, v := range readStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize select statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range insertStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize insert statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range updateStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize update statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range deleteStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize delete statement key %v, err : %v", v.key, err)
		}
	}

	*d.stmt = stmts
}
