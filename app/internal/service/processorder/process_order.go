package processorder

import (
	"context"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
)

type ProcessOrderService interface {
	GetList(context.Context, *request.GetProcessOrderReq) (*result.ProcessOrderResult, error)
	GetUnapprovedList(context.Context, *request.GetUnapprovedOrderReq) (*result.ProcessOrderUnapprovedResult, error)
	GetApprovalList(context.Context, *request.GetProcessOrderReq) (*result.ProcessOrderResult, error)
	Create(context.Context, *request.CreateProcessOrderReq) error
	Update(context.Context, *request.UpdateProcessOrderReq) error
	Delete(context.Context, int64) error
	GetByID(context.Context, int64) (*result.ProcessOrderDetailResult, error)
	Apply(context.Context, *request.ApplyProcessOrderReq) error
	ReApply(context.Context, *request.ReApplyProcessOrderReq) error
	Approve(context.Context, *request.ApproveProcessOrderReq) error
	AssignApprover(context.Context, *request.AssignApproverReq) error
}
