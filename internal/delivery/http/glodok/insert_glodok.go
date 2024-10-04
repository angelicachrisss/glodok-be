package glodok

import (
	// "bytes"
	"encoding/json"
	"fmt"
	glodokEntity "glodok-be/internal/entity/glodok"
	"glodok-be/pkg/response"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

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

		// InsertTipeTransportasi glodokEntity.TableTipeTransportasi
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
	//admin
	case "submitlogin":
		result, err = h.glodokSvc.SubmitLogin(ctx, r.FormValue("adminid"), r.FormValue("adminpass"))

	//-------------------------------------------------------------------------------------------------------------------------
	//destinasi
	case "insertdestinasi":
		// Memproses bagian dari form-data
		err := r.ParseMultipartForm(10 << 20) // Maksimum ukuran file 10MB
		if err != nil {
			fmt.Println("Error memproses bagian dari form-data:", err)
			return
		}

		// Mengambil file dari form-data
		file, _, err := r.FormFile("destinasi_gambar")
		if err != nil {
			fmt.Println("Error mengambil file dari form-data:", err)
			return
		}
		defer file.Close()

		// Membaca isi file ke dalam byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println("Error membaca isi file ke dalam byte array:", err)
			return
		}

		// Konversi string ke float64 untuk destinasi_lang
		destinasiLang, err := strconv.ParseFloat(r.FormValue("destinasi_lang"), 64)
		if err != nil {
			fmt.Println("Error konversi string ke float64 untuk destinasi_lang:", err)
			return
		}

		// Konversi string ke float64 untuk destinasi_long
		destinasiLong, err := strconv.ParseFloat(r.FormValue("destinasi_long"), 64)
		if err != nil {
			fmt.Println("Error konversi string ke float64 untuk destinasi_long:", err)
			return
		}

		// // Parse time values
		layout := "15:04:05" // Adjust layout to match your time format

		// Parse the full time string into a time.Time object
		parsedTimeJamBuka, err := time.Parse(layout, r.FormValue("destinasi_jbuka"))
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return
		}

		// Extract the time component and create a new time.Time object with a default date
		defaultDate := "0001-01-01" // Use the minimum date if the specific date is not important
		timeOnlyStr := defaultDate + " " + parsedTimeJamBuka.Format("15:04:05")

		// Define the layout for the new time string with default date
		defaultLayout := "2006-01-02 15:04:05"

		// Parse the new time string into a time.Time object with the default date
		jamBuka, err := time.Parse(defaultLayout, timeOnlyStr)
		if err != nil {
			fmt.Println("Error parsing time only:", err)
			return
		}

		parsedTimeJamTutup, err := time.Parse(layout, r.FormValue("destinasi_jtutup"))
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return
		}

		// Extract the time component and create a new time.Time object with a default date
		defaultDate2 := "0001-01-01" // Use the minimum date if the specific date is not important
		timeOnlyStr2 := defaultDate2 + " " + parsedTimeJamTutup.Format("15:04:05")

		// Define the layout for the new time string with default date
		defaultLayout2 := "2006-01-02 15:04:05"

		// Parse the new time string into a time.Time object with the default date
		jamTutup, err := time.Parse(defaultLayout2, timeOnlyStr2)
		if err != nil {
			fmt.Println("Error parsing time only:", err)
			return
		}

		// Membaca data JSON yang lain dari form-data
		TableDestinasi := glodokEntity.TableDestinasi{
			DestinasiID:     "",
			DestinasiName:   r.FormValue("destinasi_name"),
			DestinasiDesc:   r.FormValue("destinasi_desc"),
			DestinasiAlamat: r.FormValue("destinasi_alamat"),
			DestinasiLang:   destinasiLang,
			DestinasiLong:   destinasiLong,
			DestinasiHBuka:  r.FormValue("destinasi_hbuka"),
			DestinasiHTutup: r.FormValue("destinasi_htutup"),
			DestinasiJBuka:  jamBuka,
			DestinasiJTutup: jamTutup,
			DestinasiKet:    r.FormValue("destinasi_kat"),
			DestinasiHalal:  r.FormValue("destinasi_labelhalal"),
			DestinasiGambar: fileBytes, // Menyimpan byte array gambar ke struct
		}

		// Memasukkan data ke dalam database melalui layanan InsertDestinasi
		result, err = h.glodokSvc.InsertDestinasi(ctx, TableDestinasi)
		if err != nil {
			resp.SetError(err, http.StatusInternalServerError)
			resp.StatusCode = 500
			log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, err.Error())
			resp.Data = result
			return
		}

	//-------------------------------------------------------------------------------------------------------------------------
	//tipetransportasi
	case "inserttipetransportasi":
		var (
			InsertTipeTransportasi glodokEntity.TableTipeTransportasi
		)

		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &InsertTipeTransportasi)
		result, err = h.glodokSvc.InsertTipeTransportasi(ctx, InsertTipeTransportasi)

		//-------------------------------------------------------------------------------------------------------------------------
	//rutetransportasi
	case "insertrutetransportasi":
		var (
			InsertRuteTransportasi glodokEntity.TableRuteTransportasi
		)

		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &InsertRuteTransportasi)
		result, err = h.glodokSvc.InsertRuteTransportasi(ctx, InsertRuteTransportasi)
		//-------------------------------------------------------------------------------------------------------------------------

		//review
	case "insertreview":
		var (
			InsertReview glodokEntity.TableReview
		)

		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &InsertReview)
		result, err = h.glodokSvc.InsertReview(ctx, InsertReview)
		//-------------------------------------------------------------------------------------------------------------------------
	case "insertberita":
		// Memproses bagian dari form-data
		err := r.ParseMultipartForm(10 << 20) // Maksimum ukuran file 10MB
		if err != nil {
			fmt.Println("Error memproses bagian dari form-data:", err)
			return
		}

		// Mengambil file dari form-data
		file, _, err := r.FormFile("berita_foto")
		if err != nil {
			fmt.Println("Error mengambil file dari form-data:", err)
			return
		}
		defer file.Close()

		// Membaca isi file ke dalam byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println("Error membaca isi file ke dalam byte array:", err)
			return
		}

		// Membaca data JSON yang lain dari form-data
		TableBerita := glodokEntity.TableBerita{
			BeritaID: "",
			DestinasiID: r.FormValue("destinasi_id"),
			BeritaJudul: r.FormValue("berita_judul"),
			BeritaDesc: r.FormValue("berita_desc"),
			BeritaGambar: fileBytes,
			BeritaLinkSumber: r.FormValue("berita_linksumber"),
		}

		// Memasukkan data ke dalam database melalui layanan InsertBerita
		result, err = h.glodokSvc.InsertBerita(ctx, TableBerita)
		if err != nil {
			resp.SetError(err, http.StatusInternalServerError)
			resp.StatusCode = 500
			log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, err.Error())
			resp.Data = result
			return
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
