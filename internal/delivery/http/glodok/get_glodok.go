package glodok

import (
	// "internal/itoa"

	httpHelper "glodok-be/internal/delivery/http"
	"glodok-be/pkg/response"
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
	case "getrutetransportasi":
		page, _ := strconv.Atoi(r.FormValue("page"))
		length, _ := strconv.Atoi(r.FormValue("length"))
		result, metadata, err = h.glodokSvc.GetTableRuteTransportasi(ctx, page, length)
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
