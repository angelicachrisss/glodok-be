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
	// //query get
	// getKaryawan  = "GetKaryawan"
	// qGetKaryawan = `SELECT karyawanID, namaKaryawan, noTelp, keterangan FROM t_karyawan`

	// getCountKaryawan  = "GetCountKaryawan"
	// qGetCountKaryawan = `SELECT COUNT(karyawanID)  AS TotalCount FROM t_karyawan`

	// getAdmin  = "GetAdmin"
	// qGetAdmin = `SELECT admin_id, admin_pass from t_admin`

	// //query insert
	// insertKaryawan  = "InsertKaryawan"
	// qInsertKaryawan = `INSERT INTO t_karyawan (karyawanID, namaKaryawan, noTelp, keterangan) VALUES (?,?,?,?)`

	// query get
	getAdmin  = "GetAdmin"
	qGetAdmin = `SELECT admin_id, admin_name, admin_pass FROM t_admin`

	getTableAdmin = "GetTableAdmin"
	qGetTableAdmin=`SELECT admin_id, admin_name, admin_pass FROM t_admin LIMIT ?,?`

	getAdminbyID  = "GetAdminByID"
	qGetAdminByID = `SELECT admin_id, admin_name, admin_pass FROM t_admin WHERE admin_id =?`

	getCountAdmin  = "GetCountAdmin"
	qGetCountAdmin=`SELECT COUNT(admin_id) AS TotalCount FROM t_admin`

	getSearchAdmin = "GetSearchAdmin"
	qGetSearchAdmin = `SELECT admin_id, admin_name, admin_pass FROM t_admin WHERE admin_id LIKE ? LIMIT ?, ?`

	getCountSearchAdmin = "GetCountSearchAdmin"
	qGetCountSearchAdmin = `SELECT COUNT(admin_id) AS TotalCount FROM t_admin WHERE admin_id LIKE ?`

	//query insert
	insertAdmin  = "InsertAdmin"
	qInsertAdmin = `INSERT INTO t_admin (admin_id, admin_name, admin_pass) VALUES (?,?,?)`

	submitLogin  = "SubmitLogin"
	qSubmitLogin = `SELECT admin_id, admin_name, admin_pass FROM t_admin WHERE admin_id = ?`
)

var (
	readStmt = []statement{
		{getAdmin, qGetAdmin},
		{submitLogin, qSubmitLogin},
		{getAdminbyID, qGetAdminByID},
		{getTableAdmin,qGetTableAdmin},
		{getCountAdmin,qGetCountAdmin},
		{getSearchAdmin,qGetSearchAdmin},
		{getCountSearchAdmin,qGetCountSearchAdmin},

	}
	insertStmt = []statement{
		{insertAdmin, qInsertAdmin},
	}
	updateStmt = []statement{}
	deleteStmt = []statement{}
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
