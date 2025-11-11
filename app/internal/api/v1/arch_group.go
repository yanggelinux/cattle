package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/app"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	"github.com/yanggelinux/cattle/internal/service/archgroup"
	"strconv"
)

type ArchGroup struct {
}

func NewArchGroup() *ArchGroup {
	return &ArchGroup{}
}

func (a *ArchGroup) GetList(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.GetArchGroupReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := archgroup.NewArchGroupService()
	resultData, err := svc.GetList(c, req)
	app.Response(c, resultData, err)
}

func (a *ArchGroup) Create(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.CreateArchGroupReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := archgroup.NewArchGroupService()
	err = svc.Create(c, req)
	app.Response(c, data, err)
}
func (a *ArchGroup) Update(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.UpdateArchGroupReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := archgroup.NewArchGroupService()
	err = svc.Update(c, req)
	app.Response(c, data, err)
}

func (a *ArchGroup) Delete(c *gin.Context) {
	data := &result.EmptyResult{}
	strID := c.Param("id")
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := archgroup.NewArchGroupService()
	err = svc.Delete(c, id)
	app.Response(c, data, err)
}
