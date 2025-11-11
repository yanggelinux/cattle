package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/app"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	"github.com/yanggelinux/cattle/internal/service/user"
	"strconv"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (a *User) GetList(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.GetUserReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := user.NewUserService()
	resultData, err := svc.GetList(c, req)
	app.Response(c, resultData, err)
}

func (a *User) Create(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.CreateUserReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := user.NewUserService()
	err = svc.Create(c, req)
	app.Response(c, data, err)
}
func (a *User) Update(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.UpdateUserReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := user.NewUserService()
	err = svc.Update(c, req)
	app.Response(c, data, err)
}

func (a *User) Delete(c *gin.Context) {
	data := &result.EmptyResult{}
	strID := c.Param("id")
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := user.NewUserService()
	err = svc.Delete(c, id)
	app.Response(c, data, err)
}
