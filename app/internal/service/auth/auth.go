package auth

import (
	"context"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
)

type AuthService interface {
	Login(context.Context, *request.LoginReq) (*result.LoginResult, error)
	GetUserPermList(context.Context, *request.GetUserPermReq) (*result.RolePermResult, error)
	GetToken(context.Context, *request.GetTokenReq) (*result.TokenResult, error)
}
