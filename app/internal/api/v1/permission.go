package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/app"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	"github.com/yanggelinux/cattle/internal/service/permission"
	"strconv"
)

type Permission struct {
}

func NewPermission() *Permission {
	return &Permission{}
}

func (a *Permission) GetList(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.GetPermissionReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := permission.NewPermissionService()
	resultData, err := svc.GetList(c, req)
	app.Response(c, resultData, err)
}

func (a *Permission) Create(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.CreatePermissionReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := permission.NewPermissionService()
	err = svc.Create(c, req)
	app.Response(c, data, err)
}
func (a *Permission) Update(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.UpdatePermissionReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := permission.NewPermissionService()
	err = svc.Update(c, req)
	app.Response(c, data, err)
}

func (a *Permission) Delete(c *gin.Context) {
	data := &result.EmptyResult{}
	strID := c.Param("id")
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := permission.NewPermissionService()
	err = svc.Delete(c, id)
	app.Response(c, data, err)
}

func (a *Permission) GetRolePermList(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.GetRolePermReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := permission.NewPermissionService()
	permTreeData, err := svc.GetRolePermList(c, req)
	app.Response(c, permTreeData, err)
}

func (a *Permission) UpdateRolePerm(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.UpdateRolePermReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := permission.NewPermissionService()
	err = svc.UpdateRolePerm(c, req)
	app.Response(c, data, err)
}
