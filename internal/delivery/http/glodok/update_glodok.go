package glodok

import (
	// "encoding/json"
	// "io/ioutil"
	"encoding/json"
	httpHelper "glodok-be/internal/delivery/http"
	"io/ioutil"

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
