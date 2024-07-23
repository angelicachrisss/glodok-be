package glodok

import (
	"glodok-be/internal/entity"
	// "glodok-be/internal/entity/auth"
	"context"
	"errors"
	jaegerLog "glodok-be/pkg/log"

	// "time"

	glodokEntity "glodok-be/internal/entity/glodok"

	"github.com/opentracing/opentracing-go"
)

type Data interface {
	// //get
	// GetKaryawan(ctx context.Context) ([]glodokEntity.GetKaryawan, error)
	// GetCountKaryawan(ctx context.Context) (int, error)
	// //insert
	// InsertKaryawan(ctx context.Context, karyawan glodokEntity.GetKaryawan) (string, error)

	// get
	GetAdmin(ctx context.Context) ([]glodokEntity.GetAdmin, error)
	GetAdminbyID(ctx context.Context, adminid string) ([]glodokEntity.GetAdmin, error)

	//insert
	InsertAdmin(ctx context.Context, admin glodokEntity.GetAdmin) (string, error)
	SubmitLogin(ctx context.Context, adminid string, adminpass string) (string, error)
}

type Service struct {
	glodok Data
	tracer opentracing.Tracer
	logger jaegerLog.Factory
}

// New ...
// Tambahkan parameter sesuai banyak data layer yang dibutuhkan
func New(glodokData Data, tracer opentracing.Tracer, logger jaegerLog.Factory) Service {
	// Assign variable dari parameter ke object
	return Service{
		glodok: glodokData,
		tracer: tracer,
		logger: logger,
	}
}

func (s Service) checkPermission(ctx context.Context, _permissions ...string) error {
	claims := ctx.Value(entity.ContextKey("claims"))
	if claims != nil {
		actions := claims.(entity.ContextValue).Get("permissions").(map[string]interface{})
		for _, action := range actions {
			permissions := action.([]interface{})
			for _, permission := range permissions {
				for _, _permission := range _permissions {
					if permission.(string) == _permission {
						return nil
					}
				}
			}
		}
	}
	return errors.New("401 unauthorized")
}
