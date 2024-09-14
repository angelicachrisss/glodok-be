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

func (s Service) InsertDestinasi(ctx context.Context, destinasi glodokEntity.TableDestinasi) (string, error) {

	var (
		result string
	)
	result, err := s.glodok.InsertDestinasi(ctx, destinasi)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][InsertDestinasi]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) GetTableDestinasi(ctx context.Context, ket string, page int, length int) ([]glodokEntity.TableDestinasi, interface{}, error) {
	var (
		total int
		// destinasiArrayFinish []glodokEntity.TableDestinasi
	)

	metadata := make(map[string]interface{})

	offset := page * length
	destinasiArray, err := s.glodok.GetTableDestinasi(ctx, ket, offset, length)

	if err != nil {
		return destinasiArray, metadata, errors.Wrap(err, "[Service][GetTableDestinasi]")
	}

	total, err = s.glodok.GetCountDestinasi(ctx, ket)

	if err != nil {
		return destinasiArray, metadata, errors.Wrap(err, "[Service][GetCountDestinasi]")
	}
	metadata["total_data"] = total

	return destinasiArray, metadata, nil
}

func (s Service) DeleteDestinasi(ctx context.Context, destinasiid string) (string, error) {

	var (
		result string
	)
	_, err := s.glodok.DeleteDestinasi(ctx, destinasiid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][DeleteDestinasi]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) InsertTipeTransportasi(ctx context.Context, tipetransportasi glodokEntity.TableTipeTransportasi) (string, error) {

	var (
		result string
	)
	result, err := s.glodok.InsertTipeTransportasi(ctx, tipetransportasi)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][InsertTipeTransportasi]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) GetTableTipeTransportasi(ctx context.Context, page int, length int) ([]glodokEntity.TableTipeTransportasi, interface{}, error) {

	var (
		total int
	)

	metadata := make(map[string]interface{})

	offset := page * length
	tipeTransportasiArray, err := s.glodok.GetTableTipeTransportasi(ctx, offset, length)

	if err != nil {
		return tipeTransportasiArray, metadata, errors.Wrap(err, "[Service][GetTableTipeTransportasi]")
	}

	total, err = s.glodok.GetCountTableTipeTransportasi(ctx)

	if err != nil {
		return tipeTransportasiArray, metadata, errors.Wrap(err, "[Service][GetCountTableTipeTransportasi]")
	}
	metadata["total_data"] = total

	return tipeTransportasiArray, metadata, nil

}

func (s Service) GetSearchTipeTransportasi(ctx context.Context, tipetransportasiid string, tipetransportasiname string, page int, length int) ([]glodokEntity.TableTipeTransportasi, interface{}, error) {
	var (
		total int
	)

	metadata := make(map[string]interface{})

	offset := page * length
	fmt.Println("offset: ", offset)
	searchListTipeTransportasiArray, err := s.glodok.GetSearchTipeTransportasi(ctx, tipetransportasiid, tipetransportasiname, offset, length)

	if err != nil {
		return searchListTipeTransportasiArray, metadata, errors.Wrap(err, "[Service][GetSearchTipeTransportasi]")
	}

	total, err = s.glodok.GetCountSearchTipeTransportasi(ctx, tipetransportasiid, tipetransportasiname)

	if err != nil {
		return searchListTipeTransportasiArray, metadata, errors.Wrap(err, "[Service][GetCountSearchTipeTransportasi]")
	}
	metadata["total_data"] = total

	return searchListTipeTransportasiArray, metadata, nil
}

func (s Service) DeleteTipeTransportasi(ctx context.Context, tipetransportasiid string) (string, error) {

	var (
		result string
	)
	_, err := s.glodok.DeleteTipeTransportasi(ctx, tipetransportasiid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][DeleteTipeTransportasi]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) UpdateTipeTransportasi(ctx context.Context, tipetransportasi glodokEntity.TableTipeTransportasi, tipetransportasiid string) (string, error) {
	var (
		result string
	)
	result, err := s.glodok.UpdateTipeTransportasi(ctx, tipetransportasi, tipetransportasiid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][UpdateTipeTransportasi]")
	}

	result = "Berhasil"

	return result, err
}

//rutetransportasi

func (s Service) GetTipeTransportasi(ctx context.Context) ([]glodokEntity.TableTipeTransportasi, error) {

	tipeTransportasiArray, err := s.glodok.GetTipeTransportasi(ctx)

	if err != nil {
		return tipeTransportasiArray, errors.Wrap(err, "[Service][GetTipeTransportasi]")
	}

	return tipeTransportasiArray, nil

}

func (s Service) InsertRuteTransportasi(ctx context.Context, rutetransportasi glodokEntity.TableRuteTransportasi) (string, error) {

	var (
		result string
	)
	result, err := s.glodok.InsertRuteTransportasi(ctx, rutetransportasi)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][InsertRuteTransportasi]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) GetTableRuteTransportasi(ctx context.Context, page int, length int) ([]glodokEntity.TableRuteTransportasi, interface{}, error) {

	var (
		total int
	)

	metadata := make(map[string]interface{})

	offset := page * length
	ruteTransportasiArray, err := s.glodok.GetTableRuteTransportasi(ctx, offset, length)

	if err != nil {
		return ruteTransportasiArray, metadata, errors.Wrap(err, "[Service][GetTableRuteTransportasi]")
	}

	total, err = s.glodok.GetCountTableRuteTransportasi(ctx)

	if err != nil {
		return ruteTransportasiArray, metadata, errors.Wrap(err, "[Service][GetCountTableRuteTransportasi]")
	}
	metadata["total_data"] = total

	return ruteTransportasiArray, metadata, nil

}

func (s Service) DeleteRuteTransportasi(ctx context.Context, ruteid string) (string, error) {

	var (
		result string
	)
	_, err := s.glodok.DeleteRuteTransportasi(ctx, ruteid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][DeleteRuteTransportasi]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) GetSearchRuteTransportasi(ctx context.Context, tipetransportasiname string, tujuanawal string, tujuanakhir string, page int, length int) ([]glodokEntity.TableRuteTransportasi, interface{}, error) {
	var (
		total int
	)

	metadata := make(map[string]interface{})

	offset := page * length
	fmt.Println("offset: ", offset)
	searchListRuteTransportasiArray, err := s.glodok.GetSearchRuteTransportasi(ctx, tipetransportasiname, tujuanawal, tujuanakhir, offset, length)

	if err != nil {
		return searchListRuteTransportasiArray, metadata, errors.Wrap(err, "[Service][GetSearchTipeTransportasi]")
	}

	total, err = s.glodok.GetCountSearchRuteTransportasi(ctx, tipetransportasiname, tujuanawal, tujuanakhir)

	if err != nil {
		return searchListRuteTransportasiArray, metadata, errors.Wrap(err, "[Service][GetCountSearchRuteTransportasi]")
	}
	metadata["total_data"] = total

	return searchListRuteTransportasiArray, metadata, nil
}

func (s Service) UpdateRuteTransportasi(ctx context.Context, rutetransportasi glodokEntity.TableRuteTransportasi, ruteid string) (string, error) {
	var (
		result string
	)
	result, err := s.glodok.UpdateRuteTransportasi(ctx, rutetransportasi, ruteid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][UpdateRuteTransportasi]")
	}

	result = "Berhasil"

	return result, err
}

