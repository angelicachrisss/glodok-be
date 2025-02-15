package glodok

import (
	// "internal/itoa"

	"bytes"
	"fmt"
	httpHelper "glodok-be/internal/delivery/http"
	"glodok-be/pkg/response"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"strconv"

	// "strconv"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go.uber.org/zap"
)

// @Router /v1/profiles [get]
func (h *Handler) GetGlodok(w http.ResponseWriter, r *http.Request) {
	var (
		result   interface{}
		metadata interface{}
		err      error
		resp     response.Response
		types    string
	)
	defer resp.RenderJSON(w, r)

	// ptid, _ := strconv.Atoi(r.FormValue("ptid"))
	// page, _ := strconv.Atoi(r.FormValue("page"))
	// limit, _ := strconv.Atoi(r.FormValue("limit"))

	spanCtx, _ := h.tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	span := h.tracer.StartSpan("GetGlodok", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	ctx := r.Context()
	ctx = opentracing.ContextWithSpan(ctx, span)
	h.logger.For(ctx).Info("HTTP request received", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// Your code here
	types = r.FormValue("type")
	switch types {
	case "getadmin":
		result, err = h.glodokSvc.GetAdmin(ctx)
	case "getadminbyid":
		result, err = h.glodokSvc.GetAdminbyID(ctx, r.FormValue("adminid"))
	case "getjenisdestinasidropdown":
		result, err = h.glodokSvc.GetJenisDestinasi(ctx)
	case "getdestinasi":
		page, _ := strconv.Atoi(r.FormValue("page"))
		length, _ := strconv.Atoi(r.FormValue("length"))
		result, metadata, err = h.glodokSvc.GetTableAllDestinasi(ctx, page, length)
	case "getdestinasibyjenis":
		page, _ := strconv.Atoi(r.FormValue("page"))
		length, _ := strconv.Atoi(r.FormValue("length"))
		result, metadata, err = h.glodokSvc.GetTableDestinasiByJenis(ctx, r.FormValue("jenisdestinasiid"), page, length)
	case "getimagedestinasi":
		result, err = h.glodokSvc.GetImageDestinasi(ctx, r.FormValue("destinasiid"))
		if err != nil {
			http.Error(w, "Failed to get image data", http.StatusInternalServerError)
			return
		}

		// Type assertion
		imgData, ok := result.([]byte)
		if !ok {
			log.Fatal("The result is not of type []byte")
		}

		// Create a buffer from the image data
		imgBuffer := bytes.NewReader(imgData)

		// Decode the image data to get the image.Image object
		img, imgFormat, err := image.Decode(imgBuffer)
		if err != nil {
			http.Error(w, "Unable to decode image", http.StatusInternalServerError)
			return
		}

		// Set the appropriate header for the image format
		var contentType string
		switch imgFormat {
		case "png":
			contentType = "image/png"
		case "jpeg", "jpg":
			contentType = "image/jpeg"
		default:
			http.Error(w, "Unsupported image format", http.StatusUnsupportedMediaType)
			return
		}
		w.Header().Set("Content-Type", contentType)

		// Encode the image to the appropriate format and write it to the response
		switch imgFormat {
		case "png":
			err = png.Encode(w, img)
		case "jpeg", "jpg":
			err = jpeg.Encode(w, img, nil)
		default:
			http.Error(w, "Unsupported image format", http.StatusUnsupportedMediaType)
			return
		}

		if err != nil {
			http.Error(w, "Unable to encode image", http.StatusInternalServerError)
			return
		}
	case "getsearchdestinasi":
		page, _ := strconv.Atoi(r.FormValue("page"))
		length, _ := strconv.Atoi(r.FormValue("length"))
		result, metadata, err = h.glodokSvc.GetSearchTableAllDestinasi(ctx, r.FormValue("destinasiid"), r.FormValue("destinasiname"), page, length)
	case "getsearchdestinasibyjenis":
		page, _ := strconv.Atoi(r.FormValue("page"))
		length, _ := strconv.Atoi(r.FormValue("length"))
		result, metadata, err = h.glodokSvc.GetSearchTableDestinasiByJenis(ctx, r.FormValue("jenisdestinasiid"), r.FormValue("destinasiid"), r.FormValue("destinasiname"), page, length)
	case "gettipetransportasi":
		page, _ := strconv.Atoi(r.FormValue("page"))
		length, _ := strconv.Atoi(r.FormValue("length"))
		result, metadata, err = h.glodokSvc.GetTableTipeTransportasi(ctx, page, length)
	case "getsearchtipetransportasi":
		page, _ := strconv.Atoi(r.FormValue("page"))
		length, _ := strconv.Atoi(r.FormValue("length"))
		result, metadata, err = h.glodokSvc.GetSearchTipeTransportasi(ctx, r.FormValue("tipetransportasiid"), r.FormValue("tipetransportasiname"), page, length)
	case "gettipetransportasidropdown":
		result, err = h.glodokSvc.GetTipeTransportasi(ctx)
	case "gettujuandropdown":
		result, err = h.glodokSvc.GetTujuanTransportasiDropDown(ctx, r.FormValue("tipetransportasiid"))
	case "getpemberhentiandropdown":
		result, err = h.glodokSvc.GetPemberhentianDropDown(ctx, r.FormValue("tipetransportasiid"))
	case "getrutetransportasi":
		page, _ := strconv.Atoi(r.FormValue("page"))
		length, _ := strconv.Atoi(r.FormValue("length"))
		result, metadata, err = h.glodokSvc.GetTableRuteTransportasi(ctx, page, length)
	case "getsearchrutetransportasi":
		page, _ := strconv.Atoi(r.FormValue("page"))
		length, _ := strconv.Atoi(r.FormValue("length"))
		result, metadata, err = h.glodokSvc.GetSearchRuteTransportasi(ctx, r.FormValue("tipetransportasiname"), r.FormValue("tujuanawal"), r.FormValue("tujuanakhir"), page, length)
	case "getreview":
		page, _ := strconv.Atoi(r.FormValue("page"))
		length, _ := strconv.Atoi(r.FormValue("length"))
		result, metadata, err = h.glodokSvc.GetTableReview(ctx, r.FormValue("destinasiid"), r.FormValue("userid"), page, length)
	// case "getsearchreview":
	// 	page, _ := strconv.Atoi(r.FormValue("page"))
	// 	length, _ := strconv.Atoi(r.FormValue("length"))
	// 	result, metadata, err = h.glodokSvc.GetSearchReview(ctx, r.FormValue("reviewid"), r.FormValue("reviewer"), page, length)
	// case "getreviewbyrating":
	// 	rating, _ := strconv.Atoi(r.FormValue("rating"))
	// 	page, _ := strconv.Atoi(r.FormValue("page"))
	// 	length, _ := strconv.Atoi(r.FormValue("length"))
	// 	result, metadata, err = h.glodokSvc.GetTableReviewByRating(ctx, rating, page, length)
	// case "getsearchreviewbyrating":
	// 	rating, _ := strconv.Atoi(r.FormValue("rating"))
	// 	page, _ := strconv.Atoi(r.FormValue("page"))
	// 	length, _ := strconv.Atoi(r.FormValue("length"))
	// 	result, metadata, err = h.glodokSvc.GetSearchReviewByRating(ctx, rating, r.FormValue("reviewid"), r.FormValue("reviewer"), page, length)
	case "getdestinasidropdown":
		result, err = h.glodokSvc.GetDestinasi(ctx)
	case "getimageberita":
		result, err = h.glodokSvc.GetImageBerita(ctx, r.FormValue("beritaid"))
		if err != nil {
			http.Error(w, "Failed to get image data", http.StatusInternalServerError)
			return
		}

		// Type assertion
		imgData, ok := result.([]byte)
		if !ok {
			log.Fatal("The result is not of type []byte")
		}

		// Create a buffer from the image data
		imgBuffer := bytes.NewReader(imgData)

		// Decode the image data to get the image.Image object
		img, imgFormat, err := image.Decode(imgBuffer)
		if err != nil {
			http.Error(w, "Unable to decode image", http.StatusInternalServerError)
			return
		}

		// Set the appropriate header for the image format
		var contentType string
		switch imgFormat {
		case "png":
			contentType = "image/png"
		case "jpeg", "jpg":
			contentType = "image/jpeg"
		default:
			http.Error(w, "Unsupported image format", http.StatusUnsupportedMediaType)
			return
		}
		w.Header().Set("Content-Type", contentType)

		// Encode the image to the appropriate format and write it to the response
		switch imgFormat {
		case "png":
			err = png.Encode(w, img)
		case "jpeg", "jpg":
			err = jpeg.Encode(w, img, nil)
		default:
			http.Error(w, "Unsupported image format", http.StatusUnsupportedMediaType)
			return
		}

		if err != nil {
			http.Error(w, "Unable to encode image", http.StatusInternalServerError)
			return
		}
	case "gettableberita":
		page, _ := strconv.Atoi(r.FormValue("page"))
		length, _ := strconv.Atoi(r.FormValue("length"))
		result, metadata, err = h.glodokSvc.GetTableBerita(ctx, page, length)

	case "getsearchberita":
		page, _ := strconv.Atoi(r.FormValue("page"))
		length, _ := strconv.Atoi(r.FormValue("length"))
		result, metadata, err = h.glodokSvc.GetSearchBerita(ctx, r.FormValue("beritaid"), r.FormValue("destinasiname"), r.FormValue("beritajudul"), page, length)
	case "gettablejenisdestinasi":
		page, _ := strconv.Atoi(r.FormValue("page"))
		length, _ := strconv.Atoi(r.FormValue("length"))
		result, metadata, err = h.glodokSvc.GetTableJenisDestinasi(ctx, r.FormValue("jenisdestinasiid"), r.FormValue("jenisdestinasiket"), page, length)
	case "getsejarahberanda":
		result, err = h.glodokSvc.GetSejarahBeranda(ctx)
	case "getmaps":
		result, err = h.glodokSvc.GetMaps(ctx)
	case "getimagefotoberanda":
		result, err = h.glodokSvc.GetImageFotoBeranda(ctx, r.FormValue("fotoberandaid"))
		if err != nil {
			http.Error(w, "Failed to get image data", http.StatusInternalServerError)
			return
		}

		// Type assertion
		imgData, ok := result.([]byte)
		if !ok {
			log.Fatal("The result is not of type []byte")
		}

		// Create a buffer from the image data
		imgBuffer := bytes.NewReader(imgData)

		// Decode the image data to get the image.Image object
		img, imgFormat, err := image.Decode(imgBuffer)
		if err != nil {
			http.Error(w, "Unable to decode image", http.StatusInternalServerError)
			return
		}

		// Set the appropriate header for the image format
		var contentType string
		switch imgFormat {
		case "png":
			contentType = "image/png"
		case "jpeg", "jpg":
			contentType = "image/jpeg"
		default:
			http.Error(w, "Unsupported image format", http.StatusUnsupportedMediaType)
			return
		}
		w.Header().Set("Content-Type", contentType)

		// Encode the image to the appropriate format and write it to the response
		switch imgFormat {
		case "png":
			err = png.Encode(w, img)
		case "jpeg", "jpg":
			err = jpeg.Encode(w, img, nil)
		default:
			http.Error(w, "Unsupported image format", http.StatusUnsupportedMediaType)
			return
		}

		if err != nil {
			http.Error(w, "Unable to encode image", http.StatusInternalServerError)
			return
		}
	case "gettablefotoberanda":
		page, _ := strconv.Atoi(r.FormValue("page"))
		length, _ := strconv.Atoi(r.FormValue("length"))
		result, metadata, err = h.glodokSvc.GetTableFotoBeranda(ctx, r.FormValue("fotoberandaid"), page, length)
	case "gettablevideoberanda":
		page, _ := strconv.Atoi(r.FormValue("page"))
		length, _ := strconv.Atoi(r.FormValue("length"))
		result, metadata, err = h.glodokSvc.GetTableVideoBeranda(ctx, r.FormValue("videoberandaid"), page, length)
	case "gettabletujuan":
		page, _ := strconv.Atoi(r.FormValue("page"))
		length, _ := strconv.Atoi(r.FormValue("length"))
		result, metadata, err = h.glodokSvc.GetTableTujuanTransportasi(ctx, r.FormValue("tujuanid"), r.FormValue("tipetransportasiname"), r.FormValue("tujuanawal"), r.FormValue("tujuanakhir"), page, length)
	case "gettablepemberhentian":
		page, _ := strconv.Atoi(r.FormValue("page"))
		length, _ := strconv.Atoi(r.FormValue("length"))
		result, metadata, err = h.glodokSvc.GetTablePemberhentianTransportasi(ctx, r.FormValue("pemberhentianid"), r.FormValue("tipetransportasiname"), r.FormValue("pemberhentianname"), page, length)
	case "gettableuser":
		page, _ := strconv.Atoi(r.FormValue("page"))
		length, _ := strconv.Atoi(r.FormValue("length"))
		result, metadata, err = h.glodokSvc.GetTableUser(ctx, r.FormValue("userid"), r.FormValue("username"), page, length)

		//for masyarakat
	case "getdestinasibyid":
		result, err = h.glodokSvc.GetDestinasiByID(ctx, r.FormValue("destinasiid"))
	case "getalldestinasi":
		result, err = h.glodokSvc.GetAllDestinasi(ctx, r.FormValue("jenisdestinasiid"), r.FormValue("destinasiname"))
	case "getallreview":
		page, _ := strconv.Atoi(r.FormValue("page"))
		length, _ := strconv.Atoi(r.FormValue("length"))
		result, metadata, err = h.glodokSvc.GetAllReview(ctx, r.FormValue("destinasiid"), r.FormValue("rating"), page, length)
	case "getavgreview":
		result, metadata, err = h.glodokSvc.GetAvgReview(ctx, r.FormValue("destinasiid"))
	case "getfotoberandaml":
		result, err = h.glodokSvc.GetFotoBerandaML(ctx)
	case "getvideoberandaml":
		result, err = h.glodokSvc.GetVideoBerandaML(ctx)
	case "gettransportasiml":
		fmt.Println("tes", r.FormValue("perbaikanyn"))
		result, metadata, err = h.glodokSvc.GetTransportasiML(ctx, r.FormValue("perbaikanyn"))
	case "getberitaml":
		page, _ := strconv.Atoi(r.FormValue("page"))
		length, _ := strconv.Atoi(r.FormValue("length"))
		result, metadata, err = h.glodokSvc.GetBeritaML(ctx, r.FormValue("judul"), page, length)
	case "getberitabyid":
		result, err = h.glodokSvc.GetBeritaMLByID(ctx, r.FormValue("beritaid"))
	case "getjenisdestinasiml":
		result, err = h.glodokSvc.GetJenisDestinasiML(ctx)
	case "getdestinasiddml":
		result, err = h.glodokSvc.GetDestinasiDDML(ctx)
	case "getuser":
		result, err = h.glodokSvc.GetUser(ctx, r.FormValue("userid"))

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
