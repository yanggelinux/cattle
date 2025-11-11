package dashboard

import (
	"context"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	graphrepo "github.com/yanggelinux/cattle/internal/repository/archgraph"
	demandrepo "github.com/yanggelinux/cattle/internal/repository/demand"
	orderrepo "github.com/yanggelinux/cattle/internal/repository/processorder"
	"github.com/yanggelinux/cattle/pkg/util"
	"time"
)

type dashboardService struct {
	archGraphRepo    graphrepo.ArchGraphRepo
	processOrderRepo orderrepo.ProcessOrderRepo
	demandRepo       demandrepo.DemandRepo
}

func NewDashboardService() DashboardService {
	return &dashboardService{
		archGraphRepo:    graphrepo.NewArchGraphRepo(),
		processOrderRepo: orderrepo.NewProcessOrderRepo(),
		demandRepo:       demandrepo.NewDemandRepo(),
	}
}

func (s *dashboardService) GetGraphInfo(ctx context.Context) (*result.DashboardGraphResult, error) {
	resultData := &result.DashboardGraphResult{}
	filter := &graphrepo.ArchGraphFilter{}
	var status int8
	// unapproved
	status = int8(0)
	filter.Status = &status
	unapprovedCount, err := s.archGraphRepo.GetCount(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return resultData, err
	}
	// approving
	status = int8(1)
	filter.Status = &status
	approvingCount, err := s.archGraphRepo.GetCount(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return resultData, err
	}
	// success
	status = int8(2)
	filter.Status = &status
	successCount, err := s.archGraphRepo.GetCount(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return resultData, err
	}
	// failed
	status = int8(3)
	filter.Status = &status
	failedCount, err := s.archGraphRepo.GetCount(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return resultData, err
	}
	resultData.TotalCount = unapprovedCount + approvingCount + successCount + failedCount
	resultData.UnapprovedCount = unapprovedCount
	resultData.ApprovingCount = approvingCount
	resultData.SuccessCount = successCount
	resultData.FailedCount = failedCount
	return resultData, nil
}

func (s *dashboardService) GetOrderInfo(ctx context.Context) (*result.DashboardOrderResult, error) {

	resultData := &result.DashboardOrderResult{}
	filter := &orderrepo.ProcessOrderFilter{}
	var status int8
	// total
	totalCount, err := s.processOrderRepo.GetCount(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return resultData, err
	}
	// unapproved
	status = int8(0)
	filter.Status = &status
	unapprovedCount, err := s.processOrderRepo.GetCount(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return resultData, err
	}
	// approving
	status = int8(1)
	filter.Status = &status
	approvingCount, err := s.processOrderRepo.GetCount(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return resultData, err
	}
	// success
	status = int8(2)
	filter.Status = &status
	successCount, err := s.processOrderRepo.GetCount(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return resultData, err
	}
	// failed
	status = int8(3)
	filter.Status = &status
	failedCount, err := s.processOrderRepo.GetCount(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return resultData, err
	}
	var orderType int8
	orderType = 1
	graphApplyDist, err := s.getOrderDist(ctx, orderType)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return resultData, err
	}
	orderType = 2
	graphChangeDist, err := s.getOrderDist(ctx, orderType)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return resultData, err
	}
	orderType = 3
	resApplyChangeDist, err := s.getOrderDist(ctx, orderType)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return resultData, err
	}
	resultData.TotalCount = totalCount
	resultData.UnapprovedCount = unapprovedCount
	resultData.ApprovingCount = approvingCount
	resultData.SuccessCount = successCount
	resultData.FailedCount = failedCount
	resultData.GraphApplyDist = graphApplyDist
	resultData.GraphChangeDist = graphChangeDist
	resultData.ResApplyChangeDist = resApplyChangeDist
	return resultData, nil
}

func (s *dashboardService) Last30DaysOrderDist() []*result.OrderDistRet {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	var list []*result.OrderDistRet
	for i := 29; i >= 0; i-- {
		day := today.AddDate(0, 0, -i)
		list = append(list, &result.OrderDistRet{
			Count: 0,
			DT:    day.Format("2006-01-02"),
		})
	}
	return list
}

func (s *dashboardService) getOrderDist(ctx context.Context, orderType int8) ([]*result.OrderDistRet, error) {
	dists := make([]*result.OrderDistRet, 0, 0)
	recourds, err := s.processOrderRepo.GetCountByTime(ctx, orderType)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return dists, err
	}
	last30DaysOrderDist := s.Last30DaysOrderDist()
	for _, v := range last30DaysOrderDist {
		for _, record := range recourds {
			count := record.Count
			dt := util.FormatDateToString(record.DTT)
			if dt == v.DT {
				v.Count = count
			}
		}
		dists = append(dists, v)
	}
	return dists, nil
}

func (s *dashboardService) GetDemandInfo(ctx context.Context) (*result.DashboardDemandResult, error) {
	resultData := &result.DashboardDemandResult{}
	filter := &demandrepo.DemandFilter{}
	var status int8
	// total
	totalCount, err := s.demandRepo.GetCount(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return resultData, err
	}
	// unapproved
	status = int8(0)
	filter.Status = &status
	unapprovedCount, err := s.demandRepo.GetCount(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return resultData, err
	}
	// approving
	status = int8(1)
	filter.Status = &status
	approvingCount, err := s.demandRepo.GetCount(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return resultData, err
	}
	// success
	status = int8(2)
	filter.Status = &status
	successCount, err := s.demandRepo.GetCount(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return resultData, err
	}
	// failed
	status = int8(3)
	filter.Status = &status
	failedCount, err := s.demandRepo.GetCount(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return resultData, err
	}
	resultData.TotalCount = totalCount
	resultData.UnapprovedCount = unapprovedCount
	resultData.ApprovingCount = approvingCount
	resultData.SuccessCount = successCount
	resultData.FailedCount = failedCount
	return resultData, nil
}
