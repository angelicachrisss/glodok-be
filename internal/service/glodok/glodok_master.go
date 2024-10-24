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

//destinasi

func (s Service) GetJenisDestinasi(ctx context.Context) ([]glodokEntity.TableJenisDestinasi, error) {

	jenisDestinasiArray, err := s.glodok.GetJenisDestinasi(ctx)

	if err != nil {
		return jenisDestinasiArray, errors.Wrap(err, "[Service][GetJenisDestinasi]")
	}

	return jenisDestinasiArray, nil

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

func (s Service) GetTableAllDestinasi(ctx context.Context, page int, length int) ([]glodokEntity.TableDestinasi, interface{}, error) {
	var (
		total int
		// destinasiArrayFinish []glodokEntity.TableDestinasi
	)

	metadata := make(map[string]interface{})

	offset := page * length
	destinasiArray, err := s.glodok.GetTableAllDestinasi(ctx, offset, length)

	if err != nil {
		return destinasiArray, metadata, errors.Wrap(err, "[Service][GetTableAllDestinasi]")
	}

	total, err = s.glodok.GetCountTableAllDestinasi(ctx)

	if err != nil {
		return destinasiArray, metadata, errors.Wrap(err, "[Service][GetCountTableAllDestinasi]")
	}
	metadata["total_data"] = total

	return destinasiArray, metadata, nil
}

func (s Service) GetImageDestinasi(ctx context.Context, destinasiid string) ([]byte, error) {
	image, err := s.glodok.GetImageDestinasi(ctx, destinasiid)
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

func (s Service) GetTableDestinasiByJenis(ctx context.Context, jenisdestinasiid string, page int, length int) ([]glodokEntity.TableDestinasi, interface{}, error) {
	var (
		total int
		// destinasiArrayFinish []glodokEntity.TableDestinasi
	)

	metadata := make(map[string]interface{})

	offset := page * length
	destinasiArray, err := s.glodok.GetTableDestinasiByJenis(ctx, jenisdestinasiid, offset, length)

	if err != nil {
		return destinasiArray, metadata, errors.Wrap(err, "[Service][GetTableDestinasiByJenis]")
	}

	total, err = s.glodok.GetCountTableDestinasiByJenis(ctx, jenisdestinasiid)

	if err != nil {
		return destinasiArray, metadata, errors.Wrap(err, "[Service][GetCountTableDestinasiByJenis]")
	}
	metadata["total_data"] = total

	return destinasiArray, metadata, nil
}

func (s Service) GetSearchTableAllDestinasi(ctx context.Context, destinasiid string, destinasiname string, page int, length int) ([]glodokEntity.TableDestinasi, interface{}, error) {
	var (
		total int
		// destinasiArrayFinish []glodokEntity.TableDestinasi
	)

	metadata := make(map[string]interface{})

	offset := page * length
	destinasiArray, err := s.glodok.GetSearchTableAllDestinasi(ctx, destinasiid, destinasiname, offset, length)

	if err != nil {
		return destinasiArray, metadata, errors.Wrap(err, "[Service][GetSearchTableAllDestinasi]")
	}

	total, err = s.glodok.GetCountSearchTableAllDestinasi(ctx, destinasiid, destinasiname)

	if err != nil {
		return destinasiArray, metadata, errors.Wrap(err, "[Service][GetCountSearchTableAllDestinasi]")
	}
	metadata["total_data"] = total

	return destinasiArray, metadata, nil
}

func (s Service) GetSearchTableDestinasiByJenis(ctx context.Context, jenisdestinasiid string, destinasiid string, destinasiname string, page int, length int) ([]glodokEntity.TableDestinasi, interface{}, error) {
	var (
		total int
		// destinasiArrayFinish []glodokEntity.TableDestinasi
	)

	metadata := make(map[string]interface{})

	offset := page * length
	destinasiArray, err := s.glodok.GetSearchTableDestinasiByJenis(ctx, jenisdestinasiid, destinasiid, destinasiname, offset, length)

	if err != nil {
		return destinasiArray, metadata, errors.Wrap(err, "[Service][GetSearchTableDestinasiByJenis]")
	}

	total, err = s.glodok.GetCountSearchTableDestinasiByJenis(ctx, jenisdestinasiid, destinasiid, destinasiname)

	if err != nil {
		return destinasiArray, metadata, errors.Wrap(err, "[Service][GetCountSearchTableDestinasiByJenis]")
	}
	metadata["total_data"] = total

	return destinasiArray, metadata, nil
}

// func (s Service) GetSearchDestinasi(ctx context.Context, kategori string, destinasiid string, destinasiname string, page int, length int) ([]glodokEntity.TableDestinasi, interface{}, error) {
// 	var (
// 		total int
// 	)

// 	metadata := make(map[string]interface{})

// 	offset := page * length
// 	fmt.Println("offset: ", offset)
// 	searchListDestinasiArray, err := s.glodok.GetSearchDestinasi(ctx, kategori, destinasiid, destinasiname, offset, length)

// 	if err != nil {
// 		return searchListDestinasiArray, metadata, errors.Wrap(err, "[Service][GetSearchDestinasi]")
// 	}

// 	total, err = s.glodok.GetCountSearchDestinasi(ctx, kategori, destinasiid, destinasiname)

// 	if err != nil {
// 		return searchListDestinasiArray, metadata, errors.Wrap(err, "[Service][GetCountSearchDestinasi]")
// 	}
// 	metadata["total_data"] = total

// 	return searchListDestinasiArray, metadata, nil
// }

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

func (s Service) UpdateStatusDestinasi(ctx context.Context, destinasi glodokEntity.TableDestinasi, destinasiid string) (string, error) {
	var (
		result string
	)
	result, err := s.glodok.UpdateStatusDestinasi(ctx, destinasi, destinasiid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][UpdateStatusDestinasi]")
	}

	result = "Berhasil"

	return result, err
}

// tipe transportasi
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

func (s Service) GetTujuanTransportasiDropDown(ctx context.Context, tipetransportasiid string) ([]glodokEntity.TableTujuan, error) {

	tujuanArray, err := s.glodok.GetTujuanTransportasiDropDown(ctx, tipetransportasiid)

	if err != nil {
		return tujuanArray, errors.Wrap(err, "[Service][GetTujuanTransportasiDropDown]")
	}

	return tujuanArray, nil

}

func (s Service) GetPemberhentianDropDown(ctx context.Context, tipetransportasiid string) ([]glodokEntity.TablePemberhentian, error) {

	pemberhentianArray, err := s.glodok.GetPemberhentianDropDown(ctx, tipetransportasiid)

	if err != nil {
		return pemberhentianArray, errors.Wrap(err, "[Service][GetPemberhentianDropDown]")
	}

	return pemberhentianArray, nil

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

func (s Service) DeleteRuteByPemberhentian(ctx context.Context) (string, error) {

	var (
		result string
	)
	_, err := s.glodok.DeleteRuteByPemberhentian(ctx)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][DeleteRuteByPemberhentian]")
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

func (s Service) GetTableReview(ctx context.Context, destinasiid string, userid string, page int, length int) ([]glodokEntity.TableReview, interface{}, error) {

	var (
		total int
	)

	metadata := make(map[string]interface{})

	offset := page * length
	reviewArray, err := s.glodok.GetTableReview(ctx, destinasiid, userid, offset, length)

	if err != nil {
		return reviewArray, metadata, errors.Wrap(err, "[Service][GetTableReview]")
	}

	total, err = s.glodok.GetCountTableReview(ctx, destinasiid, userid)

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

// func (s Service) GetSearchReview(ctx context.Context, reviewid string, reviewer string, page int, length int) ([]glodokEntity.TableReview, interface{}, error) {
// 	var (
// 		total int
// 	)

// 	metadata := make(map[string]interface{})

// 	offset := page * length
// 	searchListReviewArray, err := s.glodok.GetSearchReview(ctx, reviewid, reviewer, offset, length)

// 	if err != nil {
// 		return searchListReviewArray, metadata, errors.Wrap(err, "[Service][GetSearchReview]")
// 	}

// 	total, err = s.glodok.GetCountSearchReview(ctx, reviewid, reviewer)

// 	if err != nil {
// 		return searchListReviewArray, metadata, errors.Wrap(err, "[Service][GetCountSearchReview]")
// 	}
// 	metadata["total_data"] = total

// 	return searchListReviewArray, metadata, nil
// }

// func (s Service) GetTableReviewByRating(ctx context.Context, rating int, page int, length int) ([]glodokEntity.TableReview, interface{}, error) {

// 	var (
// 		total int
// 	)

// 	metadata := make(map[string]interface{})

// 	offset := page * length
// 	reviewArray, err := s.glodok.GetTableReviewByRating(ctx, rating, offset, length)

// 	if err != nil {
// 		return reviewArray, metadata, errors.Wrap(err, "[Service][GetTableReviewByRating]")
// 	}

// 	total, err = s.glodok.GetCountTableReviewByRating(ctx, rating)

// 	if err != nil {
// 		return reviewArray, metadata, errors.Wrap(err, "[Service][GetCountTableReviewByRating]")
// 	}
// 	metadata["total_data"] = total

// 	return reviewArray, metadata, nil

// }

// func (s Service) GetSearchReviewByRating(ctx context.Context, rating int, reviewid string, reviewer string, page int, length int) ([]glodokEntity.TableReview, interface{}, error) {
// 	var (
// 		total int
// 	)

// 	metadata := make(map[string]interface{})

// 	offset := page * length
// 	searchListReviewArray, err := s.glodok.GetSearchReviewByRating(ctx, rating, reviewid, reviewer, offset, length)

// 	if err != nil {
// 		return searchListReviewArray, metadata, errors.Wrap(err, "[Service][GetSearchReviewByRating]")
// 	}

// 	total, err = s.glodok.GetCountSearchReviewByRating(ctx, rating, reviewid, reviewer)

// 	if err != nil {
// 		return searchListReviewArray, metadata, errors.Wrap(err, "[Service][GetCountSearchReviewByRating]")
// 	}
// 	metadata["total_data"] = total

// 	return searchListReviewArray, metadata, nil
// }

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

// jenis destinasi
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

// sejarahberanda
func (s Service) UpdateSejarahBeranda(ctx context.Context, sejarahberanda glodokEntity.TableSejarahBeranda) (string, error) {
	var (
		result string
	)
	result, err := s.glodok.UpdateSejarahBeranda(ctx, sejarahberanda)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][UpdateSejarahBeranda]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) GetSejarahBeranda(ctx context.Context) (glodokEntity.TableSejarahBeranda, error) {

	sejarahBeranda, err := s.glodok.GetSejarahBeranda(ctx)

	if err != nil {
		return sejarahBeranda, errors.Wrap(err, "[Service][GetSejarahBeranda]")
	}

	return sejarahBeranda, nil
}

func (s Service) InsertFotoBeranda(ctx context.Context, fotoberanda glodokEntity.TableFotoBeranda) (string, error) {

	var (
		result string
	)
	result, err := s.glodok.InsertFotoBeranda(ctx, fotoberanda)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][InsertFotoBeranda]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) GetTableFotoBeranda(ctx context.Context, fotoberandaid string, page int, length int) ([]glodokEntity.TableFotoBeranda, interface{}, error) {
	var (
		total int
	)

	metadata := make(map[string]interface{})

	offset := page * length
	fmt.Println("offset: ", offset)
	searchFotoBerandaArray, err := s.glodok.GetTableFotoBeranda(ctx, fotoberandaid, offset, length)

	if err != nil {
		return searchFotoBerandaArray, metadata, errors.Wrap(err, "[Service][GetTableFotoBeranda]")
	}

	total, err = s.glodok.GetCountTableFotoBeranda(ctx, fotoberandaid)

	if err != nil {
		return searchFotoBerandaArray, metadata, errors.Wrap(err, "[Service][GetCountTableFotoBeranda]")
	}
	metadata["total_data"] = total

	return searchFotoBerandaArray, metadata, nil
}

func (s Service) GetImageFotoBeranda(ctx context.Context, fotoberandaid string) ([]byte, error) {
	image, err := s.glodok.GetImageFotoBeranda(ctx, fotoberandaid)
	if err != nil {
		return image, errors.Wrap(err, "[SERVICE][GetImageFotoBeranda]")
	}

	return image, err
}

func (s Service) DeleteFotoBeranda(ctx context.Context, fotoberandaid string) (string, error) {

	var (
		result string
	)
	_, err := s.glodok.DeleteFotoBeranda(ctx, fotoberandaid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][DeleteFotoBeranda]")
	}

	result = "Berhasil"

	return result, err
}

// videoberanda
func (s Service) InsertVideoBeranda(ctx context.Context, videoberanda glodokEntity.TableVideoBeranda) (string, error) {

	var (
		result string
	)
	result, err := s.glodok.InsertVideoBeranda(ctx, videoberanda)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][InsertVideoBeranda]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) GetTableVideoBeranda(ctx context.Context, videoberandaid string, page int, length int) ([]glodokEntity.TableVideoBeranda, interface{}, error) {
	var (
		total int
	)

	metadata := make(map[string]interface{})

	offset := page * length
	listVideoBeranda, err := s.glodok.GetTableVideoBeranda(ctx, videoberandaid, offset, length)

	if err != nil {
		return listVideoBeranda, metadata, errors.Wrap(err, "[Service][GetTableVideoBeranda]")
	}

	total, err = s.glodok.GetCountTableVideoBeranda(ctx, videoberandaid)

	if err != nil {
		return listVideoBeranda, metadata, errors.Wrap(err, "[Service][GetCountTableVideoBeranda]")
	}
	metadata["total_data"] = total

	return listVideoBeranda, metadata, nil
}

func (s Service) DeleteVideoBeranda(ctx context.Context, videoberandaid string) (string, error) {

	var (
		result string
	)
	_, err := s.glodok.DeleteVideoBeranda(ctx, videoberandaid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][DeleteVideoBeranda]")
	}

	result = "Berhasil"

	return result, err
}

// tujuan
func (s Service) InsertTujuanTransportasi(ctx context.Context, tujuan glodokEntity.TableTujuan) (string, error) {

	var (
		result string
	)
	result, err := s.glodok.InsertTujuanTransportasi(ctx, tujuan)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][InsertTujuanTransportasi]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) GetTableTujuanTransportasi(ctx context.Context, tujuanid string, tipetransportasiname string, tujuanawal string, tujuanakhir string, page int, length int) ([]glodokEntity.TableTujuan, interface{}, error) {
	var (
		total int
	)

	metadata := make(map[string]interface{})

	offset := page * length
	searchListTujuan, err := s.glodok.GetTableTujuanTransportasi(ctx, tujuanid, tipetransportasiname, tujuanawal, tujuanakhir, offset, length)

	if err != nil {
		return searchListTujuan, metadata, errors.Wrap(err, "[Service][GetTableTujuanTransportasi]")
	}

	total, err = s.glodok.GetCountTableTujuanTransportasi(ctx, tujuanid, tipetransportasiname, tujuanawal, tujuanakhir)

	if err != nil {
		return searchListTujuan, metadata, errors.Wrap(err, "[Service][GetCountTableTujuanTransportasi]")
	}
	metadata["total_data"] = total

	return searchListTujuan, metadata, nil
}

func (s Service) DeleteTujuan(ctx context.Context, tujuanid string) (string, error) {

	var (
		result string
	)
	_, err := s.glodok.DeleteTujuan(ctx, tujuanid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][DeleteTujuan]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) UpdateTujuan(ctx context.Context, tujuan glodokEntity.TableTujuan, tujuanid string) (string, error) {
	var (
		result string
	)
	result, err := s.glodok.UpdateTujuan(ctx, tujuan, tujuanid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][UpdateTujuan]")
	}

	result = "Berhasil"

	return result, err
}

// pemberhentian
func (s Service) InsertPemberhentianTransportasi(ctx context.Context, pemberhentian glodokEntity.TablePemberhentian) (string, error) {

	var (
		result string
	)
	result, err := s.glodok.InsertPemberhentianTransportasi(ctx, pemberhentian)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][InsertPemberhentianTransportasi]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) GetTablePemberhentianTransportasi(ctx context.Context, pemberhentianid string, tipetransportasiname string, pemberhentianname string, page int, length int) ([]glodokEntity.TablePemberhentian, interface{}, error) {
	var (
		total int
	)

	metadata := make(map[string]interface{})

	offset := page * length
	searchListPemberhentian, err := s.glodok.GetTablePemberhentianTransportasi(ctx, pemberhentianid, tipetransportasiname, pemberhentianname, offset, length)

	if err != nil {
		return searchListPemberhentian, metadata, errors.Wrap(err, "[Service][GetTablePemberhentianTransportasi]")
	}

	total, err = s.glodok.GetCountTablePemberhentianTransportasi(ctx, pemberhentianid, tipetransportasiname, pemberhentianname)

	if err != nil {
		return searchListPemberhentian, metadata, errors.Wrap(err, "[Service][GetCountTablePemberhentianTransportasi]")
	}
	metadata["total_data"] = total

	return searchListPemberhentian, metadata, nil
}

func (s Service) DeletePemberhentian(ctx context.Context, pemberhentianid string) (string, error) {

	var (
		result string
	)
	_, err := s.glodok.DeletePemberhentian(ctx, pemberhentianid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][DeletePemberhentian]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) UpdatePemberhentian(ctx context.Context, pemberhentian glodokEntity.TablePemberhentian, tujuanid string) (string, error) {
	var (
		result string
	)
	result, err := s.glodok.UpdatePemberhentian(ctx, pemberhentian, tujuanid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][UpdatePemberhentian]")
	}

	result = "Berhasil"

	return result, err
}

// maps
func (s Service) UpdateMaps(ctx context.Context, maps glodokEntity.TableMaps, isi string) (string, error) {
	var (
		result string
	)
	result, err := s.glodok.UpdateMaps(ctx, maps, isi)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][UpdateMaps]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) GetMaps(ctx context.Context) (glodokEntity.TableMaps, error) {

	maps, err := s.glodok.GetMaps(ctx)

	if err != nil {
		return maps, errors.Wrap(err, "[Service][GetMaps]")
	}

	return maps, nil
}

func (s Service) GetTableUser(ctx context.Context, userid string, username string, page int, length int) ([]glodokEntity.TableUser, interface{}, error) {
	var (
		total int
	)

	metadata := make(map[string]interface{})

	offset := page * length
	listUser, err := s.glodok.GetTableUser(ctx, userid, username, offset, length)

	if err != nil {
		return listUser, metadata, errors.Wrap(err, "[Service][GetTableUser]")
	}

	total, err = s.glodok.GetCountTableUser(ctx, userid, username)

	if err != nil {
		return listUser, metadata, errors.Wrap(err, "[Service][GetTableUser]")
	}
	metadata["total_data"] = total

	return listUser, metadata, nil
}

func (s Service) DeleteUser(ctx context.Context, userid string) (string, error) {

	var (
		result string
	)
	_, err := s.glodok.DeleteUser(ctx, userid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][DeleteUser]")
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

func (s Service) GetAllDestinasi(ctx context.Context, jenisdestinasiid string, destinasiname string) ([]glodokEntity.TableDestinasi, error) {

	searchListDestinasiArray, err := s.glodok.GetAllDestinasi(ctx, jenisdestinasiid, destinasiname)

	if err != nil {
		return searchListDestinasiArray, errors.Wrap(err, "[Service][GetAllDestinasi]")
	}

	return searchListDestinasiArray, nil
}

func (s Service) GetAllReview(ctx context.Context, destinasiid string, rating string, page int, length int) ([]glodokEntity.TableReview, interface{}, error) {
	var (
		total int
	)

	metadata := make(map[string]interface{})

	offset := page * length
	searchListReviewArray, err := s.glodok.GetAllReview(ctx, destinasiid, rating, offset, length)

	if err != nil {
		return searchListReviewArray, metadata, errors.Wrap(err, "[Service][GetAllReview]")
	}

	total, err = s.glodok.GetCountAllReview(ctx, destinasiid, rating)

	if err != nil {
		return searchListReviewArray, metadata, errors.Wrap(err, "[Service][GetCountAllReview]")
	}
	metadata["total_data"] = total

	return searchListReviewArray, metadata, nil
}

func (s Service) GetFotoBerandaML(ctx context.Context) ([]glodokEntity.TableFotoBeranda, error) {

	searchFotoBerandaArray, err := s.glodok.GetFotoBerandaML(ctx)

	if err != nil {
		return searchFotoBerandaArray, errors.Wrap(err, "[Service][GetTableFotoBeranda]")
	}

	return searchFotoBerandaArray, nil
}

func (s Service) GetVideoBerandaML(ctx context.Context) ([]glodokEntity.TableVideoBeranda, error) {
	listVideoBeranda, err := s.glodok.GetVideoBerandaML(ctx)

	if err != nil {
		return listVideoBeranda, errors.Wrap(err, "[Service][GetTableVideoBeranda]")
	}

	return listVideoBeranda, nil
}

func (s Service) GetTransportasiML(ctx context.Context, perbaikanyn string) ([]glodokEntity.TableRuteTransportasi, interface{}, error) {

	var (
		total int
	)

	metadata := make(map[string]interface{})

	ruteTransportasiArray, err := s.glodok.GetTransportasiML(ctx, perbaikanyn)

	if err != nil {
		return ruteTransportasiArray, metadata, errors.Wrap(err, "[Service][GetTransportasiML]")
	}

	total, err = s.glodok.GetCountTransportasiML(ctx, perbaikanyn)

	if err != nil {
		return ruteTransportasiArray, metadata, errors.Wrap(err, "[Service][GetCountTransportasiML]")
	}
	metadata["total_data"] = total

	return ruteTransportasiArray, metadata, nil

}

func (s Service) GetBeritaML(ctx context.Context, judul string, page int, length int) ([]glodokEntity.TableBerita, interface{}, error) {
	var (
		total int
	)

	metadata := make(map[string]interface{})

	offset := page * length
	beritaArray, err := s.glodok.GetBeritaML(ctx, judul, offset, length)

	if err != nil {
		return beritaArray, metadata, errors.Wrap(err, "[Service][GetBeritaML]")
	}

	total, err = s.glodok.GetCountBeritaML(ctx, judul)

	if err != nil {
		return beritaArray, metadata, errors.Wrap(err, "[Service][GetCountBeritaML]")
	}
	metadata["total_data"] = total

	return beritaArray, metadata, nil
}

func (s Service) GetBeritaMLByID(ctx context.Context, beritaid string) ([]glodokEntity.TableBerita, error) {

	beritaArray, err := s.glodok.GetBeritaMLByID(ctx, beritaid)

	if err != nil {
		return beritaArray, errors.Wrap(err, "[Service][GetBeritaMLByID]")
	}

	return beritaArray, nil
}

func (s Service) GetJenisDestinasiML(ctx context.Context) ([]glodokEntity.TableJenisDestinasi, error) {

	jenisDestinasiArray, err := s.glodok.GetJenisDestinasiML(ctx)

	if err != nil {
		return jenisDestinasiArray, errors.Wrap(err, "[Service][GetJenisDestinasiML]")
	}

	return jenisDestinasiArray, nil

}

func (s Service) GetDestinasiDDML(ctx context.Context) ([]glodokEntity.TableDestinasi, error) {

	destinasiArray, err := s.glodok.GetDestinasiDDML(ctx)

	if err != nil {
		return destinasiArray, errors.Wrap(err, "[Service][GetDestinasiDDML]")
	}

	return destinasiArray, nil

}

func (s Service) GetAvgReview(ctx context.Context, destinasiid string) (float64, interface{}, error) {
	var (
		total int
	)

	metadata := make(map[string]interface{})

	avg, err := s.glodok.GetAvgReview(ctx, destinasiid)

	if err != nil {
		return avg, metadata, errors.Wrap(err, "[Service][GetAvgReview]")
	}

	total, err = s.glodok.GetCountAllReview(ctx, destinasiid, "")

	metadata["total_data"] = total

	return avg, metadata, nil

}

func (s Service) InsertUser(ctx context.Context, user glodokEntity.TableUser) (string, error) {
	var (
		result string
	)
	result, err := s.glodok.InsertUser(ctx, user)
	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][InsertUser]")
	}
	result = "Berhasil"
	return result, err
}

func (s Service) SubmitLoginML(ctx context.Context, userid string, pass string) (string, error) {
	var (
		result string
	)
	result, err := s.glodok.SubmitLoginML(ctx, userid, pass)
	if err != nil {
		result = "Gagal Login"
		fmt.Println("result", result)
		return result, errors.Wrap(err, "[Service][SubmitLoginML]")
	}
	result = "Berhasil Login"
	return result, err
}


func (s Service) GetUser(ctx context.Context,userid string) (glodokEntity.TableUser, error) {

	user, err := s.glodok.GetUser(ctx,userid)

	if err != nil {
		return user, errors.Wrap(err, "[Service][GetUser]")
	}

	return user, nil
}

func (s Service) UpdateUser(ctx context.Context, user glodokEntity.TableUser, userid string) (string, error) {
	var (
		result string
	)
	result, err := s.glodok.UpdateUser(ctx, user, userid)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][UpdateUser]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) DeleteReviewByUser(ctx context.Context) (string, error) {

	var (
		result string
	)
	_, err := s.glodok.DeleteReviewByUser(ctx)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][DeleteReviewByUser]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) DeleteRuteByTujuan(ctx context.Context) (string, error) {

	var (
		result string
	)
	_, err := s.glodok.DeleteRuteByTujuan(ctx)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][DeleteRuteByTujuan]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) DeleteRuteByTipe(ctx context.Context) (string, error) {

	var (
		result string
	)
	_, err := s.glodok.DeleteRuteByTipe(ctx)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][DeleteRuteByTipe]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) DeletePemberhentianByTipe(ctx context.Context) (string, error) {

	var (
		result string
	)
	_, err := s.glodok.DeletePemberhentianByTipe(ctx)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][DeletePemberhentianByTipe]")
	}

	result = "Berhasil"

	return result, err
}

func (s Service) DeleteTujuanByTipe(ctx context.Context) (string, error) {

	var (
		result string
	)
	_, err := s.glodok.DeleteTujuanByTipe(ctx)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][DeleteTujuanByTipe]")
	}

	result = "Berhasil"

	return result, err
}