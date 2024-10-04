package glodok

import (
	httpHelper "glodok-be/internal/delivery/http"
	"glodok-be/pkg/response"
	"log"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go.uber.org/zap"
)

func (h *Handler) DeleteGlodok(w http.ResponseWriter, r *http.Request) {
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

	types = r.FormValue("type")
	switch types {
	case "deleteadmin":
		result, err = h.glodokSvc.DeleteAdmin(ctx, (r.FormValue("adminid")))
	case "deletedestinasi":
		result, err = h.glodokSvc.DeleteDestinasi(ctx, (r.FormValue("destinasiid")))
	case "deletetipetransportasi":
		result, err = h.glodokSvc.DeleteTipeTransportasi(ctx, (r.FormValue("tipetransportasiid")))
	case "deleterutetransportasi":
		result, err = h.glodokSvc.DeleteRuteTransportasi(ctx, (r.FormValue("ruteid")))
	case "deletereview":
		result, err = h.glodokSvc.DeleteReview(ctx, (r.FormValue("reviewid")))
	case "deleteberita":
		result, err = h.glodokSvc.DeleteBerita(ctx, (r.FormValue("beritaid")))
	case "deletejenisdestinasi":
		result, err = h.glodokSvc.DeleteJenisDestinasi(ctx, (r.FormValue("jenisdestinasiid")))
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
