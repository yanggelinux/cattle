package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/app"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	"github.com/yanggelinux/cattle/internal/service/order"
	"strconv"
)

type Order struct {
}

func NewOrder() *Order {
	return &Order{}
}

func (a *Order) GetList(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.GetOrderReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := order.NewOrderService()
	resultData, err := svc.GetList(c, req)
	app.Response(c, resultData, err)
}

func (a *Order) GetDetail(c *gin.Context) {
	strID := c.Param("id")
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		app.Response(c, nil, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := order.NewOrderService()
	resultData, err := svc.GetByID(c, id)
	app.Response(c, resultData, err)
}

func (a *Order) GetNodeTypeList(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.GetOrderReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := order.NewOrderService()
	resultData, err := svc.GetNodeTypeList(c, req)
	app.Response(c, resultData, err)
}

func (a *Order) Create(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.CreateOrderReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := order.NewOrderService()
	err = svc.Create(c, req)
	app.Response(c, data, err)
}
func (a *Order) Update(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.UpdateOrderReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := order.NewOrderService()
	err = svc.Update(c, req)
	app.ResponseWithError(c, data, err)
}

func (a *Order) Delete(c *gin.Context) {
	data := &result.EmptyResult{}
	strID := c.Param("id")
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		app.Response(c, nil, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := order.NewOrderService()
	err = svc.Delete(c, id)
	app.Response(c, data, err)
}

func (a *Order) Copy(c *gin.Context) {
	data := &result.EmptyResult{}
	strID := c.Param("id")
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		app.Response(c, nil, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := order.NewOrderService()
	err = svc.Copy(c, id)
	app.Response(c, data, err)
}
