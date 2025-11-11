package permission

import (
	"context"
	"github.com/yanggelinux/cattle/internal/store/model"
)

type PermissionRepo interface {
	WithTx(*model.MDB) PermissionRepo
	GetList(context.Context, *PermissionFilter) ([]*model.Permission, int64, error)
	GetByID(context.Context, int64) (*model.Permission, error)
	GetByCode(context.Context, string) (*model.Permission, error)
	CheckDuplicateEntry(context.Context, string, string) bool
	Create(context.Context, *model.Permission) error
	Update(context.Context, int64, map[string]interface{}) error
	CreateOrUpdate(context.Context, *model.Permission) error
	DeleteByID(context.Context, int64) error

	GetPermsByRole(context.Context, int64) ([]*model.Permission, error)
	GetPermsByProject(context.Context, string) ([]*model.Permission, error)
}
