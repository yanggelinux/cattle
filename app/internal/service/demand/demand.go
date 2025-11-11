package demand

import (
	"context"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
)

type DemandService interface {
	GetList(context.Context, *request.GetDemandReq) (*result.DemandResult, error)
	GetDetail(context.Context, *request.GetDemandDetailReq) (*result.DemandRet, error)
	Create(context.Context, *request.CreateDemandReq) error
	Update(context.Context, *request.UpdateDemandReq) error
	Delete(context.Context, int64) error
	Approve(context.Context, *request.ApproveDemandReq) error
	Evaluate(context.Context, *request.EvaluateDemandReq) error
	CheckEvaluation(context.Context) error
}
