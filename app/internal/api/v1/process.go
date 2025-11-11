package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/app"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	"github.com/yanggelinux/cattle/internal/service/process"
	"strconv"
)

type Process struct {
}

func NewProcess() *Process {
	return &Process{}
}

func (a *Process) GetList(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.GetProcessReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := process.NewProcessService()
	resultData, err := svc.GetList(c, req)
	app.Response(c, resultData, err)
}

func (a *Process) GetDetail(c *gin.Context) {
	data := &result.EmptyResult{}
	strID := c.Param("id")
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := process.NewProcessService()
	resultData, err := svc.GetDetail(c, id)
	app.Response(c, resultData, err)
}

func (a *Process) Create(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.CreateProcessReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := process.NewProcessService()
	resultData, err := svc.Create(c, req)
	app.Response(c, resultData, err)
}
func (a *Process) Update(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.UpdateProcessReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := process.NewProcessService()
	err = svc.Update(c, req)
	app.ResponseWithError(c, data, err)
}

func (a *Process) Delete(c *gin.Context) {
	data := &result.EmptyResult{}
	strID := c.Param("id")
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		app.Response(c, nil, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := process.NewProcessService()
	err = svc.Delete(c, id)
	app.Response(c, data, err)
}

func (a *Process) Copy(c *gin.Context) {
	data := &result.EmptyResult{}
	strID := c.Param("id")
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		app.Response(c, nil, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := process.NewProcessService()
	err = svc.Copy(c, id)
	app.Response(c, data, err)
}
