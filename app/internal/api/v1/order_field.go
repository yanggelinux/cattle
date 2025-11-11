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

type OrderField struct {
}

func NewOrderField() *OrderField {
	return &OrderField{}
}

func (a *OrderField) GetList(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.GetOrderFieldReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := order.NewOrderFieldService()
	resultData, err := svc.GetList(c, req)
	app.Response(c, resultData, err)
}

func (a *OrderField) Create(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.CreateOrderFieldReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := order.NewOrderFieldService()
	err = svc.Create(c, req)
	app.Response(c, data, err)
}
func (a *OrderField) Update(c *gin.Context) {
	data := &result.EmptyResult{}
	req := &request.UpdateOrderFieldReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		app.Response(c, data, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := order.NewOrderFieldService()
	err = svc.Update(c, req)
	app.ResponseWithError(c, data, err)
}

func (a *OrderField) Delete(c *gin.Context) {
	data := &result.EmptyResult{}
	strID := c.Param("id")
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		app.Response(c, nil, errors.WithCode(ce.ErrorInvaildParams.Code(), "参数验证失败", err))
		return
	}
	svc := order.NewOrderFieldService()
	err = svc.Delete(c, id)
	app.Response(c, data, err)
}
