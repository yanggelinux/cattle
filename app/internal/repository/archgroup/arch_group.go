package archgroup

import (
	"context"
	"github.com/yanggelinux/cattle/internal/store/model"
)

type ArchGroupRepo interface {
	WithTx(*model.MDB) ArchGroupRepo
	GetList(context.Context, *ArchGroupFilter) ([]*model.ArchGroup, error)
	GetByID(context.Context, int64) (*model.ArchGroup, error)
	GetAll(context.Context) ([]*model.ArchGroup, error)
	GetByName(context.Context, string) (*model.ArchGroup, error)
	GetByParent(context.Context, int64) ([]*model.ArchGroup, error)
	GetCountByParent(context.Context, int64) (int64, error)
	CheckDuplicateEntry(context.Context, int64, string) bool
	CheckHas(context.Context, int64) bool
	Create(context.Context, *model.ArchGroup) error
	Update(context.Context, int64, map[string]interface{}) error
	CreateOrUpdate(context.Context, *model.ArchGroup) error
	DeleteByID(context.Context, int64) error
}
