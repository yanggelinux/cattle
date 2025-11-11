package processorder

import (
	"context"
	"github.com/yanggelinux/cattle/internal/store/model"
)

type ProcessOrderRepo interface {
	WithTx(*model.MDB) ProcessOrderRepo
	GetList(context.Context, *ProcessOrderFilter) ([]*model.ProcessOrder, int64, error)
	GetApprovalList(context.Context, *ProcessOrderFilter) ([]*model.ProcessOrder, int64, error)
	GetCount(context.Context, *ProcessOrderFilter) (int64, error)
	GetCountByTime(context.Context, int8) ([]*model.ProcessOrderDist, error)
	GetByID(context.Context, int64) (*model.ProcessOrder, error)
	GetByIDStatus(context.Context, int64, int8) ([]*model.ProcessOrder, error)
	GetByParentID(context.Context, int64) ([]*model.ProcessOrder, error)
	GetByParentIDStatus(context.Context, int64, int8) ([]*model.ProcessOrder, error)
	GetTaskOrder(ctx context.Context, orderIDs []int64) ([]*model.ProcessOrder, error)
	Create(context.Context, *model.ProcessOrder) error
	Update(context.Context, int64, map[string]interface{}) error
	Updates(context.Context, []int64, map[string]interface{}) error
	CreateOrUpdate(context.Context, *model.ProcessOrder) error
	DeleteByID(context.Context, int64) error
}
