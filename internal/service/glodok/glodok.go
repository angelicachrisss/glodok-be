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
	// get
	GetAdmin(ctx context.Context) ([]glodokEntity.GetAdmin, error)
	GetAdminbyID(ctx context.Context, adminid string) ([]glodokEntity.GetAdmin, error)
	GetTableAdmin(ctx context.Context, page int, length int) ([]glodokEntity.GetAdmin, error)
	GetCountAdmin(ctx context.Context) (int, error)
	GetSearchAdmin(ctx context.Context, adminid string, adminname string, page int, length int) ([]glodokEntity.GetAdmin, error)
	GetCountSearchAdmin(ctx context.Context, adminid string, adminname string) (int, error)

	GetTableDestinasi(ctx context.Context, ket string, page int, length int) ([]glodokEntity.TableDestinasi, error)
	GetCountDestinasi(ctx context.Context, ket string) (int, error)
	GetImageDestinasi(ctx context.Context, destinasiid string, destinasikat string) ([]byte, error)
	GetSearchDestinasi(ctx context.Context, kategori string, destinasiid string, destinasiname string, page int, length int) ([]glodokEntity.TableDestinasi, error)
	GetCountSearchDestinasi(ctx context.Context, kategori string, destinasiid string, destinasiname string) (int, error)
	UpdateDestinasi(ctx context.Context, destinasi glodokEntity.TableDestinasi, destinasiid string) (string, error)

	GetTableTipeTransportasi(ctx context.Context, page int, length int) ([]glodokEntity.TableTipeTransportasi, error)
	GetCountTableTipeTransportasi(ctx context.Context) (int, error)
	GetSearchTipeTransportasi(ctx context.Context, tipetransportasiid string, tipetransportasiname string, page int, length int) ([]glodokEntity.TableTipeTransportasi, error)
	GetCountSearchTipeTransportasi(ctx context.Context, tipetransportasiid string, tipetransportasiname string) (int, error)

	GetTipeTransportasi(ctx context.Context) ([]glodokEntity.TableTipeTransportasi, error)
	GetTableRuteTransportasi(ctx context.Context, page int, length int) ([]glodokEntity.TableRuteTransportasi, error)
	GetCountTableRuteTransportasi(ctx context.Context) (int, error)
	GetSearchRuteTransportasi(ctx context.Context, tipetransportasiname string, tujuanawal string, tujuanakhir string, page int, length int) ([]glodokEntity.TableRuteTransportasi, error)
	GetCountSearchRuteTransportasi(ctx context.Context, tipetransportasiname string, tujuanawal string, tujuanakhir string) (int, error)

	GetTableReview(ctx context.Context, page int, length int) ([]glodokEntity.TableReview, error)
	GetCountTableReview(ctx context.Context) (int, error)
	GetSearchReview(ctx context.Context, reviewid string, reviewer string, page int, length int) ([]glodokEntity.TableReview, error)
	GetCountSearchReview(ctx context.Context, reviewid string, reviewer string) (int, error)
	GetTableReviewByRating(ctx context.Context, rating int, page int, length int) ([]glodokEntity.TableReview, error)
	GetCountTableReviewByRating(ctx context.Context, rating int) (int, error)
	GetSearchReviewByRating(ctx context.Context, rating int, reviewid string, reviewer string, page int, length int) ([]glodokEntity.TableReview, error)
	GetCountSearchReviewByRating(ctx context.Context, rating int , reviewid string, reviewer string) (int, error) 

	GetDestinasi(ctx context.Context) ([]glodokEntity.TableDestinasi, error)
	GetTableBerita(ctx context.Context, page int, length int) ([]glodokEntity.TableBerita, error)
	GetImageBerita(ctx context.Context, beritaid string) ([]byte, error)
	GetCountBerita(ctx context.Context) (int, error)
	GetSearchBerita(ctx context.Context, beritaid string, destinasiname string, beritajudul string, page int, length int) ([]glodokEntity.TableBerita, error)
	GetCountSearchBerita(ctx context.Context, beritaid string, destinasiname string, beritajudul string) (int, error)

	//--------------------------------------------------------------------------------------------------
	//insert
	InsertAdmin(ctx context.Context, admin glodokEntity.GetAdmin) (string, error)
	SubmitLogin(ctx context.Context, adminid string, adminpass string) (string, error)

	InsertDestinasi(ctx context.Context, destinasi glodokEntity.TableDestinasi) (string, error)

	InsertTipeTransportasi(ctx context.Context, tipetransportasi glodokEntity.TableTipeTransportasi) (string, error)

	InsertRuteTransportasi(ctx context.Context, rutetransportasi glodokEntity.TableRuteTransportasi) (string, error)

	InsertReview(ctx context.Context, review glodokEntity.TableReview) (string, error)

	InsertBerita(ctx context.Context, berita glodokEntity.TableBerita) (string, error)

	//update
	UpdateAdmin(ctx context.Context, admin glodokEntity.GetAdmin, adminid string) (string, error)
	UpdateTipeTransportasi(ctx context.Context, tipetransportasi glodokEntity.TableTipeTransportasi, tipetransportasiid string) (string, error)
	UpdateRuteTransportasi(ctx context.Context, rutetransportasi glodokEntity.TableRuteTransportasi, ruteid string) (string, error)
	UpdateBerita(ctx context.Context, berita glodokEntity.TableBerita, beritaid string) (string, error)

	//delete
	DeleteAdmin(ctx context.Context, adminid string) (string, error)
	DeleteDestinasi(ctx context.Context, destinasiid string) (string, error)
	DeleteTipeTransportasi(ctx context.Context, tipetransportasiid string) (string, error)
	DeleteRuteTransportasi(ctx context.Context, ruteid string) (string, error)
	DeleteReview(ctx context.Context, reviewid string) (string, error)
	DeleteBerita(ctx context.Context, beritaid string) (string, error)
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
