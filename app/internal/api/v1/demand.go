package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/app"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	"github.com/yanggelinux/cattle/internal/service/demand"
	"strconv"
)

type Demand struct {
}

func NewDemand() *Demand {
	return &Demand{}
}

func (a *Demand) GetList(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.GetDemandReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := demand.NewDemandService()
	resultData, err := svc.GetList(c, req)
	app.Response(c, resultData, err)
}

func (a *Demand) GetDetail(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.GetDemandDetailReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := demand.NewDemandService()
	resultData, err := svc.GetDetail(c, req)
	app.Response(c, resultData, err)
}

func (a *Demand) Create(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.CreateDemandReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := demand.NewDemandService()
	err = svc.Create(c, req)
	app.Response(c, data, err)
}
func (a *Demand) Update(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.UpdateDemandReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := demand.NewDemandService()
	err = svc.Update(c, req)
	app.Response(c, data, err)
}

func (a *Demand) Delete(c *gin.Context) {
	data := &result.EmptyResult{}
	strID := c.Param("id")
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		app.Response(c, nil, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := demand.NewDemandService()
	err = svc.Delete(c, id)
	app.Response(c, data, err)
}
func (a *Demand) Approve(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.ApproveDemandReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := demand.NewDemandService()
	err = svc.Approve(c, req)
	app.Response(c, data, err)
}

func (a *Demand) Evaluate(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.EvaluateDemandReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := demand.NewDemandService()
	err = svc.Evaluate(c, req)
	app.Response(c, data, err)
}
