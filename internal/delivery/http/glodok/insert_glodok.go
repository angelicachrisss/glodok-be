package glodok

import (
	// "bytes"
	"encoding/json"
	glodokEntity "glodok-be/internal/entity/glodok"
	"glodok-be/pkg/response"
	"io/ioutil"
	"log"
	"net/http"

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

		InsertAdmin     glodokEntity.GetAdmin
		TableDestinasiIc glodokEntity.TableDestinasiIc
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
	case "insertdestinasi-ic":
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &TableDestinasiIc)
		result, err = h.glodokSvc.InsertDestinasiIc(ctx, TableDestinasiIc)
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
