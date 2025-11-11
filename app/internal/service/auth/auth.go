package auth

import (
	"context"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
)

type AuthService interface {
	Login(context.Context, *request.LoginReq) (*result.LoginResult, error)
	LoginByAuthorize(context.Context, *request.LoginByAuthorizeReq) (*result.LoginResult, error)
	GetUserPermList(context.Context, *request.GetUserPermReq) (*result.RolePermResult, error)
	GetAuthorization(context.Context, *request.GetAuthorizationReq) (*result.AuthorizationResult, error)
	GetToken(context.Context, *request.GetTokenReq) (*result.TokenResult, error)
}
