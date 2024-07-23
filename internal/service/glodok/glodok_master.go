package glodok

import (
	"context"
	"fmt"
	glodokEntity "glodok-be/internal/entity/glodok"
	"glodok-be/pkg/errors"
	// "encoding/json"
	// "fmt"
	// "log"
	// "strconv"
	// "time"
)

// func (s Service) GetKaryawan(ctx context.Context) ([]JoEntity.GetKaryawan, interface{}, error) {
// 	var (
// 		total int
// 	)

// 	metadata := make(map[string]interface{})
// 	karyawanArray, err := s.glodok.GetKaryawan(ctx)

// 	if err != nil {
// 		return karyawanArray, metadata, errors.Wrap(err, "[Service][GetKaryawan]")
// 	}

// 	total, err = s.glodok.GetCountKaryawan(ctx)

// 	if err != nil {
// 		return karyawanArray, metadata, errors.Wrap(err, "[Service][GetCountKaryawan]")
// 	}
// 	metadata["total_data"] = total

// 	return karyawanArray, metadata, nil

// }

// func (s Service) InsertKaryawan(ctx context.Context, karyawan JoEntity.InsertKaryawan) (string, error) {

// 	var (
// 		result string
// 	)
// 	result, err := s.glodok.InsertKaryawan(ctx, karyawan.Insertkaryawan)

// 	if err != nil {
// 		result = "Gagal"
// 		return result, errors.Wrap(err, "[Service][InsertKaryawan]")
// 	}

// 	result = "Berhasil"

// 	return result, err
// }

func (s Service) GetAdmin(ctx context.Context) ([]glodokEntity.GetAdmin, error) {

	adminArray, err := s.glodok.GetAdmin(ctx)

	if err != nil {
		return adminArray, errors.Wrap(err, "[Service][GetAdmin]")
	}

	return adminArray, nil

}

func (s Service) InsertAdmin(ctx context.Context, admin glodokEntity.InsertAdmin) (string, error) {

	var (
		result string
	)
	result, err := s.glodok.InsertAdmin(ctx, admin.InsertAdmin)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][InsertKaryawan]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) SubmitLogin(ctx context.Context, adminid string, adminpass string) (string, error) {
	var (
		result string
	)

	result, err := s.glodok.SubmitLogin(ctx, adminid, adminpass)

	if err != nil {
		result = "Gagal Login"
		fmt.Println("result", result)
		return result, errors.Wrap(err, "[Service][SubmitLogin]")
	}
	result = "Berhasil Login"

	return result, err
}

func (s Service) GetAdminbyID(ctx context.Context, adminid string) ([]glodokEntity.GetAdmin, error) {

	adminArray, err := s.glodok.GetAdminbyID(ctx, adminid)

	if err != nil {
		return adminArray, errors.Wrap(err, "[Service][GetAdminbyID]")
	}

	return adminArray, nil

}