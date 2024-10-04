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

	GetTableDestinasi(ctx context.Context, ket string, page int, length int) ([]glodokEntity.TableDestinasi, interface{}, error)
	GetImageDestinasi(ctx context.Context, destinasiid string, destinasikat string) ([]byte, error)
	GetSearchDestinasi(ctx context.Context, kategori string, destinasiid string, destinasiname string, page int, length int) ([]glodokEntity.TableDestinasi, interface{}, error)
	UpdateDestinasi(ctx context.Context, destinasi glodokEntity.TableDestinasi, destinasiid string) (string, error)

	GetTableTipeTransportasi(ctx context.Context, page int, length int) ([]glodokEntity.TableTipeTransportasi, interface{}, error)
	GetSearchTipeTransportasi(ctx context.Context, tipetransportasiid string, tipetransportasiname string, page int, length int) ([]glodokEntity.TableTipeTransportasi, interface{}, error)

	GetTipeTransportasi(ctx context.Context) ([]glodokEntity.TableTipeTransportasi, error)
	GetTableRuteTransportasi(ctx context.Context, page int, length int) ([]glodokEntity.TableRuteTransportasi, interface{}, error)
	GetSearchRuteTransportasi(ctx context.Context, tipetransportasiname string, tujuanawal string, tujuanakhir string, page int, length int) ([]glodokEntity.TableRuteTransportasi, interface{}, error)

	GetTableReview(ctx context.Context, page int, length int) ([]glodokEntity.TableReview, interface{}, error)
	GetSearchReview(ctx context.Context, reviewid string, reviewer string, page int, length int) ([]glodokEntity.TableReview, interface{}, error)
	GetTableReviewByRating(ctx context.Context, rating int, page int, length int) ([]glodokEntity.TableReview, interface{}, error)
	GetSearchReviewByRating(ctx context.Context, rating int, reviewid string, reviewer string, page int, length int) ([]glodokEntity.TableReview, interface{}, error)

	GetDestinasi(ctx context.Context) ([]glodokEntity.TableDestinasi, error)
	GetImageBerita(ctx context.Context, beritaid string) ([]byte, error)
	GetTableBerita(ctx context.Context, page int, length int) ([]glodokEntity.TableBerita, interface{}, error)
	GetSearchBerita(ctx context.Context, beritaid string, destinasiname string, beritajudul string, page int, length int) ([]glodokEntity.TableBerita, interface{}, error)
	
	GetTableJenisDestinasi(ctx context.Context, jenisdestinasiid string, jenisdestinasiket string, page int, length int) ([]glodokEntity.TableJenisDestinasi, interface{}, error)

	//--------------------------------------------------------------------------------------------------------------------------------

	//insert
	SubmitLogin(ctx context.Context, adminid string, adminpass string) (string, error)

	InsertDestinasi(ctx context.Context, destinasi glodokEntity.TableDestinasi) (string, error)

	InsertTipeTransportasi(ctx context.Context, tipetransportasi glodokEntity.TableTipeTransportasi) (string, error)

	InsertRuteTransportasi(ctx context.Context, rutetransportasi glodokEntity.TableRuteTransportasi) (string, error)

	InsertReview(ctx context.Context, review glodokEntity.TableReview) (string, error)

	InsertBerita(ctx context.Context, berita glodokEntity.TableBerita) (string, error)

	InsertJenisDestinasi(ctx context.Context, jenisdestinasi glodokEntity.TableJenisDestinasi) (string, error) 

	//--------------------------------------------------------------------------------------------------------------------------------

	//update
	UpdateAdmin(ctx context.Context, admin glodokEntity.GetAdmin, adminid string) (string, error)
	UpdateTipeTransportasi(ctx context.Context, tipetransportasi glodokEntity.TableTipeTransportasi, tipetransportasiid string) (string, error)
	UpdateRuteTransportasi(ctx context.Context, rutetransportasi glodokEntity.TableRuteTransportasi, ruteid string) (string, error)
	UpdateBerita(ctx context.Context, berita glodokEntity.TableBerita, beritaid string) (string, error)
	UpdateJenisDestinasi(ctx context.Context, jenisdestinasi glodokEntity.TableJenisDestinasi, jenisdestinasiid string) (string, error) 

	//--------------------------------------------------------------------------------------------------------------------------------

	//delete
	DeleteAdmin(ctx context.Context, adminid string) (string, error)
	DeleteDestinasi(ctx context.Context, destinasiid string) (string, error)
	DeleteTipeTransportasi(ctx context.Context, tipetransportasiid string) (string, error)
	DeleteRuteTransportasi(ctx context.Context, ruteid string) (string, error)
	DeleteReview(ctx context.Context, reviewid string) (string, error)
	DeleteBerita(ctx context.Context, beritaid string) (string, error)
	DeleteJenisDestinasi(ctx context.Context, jenisdestinasiid string) (string, error) 

	//for masyarakat
	GetDestinasiByID(ctx context.Context, destinasiid string) ([]glodokEntity.TableDestinasi, error)
	GetAllDestinasi(ctx context.Context, kategori string, labelhalal string, destinasiname string) ([]glodokEntity.TableDestinasi, error)
	GetAllReview(ctx context.Context, rating string, page int, length int) ([]glodokEntity.TableReview, interface{}, error) 
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
