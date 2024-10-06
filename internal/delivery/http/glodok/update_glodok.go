package glodok

import (
	// "encoding/json"
	// "io/ioutil"
	"encoding/json"
	"fmt"
	httpHelper "glodok-be/internal/delivery/http"
	"io/ioutil"
	"time"

	// "strconv"

	// glodokEntity "glodok-be/internal/entity/glodok"
	"glodok-be/pkg/response"
	// "encoding/json"
	// "io/ioutil"
	"log"
	"net/http"

	// "strconv"

	glodokEntity "glodok-be/internal/entity/glodok"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go.uber.org/zap"
)

func (h *Handler) UpdateGlodok(w http.ResponseWriter, r *http.Request) {
	var (
		result   interface{}
		metadata interface{}
		err      error
		resp     response.Response
		types    string
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
	case "updateadmin":
		var requestUpdate glodokEntity.GetAdmin
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &requestUpdate)
		result, err = h.glodokSvc.UpdateAdmin(ctx, requestUpdate, r.FormValue("adminid"))
	case "updatetipetransportasi":
		var tipetransportasi glodokEntity.TableTipeTransportasi
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &tipetransportasi)
		result, err = h.glodokSvc.UpdateTipeTransportasi(ctx, tipetransportasi, r.FormValue("tipetransportasiid"))
	case "updaterutetransportasi":
		var rutetransportasi glodokEntity.TableRuteTransportasi
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &rutetransportasi)
		result, err = h.glodokSvc.UpdateRuteTransportasi(ctx, rutetransportasi, r.FormValue("ruteid"))
	case "updatestatusdestinasi":
		var destinasi glodokEntity.TableDestinasi
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &destinasi)
		result, err = h.glodokSvc.UpdateStatusDestinasi(ctx, destinasi, r.FormValue("destinasiid"))
	case "updatedestinasi":
		// Memproses bagian dari form-data
		err := r.ParseMultipartForm(10 << 20) // Maksimum ukuran file 10MB
		if err != nil {
			fmt.Println("Error memproses bagian dari form-data:", err)
			return
		}

		// Mengambil file gambar dari form-data
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

		// Parse time values
		layout := "15:04:05"

		parsedTimeJamBuka, err := time.Parse(layout, r.FormValue("destinasi_jbuka"))
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return
		}
		jamBuka := time.Date(1, 1, 1, parsedTimeJamBuka.Hour(), parsedTimeJamBuka.Minute(), parsedTimeJamBuka.Second(), 0, time.UTC)

		parsedTimeJamTutup, err := time.Parse(layout, r.FormValue("destinasi_jtutup"))
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return
		}
		jamTutup := time.Date(1, 1, 1, parsedTimeJamTutup.Hour(), parsedTimeJamTutup.Minute(), parsedTimeJamTutup.Second(), 0, time.UTC)

		// Membaca data JSON yang lain dari form-data
		destinasiID := r.FormValue("destinasiid")
		TableDestinasi := glodokEntity.TableDestinasi{
			DestinasiID:      destinasiID,
			DestinasiName:    r.FormValue("destinasi_name"),
			DestinasiDesc:    r.FormValue("destinasi_desc"),
			DestinasiHBuka:   r.FormValue("destinasi_hbuka"),
			DestinasiHTutup:  r.FormValue("destinasi_htutup"),
			DestinasiJBuka:   jamBuka,
			DestinasiJTutup:  jamTutup,
			DestinasiHalal:   r.FormValue("destinasi_labelhalalyn"),
			DestinasiOtentik: r.FormValue("destinasi_otentikyn"),
			DestinasiGambar:  fileBytes, // Menyimpan byte array gambar ke struct
		}

		// Memperbarui data ke dalam database melalui layanan UpdateDestinasi
		result, err = h.glodokSvc.UpdateDestinasi(ctx, TableDestinasi, destinasiID)
		if err != nil {
			resp.SetError(err, http.StatusInternalServerError)
			resp.StatusCode = 500
			log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, err.Error())
			resp.Data = result
			return
		}

	case "updateberita":
		// Memproses bagian dari form-data
		err := r.ParseMultipartForm(10 << 20) // Maksimum ukuran file 10MB
		if err != nil {
			fmt.Println("Error memproses bagian dari form-data:", err)
			return
		}

		// Mengambil file gambar dari form-data
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
		beritaID := r.FormValue("beritaid")
		TableBerita := glodokEntity.TableBerita{
			BeritaID:         beritaID,
			DestinasiID:      r.FormValue("destinasi_id"),
			BeritaJudul:      r.FormValue("berita_judul"),
			BeritaDesc:       r.FormValue("berita_desc"),
			BeritaGambar:     fileBytes,
			BeritaLinkSumber: r.FormValue("berita_linksumber"),
		}

		// Memperbarui data ke dalam database melalui layanan UpdateDestinasi
		result, err = h.glodokSvc.UpdateBerita(ctx, TableBerita, beritaID)
		if err != nil {
			resp.SetError(err, http.StatusInternalServerError)
			resp.StatusCode = 500
			log.Printf("[ERROR] %s %s - %s\n", r.Method, r.URL, err.Error())
			resp.Data = result
			return
		}
	case "updatejenisdestinasi":
		var jenisdestinasi glodokEntity.TableJenisDestinasi
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &jenisdestinasi)
		result, err = h.glodokSvc.UpdateJenisDestinasi(ctx, jenisdestinasi, r.FormValue("jenisdestinasiid"))
	case "updatesejarahberanda":
		var sejarahberanda glodokEntity.TableSejarahBeranda
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Println("sejarahberanda: ", sejarahberanda)
		json.Unmarshal(body, &sejarahberanda)
		result, err = h.glodokSvc.UpdateSejarahBeranda(ctx, sejarahberanda)
	case "updatemaps":
		var maps glodokEntity.TableMaps
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &maps)
		result, err = h.glodokSvc.UpdateMaps(ctx, maps, r.FormValue("isi"))
	case "updatetujuan":
		var tujuan glodokEntity.TableTujuan
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &tujuan)
		result, err = h.glodokSvc.UpdateTujuan(ctx, tujuan, r.FormValue("tujuanid"))
	case "updatepemberhentian":
		var pemberhentian glodokEntity.TablePemberhentian
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &pemberhentian)
		result, err = h.glodokSvc.UpdatePemberhentian(ctx, pemberhentian, r.FormValue("pemberhentianid"))
	}

	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		h.logger.For(ctx).Error("HTTP request error", zap.String("method", r.Method), zap.Stringer("url", r.URL), zap.Error(err))
		return
	}

	resp.Data = result
	resp.Metadata = metadata
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))

}
