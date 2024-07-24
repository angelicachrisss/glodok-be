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
	// //get
	// GetKaryawan(ctx context.Context) ([]JOEntity.GetKaryawan, interface{}, error)
	// //insert
	// InsertKaryawan(ctx context.Context, karyawan JOEntity.InsertKaryawan) (string, error)

	// get
	GetAdmin(ctx context.Context) ([]glodokEntity.GetAdmin, error)
	GetAdminbyID(ctx context.Context, adminid string) ([]glodokEntity.GetAdmin, error)
	GetTableAdmin(ctx context.Context,page int, length int) ([]glodokEntity.GetAdmin,interface{}, error)
	GetSearchAdmin(ctx context.Context, adminid string, page int, length int) ([]glodokEntity.GetAdmin, interface{}, error)

	//insert
	InsertAdmin(ctx context.Context, admin glodokEntity.GetAdmin) (string, error)
	SubmitLogin(ctx context.Context, adminid string, adminpass string) (string, error)
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
