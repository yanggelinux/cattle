package permission

import (
	"context"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
)

type PermissionService interface {
	GetList(context.Context, *request.GetPermissionReq) (*result.PermissionResult, error)
	Create(context.Context, *request.CreatePermissionReq) error
	Update(context.Context, *request.UpdatePermissionReq) error
	Delete(context.Context, int64) error
	GetRolePermList(context.Context, *request.GetRolePermReq) (*result.PermTreeData, error)
	UpdateRolePerm(context.Context, *request.UpdateRolePermReq) error
}
