package role

import (
	"context"
	"github.com/yanggelinux/cattle/internal/store/model"
)

type RoleRepo interface {
	WithTx(*model.MDB) RoleRepo
	GetList(context.Context, *RoleFilter) ([]*model.Role, int64, error)
	GetByID(context.Context, int64) (*model.Role, error)
	GetByName(context.Context, string) (*model.Role, error)
	CheckDuplicateEntry(context.Context, string) bool
	Create(context.Context, *model.Role) error
	Update(context.Context, int64, map[string]interface{}) error
	CreateOrUpdate(context.Context, *model.Role) error
	DeleteByID(context.Context, int64) error
	GetRoles(context.Context, int64) ([]*model.Role, error)
	CreatePermRelInBatches(context.Context, []*model.RolePermRel) error
	DeletePermRelByIDs(context.Context, int64, []int64) error
}
