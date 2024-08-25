package glodok

import (
	// "internal/itoa"

	"bytes"
	"encoding/base64"
	"encoding/json"
	httpHelper "glodok-be/internal/delivery/http"
	glodokEntity "glodok-be/internal/entity/glodok"
	"glodok-be/pkg/response"
	"image"
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
	case "gettableadmin":
		page, _ := strconv.Atoi(r.FormValue("page"))
		length, _ := strconv.Atoi(r.FormValue("length"))
		result, metadata, err = h.glodokSvc.GetTableAdmin(ctx, page, length)
	case "getsearchadmin":
		page, _ := strconv.Atoi(r.FormValue("page"))
		length, _ := strconv.Atoi(r.FormValue("length"))
		result, metadata, err = h.glodokSvc.GetSearchAdmin(ctx, r.FormValue("adminid"), page, length)
	case "getdestinasi":
		page, _ := strconv.Atoi(r.FormValue("page"))
		length, _ := strconv.Atoi(r.FormValue("length"))
		result, metadata, err = h.glodokSvc.GetTableDestinasi(ctx, r.FormValue("ket"), page, length)

		// Assert that data is of type []Destinasi
		destinasiList, ok := result.([]glodokEntity.TableDestinasi)
		if !ok {
			http.Error(w, "Invalid data format", http.StatusInternalServerError)
			return
		}

		// Set Content-Type to application/json
		w.Header().Set("Content-Type", "application/json")

		// Prepare a slice to hold encoded data
		encodedDestinasiList := []map[string]interface{}{}

		for _, dest := range destinasiList {
			// Encode image data to PNG format and write it to a buffer
			var buf bytes.Buffer
			imgBuffer := bytes.NewReader([]byte(dest.DestinasiGambar))
			img, _, err := image.Decode(imgBuffer)
			if err != nil {
				http.Error(w, "Unable to decode image", http.StatusInternalServerError)
				return
			}

			err = png.Encode(&buf, img)
			if err != nil {
				http.Error(w, "Unable to encode image", http.StatusInternalServerError)
				return
			}

			// Convert encoded image data to Base64
			encodedImageData := base64.StdEncoding.EncodeToString(buf.Bytes())

			// Add to the response slice
			encodedDestinasi := map[string]interface{}{
				"destinasi_id":         dest.DestinasiID,
				"destinasi_name":       dest.DestinasiName,
				"destinasi_desc":       dest.DestinasiDesc,
				"destinasi_alamat":     dest.DestinasiAlamat,
				"destinasi_gambar":     dest.DestinasiGambar,
				"destinasi_lang":       dest.DestinasiLang,
				"destinasi_long":       dest.DestinasiLong,
				"destinasi_hbuka":      dest.DestinasiHBuka,
				"destinasi_htutup":     dest.DestinasiHTutup,
				"destinasi_kat":        dest.DestinasiKet,
				"destinasi_labelhalal": dest.DestinasiHalal,
				"destinasi_pic":        encodedImageData,
			}
			encodedDestinasiList = append(encodedDestinasiList, encodedDestinasi)
		}

		// Encode the entire list of destinasi as JSON
		err = json.NewEncoder(w).Encode(encodedDestinasiList)
		if err != nil {
			http.Error(w, "Unable to encode JSON response", http.StatusInternalServerError)
			return
		}
		result = encodedDestinasiList
		// --------------------------------------------------------------------------------

		// // Assert result as a slice of Destinasi
		// destinasiSlice, ok := result.([]glodok.TableDestinasi)
		// if !ok {
		// 	http.Error(w, "Invalid data type", http.StatusInternalServerError)
		// 	return
		// }

		// // fmt.Println("result", result)
		// fmt.Println("destinasiSlice")

		// // Ensure there's at least one item in the slice
		// if len(destinasiSlice) == 0 {
		// 	http.Error(w, "No images found", http.StatusNotFound)
		// 	return
		// }
		// fmt.Println("destinasiSlice2", len(destinasiSlice))
		// var responses []map[string]interface{}
		// for _, y := range destinasiSlice {
		// 	// Type assertion
		// 	imgData := y.DestinasiGambar

		// 	fmt.Println("imgData", y.DestinasiID)

		// 	// Create a buffer from the image data
		// 	imgBuffer := bytes.NewReader(imgData)

		// 	// Decode the image data to get the image.Image object
		// 	img, _, err := image.Decode(imgBuffer)
		// 	if err != nil {
		// 		http.Error(w, "Unable to decode image", http.StatusInternalServerError)
		// 		return
		// 	}

		// 	// // Set the appropriate header for PNG image
		// 	// w.Header().Set("Content-Type", "image/png")

		// 	// Create a buffer to encode the image data
		// 	var buf bytes.Buffer

		// 	// Encode the image to PNG format and write it to the response
		// 	err = png.Encode(&buf, img)
		// 	if err != nil {
		// 		http.Error(w, "Unable to encode image", http.StatusInternalServerError)
		// 		return
		// 	}
		// 	// destinasiSlice[y].DestinasiPic = img

		// 	// Convert encoded image data to Base64
		// 	encodedImageData := base64.StdEncoding.EncodeToString(buf.Bytes())

		// 	response := map[string]interface{}{
		// 		"destinasi_id":         y.DestinasiID,
		// 		"destinasi_name":       y.DestinasiName,
		// 		"destinasi_desc":       y.DestinasiDesc,
		// 		"destinasi_alamat":     y.DestinasiAlamat,
		// 		"destinasi_gambar":     y.DestinasiGambar,
		// 		"destinasi_lang":       y.DestinasiLang,
		// 		"destinasi_long":       y.DestinasiLong,
		// 		"destinasi_hbuka":      y.DestinasiHBuka,
		// 		"destinasi_htutup":     y.DestinasiHTutup,
		// 		"destinasi_kat":        y.DestinasiKet,
		// 		"destinasi_labelhalal": y.DestinasiHalal,
		// 		"destinasi_pic":        encodedImageData,
		// 	}
		// 	responses = append(responses, response)
		// }

		// // Set Content-Type header to application/json
		// w.Header().Set("Content-Type", "application/json")

		// fmt.Println("responses", responses)

		// // Encode response to JSON and write it to the response writer
		// err = json.NewEncoder(w).Encode(responses)
		// if err != nil {
		// 	http.Error(w, "Unable to encode JSON response", http.StatusInternalServerError)
		// 	return
		// }
		// result = responses
		// --------------------------------------------------------------------------------
		// // Type assertion
		// imgData := destinasiSlice[0].DestinasiGambar

		// fmt.Println("imgData", imgData)

		// // Create a buffer from the image data
		// imgBuffer := bytes.NewReader(imgData)

		// // Decode the image data to get the image.Image object
		// img, _, err := image.Decode(imgBuffer)
		// if err != nil {
		// 	http.Error(w, "Unable to decode image", http.StatusInternalServerError)
		// 	return
		// }

		// // Set the appropriate header for PNG image
		// w.Header().Set("Content-Type", "image/png")

		// // Encode the image to PNG format and write it to the response
		// err = png.Encode(w, img)
		// if err != nil {
		// 	http.Error(w, "Unable to encode image", http.StatusInternalServerError)
		// 	return
		// }
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
