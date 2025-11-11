package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/app"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	"github.com/yanggelinux/cattle/internal/service/role"
	"strconv"
)

type Role struct {
}

func NewRole() *Role {
	return &Role{}
}

func (a *Role) GetList(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.GetRoleReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := role.NewRoleService()
	resultData, err := svc.GetList(c, req)
	app.Response(c, resultData, err)
}

func (a *Role) Create(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.CreateRoleReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := role.NewRoleService()
	err = svc.Create(c, req)
	app.Response(c, data, err)
}
func (a *Role) Update(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.UpdateRoleReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := role.NewRoleService()
	err = svc.Update(c, req)
	app.Response(c, data, err)
}

func (a *Role) Delete(c *gin.Context) {
	data := &result.EmptyResult{}
	strID := c.Param("id")
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		app.Response(c, nil, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := role.NewRoleService()
	err = svc.Delete(c, id)
	app.Response(c, data, err)
}
