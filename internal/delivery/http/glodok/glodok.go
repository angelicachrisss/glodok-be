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
	GetTableAdmin(ctx context.Context, page int, length int) ([]glodokEntity.GetAdmin, interface{}, error)
	GetSearchAdmin(ctx context.Context, adminid string, adminname string,page int, length int) ([]glodokEntity.GetAdmin, interface{}, error)

	GetTableDestinasi(ctx context.Context, ket string, page int, length int) ([]glodokEntity.TableDestinasi, interface{}, error)
	GetImageDestinasi(ctx context.Context, destinasiid string, destinasikat string) ([]byte, error)
	GetSearchDestinasi(ctx context.Context, kategori string, destinasiid string, destinasiname string, page int, length int) ([]glodokEntity.TableDestinasi, interface{}, error) 
	UpdateDestinasi(ctx context.Context, destinasi glodokEntity.TableDestinasi, destinasiid string) (string, error) 
	
	GetTableTipeTransportasi(ctx context.Context, page int, length int) ([]glodokEntity.TableTipeTransportasi, interface{}, error)
	GetSearchTipeTransportasi(ctx context.Context, tipetransportasiid string, tipetransportasiname string, page int, length int) ([]glodokEntity.TableTipeTransportasi, interface{}, error)

	GetTipeTransportasi(ctx context.Context) ([]glodokEntity.TableTipeTransportasi, error)
	GetTableRuteTransportasi(ctx context.Context, page int, length int) ([]glodokEntity.TableRuteTransportasi, interface{}, error)
	GetSearchRuteTransportasi(ctx context.Context, tipetransportasiname string, tujuanawal string, tujuanakhir string, page int, length int) ([]glodokEntity.TableRuteTransportasi, interface{}, error)

	//--------------------------------------------------------------------------------------------------------------------------------

	//insert
	InsertAdmin(ctx context.Context, admin glodokEntity.GetAdmin) (string, error)
	SubmitLogin(ctx context.Context, adminid string, adminpass string) (string, error)

	InsertDestinasi(ctx context.Context, destinasi glodokEntity.TableDestinasi) (string, error)

	InsertTipeTransportasi(ctx context.Context, tipetransportasi glodokEntity.TableTipeTransportasi) (string, error)

	InsertRuteTransportasi(ctx context.Context, rutetransportasi glodokEntity.TableRuteTransportasi) (string, error)

	//--------------------------------------------------------------------------------------------------------------------------------

	//update
	UpdateAdmin(ctx context.Context, admin glodokEntity.GetAdmin, adminid string) (string, error)
	UpdateTipeTransportasi(ctx context.Context, tipetransportasi glodokEntity.TableTipeTransportasi, tipetransportasiid string) (string, error)
	UpdateRuteTransportasi(ctx context.Context, rutetransportasi glodokEntity.TableRuteTransportasi, ruteid string) (string, error)

	//--------------------------------------------------------------------------------------------------------------------------------

	//delete
	DeleteAdmin(ctx context.Context, adminid string) (string, error)
	DeleteDestinasi(ctx context.Context, destinasiid string) (string, error)
	DeleteTipeTransportasi(ctx context.Context, tipetransportasiid string) (string, error)
	DeleteRuteTransportasi(ctx context.Context, ruteid string) (string, error)
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
