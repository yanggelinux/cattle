package team

import (
	"context"
	"github.com/yanggelinux/cattle/internal/store/model"
)

type TeamRepo interface {
	WithTx(*model.MDB) TeamRepo
	GetList(context.Context, *TeamFilter) ([]*model.Team, int64, error)
	GetByID(context.Context, int64) (*model.Team, error)
	GetByName(context.Context, string) (*model.Team, error)
	GetAll(context.Context) ([]*model.Team, error)
	CheckDuplicateEntry(context.Context, string) bool
	Create(context.Context, *model.Team) error
	Update(context.Context, int64, map[string]interface{}) error
	CreateOrUpdate(context.Context, *model.Team) error
	DeleteByID(context.Context, int64) error
}
