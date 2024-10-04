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
	getTableDestinasi  = "GetTableDestinasi"
	qGetTableDestinasi = `SELECT destinasi_id, destinasi_name, destinasi_desc, destinasi_alamat, destinasi_gambar, destinasi_lang, destinasi_long,destinasi_hbuka,destinasi_htutup,destinasi_jbuka,destinasi_jtutup, destinasi_kat, destinasi_labelhalal FROM t_destinasi WHERE destinasi_kat = ? LIMIT ?,?`

	getCountDestinasi  = "GetCountDestinasi"
	qGetCountDestinasi = `SELECT COUNT(destinasi_id) AS TotalCount FROM t_destinasi WHERE destinasi_kat = ?`

	fetchLastDestinasiID  = "FetchLastDestinasiID"
	qFetchLastDestinasiID = `SELECT destinasi_id FROM t_destinasi ORDER BY destinasi_id DESC LIMIT 1`

	getImageDestinasi  = "GetImageDestinasi"
	qGetImageDestinasi = `SELECT destinasi_gambar from t_destinasi WHERE destinasi_id = ? AND destinasi_kat = ?`

	getSearchDestinasi  = "GetSearchDestinasi"
	qGetSearchDestinasi = `SELECT destinasi_id, destinasi_name, destinasi_desc, destinasi_alamat, destinasi_gambar, destinasi_lang, destinasi_long,destinasi_hbuka,destinasi_htutup,destinasi_jbuka,destinasi_jtutup, destinasi_kat, destinasi_labelhalal FROM t_destinasi WHERE destinasi_kat = ? AND( destinasi_id LIKE ? OR destinasi_name LIKE ? ) LIMIT ?,?`

	getCountSearchDestinasi  = "GetCountSearchDestinasi"
	qGetCountSearchDestinasi = `SELECT COUNT(destinasi_id) AS TotalCount FROM t_destinasi WHERE destinasi_kat = ? AND( destinasi_id LIKE ? OR destinasi_name LIKE ? )`

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

	fetchLastRuteTransportasiID  = "FetchLastRuteTransportasiID"
	qFetchLastRuteTransportasiID = `SELECT rute_id FROM t_rutetransportasi ORDER BY rute_id DESC LIMIT 1`

	getTableRuteTransportasi  = "GetTableRuteTransportasi"
	qGetTableRuteTransportasi = `SELECT 
    r.rute_id,
    r.tipetransportasi_id,
    t.tipetransportasi_name,
    r.rute_no,
    r.rute_tujuanawal,
    r.rute_tujuanakhir,
	r.rute_turun1,
	r.rute_turun2,
	r.rute_flagperbaikan1,
	r.rute_flagperbaikan2 FROM t_rutetransportasi AS r JOIN t_tipetransportasi AS t ON r.tipetransportasi_id = t.tipetransportasi_id LIMIT ?,?;`

	getCountTableRuteTransportasi  = "GetCountTableRuteTransportasi"
	qGetCountTableRuteTransportasi = "SELECT COUNT(rute_id) AS TotalCount FROM t_rutetransportasi"

	getSearchRuteTransportasi  = "GetSearchRuteTransportasi"
	qGetSearchRuteTransportasi = `SELECT 
    r.rute_id,
    r.tipetransportasi_id,
    t.tipetransportasi_name,
    r.rute_no,
    r.rute_tujuanawal,
    r.rute_tujuanakhir,
	r.rute_turun1,
	r.rute_turun2,
	r.rute_flagperbaikan1,
	r.rute_flagperbaikan2 FROM t_rutetransportasi AS r JOIN t_tipetransportasi AS t ON r.tipetransportasi_id = t.tipetransportasi_id WHERE t.tipetransportasi_name LIKE ?
    OR r.rute_tujuanawal LIKE ? OR r.rute_tujuanakhir LIKE ? LIMIT ?,?`

	getCountSearchRuteTransportasi  = "GetCountSearchRuteTransportasi"
	qGetCountSearchRuteTransportasi = `SELECT COUNT(r.rute_id) AS TotalCount FROM t_rutetransportasi AS r JOIN t_tipetransportasi AS t ON r.tipetransportasi_id = t.tipetransportasi_id WHERE t.tipetransportasi_name LIKE ? 
    OR r.rute_tujuanawal LIKE ? OR r.rute_tujuanakhir LIKE ?`

	//review
	fetchLastReviewID  = "FetchLastReviewID"
	qFetchLastReviewID = `SELECT review_id FROM t_review ORDER BY review_id DESC LIMIT 1`

	//-get all
	getTableReview  = "GetTableReview"
	qGetTableReview = `SELECT review_id, review_rating, reviewer_name, review_desc, review_date FROM t_review ORDER BY review_date DESC LIMIT ?,? `

	getCountReview  = "GetCountReview"
	qGetCountReview = `SELECT COUNT(review_id) AS TotalCount FROM t_review`

	getSearchTableReview  = "GetSearchTableReview"
	qGetSearchTableReview = `SELECT review_id, review_rating, reviewer_name, review_desc, review_date 
	FROM t_review 
	WHERE review_id LIKE ? OR reviewer_name LIKE ?
	ORDER BY review_id DESC 
	LIMIT ?,?`

	getCountSearchReview  = "GetCountSearchReview"
	qGetCountSearchReview = `SELECT COUNT(review_id) AS TotalCount FROM t_review WHERE review_id LIKE ? OR reviewer_name LIKE ?`

	//-by rating
	getTableReviewByRating  = "GetTableReviewByRating"
	qGetTableReviewByRating = `SELECT review_id, review_rating, reviewer_name, review_desc, review_date FROM t_review WHERE review_rating = ? LIMIT ?,?`

	getCountReviewByRating  = "GetCountReviewByRating"
	qGetCountReviewByRating = `SELECT COUNT(review_id) AS TotalCount FROM t_review WHERE review_rating = ?`

	getSearchReviewByRating  = "GetSearchReviewByRating"
	qGetSearchReviewByRating = `SELECT review_id, review_rating, reviewer_name, review_desc, review_date 
	FROM t_review 
	WHERE review_rating = ? AND (review_id LIKE ? OR reviewer_name LIKE ?) 
	ORDER BY review_id DESC 
	LIMIT ?,?`

	getCountSearchReviewByRating  = "GetCountSearchReviewByRating"
	qGetCountSearchReviewByRating = `SELECT COUNT(review_id) AS TotalCount
	FROM t_review WHERE review_rating = ? AND (review_id LIKE ? OR reviewer_name LIKE ?)`

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
	qGetCountTableBerita = "SELECT COUNT(berita_id) AS TotalCount FROM t_berita"

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

	//------------------------------------------------------------------------
	//query insert
	//--admin
	submitLogin  = "SubmitLogin"
	qSubmitLogin = `SELECT admin_id, admin_pass FROM t_admin WHERE admin_id = ?`

	//--destinasi
	insertDestinasi  = "InsertDestinasi"
	qInsertDestinasi = `INSERT INTO t_destinasi (destinasi_id, destinasi_name,destinasi_desc,destinasi_alamat, destinasi_gambar, destinasi_lang, destinasi_long,destinasi_hbuka, destinasi_htutup, destinasi_jbuka, destinasi_jtutup, destinasi_kat,destinasi_labelhalal)
	values (?,?,?,?,?,?,?,?,?,TIME(?),TIME(?),?,?)`

	//--tipetransportasi
	insertTipeTransportasi  = "InsertTipeTransportasi"
	qInsertTipeTransportasi = `INSERT INTO t_tipetransportasi (tipetransportasi_id, tipetransportasi_name) VALUES (?,?)`

	//--rutetransportasi
	insertRuteTransportasi  = "InsertRuteTransportasi"
	qInsertRuteTransportasi = `INSERT INTO t_rutetransportasi (rute_id, tipetransportasi_id, rute_no, rute_tujuanawal, rute_tujuanakhir, rute_turun1, rute_turun2, rute_flagperbaikan1, rute_flagperbaikan2) VALUES (?,?,?,?,?,?,?,?,?)`

	//berita
	insertBerita  = "InsertBerita"
	qInsertBerita = `INSERT INTO t_berita (berita_id, destinasi_id, berita_judul, berita_desc, berita_foto, berita_date_update, berita_linksumber) VALUES (?,?,?,?,?,CONVERT_TZ(NOW(), '+00:00', '+07:00'),?)`

	//------------------------------------------------------------------------
	//query update
	updateAdmin  = "UpdateAdmin"
	qUpdateAdmin = `UPDATE t_admin SET admin_pass =? WHERE admin_id =?`

	updateTipeTransportasi  = "UpdateTipeTransportasi"
	qUpdateTipeTransportasi = `UPDATE t_tipetransportasi SET tipetransportasi_name =?  WHERE tipetransportasi_id =?`

	updateRuteTransportasi  = "UpdateRuteTransportasi"
	qUpdateRuteTransportasi = `UPDATE t_rutetransportasi SET tipetransportasi_id = ?, rute_no = ?, rute_tujuanawal = ?, rute_tujuanakhir = ?, rute_turun1= ?, rute_turun2= ?, rute_flagperbaikan1 = ?, rute_flagperbaikan2 = ? WHERE rute_id =?`

	updateDestinasi  = "UpdateDestinasi"
	qUpdateDestinasi = `UPDATE t_destinasi SET destinasi_name =?, destinasi_desc =?, destinasi_gambar =?, destinasi_hbuka =?, destinasi_htutup =?, destinasi_jbuka =?, destinasi_jtutup =?, destinasi_kat =?, destinasi_labelhalal =? WHERE destinasi_id =?`

	updateBerita  = "UpdateBerita"
	qUpdateBerita = `UPDATE t_berita SET destinasi_id = ?, berita_judul = ?, berita_desc = ?, berita_foto= ?, berita_date_update=CONVERT_TZ(NOW(), '+00:00', '+07:00'), berita_linksumber=? WHERE berita_id=?`
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

	//FOR MASYARAKAT
	//destinasi
	getDestinasiByID  = "GetDestinasiByID"
	qGetDestinasiByID = `SELECT destinasi_id, destinasi_name, destinasi_desc, destinasi_alamat, destinasi_gambar, destinasi_lang, destinasi_long,destinasi_hbuka, destinasi_htutup, destinasi_jbuka, destinasi_jtutup, destinasi_kat, destinasi_labelhalal FROM t_destinasi WHERE destinasi_id = ?`

	getAllDestinasi  = "GetAllDestinasi"
	qGetAllDestinasi = `SELECT destinasi_id, destinasi_name, destinasi_desc, destinasi_alamat, destinasi_gambar, destinasi_lang, destinasi_long,destinasi_hbuka, destinasi_htutup, destinasi_jbuka, destinasi_jtutup, destinasi_kat, destinasi_labelhalal FROM t_destinasi WHERE destinasi_kat = ? AND destinasi_labelhalal LIKE ? AND (destinasi_name LIKE ?)`

	//review
	insertReview  = "InsertReview"
	qInsertReview = `INSERT INTO t_review (review_id, review_rating, reviewer_name, review_desc, review_date) VALUES (?,?,?,?,CONVERT_TZ(NOW(), '+00:00', '+07:00'))`

	getAllReview  = "GetAllReview"
	qGetAllReview = `SELECT review_id, review_rating, reviewer_name, review_desc, review_date 
	FROM t_review WHERE review_rating LIKE ? ORDER BY review_id DESC LIMIT ?,? `

	getCountAllReview  = `GetCountAllReview`
	qGetCountAllReview = `SELECT COUNT(review_id) AS TotalCount FROM t_review WHERE review_rating LIKE ?`
)

var (
	readStmt = []statement{
		//---admin
		{getAdmin, qGetAdmin},
		{submitLogin, qSubmitLogin},
		{getAdminbyID, qGetAdminByID},

		//---destinasi
		{fetchLastDestinasiID, qFetchLastDestinasiID},
		{getTableDestinasi, qGetTableDestinasi},
		{getCountDestinasi, qGetCountDestinasi},
		{getImageDestinasi, qGetImageDestinasi},
		{getSearchDestinasi, qGetSearchDestinasi},
		{getCountSearchDestinasi, qGetCountSearchDestinasi},

		//---tipetransportasi
		{fetchLastTipeTransportasiID, qFetchLastTipeTransportasiID},
		{getTableTipeTransportasi, qGetTableTipeTransportasi},
		{getCountTableTipeTransportasi, qGetCountTableTipeTransportasi},
		{getSearchTipeTransportasi, qGetSearchTipeTransportasi},
		{getCountSearchTipeTransportasi, qGetCountSearchTipeTransportasi},

		//--rutetransportasi
		{getTipeTransportasi, qGetTipeTransportasi},
		{fetchLastRuteTransportasiID, qFetchLastRuteTransportasiID},
		{getTableRuteTransportasi, qGetTableRuteTransportasi},
		{getCountTableRuteTransportasi, qGetCountTableRuteTransportasi},
		{getSearchRuteTransportasi, qGetSearchRuteTransportasi},
		{getCountSearchRuteTransportasi, qGetCountSearchRuteTransportasi},

		//review
		{fetchLastReviewID, qFetchLastReviewID},
		{getTableReview, qGetTableReview},
		{getCountReview, qGetCountReview},
		{getSearchTableReview, qGetSearchTableReview},
		{getCountSearchReview, qGetCountSearchReview},
		{getTableReviewByRating, qGetTableReviewByRating},
		{getCountReviewByRating, qGetCountReviewByRating},
		{getSearchReviewByRating, qGetSearchReviewByRating},
		{getCountSearchReviewByRating, qGetCountSearchReviewByRating},

		//berita
		{getDestinasi, qGetDestinasi},
		{fetchLastBeritaID, qFetchLasBeritaiID},
		{getTableBerita, qGetTableBerita},
		{getCountTableBerita, qGetCountTableBerita},
		{getImageBerita, qGetImageBerita},
		{getSearchBerita, qSearchBerita},
		{getCountSearchBerita, qCountSearchBerita},

		//for masyarakat
		{getDestinasiByID, qGetDestinasiByID},
		{getAllDestinasi, qGetAllDestinasi},
		{getAllReview, qGetAllReview},
		{getCountAllReview, qGetCountAllReview},
	}
	insertStmt = []statement{
		{insertDestinasi, qInsertDestinasi},
		{insertTipeTransportasi, qInsertTipeTransportasi},
		{insertRuteTransportasi, qInsertRuteTransportasi},
		{insertReview, qInsertReview},
		{insertBerita, qInsertBerita},
	}
	updateStmt = []statement{
		{updateAdmin, qUpdateAdmin},
		{updateTipeTransportasi, qUpdateTipeTransportasi},
		{updateRuteTransportasi, qUpdateRuteTransportasi},
		{updateDestinasi, qUpdateDestinasi},
		{updateBerita, qUpdateBerita},
	}
	deleteStmt = []statement{
		{deleteAdmin, qDeleteAdmin},
		{deleteDestinasi, qDeleteDestinasi},
		{deleteTipeTransportasi, qDeleteTipeTransportasi},
		{deleteRuteTransportasi, qDeleteRuteTransportasi},
		{deleteReview, qDeleteReview},
		{deleteBerita, qDeleteBerita},
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
