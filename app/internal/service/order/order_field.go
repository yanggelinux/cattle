package order

import (
	"context"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
)

type OrderFieldService interface {
	GetList(context.Context, *request.GetOrderFieldReq) (*result.OrderFieldResult, error)
	Create(context.Context, *request.CreateOrderFieldReq) error
	Update(context.Context, *request.UpdateOrderFieldReq) error
	Delete(context.Context, int64) error
}
