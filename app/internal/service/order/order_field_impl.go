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
	"strings"
)

type orderFieldService struct {
	orderRepo      orderrepo.OrderRepo
	orderFieldRepo orderrepo.OrderFieldRepo
}

func NewOrderFieldService() OrderFieldService {
	return &orderFieldService{
		orderRepo:      orderrepo.NewOrderRepo(),
		orderFieldRepo: orderrepo.NewOrderFieldRepo(),
	}
}

func (s *orderFieldService) GetList(ctx context.Context, req *request.GetOrderFieldReq) (*result.OrderFieldResult, error) {

	filter := &orderrepo.OrderFieldFilter{
		Name:     req.Name,
		OrderID:  req.OrderID,
		Status:   req.Status,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	records, total, err := s.orderFieldRepo.GetList(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	orderIDNames, _ := s.getOrderIDName(ctx)
	retList := make([]*result.OrderFieldRet, 0, len(records))
	for _, record := range records {
		ret := &result.OrderFieldRet{}
		orderName := orderIDNames[record.OrderID]
		ret.ID = record.ID
		ret.OrderID = record.OrderID
		ret.OrderName = orderName
		ret.Name = record.Name
		ret.Key = record.Key
		ret.Component = record.Component
		ret.Placeholder = record.Placeholder
		ret.Description = record.Description
		ret.VerRule = record.VerRule
		ret.DefaultVal = record.DefaultVal
		ret.IsRequired = record.IsRequired
		ret.IsTitle = record.IsTitle
		ret.IsEdit = record.IsEdit
		ret.IsClear = record.IsClear
		ret.DisplayField = record.DisplayField
		ret.DisplayVal = record.DisplayVal
		ret.Enum = record.Enum
		ret.GroupName = record.GroupName
		ret.Sort = record.Sort
		ret.Status = record.Status
		ret.UpdatedTime = util.FormatTimeToString(record.UpdatedTime)
		ret.CreatedTime = util.FormatTimeToString(record.CreatedTime)
		retList = append(retList, ret)
	}
	resultData := &result.OrderFieldResult{}
	resultData.Total = total
	resultData.RetList = retList
	return resultData, nil
}

func (s *orderFieldService) getOrderIDName(ctx context.Context) (map[int64]string, error) {
	orderIDNames := make(map[int64]string)
	records, err := s.orderRepo.GetAll(ctx)
	if err != nil {
		return orderIDNames, err
	}
	for _, record := range records {
		orderIDNames[record.ID] = record.Name
	}
	return orderIDNames, nil
}

func (s *orderFieldService) Create(ctx context.Context, req *request.CreateOrderFieldReq) error {
	orderField := model.NewOrderField()
	orderField.Name = strings.Trim(*req.Name, " ")
	orderField.OrderID = *req.OrderID
	orderField.Key = strings.Trim(*req.Key, " ")
	orderField.Component = *req.Component
	orderField.VerRule = *req.VerRule
	orderField.Status = *req.Status
	orderField.Sort = *req.Sort
	if req.Placeholder != nil {
		orderField.Placeholder = *req.Placeholder
	}
	if req.DefaultVal != nil {
		orderField.DefaultVal = *req.DefaultVal
	}
	if req.IsRequired != nil {
		orderField.IsRequired = *req.IsRequired
	}
	if req.IsTitle != nil {
		orderField.IsTitle = *req.IsTitle
	}
	if req.IsEdit != nil {
		orderField.IsEdit = *req.IsEdit
	}
	if req.IsClear != nil {
		orderField.IsClear = *req.IsClear
	}
	if req.DisplayField != nil {
		orderField.DisplayField = *req.DisplayField
	}
	if req.DisplayVal != nil {
		orderField.DisplayVal = *req.DisplayVal
	}
	if req.Description != nil {
		orderField.Description = *req.Description
	}
	if req.Enum != nil {
		orderField.Enum = strings.Trim(*req.Enum, "\n")
	}
	if req.GroupName != nil {
		orderField.GroupName = *req.GroupName
	}
	checked := s.orderFieldRepo.CheckDuplicateEntry(ctx, *req.Key, *req.OrderID)
	if checked {
		return errors.WithCodeError(ce.ErrorDBDuplicateEntry.Code(), errors.New("Duplicate Entry error"))
	}
	err := s.orderFieldRepo.Create(ctx, orderField)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return err
}

func (s *orderFieldService) Update(ctx context.Context, req *request.UpdateOrderFieldReq) error {
	id := *req.ID
	orderField, err := s.orderFieldRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	updateField := make(map[string]interface{})
	if req.OrderID != nil {
		updateField[orderField.OrderIDField()] = *req.OrderID
	}
	if req.Name != nil {
		updateField[orderField.NameField()] = strings.Trim(*req.Name, " ")
	}
	if req.Key != nil {
		updateField[orderField.KeyField()] = strings.Trim(*req.Key, " ")
	}
	if req.Component != nil {
		updateField[orderField.ComponentField()] = *req.Component
	}
	if req.Placeholder != nil {
		updateField[orderField.PlaceholderField()] = *req.Placeholder
	}
	if req.VerRule != nil {
		updateField[orderField.VerRuleField()] = *req.VerRule
	}
	if req.DefaultVal != nil {
		updateField[orderField.DefaultValField()] = *req.DefaultVal
	}
	if req.IsRequired != nil {
		updateField[orderField.IsRequiredField()] = *req.IsRequired
	}
	if req.IsTitle != nil {
		updateField[orderField.IsTitleField()] = *req.IsTitle
	}
	if req.IsEdit != nil {
		updateField[orderField.IsEditField()] = *req.IsEdit
	}
	if req.IsClear != nil {
		updateField[orderField.IsClearField()] = *req.IsClear
	}
	if req.DisplayField != nil {
		updateField[orderField.DisplayFieldField()] = *req.DisplayField
	}
	if req.DisplayVal != nil {
		updateField[orderField.DisplayValField()] = *req.DisplayVal
	}
	if req.Description != nil {
		updateField[orderField.DescriptionField()] = *req.Description
	}
	if req.Enum != nil {
		updateField[orderField.EnumField()] = strings.Trim(*req.Enum, "\n")

	}
	if req.GroupName != nil {
		updateField[orderField.GroupNameField()] = *req.GroupName
	}
	if req.Status != nil {
		updateField[orderField.StatusField()] = *req.Status
	}
	if req.Sort != nil {
		updateField[orderField.SortField()] = *req.Sort
	}
	err = s.orderFieldRepo.Update(ctx, id, updateField)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

func (s *orderFieldService) Delete(ctx context.Context, id int64) error {
	_, err := s.orderFieldRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	err = s.orderFieldRepo.DeleteByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}
