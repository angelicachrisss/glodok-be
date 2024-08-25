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

	// for _, y := range destinasiArray {
	// 	// Encode image data to PNG format
	// 	var buf bytes.Buffer
	// 	imgBuffer := bytes.NewReader(y.DestinasiGambar)
	// 	img, _, err := image.Decode(imgBuffer)
	// 	if err != nil {
	// 		return destinasiArray, metadata, errors.Wrap(err, "[Service][GetCountDestinasi]")
	// 	}

	// 	err = png.Encode(&buf, img)
	// 	if err != nil {
	// 		return destinasiArray, metadata, errors.Wrap(err, "[Service][GetCountDestinasi]")
	// 	}

	// 	// Convert encoded image data to Base64
	// 	encodedImageData := base64.StdEncoding.EncodeToString(buf.Bytes())

	// 	oneObj := glodokEntity.TableDestinasi{
	// 		DestinasiID:     y.DestinasiID,
	// 		DestinasiName:   y.DestinasiName,
	// 		DestinasiDesc:   y.DestinasiDesc,
	// 		DestinasiAlamat: y.DestinasiAlamat,
	// 		// DestinasiGambar: y.DestinasiGambar,
	// 		DestinasiLang:   y.DestinasiLang,
	// 		DestinasiLong:   y.DestinasiLong,
	// 		DestinasiHBuka:  y.DestinasiHBuka,
	// 		DestinasiHTutup: y.DestinasiHTutup,
	// 		DestinasiKet:    y.DestinasiKet,
	// 		DestinasiHalal:  y.DestinasiHalal,
	// 		DestinasiPic:    encodedImageData,
	// 	}
	// 	destinasiArrayFinish = append(destinasiArrayFinish, oneObj)
	// 	// 	responses = append(responses, response)
	// }

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

// func (s Service) GetSearchDestinasiIc(ctx context.Context, destinasiname string, page int, length int) ([]glodokEntity.TableDestinasiIc, interface{}, error) {
// 	var (
// 		total int
// 	)

// 	metadata := make(map[string]interface{})

// 	offset := page * length
// 	fmt.Println("offset: ", offset)
// 	searchListDestinasiIcArray, err := s.glodok.GetSearchDestinasiIc(ctx, destinasiname, offset, length)

// 	if err != nil {
// 		return searchListDestinasiIcArray, metadata, errors.Wrap(err, "[Service][GetSearchDestinasiIc]")
// 	}

// 	total, err = s.glodok.GetCountSearchDestinasiIc(ctx, destinasiname)

// 	if err != nil {
// 		return searchListDestinasiIcArray, metadata, errors.Wrap(err, "[Service][GetSearchDestinasiIc]")
// 	}
// 	metadata["total_data"] = total

// 	return searchListDestinasiIcArray, metadata, nil
// }
