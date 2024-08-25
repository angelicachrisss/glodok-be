package glodok

import (
	"context"
	"database/sql"
	"fmt"
	"glodok-be/pkg/errors"
	"log"
	"strconv"
	"strings"

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
	fmt.Println("hashedpass", len(string(hashedPass)))
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

	// Check if the adminid is "ADMIN" or "admin"
	if strings.EqualFold(adminid, "ADMIN") {
		// Compare the plain text password directly
		if admin.AdminPass == adminpass {
			result = "Login successful"
		} else {
			result = "Invalid password"
			err = errors.New("password does not match")
			return result, errors.Wrap(err, "[DATA] [SubmitLogin]")
		}
	} else {
		// Proceed with hashed password comparison for other adminids
		err = bcrypt.CompareHashAndPassword([]byte(admin.AdminPass), []byte(adminpass))
		if err != nil {
			result = "Invalid password"
			fmt.Println("error", err)
			return result, errors.Wrap(err, "[DATA] [SubmitLogin]")
		}

		result = "Login successful"
	}

	return result, nil
}

func (d Data) GetAdminbyID(ctx context.Context, adminid string) ([]glodokEntity.GetAdmin, error) {
	var (
		admin      glodokEntity.GetAdmin
		adminArray []glodokEntity.GetAdmin
		err        error
	)

	rows, err := (*d.stmt)[getAdminbyID].QueryxContext(ctx, adminid)
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

func (d Data) GetTableAdmin(ctx context.Context, page int, length int) ([]glodokEntity.GetAdmin, error) {
	var (
		admin      glodokEntity.GetAdmin
		adminArray []glodokEntity.GetAdmin
		err        error
	)

	rows, err := (*d.stmt)[getTableAdmin].QueryxContext(ctx, page, length)
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
		err        error
	)

	rows, err := (*d.stmt)[getSearchAdmin].QueryxContext(ctx, "%"+adminid+"%", page, length)
	fmt.Println("pagelength", page, length)
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

func (d Data) DeleteAdmin(ctx context.Context, adminid string) (string, error) {
	var (
		err    error
		result string
	)

	_, err = (*d.stmt)[deleteAdmin].ExecContext(ctx, adminid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][DeleteAdmin]")
	}

	result = "Berhasil"

	return result, err
}

func (d Data) UpdateAdmin(ctx context.Context, admin glodokEntity.GetAdmin, adminid string) (string, error) {
	var (
		result string
		err    error
	)

	// Generate a hashed password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(admin.AdminPass), bcrypt.DefaultCost)
	if err != nil {
		result = "Failed to generate hashed password"
		return result, errors.Wrap(err, "[DATA][UpdateAdmin]")
	}

	// Update the hashed password in the database
	_, err = (*d.stmt)[updateAdmin].ExecContext(ctx, admin.AdminNama, string(hashedPass), adminid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][UpdateAdmin]")
	}

	result = "Berhasil"
	return result, err
}

// destinasi
func (d Data) InsertDestinasi(ctx context.Context, destinasi glodokEntity.TableDestinasi) (string, error) {
	var (
		err    error
		result string
		lastID string
		newID  string
	)

	// Fetch the last inserted DestinasiID
	err = (*d.stmt)[fetchLastDestinasiID].QueryRowxContext(ctx).Scan(&lastID)
	if err != nil && err != sql.ErrNoRows {
		result = "Gagal mengambil ID terakhir"
		return result, errors.Wrap(err, "[DATA][InsertDestinasiWarisan]")
	}

	// Generate the new DestinasiID
	if lastID != "" {
		// Extract the numeric part from lastID and increment it
		num, _ := strconv.Atoi(lastID[1:])
		newID = fmt.Sprintf("D%04d", num+1)
	} else {
		// If there are no previous records, start with D0001
		newID = "D0001"
	}

	// Assign the new DestinasiID
	destinasi.DestinasiID = newID

	// Proceed with the insertion
	_, err = (*d.stmt)[insertDestinasi].ExecContext(ctx,
		destinasi.DestinasiID,
		destinasi.DestinasiName,
		destinasi.DestinasiDesc,
		destinasi.DestinasiAlamat,
		destinasi.DestinasiGambar,
		destinasi.DestinasiLang,
		destinasi.DestinasiLong,
		destinasi.DestinasiHBuka,
		destinasi.DestinasiHTutup,
		destinasi.DestinasiKet,
		destinasi.DestinasiHalal,
	)

	log.Println("data user object", destinasi)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][InsertDestinasi]")
	}

	result = "Berhasil"
	return result, nil
}

func (d Data) GetTableDestinasi(ctx context.Context, ket string, page int, length int) ([]glodokEntity.TableDestinasi, error) {
	var (
		destinasi      glodokEntity.TableDestinasi
		destinasiArray []glodokEntity.TableDestinasi
		err            error
	)

	rows, err := (*d.stmt)[getTableDestinasi].QueryxContext(ctx, ket, page, length)
	if err != nil {
		return destinasiArray, errors.Wrap(err, "[DATA] [GetTableDestinasi]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&destinasi); err != nil {
			return destinasiArray, errors.Wrap(err, "[DATA] [GetTableDestinasi]")
		}
		destinasiArray = append(destinasiArray, destinasi)
	}
	return destinasiArray, err

}

func (d Data) GetCountDestinasi(ctx context.Context, ket string) (int, error) {
	var (
		err   error
		total int
	)

	rows, err := (*d.stmt)[getCountDestinasi].QueryxContext(ctx, ket)
	if err != nil {
		return total, errors.Wrap(err, "[DATA] [GetCountDestinasi]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&total); err != nil {
			return total, errors.Wrap(err, "[DATA] [GetCountDestinasi]")
		}

	}
	return total, err
}

func (d Data) DeleteDestinasi(ctx context.Context, destinasiid string) (string, error) {
	var (
		err    error
		result string
	)

	_, err = (*d.stmt)[deleteDestinasi].ExecContext(ctx, destinasiid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][DeleteDestinasi]")
	}

	result = "Berhasil"

	return result, err
}

// func (d Data) GetSearchDestinasiIc(ctx context.Context, destinasiname string, page int, length int) ([]glodokEntity.TableDestinasiIc, error) {
// 	var (
// 		destinasi      glodokEntity.TableDestinasiIc
// 		destinasiArray []glodokEntity.TableDestinasiIc
// 		err        error
// 	)

// 	rows, err := (*d.stmt)[getSearchDestinasiIc].QueryxContext(ctx, "%"+destinasiname+"%", page, length)
// 	fmt.Println("pagelength", page, length)
// 	if err != nil {
// 		return destinasiArray, errors.Wrap(err, "[DATA] [GetSearchDestinasiIc]")
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		if err = rows.StructScan(&destinasi); err != nil {
// 			return destinasiArray, errors.Wrap(err, "[DATA] [GetSearchDestinasiIc]")
// 		}
// 		destinasiArray = append(destinasiArray, destinasi)
// 	}
// 	return destinasiArray, err
// }

// func (d Data) GetCountSearchDestinasiIc(ctx context.Context, destinasiname string) (int, error) {
// 	var (
// 		err   error
// 		total int
// 	)

// 	rows, err := (*d.stmt)[getCountSearchDestinasiIc].QueryxContext(ctx, "%"+destinasiname+"%")
// 	if err != nil {
// 		return total, errors.Wrap(err, "[DATA] [GetCountSearchDestinasiIc]")
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		if err = rows.Scan(&total); err != nil {
// 			return total, errors.Wrap(err, "[DATA] [GetCountSearchDestinasiIc]")
// 		}

// 	}
// 	return total, err
// }
