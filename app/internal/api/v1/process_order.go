package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/app"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	"github.com/yanggelinux/cattle/internal/service/processorder"
	"strconv"
)

type ProcessOrder struct {
}

func NewProcessOrder() *ProcessOrder {
	return &ProcessOrder{}
}

func (a *ProcessOrder) GetList(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.GetProcessOrderReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := processorder.NewProcessOrderService()
	resultData, err := svc.GetList(c, req)
	app.Response(c, resultData, err)
}

func (a *ProcessOrder) GetApprovalList(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.GetProcessOrderReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := processorder.NewProcessOrderService()
	resultData, err := svc.GetApprovalList(c, req)
	app.Response(c, resultData, err)
}

func (a *ProcessOrder) GetUnapprovedList(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.GetUnapprovedOrderReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := processorder.NewProcessOrderService()
	resultData, err := svc.GetUnapprovedList(c, req)
	app.Response(c, resultData, err)
}

func (a *ProcessOrder) Create(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.CreateProcessOrderReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := processorder.NewProcessOrderService()
	err = svc.Create(c, req)
	app.ResponseWithError(c, data, err)
}

func (a *ProcessOrder) Update(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.UpdateProcessOrderReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := processorder.NewProcessOrderService()
	err = svc.Update(c, req)
	app.ResponseWithError(c, data, err)
}

func (a *ProcessOrder) Delete(c *gin.Context) {
	data := &result.EmptyResult{}
	strID := c.Param("id")
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := processorder.NewProcessOrderService()
	err = svc.Delete(c, id)
	app.Response(c, data, err)
}

func (a *ProcessOrder) GetDetail(c *gin.Context) {
	data := &result.EmptyResult{}
	strID := c.Param("id")
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := processorder.NewProcessOrderService()
	resultData, err := svc.GetByID(c, id)
	app.Response(c, resultData, err)
}

func (a *ProcessOrder) Apply(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.ApplyProcessOrderReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := processorder.NewProcessOrderService()
	err = svc.Apply(c, req)
	app.Response(c, data, err)
}

func (a *ProcessOrder) ReApply(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.ReApplyProcessOrderReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := processorder.NewProcessOrderService()
	err = svc.ReApply(c, req)
	app.Response(c, data, err)
}

func (a *ProcessOrder) Approve(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.ApproveProcessOrderReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := processorder.NewProcessOrderService()
	err = svc.Approve(c, req)
	app.Response(c, data, err)
}

func (a *ProcessOrder) AssignApprover(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.AssignApproverReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := processorder.NewProcessOrderService()
	err = svc.AssignApprover(c, req)
	app.Response(c, data, err)
}
