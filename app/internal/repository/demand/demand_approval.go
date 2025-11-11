package demand

import (
	"context"
	"github.com/yanggelinux/cattle/internal/store/model"
)

type DemandApprovalRepo interface {
	WithTx(*model.MDB) DemandApprovalRepo
	GetByID(context.Context, int64) (*model.DemandApproval, error)
	Create(context.Context, *model.DemandApproval) error
	Update(context.Context, int64, map[string]interface{}) error
	CreateOrUpdate(context.Context, *model.DemandApproval) error
	DeleteByID(context.Context, int64) error
	GetByDemand(context.Context, int64) ([]*model.DemandApproval, error)
}
