package process

import (
	"context"
	"github.com/yanggelinux/cattle/internal/store/model"
)

type ProcessRepo interface {
	WithTx(*model.MDB) ProcessRepo
	GetList(context.Context, *ProcessFilter) ([]*model.Process, int64, error)
	GetAll(context.Context) ([]*model.Process, error)
	GetByID(context.Context, int64) (*model.Process, error)
	CheckDuplicateEntry(context.Context, string) bool
	Create(context.Context, *model.Process) error
	Update(context.Context, int64, map[string]interface{}) error
	CreateOrUpdate(context.Context, *model.Process) error
	DeleteByID(context.Context, int64) error
}
