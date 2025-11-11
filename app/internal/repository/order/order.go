package order

import (
	"context"
	"github.com/yanggelinux/cattle/internal/store/model"
)

type OrderRepo interface {
	WithTx(*model.MDB) OrderRepo
	GetList(context.Context, *OrderFilter) ([]*model.Order, int64, error)
	GetAll(context.Context) ([]*model.Order, error)
	GetByID(context.Context, int64) (*model.Order, error)
	GetByName(context.Context, string) (*model.Order, error)
	GetByLabel(context.Context, string) ([]*model.Order, error)
	GetByLabels(context.Context, []string) ([]*model.Order, error)
	CheckDuplicateEntry(context.Context, string) bool
	CheckDuplicateLabel(context.Context, string) bool
	Create(context.Context, *model.Order) error
	Update(context.Context, int64, map[string]interface{}) error
	CreateOrUpdate(context.Context, *model.Order) error
	DeleteByID(context.Context, int64) error
}
