package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/app"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	"github.com/yanggelinux/cattle/internal/service/auth"
)

type Auth struct {
}

func NewAuth() *Auth {
	return &Auth{}
}

func (a *Auth) Login(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.LoginReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := auth.NewAuthService()
	authResult, err := svc.Login(c, req)
	app.Response(c, authResult, err)
}

func (a *Auth) GetUserPermList(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.GetUserPermReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := auth.NewAuthService()
	authResult, err := svc.GetUserPermList(c, req)
	app.Response(c, authResult, err)
}
