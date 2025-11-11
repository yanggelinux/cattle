package dashboard

import (
	"context"
	"github.com/yanggelinux/cattle/internal/dto/result"
)

type DashboardService interface {
	GetGraphInfo(context.Context) (*result.DashboardGraphResult, error)
	GetOrderInfo(context.Context) (*result.DashboardOrderResult, error)
	GetDemandInfo(ctx context.Context) (*result.DashboardDemandResult, error)
}
