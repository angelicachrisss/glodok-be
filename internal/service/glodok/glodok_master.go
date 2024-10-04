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

func (s Service) GetImageDestinasi(ctx context.Context, destinasiid string, destinasikat string) ([]byte, error) {
	image, err := s.glodok.GetImageDestinasi(ctx, destinasiid, destinasikat)
	if err != nil {
		return image, errors.Wrap(err, "[SERVICE][GetImageDestinasi]")
	}

	return image, err
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

func (s Service) GetSearchDestinasi(ctx context.Context, kategori string, destinasiid string, destinasiname string, page int, length int) ([]glodokEntity.TableDestinasi, interface{}, error) {
	var (
		total int
	)

	metadata := make(map[string]interface{})

	offset := page * length
	fmt.Println("offset: ", offset)
	searchListDestinasiArray, err := s.glodok.GetSearchDestinasi(ctx, kategori, destinasiid, destinasiname, offset, length)

	if err != nil {
		return searchListDestinasiArray, metadata, errors.Wrap(err, "[Service][GetSearchDestinasi]")
	}

	total, err = s.glodok.GetCountSearchDestinasi(ctx, kategori, destinasiid, destinasiname)

	if err != nil {
		return searchListDestinasiArray, metadata, errors.Wrap(err, "[Service][GetCountSearchDestinasi]")
	}
	metadata["total_data"] = total

	return searchListDestinasiArray, metadata, nil
}

func (s Service) UpdateDestinasi(ctx context.Context, destinasi glodokEntity.TableDestinasi, destinasiid string) (string, error) {
	var (
		result string
	)
	result, err := s.glodok.UpdateDestinasi(ctx, destinasi, destinasiid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][UpdateDestinasi]")
	}

	result = "Berhasil"

	return result, err
}

//tipe transportasi
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

// review
func (s Service) InsertReview(ctx context.Context, review glodokEntity.TableReview) (string, error) {

	var (
		result string
	)
	result, err := s.glodok.InsertReview(ctx, review)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][InsertReview]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) GetTableReview(ctx context.Context, page int, length int) ([]glodokEntity.TableReview, interface{}, error) {

	var (
		total int
	)

	metadata := make(map[string]interface{})

	offset := page * length
	reviewArray, err := s.glodok.GetTableReview(ctx, offset, length)

	if err != nil {
		return reviewArray, metadata, errors.Wrap(err, "[Service][GetTableReview]")
	}

	total, err = s.glodok.GetCountTableReview(ctx)

	if err != nil {
		return reviewArray, metadata, errors.Wrap(err, "[Service][GetCountTableReview]")
	}
	metadata["total_data"] = total

	return reviewArray, metadata, nil

}

func (s Service) DeleteReview(ctx context.Context, reviewid string) (string, error) {

	var (
		result string
	)
	_, err := s.glodok.DeleteReview(ctx, reviewid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][DeleteReview]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) GetSearchReview(ctx context.Context, reviewid string, reviewer string, page int, length int) ([]glodokEntity.TableReview, interface{}, error) {
	var (
		total int
	)

	metadata := make(map[string]interface{})

	offset := page * length
	searchListReviewArray, err := s.glodok.GetSearchReview(ctx, reviewid, reviewer, offset, length)

	if err != nil {
		return searchListReviewArray, metadata, errors.Wrap(err, "[Service][GetSearchReview]")
	}

	total, err = s.glodok.GetCountSearchReview(ctx, reviewid, reviewer)

	if err != nil {
		return searchListReviewArray, metadata, errors.Wrap(err, "[Service][GetCountSearchReview]")
	}
	metadata["total_data"] = total

	return searchListReviewArray, metadata, nil
}

func (s Service) GetTableReviewByRating(ctx context.Context, rating int, page int, length int) ([]glodokEntity.TableReview, interface{}, error) {

	var (
		total int
	)

	metadata := make(map[string]interface{})

	offset := page * length
	reviewArray, err := s.glodok.GetTableReviewByRating(ctx, rating, offset, length)

	if err != nil {
		return reviewArray, metadata, errors.Wrap(err, "[Service][GetTableReviewByRating]")
	}

	total, err = s.glodok.GetCountTableReviewByRating(ctx, rating)

	if err != nil {
		return reviewArray, metadata, errors.Wrap(err, "[Service][GetCountTableReviewByRating]")
	}
	metadata["total_data"] = total

	return reviewArray, metadata, nil

}

func (s Service) GetSearchReviewByRating(ctx context.Context, rating int, reviewid string, reviewer string, page int, length int) ([]glodokEntity.TableReview, interface{}, error) {
	var (
		total int
	)

	metadata := make(map[string]interface{})

	offset := page * length
	searchListReviewArray, err := s.glodok.GetSearchReviewByRating(ctx, rating, reviewid, reviewer, offset, length)

	if err != nil {
		return searchListReviewArray, metadata, errors.Wrap(err, "[Service][GetSearchReviewByRating]")
	}

	total, err = s.glodok.GetCountSearchReviewByRating(ctx, rating, reviewid, reviewer)

	if err != nil {
		return searchListReviewArray, metadata, errors.Wrap(err, "[Service][GetCountSearchReviewByRating]")
	}
	metadata["total_data"] = total

	return searchListReviewArray, metadata, nil
}

// berita
func (s Service) GetDestinasi(ctx context.Context) ([]glodokEntity.TableDestinasi, error) {

	destinasiArray, err := s.glodok.GetDestinasi(ctx)

	if err != nil {
		return destinasiArray, errors.Wrap(err, "[Service][GetDestinasi]")
	}

	return destinasiArray, nil

}

func (s Service) InsertBerita(ctx context.Context, berita glodokEntity.TableBerita) (string, error) {

	var (
		result string
	)
	result, err := s.glodok.InsertBerita(ctx, berita)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][InsertBerita]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) DeleteBerita(ctx context.Context, beritaid string) (string, error) {

	var (
		result string
	)
	_, err := s.glodok.DeleteBerita(ctx, beritaid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][DeleteBerita]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) GetImageBerita(ctx context.Context, beritaid string) ([]byte, error) {
	image, err := s.glodok.GetImageBerita(ctx, beritaid)
	if err != nil {
		return image, errors.Wrap(err, "[SERVICE][GetImageBerita]")
	}

	return image, err
}

func (s Service) GetTableBerita(ctx context.Context, page int, length int) ([]glodokEntity.TableBerita, interface{}, error) {
	var (
		total int
	)

	metadata := make(map[string]interface{})

	offset := page * length
	beritaArray, err := s.glodok.GetTableBerita(ctx, offset, length)

	if err != nil {
		return beritaArray, metadata, errors.Wrap(err, "[Service][GetTableBerita]")
	}

	total, err = s.glodok.GetCountBerita(ctx)

	if err != nil {
		return beritaArray, metadata, errors.Wrap(err, "[Service][GetTableBerita]")
	}
	metadata["total_data"] = total

	return beritaArray, metadata, nil
}

func (s Service) UpdateBerita(ctx context.Context, berita glodokEntity.TableBerita, beritaid string) (string, error) {
	var (
		result string
	)
	result, err := s.glodok.UpdateBerita(ctx, berita, beritaid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][UpdateBerita]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) GetSearchBerita(ctx context.Context, beritaid string, destinasiname string, beritajudul string, page int, length int) ([]glodokEntity.TableBerita, interface{}, error) {
	var (
		total int
	)

	metadata := make(map[string]interface{})

	offset := page * length
	searchListDestinasiArray, err := s.glodok.GetSearchBerita(ctx, beritaid, destinasiname, beritajudul, offset, length)

	if err != nil {
		return searchListDestinasiArray, metadata, errors.Wrap(err, "[Service][GetSearchBeria]")
	}

	total, err = s.glodok.GetCountSearchBerita(ctx, beritaid, destinasiname, beritajudul)

	if err != nil {
		return searchListDestinasiArray, metadata, errors.Wrap(err, "[Service][GetCountSearchBerita]")
	}
	metadata["total_data"] = total

	return searchListDestinasiArray, metadata, nil
}

//jenis destinasi
func (s Service) InsertJenisDestinasi(ctx context.Context, jenisdestinasi glodokEntity.TableJenisDestinasi) (string, error) {

	var (
		result string
	)
	result, err := s.glodok.InsertJenisDestinasi(ctx, jenisdestinasi)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][InsertJenisDestinasi]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) GetTableJenisDestinasi(ctx context.Context, jenisdestinasiid string, jenisdestinasiket string, page int, length int) ([]glodokEntity.TableJenisDestinasi, interface{}, error) {
	var (
		total int
	)

	metadata := make(map[string]interface{})

	offset := page * length
	listJenisDestinasi, err := s.glodok.GetTableJenisDestinasi(ctx, jenisdestinasiid, jenisdestinasiket, offset, length)

	if err != nil {
		return listJenisDestinasi, metadata, errors.Wrap(err, "[Service][GetTableJenisDestinasi]")
	}

	total, err = s.glodok.GetCountTableJenisDestinasi(ctx, jenisdestinasiid, jenisdestinasiket)

	if err != nil {
		return listJenisDestinasi, metadata, errors.Wrap(err, "[Service][GetCountTableJenisDestinasi]")
	}
	metadata["total_data"] = total

	return listJenisDestinasi, metadata, nil
}

func (s Service) UpdateJenisDestinasi(ctx context.Context, jenisdestinasi glodokEntity.TableJenisDestinasi, jenisdestinasiid string) (string, error) {
	var (
		result string
	)
	result, err := s.glodok.UpdateJenisDestinasi(ctx, jenisdestinasi, jenisdestinasiid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][UpdateJenisDestinasi]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) DeleteJenisDestinasi(ctx context.Context, jenisdestinasiid string) (string, error) {

	var (
		result string
	)
	_, err := s.glodok.DeleteJenisDestinasi(ctx, jenisdestinasiid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][DeleteJenisDestinasi]")
	}

	result = "Berhasil"

	return result, err
}

// for masyarakat
func (s Service) GetDestinasiByID(ctx context.Context, destinasiid string) ([]glodokEntity.TableDestinasi, error) {

	searchListDestinasiArray, err := s.glodok.GetDestinasiByID(ctx, destinasiid)

	if err != nil {
		return searchListDestinasiArray, errors.Wrap(err, "[Service][GetDestinasiByID]")
	}

	return searchListDestinasiArray, nil
}

func (s Service) GetAllDestinasi(ctx context.Context, kategori string, labelhalal string, destinasiname string) ([]glodokEntity.TableDestinasi, error) {

	searchListDestinasiArray, err := s.glodok.GetAllDestinasi(ctx, kategori, labelhalal, destinasiname)

	if err != nil {
		return searchListDestinasiArray, errors.Wrap(err, "[Service][GetAllDestinasi]")
	}

	return searchListDestinasiArray, nil
}

func (s Service) GetAllReview(ctx context.Context, rating string, page int, length int) ([]glodokEntity.TableReview, interface{}, error) {
	var (
		total int
	)

	metadata := make(map[string]interface{})

	offset := page * length
	searchListReviewArray, err := s.glodok.GetAllReview(ctx, rating, offset, length)

	if err != nil {
		return searchListReviewArray, metadata, errors.Wrap(err, "[Service][GetAllReview]")
	}

	total, err = s.glodok.GetCountAllReview(ctx, rating)

	if err != nil {
		return searchListReviewArray, metadata, errors.Wrap(err, "[Service][GetCountAllReview]")
	}
	metadata["total_data"] = total

	return searchListReviewArray, metadata, nil
}
