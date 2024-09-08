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
	//--admin
	GetAdmin(ctx context.Context) ([]glodokEntity.GetAdmin, error)
	GetAdminbyID(ctx context.Context, adminid string) ([]glodokEntity.GetAdmin, error)
	GetTableAdmin(ctx context.Context, page int, length int) ([]glodokEntity.GetAdmin, interface{}, error)
	GetSearchAdmin(ctx context.Context, adminid string, page int, length int) ([]glodokEntity.GetAdmin, interface{}, error)

	//--destinasi
	GetTableDestinasi(ctx context.Context, ket string, page int, length int) ([]glodokEntity.TableDestinasi, interface{}, error)

	//--tipetransportasi
	GetTableTipeTransportasi(ctx context.Context, page int, length int) ([]glodokEntity.TableTipeTransportasi, interface{}, error)
	GetSearchTipeTransportasi(ctx context.Context, tipetransportasiid string, tipetransportasiname string, page int, length int) ([]glodokEntity.TableTipeTransportasi, interface{}, error)

	//--rutetransportasi
	GetTipeTransportasi(ctx context.Context) ([]glodokEntity.TableTipeTransportasi, error)
	GetTableRuteTransportasi(ctx context.Context, page int, length int) ([]glodokEntity.TableRuteTransportasi, interface{}, error)

	//--------------------------------------------------------------------------------------------------------------------------------

	//insert
	//--admin
	InsertAdmin(ctx context.Context, admin glodokEntity.GetAdmin) (string, error)
	SubmitLogin(ctx context.Context, adminid string, adminpass string) (string, error)

	//--destinasi
	InsertDestinasi(ctx context.Context, destinasi glodokEntity.TableDestinasi) (string, error)

	//--tipetransportasi
	InsertTipeTransportasi(ctx context.Context, tipetransportasi glodokEntity.TableTipeTransportasi) (string, error)

	//--rutetransportasi
	InsertRuteTransportasi(ctx context.Context, rutetransportasi glodokEntity.TableRuteTransportasi) (string, error)

	//--------------------------------------------------------------------------------------------------------------------------------

	//update
	//--admin
	UpdateAdmin(ctx context.Context, admin glodokEntity.GetAdmin, adminid string) (string, error)

	//--tipetransportasi
	UpdateTipeTransportasi(ctx context.Context, tipetransportasi glodokEntity.TableTipeTransportasi, tipetransportasiid string) (string, error)

	//--------------------------------------------------------------------------------------------------------------------------------

	//delete
	//--admin
	DeleteAdmin(ctx context.Context, adminid string) (string, error)

	//destinasi
	DeleteDestinasi(ctx context.Context, destinasiid string) (string, error)

	//tipetransportasi
	DeleteTipeTransportasi(ctx context.Context, tipetransportasiid string) (string, error)

	//rutetransportasi
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
