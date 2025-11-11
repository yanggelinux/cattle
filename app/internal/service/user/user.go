package user

import (
	"context"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
)

type UserService interface {
	GetList(context.Context, *request.GetUserReq) (*result.UserResult, error)
	Create(context.Context, *request.CreateUserReq) error
	Update(context.Context, *request.UpdateUserReq) error
	Delete(context.Context, int64) error
	CreateUserRole(context.Context, int64, string) error
}
