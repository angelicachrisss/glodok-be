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

func (s Service) GetAdmin(ctx context.Context) ([]glodokEntity.GetAdmin, error) {

	adminArray, err := s.glodok.GetAdmin(ctx)

	if err != nil {
		return adminArray, errors.Wrap(err, "[Service][GetAdmin]")
	}

	return adminArray, nil

}

func (s Service) InsertAdmin(ctx context.Context, admin glodokEntity.GetAdmin) (string, error) {

	var (
		result string
	)
	result, err := s.glodok.InsertAdmin(ctx, admin)

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

func (s Service) GetTableAdmin(ctx context.Context, page int, length int) ([]glodokEntity.GetAdmin, interface{}, error) {

	var (
		total int
	)

	metadata := make(map[string]interface{})

	offset := page * length
	adminArray, err := s.glodok.GetTableAdmin(ctx, offset, length)

	if err != nil {
		return adminArray, metadata, errors.Wrap(err, "[Service][GetTableAdmin]")
	}

	total, err = s.glodok.GetCountAdmin(ctx)

	if err != nil {
		return adminArray, metadata, errors.Wrap(err, "[Service][GetCountAdmin]")
	}
	metadata["total_data"] = total

	return adminArray, metadata, nil

}

func (s Service) GetSearchAdmin(ctx context.Context, adminid string, page int, length int) ([]glodokEntity.GetAdmin, interface{}, error) {
	var (
		total int
	)

	metadata := make(map[string]interface{})

	offset := page * length
	fmt.Println("offset: ", offset)
	searchListAdminArray, err := s.glodok.GetSearchAdmin(ctx, adminid, offset, length)

	if err != nil {
		return searchListAdminArray, metadata, errors.Wrap(err, "[Service][GetSearchAdmin]")
	}

	total, err = s.glodok.GetCountSearchAdmin(ctx, adminid)

	if err != nil {
		return searchListAdminArray, metadata, errors.Wrap(err, "[Service][GetSearchAdmin]")
	}
	metadata["total_data"] = total

	return searchListAdminArray, metadata, nil
}

func (s Service) DeleteAdmin(ctx context.Context, adminid string) (string, error) {

	var (
		result string
	)
	_, err := s.glodok.DeleteAdmin(ctx, adminid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][DeleteAdmin]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) UpdateAdmin(ctx context.Context, admin glodokEntity.GetAdmin, adminid string) (string, error) {
	var (
		result string
	)
	result, err := s.glodok.UpdateAdmin(ctx, admin, adminid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][UpdateAdmin]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) InsertDestinasiIc(ctx context.Context, destinasi glodokEntity.TableDestinasiIc) (string, error) {

	var (
		result string
	)
	result, err := s.glodok.InsertDestinasiIc(ctx, destinasi)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][InsertDestinasiIc]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) GetTableDestinasiIc(ctx context.Context, page int, length int) ([]glodokEntity.TableDestinasiIc, interface{}, error) {

	var (
		total int
	)

	metadata := make(map[string]interface{})

	offset := page * length
	destinasiArray, err := s.glodok.GetTableDestinasiIc(ctx, offset, length)

	if err != nil {
		return destinasiArray, metadata, errors.Wrap(err, "[Service][GetTableDestinasiIc]")
	}

	total, err = s.glodok.GetCounDestinasiIc(ctx)

	if err != nil {
		return destinasiArray, metadata, errors.Wrap(err, "[Service][GetTableDestinasiIc]")
	}
	metadata["total_data"] = total

	return destinasiArray, metadata, nil

}
