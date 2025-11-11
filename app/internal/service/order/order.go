package order

import (
	"context"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
)

type OrderService interface {
	GetList(context.Context, *request.GetOrderReq) (*result.OrderResult, error)
	GetNodeTypeList(context.Context, *request.GetOrderReq) (map[string][]*result.OrderNodeRet, error)
	GetByID(context.Context, int64) (*result.OrderRet, error)
	Create(context.Context, *request.CreateOrderReq) error
	Update(context.Context, *request.UpdateOrderReq) error
	Delete(context.Context, int64) error
	Copy(context.Context, int64) error
}
