package user

import (
	"context"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/store/model"
)

type UserRepo interface {
	WithTx(*model.MDB) UserRepo
	GetList(context.Context, *UserFilter) ([]*model.User, int64, error)
	GetUsers(context.Context) ([]*model.User, error)
	GetUsersByRole(context.Context, string) ([]*model.User, error)
	GetUsersByRoles(context.Context, []string) ([]*model.User, error)
	GetByID(context.Context, int64) (*model.User, error)
	GetByName(context.Context, string) (*model.User, error)
	GetBySysName(context.Context, string) (*model.User, error)
	GetByLdapName(context.Context, string) (*model.User, error)
	CheckDuplicateEntry(context.Context, string) bool
	Create(context.Context, *model.User) error
	Update(context.Context, int64, map[string]interface{}) error
	CreateOrUpdate(context.Context, *model.User) error
	DeleteByID(context.Context, int64) error

	GetUserRoleRels(context.Context) ([]*result.UserRoleRet, error)
	GetRoleRels(context.Context, int64) ([]*model.UserRoleRel, error)
	GetRoleRelByID(context.Context, int64, int64) (*model.UserRoleRel, error)
	CreateRoleRel(context.Context, *model.UserRoleRel) error
	CreateRoleRelInBatches(context.Context, []*model.UserRoleRel) error
	DeleteRoleRelByIDs(context.Context, int64, []int64) error
}
