package glodok

import (
	// "bytes"
	"encoding/json"
	glodokEntity "glodok-be/internal/entity/glodok"
	"glodok-be/pkg/response"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go.uber.org/zap"
)

func (h *Handler) InsertGlodok(w http.ResponseWriter, r *http.Request) {
	var (
		result   interface{}
		metadata interface{}
		err      error

		resp  response.Response
		types string

		InsertAdmin    glodokEntity.GetAdmin
		// TableDestinasi glodokEntity.TableDestinasi
	)
	defer resp.RenderJSON(w, r)

	spanCtx, _ := h.tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	span := h.tracer.StartSpan("Getglodok", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	ctx := r.Context()
	ctx = opentracing.ContextWithSpan(ctx, span)
	h.logger.For(ctx).Info("HTTP request received", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Your code here
	types = r.FormValue("type")
	switch types {

	case "insertadmin":
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &InsertAdmin)
		result, err = h.glodokSvc.InsertAdmin(ctx, InsertAdmin)

	case "submitlogin":
		result, err = h.glodokSvc.SubmitLogin(ctx, r.FormValue("adminid"), r.FormValue("adminpass"))
		// case "insertdestinasi":
		// 	body, _ := ioutil.ReadAll(r.Body)
		// 	fmt.Println("body", body)
		// 	// file, _, err := r.FormFile("destinasi_gambar")
		// 	// // fmt.Println("files", file)
		// 	// if err != nil {
		// 	// 	http.Error(w, "Unable to get file", http.StatusBadRequest)
		// 	// 	return
		// 	// }
		// 	// defer file.Close()

		// 	// fileBytes, err := ioutil.ReadAll(file)
		// 	// if err != nil {
		// 	// 	http.Error(w, "Unable to read file", http.StatusInternalServerError)
		// 	// 	return
		// 	// }

		// 	// fmt.Println("fileBytes", fileBytes)

		// 	// fmt.Println("files2", file)
		// 	json.Unmarshal(body, &TableDestinasi)
		// 	// TableDestinasi.DestinasiGambar = fileBytes
		// 	result, err = h.glodokSvc.InsertDestinasi(ctx, TableDestinasi)\
	case "insertdestinasi":
		// Memproses bagian dari form-data
		err := r.ParseMultipartForm(10 << 20) // Maksimum ukuran file 10MB
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		// Mengambil file dari form-data
		file, _, err := r.FormFile("destinasi_gambar")
		if err != nil {
			http.Error(w, "Unable to get file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Membaca isi file ke dalam byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, "Unable to read file", http.StatusInternalServerError)
			return
		}

		// Konversi string ke float64 untuk destinasi_lang
		destinasiLang, err := strconv.ParseFloat(r.FormValue("destinasi_lang"), 64)
		if err != nil {
			http.Error(w, "Invalid format for destinasi_lang", http.StatusBadRequest)
			return
		}

		// Konversi string ke float64 untuk destinasi_long
		destinasiLong, err := strconv.ParseFloat(r.FormValue("destinasi_long"), 64)
		if err != nil {
			http.Error(w, "Invalid format for destinasi_long", http.StatusBadRequest)
			return
		}

		// Membaca data JSON yang lain dari form-data
		TableDestinasi := glodokEntity.TableDestinasi{
			DestinasiID:     "", // Asumsikan ID di-generate otomatis, bisa disesuaikan dengan kebutuhan
			DestinasiName:   r.FormValue("destinasi_name"),
			DestinasiDesc:   r.FormValue("destinasi_desc"),
			DestinasiAlamat: r.FormValue("destinasi_alamat"),
			DestinasiLang:   destinasiLang,
			DestinasiLong:   destinasiLong,
			DestinasiHBuka:  r.FormValue("destinasi_hbuka"),
			DestinasiHTutup: r.FormValue("destinasi_htutup"),
			DestinasiKet:    r.FormValue("destinasi_kat"),
			DestinasiHalal:  r.FormValue("destinasi_labelhalal"),
			DestinasiPic:    r.FormValue("destinasi_pic"),
			DestinasiGambar: fileBytes, // Menyimpan byte array gambar ke struct
		}

		// Memasukkan data ke dalam database melalui layanan InsertDestinasi
		result, err = h.glodokSvc.InsertDestinasi(ctx, TableDestinasi)
		if err != nil {
			log.Println("err", err)
		}	
	}

	if err != nil {
		resp.SetError(err, http.StatusInternalServerError)
		resp.StatusCode = 500
		log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, err.Error())
		resp.Data = result
		return
	}

	resp.Data = result
	resp.Metadata = metadata
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

}
