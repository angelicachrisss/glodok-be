package glodok

import (
	"context"
	"database/sql"
	"fmt"
	"glodok-be/pkg/errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"time"

	"golang.org/x/crypto/bcrypt"

	glodokEntity "glodok-be/internal/entity/glodok"
)

func saveImageToFile(imageBytes []byte, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = file.Write(imageBytes)
	if err != nil {
		return fmt.Errorf("failed to write image to file: %w", err)
	}

	return nil
}

// Fungsi untuk menghasilkan URL gambar
func generateImageURL(id string, ket string) string {
	// var url = "https://whole-doors-clap.loca.lt"
	var url = "http://localhost:8080"
	return fmt.Sprintf(url+"/glodok/v1/data?type=getimagedestinasi&destinasiid=%s&ket=%s", id, ket)

	// http://localhost:8080/glodok/v1/data?type=getimagedestinasi&destinasiid=D0002&ket=W
}

func generateImageURLBerita(id string) string {
	// var url = "https://whole-doors-clap.loca.lt"
	var url = "http://localhost:8080"
	return fmt.Sprintf(url+"/glodok/v1/data?type=getimageberita&beritaid=%s", id)
}

func EnsureDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

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

// func (d Data) SubmitLogin(ctx context.Context, adminid string, adminpass string) (string, error) {
// 	var (
// 		err    error
// 		result string
// 		admin  glodokEntity.GetAdmin
// 	)

// 	// Query the admin data
// 	err = (*d.stmt)[submitLogin].QueryRowxContext(ctx, adminid).StructScan(&admin)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			result = "Admin not found"
// 		} else {
// 			result = "Failed to query admin"
// 		}
// 		return result, errors.Wrap(err, "[DATA] [SubmitLogin]")
// 	}

// 	// Check if the stored password is hashed or plaintext
// 	if len(admin.AdminPass) > 0 && admin.AdminPass[0] == '$' {
// 		// Assuming a hashed password starts with '$'
// 		// Proceed with hashed password comparison
// 		err = bcrypt.CompareHashAndPassword([]byte(admin.AdminPass), []byte(adminpass))
// 		if err != nil {
// 			result = "Invalid password"
// 			fmt.Println("error", err)
// 			return result, errors.Wrap(err, "[DATA] [SubmitLogin]")
// 		}
// 	} else {
// 		// Compare plaintext password directly
// 		if admin.AdminPass == adminpass {
// 			// Login successful
// 			result = "Login successful"
// 			return result, nil
// 		} else {
// 			result = "Invalid password"
// 			err = errors.New("password does not match")
// 			return result, errors.Wrap(err, "[DATA] [SubmitLogin]")
// 		}
// 	}

// 	result = "Login successful"
// 	return result, nil
// }

func (d Data) SubmitLogin(ctx context.Context, adminid string, adminpass string) (string, error) {
	var (
		err    error
		result string
		admin  glodokEntity.GetAdmin
	)

	// Query the admin data
	err = (*d.stmt)[submitLogin].QueryRowxContext(ctx, adminid).StructScan(&admin)
	if err != nil {
		if err == sql.ErrNoRows {
			result = "Admin not found"
		} else {
			result = "Failed to query admin"
		}
		return result, errors.Wrap(err, "[DATA] [SubmitLogin]")
	}

	// Check if the stored password is hashed or plaintext
	if strings.HasPrefix(admin.AdminPass, "$2a$") || strings.HasPrefix(admin.AdminPass, "$2b$") || strings.HasPrefix(admin.AdminPass, "$2y$") {
		// Proceed with hashed password comparison
		err = bcrypt.CompareHashAndPassword([]byte(admin.AdminPass), []byte(adminpass))
		if err != nil {
			result = "Invalid password"
			fmt.Println("error", err)
			return result, errors.Wrap(err, "[DATA] [SubmitLogin]")
		}
	} else {
		// Compare plaintext password directly
		if admin.AdminPass == adminpass {
			result = "Login successful"
			return result, nil
		} else {
			result = "Invalid password"
			err = errors.New("password does not match")
			return result, errors.Wrap(err, "[DATA] [SubmitLogin]")
		}
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

// func (d Data) UpdateAdmin(ctx context.Context, admin glodokEntity.GetAdmin, adminid string) (string, error) {
// 	var (
// 		result string
// 		err    error
// 	)

// 	// Generate a hashed password
// 	hashedPass, err := bcrypt.GenerateFromPassword([]byte(admin.AdminPass), bcrypt.DefaultCost)
// 	if err != nil {
// 		result = "Failed to generate hashed password"
// 		return result, errors.Wrap(err, "[DATA][UpdateAdmin]")
// 	}

// 	// Update the hashed password in the database
// 	_, err = (*d.stmt)[updateAdmin].ExecContext(ctx, string(hashedPass), adminid)

// 	if err != nil {
// 		result = "Gagal"
// 		return result, errors.Wrap(err, "[DATA][UpdateAdmin]")
// 	}

// 	result = "Berhasil"
// 	return result, err
// }

func (d Data) UpdateAdmin(ctx context.Context, admin glodokEntity.GetAdmin, adminid string) (string, error) {
	var result string

	// Check if the password is "admin" or "ADMIN"
	if admin.AdminPass == "admin" || admin.AdminPass == "ADMIN" {
		// Update the password as-is (no hashing)
		_, err := (*d.stmt)[updateAdmin].ExecContext(ctx, admin.AdminPass, adminid)
		if err != nil {
			result = "Gagal"
			return result, errors.Wrap(err, "[DATA][UpdateAdmin]")
		}
	} else {
		// Generate a hashed password
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(admin.AdminPass), bcrypt.DefaultCost)
		if err != nil {
			result = "Failed to generate hashed password"
			return result, errors.Wrap(err, "[DATA][UpdateAdmin]")
		}

		// Update the hashed password in the database
		_, err = (*d.stmt)[updateAdmin].ExecContext(ctx, string(hashedPass), adminid)
		if err != nil {
			result = "Gagal"
			return result, errors.Wrap(err, "[DATA][UpdateAdmin]")
		}
	}

	result = "Berhasil"
	return result, nil
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
		return result, errors.Wrap(err, "[DATA][InsertDestinasi]")
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

	jambuka := destinasi.DestinasiJBuka.Format("15:04:05")
	jamtutup := destinasi.DestinasiJTutup.Format("15:04:05")

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
		jambuka,
		jamtutup,
		destinasi.DestinasiKet,
		destinasi.DestinasiHalal,
	)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][InsertDestinasi]")
	}

	result = "Berhasil"
	return result, nil
}

func (d Data) GetTableDestinasi(ctx context.Context, ket string, page int, length int) ([]glodokEntity.TableDestinasi, error) {
	var (
		destinasiArray []glodokEntity.TableDestinasi
		err            error
	)

	rows, err := (*d.stmt)[getTableDestinasi].QueryxContext(ctx, ket, page, length)
	if err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetTableDestinasi] - Query failed")
	}
	defer rows.Close()

	// Ensure the directory exists
	imageDir := filepath.Join("public", "images")
	if err := EnsureDirectory(imageDir); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetTableDestinasi] - Failed to ensure directory")
	}

	for rows.Next() {
		var destinasi glodokEntity.TableDestinasi
		var jbukaStr, jtutupStr string

		if err = rows.Scan(&destinasi.DestinasiID, &destinasi.DestinasiName, &destinasi.DestinasiDesc,
			&destinasi.DestinasiAlamat, &destinasi.DestinasiGambarURL,
			&destinasi.DestinasiLang, &destinasi.DestinasiLong, &destinasi.DestinasiHBuka,
			&destinasi.DestinasiHTutup, &jbukaStr, &jtutupStr, &destinasi.DestinasiKet,
			&destinasi.DestinasiHalal); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetTableDestinasi] - Failed to scan row")
		}

		if jbukaStr != "" {
			jbukaStr = "0001-01-01 " + jbukaStr
			// destinasi.DestinasiJBuka, err = time.Parse("15:04:05", jbukaStr)
			destinasi.DestinasiJBuka, err = time.Parse("2006-01-02 15:04:05", jbukaStr)
			if err != nil {
				return nil, errors.Wrap(err, "[DATA] [GetTableDestinasi] - Failed to parse destinasi_jbuka")
			}
		}

		if jtutupStr != "" {
			jtutupStr = "0001-01-01 " + jtutupStr
			// destinasi.DestinasiJTutup, err = time.Parse("15:04:05", jtutupStr)
			destinasi.DestinasiJTutup, err = time.Parse("2006-01-02 15:04:05", jtutupStr)
			if err != nil {
				return nil, errors.Wrap(err, "[DATA] [GetTableDestinasi] - Failed to parse destinasi_jtutup")
			}
		}

		// Save image and generate URL
		filePath := filepath.Join(imageDir, destinasi.DestinasiID+".jpg")
		if err := saveImageToFile(destinasi.DestinasiGambar, filePath); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetTableDestinasi] - Failed to save image")
		}

		destinasi.DestinasiGambarURL = generateImageURL(destinasi.DestinasiID, destinasi.DestinasiKet)
		destinasiArray = append(destinasiArray, destinasi)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetTableDestinasi] - Row iteration error")
	}

	return destinasiArray, nil
}

func (d Data) GetImageDestinasi(ctx context.Context, destinasiid string, destinasikat string) ([]byte, error) {
	var image []byte
	if err := (*d.stmt)[getImageDestinasi].QueryRowxContext(ctx, destinasiid, destinasikat).Scan(&image); err != nil {
		return image, errors.Wrap(err, "[DATA][GetImageDestinasi]")
	}

	return image, nil
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

func (d Data) GetSearchDestinasi(ctx context.Context, kategori string, destinasiid string, destinasiname string, page int, length int) ([]glodokEntity.TableDestinasi, error) {
	var (
		// destinasi      glodokEntity.TableDestinasi
		destinasiArray []glodokEntity.TableDestinasi
		err            error
	)

	rows, err := (*d.stmt)[getSearchDestinasi].QueryxContext(ctx, kategori, "%"+destinasiid+"%", "%"+destinasiname+"%", page, length)
	if err != nil {
		return destinasiArray, errors.Wrap(err, "[DATA] [GetSearchDestinasi]")
	}

	defer rows.Close()

	// Ensure the directory exists
	imageDir := filepath.Join("public", "images")
	if err := EnsureDirectory(imageDir); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetTableDestinasi] - Failed to ensure directory")
	}

	for rows.Next() {
		var destinasi glodokEntity.TableDestinasi
		var jbukaStr, jtutupStr string

		if err = rows.Scan(&destinasi.DestinasiID, &destinasi.DestinasiName, &destinasi.DestinasiDesc,
			&destinasi.DestinasiAlamat, &destinasi.DestinasiGambarURL,
			&destinasi.DestinasiLang, &destinasi.DestinasiLong, &destinasi.DestinasiHBuka,
			&destinasi.DestinasiHTutup, &jbukaStr, &jtutupStr, &destinasi.DestinasiKet,
			&destinasi.DestinasiHalal); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetTableDestinasi] - Failed to scan row")
		}

		if jbukaStr != "" {
			destinasi.DestinasiJBuka, err = time.Parse("15:04:05", jbukaStr)
			if err != nil {
				return nil, errors.Wrap(err, "[DATA] [GetTableDestinasi] - Failed to parse destinasi_jbuka")
			}
		}

		if jtutupStr != "" {
			destinasi.DestinasiJTutup, err = time.Parse("15:04:05", jtutupStr)
			if err != nil {
				return nil, errors.Wrap(err, "[DATA] [GetTableDestinasi] - Failed to parse destinasi_jtutup")
			}
		}

		// Save image and generate URL
		filePath := filepath.Join(imageDir, destinasi.DestinasiID+".jpg")
		if err := saveImageToFile(destinasi.DestinasiGambar, filePath); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetTableDestinasi] - Failed to save image")
		}

		destinasi.DestinasiGambarURL = generateImageURL(destinasi.DestinasiID, destinasi.DestinasiKet)
		destinasiArray = append(destinasiArray, destinasi)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetTableDestinasi] - Row iteration error")
	}

	return destinasiArray, nil

}

func (d Data) GetCountSearchDestinasi(ctx context.Context, kategori string, destinasiid string, destinasiname string) (int, error) {
	var (
		err   error
		total int
	)

	rows, err := (*d.stmt)[getCountSearchDestinasi].QueryxContext(ctx, kategori, "%"+destinasiid+"%", "%"+destinasiname+"%")
	if err != nil {
		return total, errors.Wrap(err, "[DATA] [GetCountSearchDestinasi]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&total); err != nil {
			return total, errors.Wrap(err, "[DATA] [GetCountSearchDestinasi]")
		}

	}
	return total, err
}

func (d Data) UpdateDestinasi(ctx context.Context, destinasi glodokEntity.TableDestinasi, destinasiid string) (string, error) {
	var (
		result string
		err    error
	)

	jambuka := destinasi.DestinasiJBuka.Format("15:04:05")
	jamtutup := destinasi.DestinasiJTutup.Format("15:04:05")

	_, err = (*d.stmt)[updateDestinasi].ExecContext(ctx,
		destinasi.DestinasiName,
		destinasi.DestinasiDesc,
		destinasi.DestinasiGambar,
		destinasi.DestinasiHBuka,
		destinasi.DestinasiHTutup,
		jambuka,
		jamtutup,
		destinasi.DestinasiKet,
		destinasi.DestinasiHalal,
		destinasiid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][UpdateDestinasi]")
	}

	result = "Berhasil"
	return result, nil
}

// tipetransportasi
func (d Data) InsertTipeTransportasi(ctx context.Context, tipetransportasi glodokEntity.TableTipeTransportasi) (string, error) {
	var (
		err    error
		result string
		lastID string
		newID  string
	)

	err = (*d.stmt)[fetchLastTipeTransportasiID].QueryRowxContext(ctx).Scan(&lastID)
	if err != nil && err != sql.ErrNoRows {
		result = "Gagal mengambil ID terakhir"
		return result, errors.Wrap(err, "[DATA][InsertTipeTransportasi]")
	}

	if lastID != "" {
		// Extract the numeric part from lastID and increment it
		num, _ := strconv.Atoi(lastID[2:])
		newID = fmt.Sprintf("TT%03d", num+1)

	} else {
		newID = "TT001"

	}

	tipetransportasi.TipeTransportasiID = newID

	// Proceed with the insertion
	_, err = (*d.stmt)[insertTipeTransportasi].ExecContext(ctx,
		tipetransportasi.TipeTransportasiID,
		tipetransportasi.TipeTransportasiName,
	)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][InsertTipeTransportasi]")
	}

	result = "Berhasil"
	return result, nil
}

func (d Data) GetTableTipeTransportasi(ctx context.Context, page int, length int) ([]glodokEntity.TableTipeTransportasi, error) {
	var (
		tipeTransportasi      glodokEntity.TableTipeTransportasi
		tipeTransportasiArray []glodokEntity.TableTipeTransportasi
		err                   error
	)

	rows, err := (*d.stmt)[getTableTipeTransportasi].QueryxContext(ctx, page, length)
	if err != nil {
		return tipeTransportasiArray, errors.Wrap(err, "[DATA] [GetTableTipeTransportasi]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&tipeTransportasi); err != nil {
			return tipeTransportasiArray, errors.Wrap(err, "[DATA] [GetTableTipeTransportasi]")
		}
		tipeTransportasiArray = append(tipeTransportasiArray, tipeTransportasi)
	}
	return tipeTransportasiArray, err

}

func (d Data) GetCountTableTipeTransportasi(ctx context.Context) (int, error) {
	var (
		err   error
		total int
	)

	rows, err := (*d.stmt)[getCountTableTipeTransportasi].QueryxContext(ctx)
	if err != nil {
		return total, errors.Wrap(err, "[DATA] [GetCountTableTipeTransportasi]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&total); err != nil {
			return total, errors.Wrap(err, "[DATA] [GetCountTableTipeTransportasi]")
		}

	}
	return total, err
}

func (d Data) GetSearchTipeTransportasi(ctx context.Context, tipetransportasiid string, tipetransportasiname string, page int, length int) ([]glodokEntity.TableTipeTransportasi, error) {
	var (
		tipeTransportasi      glodokEntity.TableTipeTransportasi
		tipeTransportasiArray []glodokEntity.TableTipeTransportasi
		err                   error
	)

	rows, err := (*d.stmt)[getSearchTipeTransportasi].QueryxContext(ctx, "%"+tipetransportasiid+"%", "%"+tipetransportasiname+"%", page, length)
	fmt.Println("pagelength", page, length)
	if err != nil {
		return tipeTransportasiArray, errors.Wrap(err, "[DATA] [GetSearchTipeTransportasi]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&tipeTransportasi); err != nil {
			return tipeTransportasiArray, errors.Wrap(err, "[DATA] [GetSearchTipeTransportasi]")
		}
		tipeTransportasiArray = append(tipeTransportasiArray, tipeTransportasi)
	}
	return tipeTransportasiArray, err
}

func (d Data) GetCountSearchTipeTransportasi(ctx context.Context, tipetransportasiid string, tipetransportasiname string) (int, error) {
	var (
		err   error
		total int
	)

	rows, err := (*d.stmt)[getCountSearchTipeTransportasi].QueryxContext(ctx, "%"+tipetransportasiid+"%", "%"+tipetransportasiname+"%")
	if err != nil {
		return total, errors.Wrap(err, "[DATA] [GetCountSearchTipeTransportasi]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&total); err != nil {
			return total, errors.Wrap(err, "[DATA] [GetCountSearchTipeTransportasi]")
		}

	}
	return total, err
}

func (d Data) DeleteTipeTransportasi(ctx context.Context, tipetransportasiid string) (string, error) {
	var (
		err    error
		result string
	)

	_, err = (*d.stmt)[deleteTipeTransportasi].ExecContext(ctx, tipetransportasiid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][DeleteTipeTransportasi]")
	}

	result = "Berhasil"

	return result, err
}

func (d Data) UpdateTipeTransportasi(ctx context.Context, tipetransportasi glodokEntity.TableTipeTransportasi, tipetransportasiid string) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[updateTipeTransportasi].ExecContext(ctx, tipetransportasi.TipeTransportasiName, tipetransportasiid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][UpdateTipeTransportasi]")
	}

	result = "Berhasil"
	return result, err
}

// rutetransportasi
func (d Data) GetTipeTransportasi(ctx context.Context) ([]glodokEntity.TableTipeTransportasi, error) {
	var (
		tipeTransportasi      glodokEntity.TableTipeTransportasi
		tipeTransportasiArray []glodokEntity.TableTipeTransportasi
		err                   error
	)

	rows, err := (*d.stmt)[getTipeTransportasi].QueryxContext(ctx)
	if err != nil {
		return tipeTransportasiArray, errors.Wrap(err, "[DATA] [GetTipeTransportasi]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&tipeTransportasi); err != nil {
			return tipeTransportasiArray, errors.Wrap(err, "[DATA] [GetTipeTransportasi]")
		}
		tipeTransportasiArray = append(tipeTransportasiArray, tipeTransportasi)
	}
	return tipeTransportasiArray, err

}

func (d Data) InsertRuteTransportasi(ctx context.Context, rutetransportasi glodokEntity.TableRuteTransportasi) (string, error) {
	var (
		err    error
		result string
		lastID string
		newID  string
	)

	err = (*d.stmt)[fetchLastRuteTransportasiID].QueryRowxContext(ctx).Scan(&lastID)
	if err != nil && err != sql.ErrNoRows {
		result = "Gagal mengambil ID terakhir"
		return result, errors.Wrap(err, "[DATA][fetchLastRuteTransportasiID]")
	}

	if lastID != "" {
		// Extract the numeric part from lastID and increment it
		num, _ := strconv.Atoi(lastID[1:])
		newID = fmt.Sprintf("R%04d", num+1)

	} else {
		newID = "R0001"

	}

	rutetransportasi.RuteID = newID

	// Proceed with the insertion
	_, err = (*d.stmt)[insertRuteTransportasi].ExecContext(ctx,
		rutetransportasi.RuteID,
		rutetransportasi.TipeTransportasiID,
		rutetransportasi.RuteNoBus,
		rutetransportasi.RuteTujuanAwal,
		rutetransportasi.RuteTujuanAkhir,
		rutetransportasi.RuteTurun1,
		rutetransportasi.RuteTurun2,
		rutetransportasi.RuteFlagPerbaikan1,
		rutetransportasi.RuteFlagPerbaikan2,
	)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][InsertRuteTransportasi]")
	}

	result = "Berhasil"
	return result, nil
}

func (d Data) GetTableRuteTransportasi(ctx context.Context, page int, length int) ([]glodokEntity.TableRuteTransportasi, error) {
	var (
		ruteTransportasi      glodokEntity.TableRuteTransportasi
		ruteTransportasiArray []glodokEntity.TableRuteTransportasi
		err                   error
	)

	rows, err := (*d.stmt)[getTableRuteTransportasi].QueryxContext(ctx, page, length)
	if err != nil {
		return ruteTransportasiArray, errors.Wrap(err, "[DATA] [GetTableRuteTransportasi]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&ruteTransportasi); err != nil {
			return ruteTransportasiArray, errors.Wrap(err, "[DATA] [GetTableRuteTransportasi]")
		}
		ruteTransportasiArray = append(ruteTransportasiArray, ruteTransportasi)
	}
	return ruteTransportasiArray, err

}

func (d Data) GetCountTableRuteTransportasi(ctx context.Context) (int, error) {
	var (
		err   error
		total int
	)

	rows, err := (*d.stmt)[getCountTableRuteTransportasi].QueryxContext(ctx)
	if err != nil {
		return total, errors.Wrap(err, "[DATA] [GetCountTableRuteTransportasi]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&total); err != nil {
			return total, errors.Wrap(err, "[DATA] [GetCountTableRuteTransportasi]")
		}

	}
	return total, err
}

func (d Data) DeleteRuteTransportasi(ctx context.Context, ruteid string) (string, error) {
	var (
		err    error
		result string
	)

	_, err = (*d.stmt)[deleteRuteTransportasi].ExecContext(ctx, ruteid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][DeleteRuteTransportasi]")
	}

	result = "Berhasil"

	return result, err
}

func (d Data) GetSearchRuteTransportasi(ctx context.Context, tipetransportasiname string, tujuanawal string, tujuanakhir string, page int, length int) ([]glodokEntity.TableRuteTransportasi, error) {
	var (
		ruteTransportasi      glodokEntity.TableRuteTransportasi
		ruteTransportasiArray []glodokEntity.TableRuteTransportasi
		err                   error
	)

	rows, err := (*d.stmt)[getSearchRuteTransportasi].QueryxContext(ctx, "%"+tipetransportasiname+"%", "%"+tujuanawal+"%", "%"+tujuanakhir+"%", page, length)
	if err != nil {
		return ruteTransportasiArray, errors.Wrap(err, "[DATA] [GetSearchRuteTransportasi]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&ruteTransportasi); err != nil {
			return ruteTransportasiArray, errors.Wrap(err, "[DATA] [GetSearchRuteTransportasi]")
		}
		ruteTransportasiArray = append(ruteTransportasiArray, ruteTransportasi)
	}
	return ruteTransportasiArray, err
}

func (d Data) GetCountSearchRuteTransportasi(ctx context.Context, tipetransportasiname string, tujuanawal string, tujuanakhir string) (int, error) {
	var (
		err   error
		total int
	)

	rows, err := (*d.stmt)[getCountSearchRuteTransportasi].QueryxContext(ctx, "%"+tipetransportasiname+"%", "%"+tujuanawal+"%", "%"+tujuanakhir+"%")
	if err != nil {
		return total, errors.Wrap(err, "[DATA] [GetCountSearchRuteTransportasi]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&total); err != nil {
			return total, errors.Wrap(err, "[DATA] [GetCountSearchRuteTransportasi]")
		}

	}
	return total, err
}

func (d Data) UpdateRuteTransportasi(ctx context.Context, rutetransportasi glodokEntity.TableRuteTransportasi, ruteid string) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[updateRuteTransportasi].ExecContext(ctx,
		rutetransportasi.TipeTransportasiID,
		rutetransportasi.RuteNoBus,
		rutetransportasi.RuteTujuanAwal,
		rutetransportasi.RuteTujuanAkhir,
		rutetransportasi.RuteTurun1,
		rutetransportasi.RuteTurun2,
		rutetransportasi.RuteFlagPerbaikan1,
		rutetransportasi.RuteFlagPerbaikan2,
		ruteid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][UpdateRuteTransportasi]")
	}

	result = "Berhasil"
	return result, err
}

// review
func (d Data) InsertReview(ctx context.Context, review glodokEntity.TableReview) (string, error) {
	var (
		err    error
		result string
		lastID string
		newID  string
	)

	err = (*d.stmt)[fetchLastReviewID].QueryRowxContext(ctx).Scan(&lastID)
	if err != nil && err != sql.ErrNoRows {
		result = "Gagal mengambil ID terakhir"
		return result, errors.Wrap(err, "[DATA][InsertReview]")
	}

	if lastID != "" {
		// Extract the numeric part from lastID and increment it
		num, _ := strconv.Atoi(lastID[1:])
		newID = fmt.Sprintf("U%04d", num+1)

	} else {
		newID = "U0001"

	}

	review.ReviewID = newID

	// Proceed with the insertion
	_, err = (*d.stmt)[insertReview].ExecContext(ctx,
		review.ReviewID,
		review.ReviewRating,
		review.Reviewer,
		review.ReviewDesc,
	)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][InsertReview]")
	}

	result = "Berhasil"
	return result, nil
}

func (d Data) GetTableReview(ctx context.Context, page int, length int) ([]glodokEntity.TableReview, error) {
	var (
		review      glodokEntity.TableReview
		reviewArray []glodokEntity.TableReview
		err         error
	)

	rows, err := (*d.stmt)[getTableReview].QueryxContext(ctx, page, length)
	if err != nil {
		return reviewArray, errors.Wrap(err, "[DATA] [GetTableReview]")
	}

	defer rows.Close()

	for rows.Next() {
		var reviewID string
		var reviewRating int
		var reviewerName string
		var reviewDesc string
		var reviewDateRaw []byte // Raw byte slice for date

		if err = rows.Scan(&reviewID, &reviewRating, &reviewerName, &reviewDesc, &reviewDateRaw); err != nil {
			return reviewArray, errors.Wrap(err, "[DATA] [GetTableReview]")
		}

		// Convert raw date to time.Time
		var reviewDate time.Time
		if reviewDateRaw != nil {
			reviewDate, err = time.Parse("2006-01-02 15:04:05", string(reviewDateRaw))
			if err != nil {
				return reviewArray, errors.Wrap(err, "[DATA] [GetTableReview] parsing date")
			}
		}

		review = glodokEntity.TableReview{
			ReviewID:     reviewID,
			ReviewRating: reviewRating,
			Reviewer:     reviewerName,
			ReviewDesc:   reviewDesc,
			ReviewDate:   reviewDate,
		}

		reviewArray = append(reviewArray, review)
	}
	return reviewArray, err
}

func (d Data) GetCountTableReview(ctx context.Context) (int, error) {
	var (
		err   error
		total int
	)

	rows, err := (*d.stmt)[getCountReview].QueryxContext(ctx)
	if err != nil {
		return total, errors.Wrap(err, "[DATA] [GetCountTableReview]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&total); err != nil {
			return total, errors.Wrap(err, "[DATA] [GetCountTableReview]")
		}

	}
	return total, err
}

func (d Data) DeleteReview(ctx context.Context, reviewid string) (string, error) {
	var (
		err    error
		result string
	)

	_, err = (*d.stmt)[deleteReview].ExecContext(ctx, reviewid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][DeleteReview]")
	}

	result = "Berhasil"

	return result, err
}

func (d Data) GetSearchReview(ctx context.Context, reviewid string, reviewer string, page int, length int) ([]glodokEntity.TableReview, error) {
	var (
		review      glodokEntity.TableReview
		reviewArray []glodokEntity.TableReview
		err         error
	)

	rows, err := (*d.stmt)[getSearchTableReview].QueryxContext(ctx, "%"+reviewid+"%", "%"+reviewer+"%", page, length)
	fmt.Println("pagelength", page, length)
	if err != nil {
		return reviewArray, errors.Wrap(err, "[DATA] [GetSearchReview]")
	}

	defer rows.Close()

	for rows.Next() {
		var reviewID string
		var reviewRating int
		var reviewer string
		var reviewDesc string
		var reviewDateRaw []byte // Raw byte slice for date

		if err = rows.Scan(&reviewID, &reviewRating, &reviewer, &reviewDesc, &reviewDateRaw); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetSearchReview]")
		}

		//convert
		var reviewDate time.Time
		if reviewDateRaw != nil {
			reviewDate, err = time.Parse("2006-01-02 15:04:05", string(reviewDateRaw))
			if err != nil {
				return reviewArray, errors.Wrap(err, "[DATA] [GetTableReview] parsing date")
			}
		}

		review = glodokEntity.TableReview{
			ReviewID:     reviewID,
			ReviewRating: reviewRating,
			Reviewer:     reviewer,
			ReviewDesc:   reviewDesc,
			ReviewDate:   reviewDate,
		}

		reviewArray = append(reviewArray, review)
	}
	return reviewArray, err
}

func (d Data) GetCountSearchReview(ctx context.Context, reviewid string, reviewer string) (int, error) {
	var (
		err   error
		total int
	)

	rows, err := (*d.stmt)[getCountSearchReview].QueryxContext(ctx, "%"+reviewid+"%", "%"+reviewer+"%")
	if err != nil {
		return total, errors.Wrap(err, "[DATA] [GetCountSearchReview]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&total); err != nil {
			return total, errors.Wrap(err, "[DATA] [GetCountSearchReview]")
		}

	}
	return total, err
}

func (d Data) GetTableReviewByRating(ctx context.Context, rating int, page int, length int) ([]glodokEntity.TableReview, error) {
	var (
		review      glodokEntity.TableReview
		reviewArray []glodokEntity.TableReview
		err         error
	)

	rows, err := (*d.stmt)[getTableReviewByRating].QueryxContext(ctx, rating, page, length)
	if err != nil {
		return reviewArray, errors.Wrap(err, "[DATA] [GetTableReviewByRating]")
	}

	defer rows.Close()

	for rows.Next() {
		var reviewID string
		var reviewRating int
		var reviewerName string
		var reviewDesc string
		var reviewDateRaw []byte // Raw byte slice for date

		if err = rows.Scan(&reviewID, &reviewRating, &reviewerName, &reviewDesc, &reviewDateRaw); err != nil {
			return reviewArray, errors.Wrap(err, "[DATA] [GetTableReview]")
		}

		// Convert raw date to time.Time
		var reviewDate time.Time
		if reviewDateRaw != nil {
			reviewDate, err = time.Parse("2006-01-02 15:04:05", string(reviewDateRaw))
			if err != nil {
				return reviewArray, errors.Wrap(err, "[DATA] [GetTableReview] parsing date")
			}
		}

		review = glodokEntity.TableReview{
			ReviewID:     reviewID,
			ReviewRating: reviewRating,
			Reviewer:     reviewerName,
			ReviewDesc:   reviewDesc,
			ReviewDate:   reviewDate,
		}

		reviewArray = append(reviewArray, review)
	}
	return reviewArray, err
}

func (d Data) GetCountTableReviewByRating(ctx context.Context, rating int) (int, error) {
	var (
		err   error
		total int
	)

	rows, err := (*d.stmt)[getCountReviewByRating].QueryxContext(ctx, rating)
	if err != nil {
		return total, errors.Wrap(err, "[DATA] [GetCountTableReviewByRating]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&total); err != nil {
			return total, errors.Wrap(err, "[DATA] [GetCountTableReviewByRating]")
		}

	}
	return total, err
}

func (d Data) GetSearchReviewByRating(ctx context.Context, rating int, reviewid string, reviewer string, page int, length int) ([]glodokEntity.TableReview, error) {
	var (
		review      glodokEntity.TableReview
		reviewArray []glodokEntity.TableReview
		err         error
	)

	rows, err := (*d.stmt)[getSearchReviewByRating].QueryxContext(ctx, rating, "%"+reviewid+"%", "%"+reviewer+"%", page, length)
	if err != nil {
		return reviewArray, errors.Wrap(err, "[DATA] [GetSearchReviewByRating]")
	}

	defer rows.Close()

	for rows.Next() {
		var reviewID string
		var reviewRating int
		var reviewer string
		var reviewDesc string
		var reviewDateRaw []byte // Raw byte slice for date

		if err = rows.Scan(&reviewID, &reviewRating, &reviewer, &reviewDesc, &reviewDateRaw); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetSearchReview]")
		}

		//convert
		var reviewDate time.Time
		if reviewDateRaw != nil {
			reviewDate, err = time.Parse("2006-01-02 15:04:05", string(reviewDateRaw))
			if err != nil {
				return reviewArray, errors.Wrap(err, "[DATA] [GetTableReview] parsing date")
			}
		}

		review = glodokEntity.TableReview{
			ReviewID:     reviewID,
			ReviewRating: reviewRating,
			Reviewer:     reviewer,
			ReviewDesc:   reviewDesc,
			ReviewDate:   reviewDate,
		}

		reviewArray = append(reviewArray, review)
	}
	return reviewArray, err
}

func (d Data) GetCountSearchReviewByRating(ctx context.Context, rating int, reviewid string, reviewer string) (int, error) {
	var (
		err   error
		total int
	)

	rows, err := (*d.stmt)[getCountSearchReviewByRating].QueryxContext(ctx, rating, "%"+reviewid+"%", "%"+reviewer+"%")
	if err != nil {
		return total, errors.Wrap(err, "[DATA] [GetCountSearchReviewByRating]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&total); err != nil {
			return total, errors.Wrap(err, "[DATA] [GetCountSearchReviewByRating]")
		}

	}
	return total, err
}

//berita

func (d Data) GetDestinasi(ctx context.Context) ([]glodokEntity.TableDestinasi, error) {
	var (
		destinasi      glodokEntity.TableDestinasi
		destinasiArray []glodokEntity.TableDestinasi
		err            error
	)

	rows, err := (*d.stmt)[getDestinasi].QueryxContext(ctx)
	if err != nil {
		return destinasiArray, errors.Wrap(err, "[DATA] [GetDestinasi]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&destinasi); err != nil {
			return destinasiArray, errors.Wrap(err, "[DATA] [GetDestinasi]")
		}
		destinasiArray = append(destinasiArray, destinasi)
	}
	return destinasiArray, err

}

func (d Data) InsertBerita(ctx context.Context, berita glodokEntity.TableBerita) (string, error) {
	var (
		err    error
		result string
		lastID string
		newID  string
	)

	// Fetch the last inserted DestinasiID
	err = (*d.stmt)[fetchLastBeritaID].QueryRowxContext(ctx).Scan(&lastID)
	if err != nil && err != sql.ErrNoRows {
		result = "Gagal mengambil ID terakhir"
		return result, errors.Wrap(err, "[DATA][fetchLastBeritaID]")
	}

	// Generate the new DestinasiID
	if lastID != "" {
		// Extract the numeric part from lastID and increment it
		num, _ := strconv.Atoi(lastID[1:])
		newID = fmt.Sprintf("B%04d", num+1)

	} else {
		// If there are no previous records, start with D0001
		newID = "B0001"

	}

	// Assign the new DestinasiID
	berita.BeritaID = newID

	// Proceed with the insertion
	_, err = (*d.stmt)[insertBerita].ExecContext(ctx,
		berita.BeritaID,
		berita.DestinasiID,
		berita.BeritaJudul,
		berita.BeritaDesc,
		berita.BeritaGambar,
		berita.BeritaLinkSumber,
	)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][InsertBerita]")
	}

	result = "Berhasil"
	return result, nil
}

func (d Data) DeleteBerita(ctx context.Context, beritaid string) (string, error) {
	var (
		err    error
		result string
	)

	_, err = (*d.stmt)[deleteBerita].ExecContext(ctx, beritaid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][DeleteBerita]")
	}

	result = "Berhasil"

	return result, err
}

func (d Data) GetTableBerita(ctx context.Context, page int, length int) ([]glodokEntity.TableBerita, error) {
	var (
		berita      glodokEntity.TableBerita
		beritaArray []glodokEntity.TableBerita
		err         error
	)

	rows, err := (*d.stmt)[getTableBerita].QueryxContext(ctx, page, length)
	if err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetTableBerita] - Query failed")
	}
	defer rows.Close()

	// Ensure the directory exists
	imageDir := filepath.Join("public", "images")
	if err := EnsureDirectory(imageDir); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetTableBerita] - Failed to ensure directory")
	}

	for rows.Next() {
		var beritaID string
		var destinasiID string
		var destinasiName string
		var beritaJudul string
		var beritaDesc string
		var beritaGambarUrl string
		var beritaDateRaw []byte // Raw byte slice for date
		var beritaLinkSumber string

		if err = rows.Scan(&beritaID, &destinasiID, &destinasiName, &beritaJudul, &beritaDesc, &beritaGambarUrl, &beritaDateRaw, &beritaLinkSumber); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetTableBerita] - Failed to scan row")
		}

		// Convert raw date to time.Time
		var beritaDate time.Time
		if beritaDateRaw != nil {
			beritaDate, err = time.Parse("2006-01-02 15:04:05", string(beritaDateRaw))
			if err != nil {
				return beritaArray, errors.Wrap(err, "[DATA] [GetTableReview] parsing date")
			}
		}

		// Save image and generate URL
		filePath := filepath.Join(imageDir, berita.BeritaID+".jpg")
		if err := saveImageToFile(berita.BeritaGambar, filePath); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetTableBerita] - Failed to save image")
		}

		berita.BeritaGambarURL = generateImageURLBerita(beritaID)

		berita = glodokEntity.TableBerita{
			BeritaID:         beritaID,
			DestinasiID:      destinasiID,
			DestinasiName:    destinasiName,
			BeritaJudul:      beritaJudul,
			BeritaDesc:       beritaDesc,
			BeritaGambarURL:  berita.BeritaGambarURL,
			BeritaDate:       beritaDate,
			BeritaLinkSumber: beritaLinkSumber,
		}

		beritaArray = append(beritaArray, berita)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetTableBerita] - Row iteration error")
	}

	return beritaArray, nil
}

func (d Data) GetImageBerita(ctx context.Context, beritaid string) ([]byte, error) {
	var image []byte
	if err := (*d.stmt)[getImageBerita].QueryRowxContext(ctx, beritaid).Scan(&image); err != nil {
		return image, errors.Wrap(err, "[DATA][GetImageBerita]")
	}

	return image, nil
}

func (d Data) GetCountBerita(ctx context.Context) (int, error) {
	var (
		err   error
		total int
	)

	rows, err := (*d.stmt)[getCountTableBerita].QueryxContext(ctx)
	if err != nil {
		return total, errors.Wrap(err, "[DATA] [GetCountBerita]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&total); err != nil {
			return total, errors.Wrap(err, "[DATA] [GetCountBerita]")
		}

	}
	return total, err
}

func (d Data) UpdateBerita(ctx context.Context, berita glodokEntity.TableBerita, beritaid string) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[updateBerita].ExecContext(ctx,
		berita.DestinasiID,
		berita.BeritaJudul,
		berita.BeritaDesc,
		berita.BeritaGambar,
		berita.BeritaLinkSumber,
		beritaid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][UpdateBerita]")
	}

	result = "Berhasil"
	return result, nil
}

func (d Data) GetSearchBerita(ctx context.Context, beritaid string, destinasiname string, beritajudul string, page int, length int) ([]glodokEntity.TableBerita, error) {
	var (
		berita      glodokEntity.TableBerita
		beritaArray []glodokEntity.TableBerita
		err         error
	)

	rows, err := (*d.stmt)[getSearchBerita].QueryxContext(ctx, "%"+beritaid+"%", "%"+destinasiname+"%", "%"+beritajudul+"%", page, length)
	if err != nil {
		return beritaArray, errors.Wrap(err, "[DATA] [GetSearchBerita]")
	}

	defer rows.Close()

	// Ensure the directory exists
	imageDir := filepath.Join("public", "images")
	if err := EnsureDirectory(imageDir); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetSearchBerita] - Failed to ensure directory")
	}

	for rows.Next() {
		var beritaID string
		var destinasiID string
		var destinasiName string
		var beritaJudul string
		var beritaDesc string
		var beritaGambarUrl string
		var beritaDateRaw []byte // Raw byte slice for date
		var beritaLinkSumber string

		if err = rows.Scan(&beritaID, &destinasiID, &destinasiName, &beritaJudul, &beritaDesc, &beritaGambarUrl, &beritaDateRaw, &beritaLinkSumber); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetTableBerita] - Failed to scan row")
		}

		// Convert raw date to time.Time
		var beritaDate time.Time
		if beritaDateRaw != nil {
			beritaDate, err = time.Parse("2006-01-02 15:04:05", string(beritaDateRaw))
			if err != nil {
				return beritaArray, errors.Wrap(err, "[DATA] [GetTableReview] parsing date")
			}
		}

		// Save image and generate URL
		filePath := filepath.Join(imageDir, berita.BeritaID+".jpg")
		if err := saveImageToFile(berita.BeritaGambar, filePath); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetTableBerita] - Failed to save image")
		}

		berita.BeritaGambarURL = generateImageURLBerita(beritaID)

		berita = glodokEntity.TableBerita{
			BeritaID:         beritaID,
			DestinasiID:      destinasiID,
			DestinasiName:    destinasiName,
			BeritaJudul:      beritaJudul,
			BeritaDesc:       beritaDesc,
			BeritaGambarURL:  berita.BeritaGambarURL,
			BeritaDate:       beritaDate,
			BeritaLinkSumber: beritaLinkSumber,
		}

		beritaArray = append(beritaArray, berita)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetSearchBerita] - Row iteration error")
	}

	return beritaArray, nil

}

func (d Data) GetCountSearchBerita(ctx context.Context, beritaid string, destinasiname string, beritajudul string) (int, error) {
	var (
		err   error
		total int
	)

	rows, err := (*d.stmt)[getCountSearchBerita].QueryxContext(ctx, "%"+beritaid+"%", "%"+destinasiname+"%", "%"+beritajudul+"%")
	if err != nil {
		return total, errors.Wrap(err, "[DATA] [GetCountSearchBerita]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&total); err != nil {
			return total, errors.Wrap(err, "[DATA] [GetCountSearchBerita]")
		}

	}
	return total, err
}

//jenis destinasi
func (d Data) InsertJenisDestinasi(ctx context.Context, jenisdestinasi glodokEntity.TableJenisDestinasi) (string, error) {
	var (
		err    error
		result string
		lastID string
		newID  string
	)

	err = (*d.stmt)[fetchJenisDestinasiID].QueryRowxContext(ctx).Scan(&lastID)
	if err != nil && err != sql.ErrNoRows {
		result = "Gagal mengambil ID terakhir"
		return result, errors.Wrap(err, "[DATA][InsertJenisDestinasi]")
	}

	if lastID != "" {
		// Extract the numeric part from lastID and increment it
		num, _ := strconv.Atoi(lastID[1:])
		newID = fmt.Sprintf("J%04d", num+1)

	} else {
		newID = "J0001"

	}

	jenisdestinasi.JenisDestinasiID = newID

	// Proceed with the insertion
	_, err = (*d.stmt)[insertJenisDestinasi].ExecContext(ctx,
		jenisdestinasi.JenisDestinasiID,
		jenisdestinasi.JenisDestinasiKat,
	)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][InsertJenisDestinasi]")
	}

	result = "Berhasil"
	return result, nil
}

func (d Data) GetTableJenisDestinasi(ctx context.Context, jenisdestinasiid string, jenisdestinasiket string, page int, length int) ([]glodokEntity.TableJenisDestinasi, error) {
	var (
		jenisDestinasi      glodokEntity.TableJenisDestinasi
		jenisDestinasiArray []glodokEntity.TableJenisDestinasi
		err                   error
	)

	rows, err := (*d.stmt)[getTableJenisDestinasi].QueryxContext(ctx, "%"+jenisdestinasiid+"%", "%"+jenisdestinasiket+"%", page, length)
	fmt.Println("pagelength", page, length)
	if err != nil {
		return jenisDestinasiArray, errors.Wrap(err, "[DATA] [GetTableJenisDestinasi]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&jenisDestinasi); err != nil {
			return jenisDestinasiArray, errors.Wrap(err, "[DATA] [GetTableJenisDestinasi]")
		}
		jenisDestinasiArray = append(jenisDestinasiArray, jenisDestinasi)
	}
	return jenisDestinasiArray, err
}

func (d Data) GetCountTableJenisDestinasi(ctx context.Context, jenisdestinasiid string, jenisdestinasiket string) (int, error) {
	var (
		err   error
		total int
	)

	rows, err := (*d.stmt)[getCountTableJenisDestinasi].QueryxContext(ctx, "%"+jenisdestinasiid+"%", "%"+jenisdestinasiket+"%")
	if err != nil {
		return total, errors.Wrap(err, "[DATA] [GetCountTableJenisDestinasi]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&total); err != nil {
			return total, errors.Wrap(err, "[DATA] [GetCountSearchTipeTransportasi]")
		}

	}
	return total, err
}

func (d Data) UpdateJenisDestinasi(ctx context.Context, jenisdestinasi glodokEntity.TableJenisDestinasi, jenisdestinasiid string) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[updateJenisDestinasi].ExecContext(ctx, jenisdestinasi.JenisDestinasiKat, jenisdestinasiid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][UpdateJenisDestinasi]")
	}

	result = "Berhasil"
	return result, err
}

func (d Data) DeleteJenisDestinasi(ctx context.Context, destinasiid string) (string, error) {
	var (
		err    error
		result string
	)

	_, err = (*d.stmt)[deleteJenisDestinasi].ExecContext(ctx, destinasiid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][DeleteJenisDestinasi]")
	}

	result = "Berhasil"

	return result, err
}

// FOR MASYARAKAT
func (d Data) GetDestinasiByID(ctx context.Context, destinasiid string) ([]glodokEntity.TableDestinasi, error) {
	var (
		// destinasi      glodokEntity.TableDestinasi
		destinasiArray []glodokEntity.TableDestinasi
		err            error
	)

	rows, err := (*d.stmt)[getDestinasiByID].QueryxContext(ctx, destinasiid)
	if err != nil {
		return destinasiArray, errors.Wrap(err, "[DATA] [GetDestinasiByID]")
	}

	defer rows.Close()

	// Ensure the directory exists
	imageDir := filepath.Join("public", "images")
	if err := EnsureDirectory(imageDir); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetDestinasiByID] - Failed to ensure directory")
	}

	for rows.Next() {
		var destinasi glodokEntity.TableDestinasi
		var jbukaStr, jtutupStr string

		if err = rows.Scan(&destinasi.DestinasiID, &destinasi.DestinasiName, &destinasi.DestinasiDesc,
			&destinasi.DestinasiAlamat, &destinasi.DestinasiGambarURL,
			&destinasi.DestinasiLang, &destinasi.DestinasiLong, &destinasi.DestinasiHBuka,
			&destinasi.DestinasiHTutup, &jbukaStr, &jtutupStr, &destinasi.DestinasiKet,
			&destinasi.DestinasiHalal); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetDestinasiByID] - Failed to scan row")
		}

		if jbukaStr != "" {
			destinasi.DestinasiJBuka, err = time.Parse("15:04:05", jbukaStr)
			if err != nil {
				return nil, errors.Wrap(err, "[DATA] [GetDestinasiByID] - Failed to parse destinasi_jbuka")
			}
		}

		if jtutupStr != "" {
			destinasi.DestinasiJTutup, err = time.Parse("15:04:05", jtutupStr)
			if err != nil {
				return nil, errors.Wrap(err, "[DATA] [GetDestinasiByID] - Failed to parse destinasi_jtutup")
			}
		}

		// Save image and generate URL
		filePath := filepath.Join(imageDir, destinasi.DestinasiID+".jpg")
		if err := saveImageToFile(destinasi.DestinasiGambar, filePath); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetDestinasiByID] - Failed to save image")
		}

		destinasi.DestinasiGambarURL = generateImageURL(destinasi.DestinasiID, destinasi.DestinasiKet)
		destinasiArray = append(destinasiArray, destinasi)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetDestinasiByID] - Row iteration error")
	}

	return destinasiArray, nil

}

func (d Data) GetAllDestinasi(ctx context.Context, kategori string, labelhalal string, destinasiname string) ([]glodokEntity.TableDestinasi, error) {
	var (
		// destinasi      glodokEntity.TableDestinasi
		destinasiArray []glodokEntity.TableDestinasi
		err            error
	)

	rows, err := (*d.stmt)[getAllDestinasi].QueryxContext(ctx, kategori, "%"+labelhalal+"%", "%"+destinasiname+"%")
	if err != nil {
		return destinasiArray, errors.Wrap(err, "[DATA] [GetAllDestinasi]")
	}

	defer rows.Close()

	// Ensure the directory exists
	imageDir := filepath.Join("public", "images")
	if err := EnsureDirectory(imageDir); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetAllDestinasi] - Failed to ensure directory")
	}

	for rows.Next() {
		var destinasi glodokEntity.TableDestinasi
		var jbukaStr, jtutupStr string

		if err = rows.Scan(&destinasi.DestinasiID, &destinasi.DestinasiName, &destinasi.DestinasiDesc,
			&destinasi.DestinasiAlamat, &destinasi.DestinasiGambarURL,
			&destinasi.DestinasiLang, &destinasi.DestinasiLong, &destinasi.DestinasiHBuka,
			&destinasi.DestinasiHTutup, &jbukaStr, &jtutupStr, &destinasi.DestinasiKet,
			&destinasi.DestinasiHalal); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetAllDestinasi] - Failed to scan row")
		}

		if jbukaStr != "" {
			destinasi.DestinasiJBuka, err = time.Parse("15:04:05", jbukaStr)
			if err != nil {
				return nil, errors.Wrap(err, "[DATA] [GetAllDestinasi] - Failed to parse destinasi_jbuka")
			}
		}

		if jtutupStr != "" {
			destinasi.DestinasiJTutup, err = time.Parse("15:04:05", jtutupStr)
			if err != nil {
				return nil, errors.Wrap(err, "[DATA] [GetAllDestinasi] - Failed to parse destinasi_jtutup")
			}
		}

		// Save image and generate URL
		filePath := filepath.Join(imageDir, destinasi.DestinasiID+".jpg")
		if err := saveImageToFile(destinasi.DestinasiGambar, filePath); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetAllDestinasi] - Failed to save image")
		}

		destinasi.DestinasiGambarURL = generateImageURL(destinasi.DestinasiID, destinasi.DestinasiKet)
		destinasiArray = append(destinasiArray, destinasi)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetAllDestinasi] - Row iteration error")
	}

	return destinasiArray, nil

}

func (d Data) GetAllReview(ctx context.Context, rating string, page int, length int) ([]glodokEntity.TableReview, error) {
	var (
		review      glodokEntity.TableReview
		reviewArray []glodokEntity.TableReview
		err         error
	)

	// // Convert the integer rating to a string before concatenation
	// ratingStr := strconv.Itoa(rating)

	rows, err := (*d.stmt)[getAllReview].QueryxContext(ctx, "%"+rating+"%", page, length)
	if err != nil {
		return reviewArray, errors.Wrap(err, "[DATA] [GetAllReview]")
	}

	defer rows.Close()

	for rows.Next() {
		var reviewID string
		var reviewRating int
		var reviewerName string
		var reviewDesc string
		var reviewDateRaw []byte // Raw byte slice for date

		if err = rows.Scan(&reviewID, &reviewRating, &reviewerName, &reviewDesc, &reviewDateRaw); err != nil {
			return reviewArray, errors.Wrap(err, "[DATA] [GetAllReview]")
		}

		// Convert raw date to time.Time
		var reviewDate time.Time
		if reviewDateRaw != nil {
			reviewDate, err = time.Parse("2006-01-02 15:04:05", string(reviewDateRaw))
			if err != nil {
				return reviewArray, errors.Wrap(err, "[DATA] [GetAllReview] parsing date")
			}
		}

		review = glodokEntity.TableReview{
			ReviewID:     reviewID,
			ReviewRating: reviewRating,
			Reviewer:     reviewerName,
			ReviewDesc:   reviewDesc,
			ReviewDate:   reviewDate,
		}

		reviewArray = append(reviewArray, review)
	}
	return reviewArray, err
}

func (d Data) GetCountAllReview(ctx context.Context, rating string) (int, error) {
	var (
		err   error
		total int
	)

	// // Convert the integer rating to a string before concatenation
	// ratingStr := strconv.Itoa(rating)

	rows, err := (*d.stmt)[getCountAllReview].QueryxContext(ctx, "%"+rating+"%")
	if err != nil {
		return total, errors.Wrap(err, "[DATA] [GetCountAllReview]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&total); err != nil {
			return total, errors.Wrap(err, "[DATA] [GetCountAllReview]")
		}

	}
	return total, err
}
