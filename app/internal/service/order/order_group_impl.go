package order

import (
	"context"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	orderrepo "github.com/yanggelinux/cattle/internal/repository/order"
	"github.com/yanggelinux/cattle/internal/store/model"
	"github.com/yanggelinux/cattle/pkg/util"
)

type orderGroupService struct {
	orderGroupRepo orderrepo.OrderGroupRepo
}

func NewOrderGroupService() OrderGroupService {
	return &orderGroupService{orderGroupRepo: orderrepo.NewOrderGroupRepo()}
}

func (s *orderGroupService) GetList(ctx context.Context, req *request.GetOrderGroupReq) (*result.OrderGroupResult, error) {

	filter := &orderrepo.OrderGroupFilter{
		Name:     req.Name,
		Status:   req.Status,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	records, total, err := s.orderGroupRepo.GetList(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	retList := make([]*result.OrderGroupRet, 0, len(records))
	for _, record := range records {
		ret := &result.OrderGroupRet{}
		ret.ID = record.ID
		ret.Name = record.Name
		ret.Sort = record.Sort
		ret.Status = record.Status
		ret.UpdatedTime = util.FormatTimeToString(record.UpdatedTime)
		ret.CreatedTime = util.FormatTimeToString(record.CreatedTime)
		retList = append(retList, ret)
	}
	resultData := &result.OrderGroupResult{}
	resultData.Total = total
	resultData.RetList = retList
	return resultData, nil
}

func (s *orderGroupService) Create(ctx context.Context, req *request.CreateOrderGroupReq) error {
	orderGroup := model.NewOrderGroup()
	orderGroup.Name = *req.Name
	orderGroup.Status = *req.Status
	orderGroup.Sort = *req.Sort
	checked := s.orderGroupRepo.CheckDuplicateEntry(ctx, *req.Name)
	if checked {
		return errors.WithCodeError(ce.ErrorDBDuplicateEntry.Code(), errors.New("Duplicate Entry error"))
	}
	err := s.orderGroupRepo.Create(ctx, orderGroup)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return err
}

func (s *orderGroupService) Update(ctx context.Context, req *request.UpdateOrderGroupReq) error {
	id := *req.ID
	orderGroup, err := s.orderGroupRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	updateField := make(map[string]interface{})
	if req.Name != nil {
		updateField[orderGroup.NameField()] = *req.Name
	}
	if req.Status != nil {
		updateField[orderGroup.StatusField()] = *req.Status
	}
	if req.Sort != nil {
		updateField[orderGroup.SortField()] = *req.Sort
	}
	err = s.orderGroupRepo.Update(ctx, id, updateField)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

func (s *orderGroupService) Delete(ctx context.Context, id int64) error {
	_, err := s.orderGroupRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	err = s.orderGroupRepo.DeleteByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}
