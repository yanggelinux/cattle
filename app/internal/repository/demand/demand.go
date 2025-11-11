package demand

import (
	"context"
	"github.com/yanggelinux/cattle/internal/store/model"
)

type DemandRepo interface {
	WithTx(*model.MDB) DemandRepo
	GetList(context.Context, *DemandFilter) ([]*model.Demand, int64, error)
	GetCount(context.Context, *DemandFilter) (int64, error)
	GetByID(context.Context, int64) (*model.Demand, error)
	GetByName(context.Context, string) (*model.Demand, error)
	GetByStatus(context.Context, int8) ([]*model.Demand, error)
	CheckDuplicateEntry(context.Context, string) bool
	Create(context.Context, *model.Demand) error
	Update(context.Context, int64, map[string]interface{}) error
	CreateOrUpdate(context.Context, *model.Demand) error
	DeleteByID(context.Context, int64) error
}
