package order

import (
	"context"
	"github.com/yanggelinux/cattle/internal/store/model"
)

type OrderGroupRepo interface {
	WithTx(*model.MDB) OrderGroupRepo
	GetList(context.Context, *OrderGroupFilter) ([]*model.OrderGroup, int64, error)
	GetAll(context.Context) ([]*model.OrderGroup, error)
	GetByID(context.Context, int64) (*model.OrderGroup, error)
	CheckDuplicateEntry(context.Context, string) bool
	Create(context.Context, *model.OrderGroup) error
	Update(context.Context, int64, map[string]interface{}) error
	CreateOrUpdate(context.Context, *model.OrderGroup) error
	DeleteByID(context.Context, int64) error
}
