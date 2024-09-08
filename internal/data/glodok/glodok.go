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
	qGetAdmin = `SELECT admin_id, admin_name, admin_pass FROM t_admin`

	getTableAdmin  = "GetTableAdmin"
	qGetTableAdmin = `SELECT admin_id, admin_name, admin_pass FROM t_admin LIMIT ?,?`

	getAdminbyID  = "GetAdminByID"
	qGetAdminByID = `SELECT admin_id, admin_name, admin_pass FROM t_admin WHERE admin_id =?`

	getCountAdmin  = "GetCountAdmin"
	qGetCountAdmin = `SELECT COUNT(admin_id) AS TotalCount FROM t_admin`

	getSearchAdmin  = "GetSearchAdmin"
	qGetSearchAdmin = `SELECT admin_id, admin_name, admin_pass FROM t_admin WHERE admin_id LIKE ? LIMIT ?, ?`

	getCountSearchAdmin  = "GetCountSearchAdmin"
	qGetCountSearchAdmin = `SELECT COUNT(admin_id) AS TotalCount FROM t_admin WHERE admin_id LIKE ?`

	//--destinasi
	getTableDestinasi  = "GetTableDestinasi"
	qGetTableDestinasi = `SELECT destinasi_id, destinasi_name, destinasi_desc, destinasi_alamat, destinasi_gambar, destinasi_lang, destinasi_long,destinasi_hbuka,destinasi_htutup,destinasi_jbuka,destinasi_jtutup, destinasi_kat, destinasi_labelhalal FROM t_destinasi WHERE destinasi_kat = ? LIMIT ?,?`

	getCountDestinasi  = "GetCountDestinasi"
	qGetCountDestinasi = `SELECT COUNT(destinasi_id) AS TotalCount FROM t_destinasi WHERE destinasi_kat = ?`

	fetchLastDestinasiID  = "FetchLastDestinasiID"
	qFetchLastDestinasiID = `SELECT destinasi_id FROM t_destinasi ORDER BY destinasi_id DESC LIMIT 1`

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

	//------------------------------------------------------------------------
	//query insert
	//--admin
	insertAdmin  = "InsertAdmin"
	qInsertAdmin = `INSERT INTO t_admin (admin_id, admin_name, admin_pass) VALUES (?,?,?)`

	submitLogin  = "SubmitLogin"
	qSubmitLogin = `SELECT admin_id, admin_name, admin_pass FROM t_admin WHERE admin_id = ?`

	//--destinasi
	insertDestinasi  = "InsertDestinasi"
	qInsertDestinasi = `INSERT INTO t_destinasi (destinasi_id, destinasi_name,destinasi_desc,destinasi_alamat, destinasi_gambar, destinasi_lang, destinasi_long,destinasi_hbuka, destinasi_htutup, destinasi_jbuka, destinasi_jtutup, destinasi_kat,destinasi_labelhalal)
	values (?,?,?,?,?,?,?,?,?,TIME(?),TIME(?),?,?)`

	//--tipetransportasi
	insertTipeTransportasi  = "InsertTipeTransportasi"
	qInsertTipeTransportasi = `INSERT INTO t_tipetransportasi (tipetransportasi_id, tipetransportasi_name) VALUES (?,?)`

	//--rutetransportasi
	insertRuteTransportasi  = "InsertRuteTransportasi"
	qInsertRuteTransportasi = `INSERT INTO t_rutetransportasi (rute_id, tipetransportasi_id, rute_no, rute_tujuanawal, rute_tujuanakhir) VALUES (?,?,?,?,?)`

	getTableRuteTransportasi  = "GetTableRuteTransportasi"
	qGetTableRuteTransportasi = `SELECT 
    r.rute_id,
    r.tipetransportasi_id,
    t.tipetransportasi_name,
    r.rute_no,
    r.rute_tujuanawal,
    r.rute_tujuanakhir FROM t_rutetransportasi AS r JOIN t_tipetransportasi AS t ON r.tipetransportasi_id = t.tipetransportasi_id LIMIT ?,?;`

	getCountTableRuteTransportasi  = "GetCountTableRuteTransportasi"
	qGetCountTableRuteTransportasi = "SELECT COUNT(rute_id) AS TotalCount FROM t_rutetransportasi"

	//------------------------------------------------------------------------
	//query update
	updateAdmin  = "UpdateAdmin"
	qUpdateAdmin = `UPDATE t_admin SET admin_name =?, admin_pass =? WHERE admin_id =?`

	updateTipeTransportasi  = "UpdateTipeTransportasi"
	qUpdateTipeTransportasi = `UPDATE t_tipetransportasi SET tipetransportasi_name =?  WHERE tipetransportasi_id =?`

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
)

var (
	readStmt = []statement{
		//---admin
		{getAdmin, qGetAdmin},
		{submitLogin, qSubmitLogin},
		{getAdminbyID, qGetAdminByID},
		{getTableAdmin, qGetTableAdmin},
		{getCountAdmin, qGetCountAdmin},
		{getSearchAdmin, qGetSearchAdmin},
		{getCountSearchAdmin, qGetCountSearchAdmin},

		//---destinasi
		{fetchLastDestinasiID, qFetchLastDestinasiID},
		{getTableDestinasi, qGetTableDestinasi},
		{getCountDestinasi, qGetCountDestinasi},

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
	}
	insertStmt = []statement{
		//--admin
		{insertAdmin, qInsertAdmin},
		//--destinasi
		{insertDestinasi, qInsertDestinasi},
		//--tipetransportasi
		{insertTipeTransportasi, qInsertTipeTransportasi},
		//rutetransportasi
		{insertRuteTransportasi, qInsertRuteTransportasi},
	}
	updateStmt = []statement{
		{updateAdmin, qUpdateAdmin},
		{updateTipeTransportasi, qUpdateTipeTransportasi},
	}
	deleteStmt = []statement{
		{deleteAdmin, qDeleteAdmin},
		{deleteDestinasi, qDeleteDestinasi},
		{deleteTipeTransportasi, qDeleteTipeTransportasi},
		{deleteRuteTransportasi, qDeleteRuteTransportasi},
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
