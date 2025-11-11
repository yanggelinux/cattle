package processorder

import (
	"context"
	"github.com/yanggelinux/cattle/internal/store/model"
)

type ProcessApprovalRepo interface {
	WithTx(*model.MDB) ProcessApprovalRepo
	GetByID(context.Context, int64) (*model.ProcessApproval, error)
	Create(context.Context, *model.ProcessApproval) error
	Update(context.Context, int64, map[string]interface{}) error
	CreateOrUpdate(context.Context, *model.ProcessApproval) error
	DeleteByID(context.Context, int64) error
	GetByOrder(context.Context, int64) ([]*model.ProcessApproval, error)
	GetByApprover(context.Context, string) ([]*model.ProcessApproval, error)
}
