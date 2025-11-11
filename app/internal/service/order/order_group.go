package order

import (
	"context"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
)

type OrderGroupService interface {
	GetList(context.Context, *request.GetOrderGroupReq) (*result.OrderGroupResult, error)
	Create(context.Context, *request.CreateOrderGroupReq) error
	Update(context.Context, *request.UpdateOrderGroupReq) error
	Delete(context.Context, int64) error
}
