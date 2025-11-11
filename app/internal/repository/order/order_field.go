package order

import (
	"context"
	"github.com/yanggelinux/cattle/internal/store/model"
)

type OrderFieldRepo interface {
	WithTx(*model.MDB) OrderFieldRepo
	GetList(context.Context, *OrderFieldFilter) ([]*model.OrderField, int64, error)
	GetByID(context.Context, int64) (*model.OrderField, error)
	GetByOrderIDAndKey(context.Context, int64, string) (*model.OrderField, error)
	GetByOrder(context.Context, int64) ([]*model.OrderField, error)
	CheckDuplicateEntry(context.Context, string, int64) bool
	Create(context.Context, *model.OrderField) error
	CreateInBatches(context.Context, []*model.OrderField) error
	Update(context.Context, int64, map[string]interface{}) error
	CreateOrUpdate(context.Context, *model.OrderField) error
	DeleteByID(context.Context, int64) error
}
