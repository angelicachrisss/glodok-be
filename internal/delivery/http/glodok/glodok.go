package glodok

import (
	// "bytes"
	"context"
	glodokEntity "glodok-be/internal/entity/glodok"
	jaegerLog "glodok-be/pkg/log"

	// "context"

	"github.com/opentracing/opentracing-go"
)

type IglodokSvc interface {
	// get
	GetAdmin(ctx context.Context) ([]glodokEntity.GetAdmin, error)
	GetAdminbyID(ctx context.Context, adminid string) ([]glodokEntity.GetAdmin, error)

	GetJenisDestinasi(ctx context.Context) ([]glodokEntity.TableJenisDestinasi, error)
	// GetTableDestinasi(ctx context.Context, ket string, page int, length int) ([]glodokEntity.TableDestinasi, interface{}, error)
	GetImageDestinasi(ctx context.Context, destinasiid string) ([]byte, error)
	// GetSearchDestinasi(ctx context.Context, kategori string, destinasiid string, destinasiname string, page int, length int) ([]glodokEntity.TableDestinasi, interface{}, error)
	GetTableAllDestinasi(ctx context.Context, page int, length int) ([]glodokEntity.TableDestinasi, interface{}, error)
	GetTableDestinasiByJenis(ctx context.Context, jenisdestinasiid string, page int, length int) ([]glodokEntity.TableDestinasi, interface{}, error)
	GetSearchTableAllDestinasi(ctx context.Context, destinasiid string, destinasiname string, page int, length int) ([]glodokEntity.TableDestinasi, interface{}, error)
	GetSearchTableDestinasiByJenis(ctx context.Context, jenisdestinasiid string, destinasiid string, destinasiname string, page int, length int) ([]glodokEntity.TableDestinasi, interface{}, error)

	GetTableTipeTransportasi(ctx context.Context, page int, length int) ([]glodokEntity.TableTipeTransportasi, interface{}, error)
	GetSearchTipeTransportasi(ctx context.Context, tipetransportasiid string, tipetransportasiname string, page int, length int) ([]glodokEntity.TableTipeTransportasi, interface{}, error)

	GetTipeTransportasi(ctx context.Context) ([]glodokEntity.TableTipeTransportasi, error)
	GetTujuanTransportasiDropDown(ctx context.Context, tipetransportasiid string) ([]glodokEntity.TableTujuan, error)
	GetPemberhentianDropDown(ctx context.Context, tipetransportasiid string) ([]glodokEntity.TablePemberhentian, error)
	GetTableRuteTransportasi(ctx context.Context, page int, length int) ([]glodokEntity.TableRuteTransportasi, interface{}, error)
	GetSearchRuteTransportasi(ctx context.Context, tipetransportasiname string, tujuanawal string, tujuanakhir string, page int, length int) ([]glodokEntity.TableRuteTransportasi, interface{}, error)

	GetTableReview(ctx context.Context, destinasiid string, userid string, page int, length int) ([]glodokEntity.TableReview, interface{}, error)
	// GetSearchReview(ctx context.Context, reviewid string, reviewer string, page int, length int) ([]glodokEntity.TableReview, interface{}, error)
	// GetTableReviewByRating(ctx context.Context, rating int, page int, length int) ([]glodokEntity.TableReview, interface{}, error)
	// GetSearchReviewByRating(ctx context.Context, rating int, reviewid string, reviewer string, page int, length int) ([]glodokEntity.TableReview, interface{}, error)

	GetDestinasi(ctx context.Context) ([]glodokEntity.TableDestinasi, error)
	GetImageBerita(ctx context.Context, beritaid string) ([]byte, error)
	GetTableBerita(ctx context.Context, page int, length int) ([]glodokEntity.TableBerita, interface{}, error)
	GetSearchBerita(ctx context.Context, beritaid string, destinasiname string, beritajudul string, page int, length int) ([]glodokEntity.TableBerita, interface{}, error)

	GetTableJenisDestinasi(ctx context.Context, jenisdestinasiid string, jenisdestinasiket string, page int, length int) ([]glodokEntity.TableJenisDestinasi, interface{}, error)

	GetSejarahBeranda(ctx context.Context) (glodokEntity.TableSejarahBeranda, error)

	GetMaps(ctx context.Context) (glodokEntity.TableMaps, error)

	GetTableFotoBeranda(ctx context.Context, fotoberandaid string, page int, length int) ([]glodokEntity.TableFotoBeranda, interface{}, error)
	GetImageFotoBeranda(ctx context.Context, fotoberandaid string) ([]byte, error)

	GetTableVideoBeranda(ctx context.Context, videoberandaid string, page int, length int) ([]glodokEntity.TableVideoBeranda, interface{}, error)

	GetTableTujuanTransportasi(ctx context.Context, tujuanid string, tipetransportasiname string, tujuanawal string, tujuanakhir string, page int, length int) ([]glodokEntity.TableTujuan, interface{}, error)

	GetTablePemberhentianTransportasi(ctx context.Context, pemberhentianid string, tipetransportasiname string, pemberhentianname string, page int, length int) ([]glodokEntity.TablePemberhentian, interface{}, error)

	GetTableUser(ctx context.Context, userid string, username string, page int, length int) ([]glodokEntity.TableUser, interface{}, error)

	//--------------------------------------------------------------------------------------------------------------------------------

	//insert
	SubmitLogin(ctx context.Context, adminid string, adminpass string) (string, error)
	InsertDestinasi(ctx context.Context, destinasi glodokEntity.TableDestinasi) (string, error)
	InsertTipeTransportasi(ctx context.Context, tipetransportasi glodokEntity.TableTipeTransportasi) (string, error)
	InsertRuteTransportasi(ctx context.Context, rutetransportasi glodokEntity.TableRuteTransportasi) (string, error)
	InsertReview(ctx context.Context, review glodokEntity.TableReview) (string, error)
	InsertBerita(ctx context.Context, berita glodokEntity.TableBerita) (string, error)
	InsertJenisDestinasi(ctx context.Context, jenisdestinasi glodokEntity.TableJenisDestinasi) (string, error)
	InsertFotoBeranda(ctx context.Context, fotoberanda glodokEntity.TableFotoBeranda) (string, error)
	InsertVideoBeranda(ctx context.Context, videoberanda glodokEntity.TableVideoBeranda) (string, error)
	InsertTujuanTransportasi(ctx context.Context, tujuan glodokEntity.TableTujuan) (string, error)
	InsertPemberhentianTransportasi(ctx context.Context, pemberhentian glodokEntity.TablePemberhentian) (string, error)
	InsertUser(ctx context.Context, user glodokEntity.TableUser) (string, error)

	//--------------------------------------------------------------------------------------------------------------------------------

	//update
	UpdateAdmin(ctx context.Context, admin glodokEntity.GetAdmin, adminid string) (string, error)
	UpdateDestinasi(ctx context.Context, destinasi glodokEntity.TableDestinasi, destinasiid string) (string, error)
	UpdateStatusDestinasi(ctx context.Context, destinasi glodokEntity.TableDestinasi, destinasiid string) (string, error)
	UpdateTipeTransportasi(ctx context.Context, tipetransportasi glodokEntity.TableTipeTransportasi, tipetransportasiid string) (string, error)
	UpdateRuteTransportasi(ctx context.Context, rutetransportasi glodokEntity.TableRuteTransportasi, ruteid string) (string, error)
	UpdateBerita(ctx context.Context, berita glodokEntity.TableBerita, beritaid string) (string, error)
	UpdateJenisDestinasi(ctx context.Context, jenisdestinasi glodokEntity.TableJenisDestinasi, jenisdestinasiid string) (string, error)
	UpdateSejarahBeranda(ctx context.Context, sejarahberanda glodokEntity.TableSejarahBeranda) (string, error)
	UpdateMaps(ctx context.Context, maps glodokEntity.TableMaps, isi string) (string, error)
	UpdateTujuan(ctx context.Context, tujuan glodokEntity.TableTujuan, tujuanid string) (string, error)
	UpdatePemberhentian(ctx context.Context, pemberhentian glodokEntity.TablePemberhentian, tujuanid string) (string, error)

	//--------------------------------------------------------------------------------------------------------------------------------

	//delete
	DeleteAdmin(ctx context.Context, adminid string) (string, error)
	DeleteDestinasi(ctx context.Context, destinasiid string) (string, error)
	DeleteTipeTransportasi(ctx context.Context, tipetransportasiid string) (string, error)
	DeleteRuteTransportasi(ctx context.Context, ruteid string) (string, error)
	DeleteReview(ctx context.Context, reviewid string) (string, error)
	DeleteBerita(ctx context.Context, beritaid string) (string, error)
	DeleteJenisDestinasi(ctx context.Context, jenisdestinasiid string) (string, error)
	DeleteFotoBeranda(ctx context.Context, fotoberandaid string) (string, error)
	DeleteVideoBeranda(ctx context.Context, videoberandaid string) (string, error)
	DeleteTujuan(ctx context.Context, tujuanid string) (string, error)
	DeletePemberhentian(ctx context.Context, pemberhentianid string) (string, error)
	DeleteRuteByPemberhentian(ctx context.Context) (string, error)
	DeleteUser(ctx context.Context, userid string) (string, error)
	DeleteRuteByTujuan(ctx context.Context) (string, error)
	DeleteRuteByTipe(ctx context.Context) (string, error)
	DeletePemberhentianByTipe(ctx context.Context) (string, error)
	DeleteTujuanByTipe(ctx context.Context) (string, error)

	//for masyarakat
	GetDestinasiByID(ctx context.Context, destinasiid string) ([]glodokEntity.TableDestinasi, error)
	GetAllDestinasi(ctx context.Context, jenisdestinasiid string, destinasiname string) ([]glodokEntity.TableDestinasi, error)
	GetAllReview(ctx context.Context, destinasiid string, rating string, page int, length int) ([]glodokEntity.TableReview, interface{}, error)
	GetAvgReview(ctx context.Context, destinasiid string) (float64, interface{}, error)
	GetFotoBerandaML(ctx context.Context) ([]glodokEntity.TableFotoBeranda, error)
	GetVideoBerandaML(ctx context.Context) ([]glodokEntity.TableVideoBeranda, error)
	GetTransportasiML(ctx context.Context, perbaikanyn string) ([]glodokEntity.TableRuteTransportasi, interface{}, error)
	GetBeritaML(ctx context.Context, judul string, page int, length int) ([]glodokEntity.TableBerita, interface{}, error)
	GetBeritaMLByID(ctx context.Context, beritaid string) ([]glodokEntity.TableBerita, error)
	GetJenisDestinasiML(ctx context.Context) ([]glodokEntity.TableJenisDestinasi, error)
	GetDestinasiDDML(ctx context.Context) ([]glodokEntity.TableDestinasi, error)
	SubmitLoginML(ctx context.Context, userid string, pass string) (string, error)
	GetUser(ctx context.Context, userid string) (glodokEntity.TableUser, error)
	UpdateUser(ctx context.Context, user glodokEntity.TableUser, userid string) (string, error)
	DeleteReviewByUser(ctx context.Context) (string, error)
}

type (
	// Handler ...
	Handler struct {
		glodokSvc IglodokSvc
		tracer    opentracing.Tracer
		logger    jaegerLog.Factory
	}
)

// New for bridging product handler initialization
func New(is IglodokSvc, tracer opentracing.Tracer, logger jaegerLog.Factory) *Handler {
	return &Handler{
		glodokSvc: is,
		tracer:    tracer,
		logger:    logger,
	}
}
