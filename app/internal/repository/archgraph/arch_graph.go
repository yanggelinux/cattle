package archgraph

import (
	"context"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/store/model"
)

type ArchGraphRepo interface {
	WithTx(*model.MDB) ArchGraphRepo
	GetList(context.Context, *ArchGraphFilter) ([]*result.ArchGraphRecord, int64, error)
	GetCount(context.Context, *ArchGraphFilter) (int64, error)
	GetByID(context.Context, int64) (*model.ArchGraph, error)
	GetAll(context.Context) ([]*model.ArchGraph, error)
	GetSomeAll(context.Context, string) ([]*model.ArchGraph, error)
	GetByName(context.Context, string) (*model.ArchGraph, error)
	GetByLabel(context.Context, string) (*model.ArchGraph, error)
	GetCountByGroup(context.Context, int64) (int64, error)
	GetEnabledList(context.Context, string) ([]*model.ArchGraph, error)
	CheckDuplicateEntry(context.Context, int64, string) bool
	CheckHas(context.Context, int64) bool
	Create(context.Context, *model.ArchGraph) error
	Update(context.Context, int64, map[string]interface{}) error
	UpdateByIDs(context.Context, []int64, map[string]interface{}) error
	CreateOrUpdate(context.Context, *model.ArchGraph) error
	DeleteByID(context.Context, int64) error
}
