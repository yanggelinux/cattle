package processorder

import (
	"context"
	"github.com/yanggelinux/cattle/internal/store/model"
)

type ProcessArchRepo interface {
	WithTx(*model.MDB) ProcessArchRepo
	GetByID(context.Context, int64) (*model.ProcessArch, error)
	Create(context.Context, *model.ProcessArch) error
	Update(context.Context, int64, map[string]interface{}) error
	CreateOrUpdate(context.Context, *model.ProcessArch) error
	DeleteByID(context.Context, int64) error
	GetImageHash(context.Context, string) (*model.ProcessArch, error)
	GetEnabledImageHash(context.Context, string) (*model.ProcessArch, error)
}
