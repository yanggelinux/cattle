package role

import (
	"context"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
)

type RoleService interface {
	GetList(context.Context, *request.GetRoleReq) (*result.RoleResult, error)
	Create(context.Context, *request.CreateRoleReq) error
	Update(context.Context, *request.UpdateRoleReq) error
	Delete(context.Context, int64) error
}
