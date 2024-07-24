package glodok

import (
	"context"
	"database/sql"
	"fmt"
	"glodok-be/pkg/errors"

	"golang.org/x/crypto/bcrypt"

	glodokEntity "glodok-be/internal/entity/glodok"
)

// "encoding/json"
// "log"
// "strconv"
// "strings"
// "time"

func (d Data) GetAdmin(ctx context.Context) ([]glodokEntity.GetAdmin, error) {
	var (
		admin      glodokEntity.GetAdmin
		adminArray []glodokEntity.GetAdmin
		err        error
	)

	rows, err := (*d.stmt)[getAdmin].QueryxContext(ctx)
	if err != nil {
		return adminArray, errors.Wrap(err, "[DATA] [GetAdmin]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&admin); err != nil {
			return adminArray, errors.Wrap(err, "[DATA] [GetAdmin]")
		}
		adminArray = append(adminArray, admin)
	}
	return adminArray, err
}

func (d Data) InsertAdmin(ctx context.Context, admin glodokEntity.GetAdmin) (string, error) {
	var (
		err    error
		result string
	)

	// Generate a hashed password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(admin.AdminPass), bcrypt.DefaultCost)
	fmt.Println("hashedpass",len(string(hashedPass)))
	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][InsertAdmin] Failed to generate hashed password")
	}

	// Insert the hashed password into the database
	_, err = (*d.stmt)[insertAdmin].ExecContext(ctx,
		admin.AdminID,
		admin.AdminNama,
		string(hashedPass), // Convert the hashed password to a string
	)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][InsertAdmin]")
	}

	result = "Berhasil"

	return result, err
}

func (d Data) SubmitLogin(ctx context.Context, adminid string, adminpass string) (string, error) {
	var (
		err    error
		result string
		admin  glodokEntity.GetAdmin
	)

	err = (*d.stmt)[submitLogin].QueryRowxContext(ctx, adminid).StructScan(&admin)
	if err != nil {
		if err == sql.ErrNoRows {
			result = "Admin not found"
		} else {
			result = "Failed to query admin"
		}
		return result, errors.Wrap(err, "[DATA] [SubmitLogin]")
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.AdminPass), []byte(adminpass))
	if err != nil {
		result = "Invalid password"
		fmt.Println("error",err)
		return result, errors.Wrap(err, "[DATA] [SubmitLogin]")
	}

	result = "Login successful"
	return result, nil
}

func (d Data) GetAdminbyID(ctx context.Context, adminid string) ([]glodokEntity.GetAdmin, error) {
	var (
		admin      glodokEntity.GetAdmin
		adminArray []glodokEntity.GetAdmin
		err        error
	)

	rows, err := (*d.stmt)[getAdminbyID].QueryxContext(ctx,adminid)
	if err != nil {
		return adminArray, errors.Wrap(err, "[DATA] [GetAdminbyID]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&admin); err != nil {
			return adminArray, errors.Wrap(err, "[DATA] [GetAdminbyID]")
		}
		adminArray = append(adminArray, admin)
	}
	return adminArray, err
}

func (d Data) GetTableAdmin(ctx context.Context,page int, length int) ([]glodokEntity.GetAdmin, error) {
	var (
		admin      glodokEntity.GetAdmin
		adminArray []glodokEntity.GetAdmin
		err        error
	)

	rows, err := (*d.stmt)[getTableAdmin].QueryxContext(ctx,page,length)
	if err != nil {
		return adminArray, errors.Wrap(err, "[DATA] [GetTableAdmin]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&admin); err != nil {
			return adminArray, errors.Wrap(err, "[DATA] [GetTableAdmin]")
		}
		adminArray = append(adminArray, admin)
	}
	return adminArray, err
}

func (d Data) GetCountAdmin(ctx context.Context) (int, error) {
	var (
		err   error
		total int
	)

	rows, err := (*d.stmt)[getCountAdmin].QueryxContext(ctx)
	if err != nil {
		return total, errors.Wrap(err, "[DATA] [GetCountAdmin]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&total); err != nil {
			return total, errors.Wrap(err, "[DATA] [GetCountAdmin]")
		}

	}
	return total, err
}

func (d Data) GetSearchAdmin(ctx context.Context, adminid string, page int, length int) ([]glodokEntity.GetAdmin, error) {
	var (
		admin      glodokEntity.GetAdmin
		adminArray []glodokEntity.GetAdmin
		err             error
	)

	rows, err := (*d.stmt)[getSearchAdmin].QueryxContext(ctx, "%"+adminid+"%", page, length)
	fmt.Println("pagelength",page,length)
	if err != nil {
		return adminArray, errors.Wrap(err, "[DATA] [GetSearchAdmin]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&admin); err != nil {
			return adminArray, errors.Wrap(err, "[DATA] [GetSearchAdmin]")
		}
		adminArray = append(adminArray, admin)
	}
	return adminArray, err
}

func (d Data) GetCountSearchAdmin(ctx context.Context, adminid string) (int, error) {
	var (
		err   error
		total int
	)

	rows, err := (*d.stmt)[getCountSearchAdmin].QueryxContext(ctx, "%"+adminid+"%")
	if err != nil {
		return total, errors.Wrap(err, "[DATA] [GetCountSearchAdmin]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&total); err != nil {
			return total, errors.Wrap(err, "[DATA] [GetCountSearchAdmin]")
		}

	}
	return total, err
}

