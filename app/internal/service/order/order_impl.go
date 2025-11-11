package order

import (
	"context"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	orderrepo "github.com/yanggelinux/cattle/internal/repository/order"
	processrepo "github.com/yanggelinux/cattle/internal/repository/process"
	"github.com/yanggelinux/cattle/internal/store/model"
	"github.com/yanggelinux/cattle/pkg/util"
	"strings"
)

type orderService struct {
	db             *model.MDB
	orderRepo      orderrepo.OrderRepo
	orderGroupRepo orderrepo.OrderGroupRepo
	orderFieldRepo orderrepo.OrderFieldRepo
	processRepo    processrepo.ProcessRepo
}

func NewOrderService() OrderService {
	return &orderService{
		db:             model.GetDB(),
		orderRepo:      orderrepo.NewOrderRepo(),
		orderGroupRepo: orderrepo.NewOrderGroupRepo(),
		orderFieldRepo: orderrepo.NewOrderFieldRepo(),
		processRepo:    processrepo.NewProcessRepo(),
	}
}

func (s *orderService) GetList(ctx context.Context, req *request.GetOrderReq) (*result.OrderResult, error) {

	filter := &orderrepo.OrderFilter{
		Name:     req.Name,
		Status:   req.Status,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	records, total, err := s.orderRepo.GetList(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	retList := make([]*result.OrderRet, 0, len(records))
	orderGrupIDNames, err := s.genOrderGroupIDNames(ctx)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	processIDNames, err := s.genProcessIDNames(ctx)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	for _, record := range records {
		ret := &result.OrderRet{}
		var isTask int8
		groupName := orderGrupIDNames[record.GroupID]
		processName := processIDNames[record.ProcessID]
		if (len(record.TaskUrl) > 0) && (len(record.TaskMethod) > 0) {
			isTask = 1
		}
		var nodeType []string
		if len(record.NodeType) > 0 {
			nodeType = strings.Split(record.NodeType, ",")
		}
		ret.ID = record.ID
		ret.Name = record.Name
		ret.GroupID = record.GroupID
		ret.ProcessID = record.ProcessID
		ret.ProcessName = processName
		ret.GroupName = groupName
		ret.OrderType = record.OrderType
		ret.NodeType = nodeType
		ret.Label = record.Label
		ret.Layout = record.Layout
		ret.IsTask = isTask
		ret.TaskUrl = record.TaskUrl
		ret.TaskMethod = record.TaskMethod
		ret.Sort = record.Sort
		ret.Status = record.Status
		ret.UpdatedTime = util.FormatTimeToString(record.UpdatedTime)
		ret.CreatedTime = util.FormatTimeToString(record.CreatedTime)
		retList = append(retList, ret)
	}
	resultData := &result.OrderResult{}
	resultData.Total = total
	resultData.RetList = retList
	return resultData, nil
}

func (s *orderService) GetByID(ctx context.Context, id int64) (*result.OrderRet, error) {

	record, err := s.orderRepo.GetByID(ctx, id)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	resultData := &result.OrderRet{}
	var isTask int8
	if (len(record.TaskUrl) > 0) && (len(record.TaskMethod) > 0) {
		isTask = 1
	}
	var nodeType []string
	if len(record.NodeType) > 0 {
		nodeType = strings.Split(record.NodeType, ",")
	}
	resultData.ID = record.ID
	resultData.Name = record.Name
	resultData.GroupID = record.GroupID
	resultData.ProcessID = record.ProcessID
	resultData.OrderType = record.OrderType
	resultData.NodeType = nodeType
	resultData.Label = record.Label
	resultData.Layout = record.Layout
	resultData.IsTask = isTask
	resultData.TaskUrl = record.TaskUrl
	resultData.TaskMethod = record.TaskMethod
	resultData.Sort = record.Sort
	resultData.Status = record.Status
	resultData.UpdatedTime = util.FormatTimeToString(record.UpdatedTime)
	resultData.CreatedTime = util.FormatTimeToString(record.CreatedTime)
	return resultData, nil
}

func (s *orderService) genOrderGroupIDNames(ctx context.Context) (map[int64]string, error) {
	idNames := make(map[int64]string)
	records, err := s.orderGroupRepo.GetAll(ctx)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	for _, record := range records {
		idNames[record.ID] = record.Name
	}
	return idNames, nil
}

func (s *orderService) genProcessIDNames(ctx context.Context) (map[int64]string, error) {
	idNames := make(map[int64]string)
	records, err := s.processRepo.GetAll(ctx)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	for _, record := range records {
		idNames[record.ID] = record.Name
	}
	return idNames, nil
}

func (s *orderService) GetNodeTypeList(ctx context.Context, req *request.GetOrderReq) (map[string][]*result.OrderNodeRet, error) {
	filter := &orderrepo.OrderFilter{
		OrderType: req.OrderType,
	}
	records, _, err := s.orderRepo.GetList(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	nodeTypeResult := make(map[string][]*result.OrderNodeRet)
	for _, record := range records {
		nodeType := record.NodeType
		if len(nodeType) == 0 {
			continue
		}
		nodeTypeList := strings.Split(nodeType, ",")
		orderNode := &result.OrderNodeRet{}
		orderNode.Label = record.Name
		orderNode.Value = record.ID
		for _, nType := range nodeTypeList {
			if orderNodeRets, ok := nodeTypeResult[nType]; !ok {
				nodeTypeResult[nType] = []*result.OrderNodeRet{orderNode}
			} else {
				nodeTypeResult[nType] = append(orderNodeRets, orderNode)
			}
		}
	}
	return nodeTypeResult, nil
}

func (s *orderService) Create(ctx context.Context, req *request.CreateOrderReq) error {
	order := model.NewOrder()
	order.Name = *req.Name
	order.ProcessID = *req.ProcessID
	order.GroupID = *req.GroupID
	order.OrderType = *req.OrderType
	order.Sort = *req.Sort
	order.Status = *req.Status
	if len(req.NodeType) != 0 {
		order.NodeType = strings.Join(req.NodeType, ",")
	}
	if req.Label != nil {
		order.Label = *req.Label
		//检查标签是否重复
		//checked := s.orderRepo.CheckDuplicateLabel(ctx, *req.Label)
		//if checked {
		//	return errors.WithCodeError(ce.ErrorOrderDuplicateLabel.Code(), errors.New("duplicate order label error"))
		//}
	}
	if req.Layout != nil {
		order.Layout = *req.Layout
	}
	if req.TaskUrl != nil {
		order.TaskUrl = *req.TaskUrl
	}
	if req.TaskMethod != nil {
		order.TaskMethod = *req.TaskMethod
	}
	checked := s.orderRepo.CheckDuplicateEntry(ctx, *req.Name)
	if checked {
		return errors.WithCodeError(ce.ErrorDBDuplicateEntry.Code(), errors.New("Duplicate Entry error"))
	}
	err := s.orderRepo.Create(ctx, order)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return err
}

func (s *orderService) Update(ctx context.Context, req *request.UpdateOrderReq) error {
	id := *req.ID
	order, err := s.orderRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	updateField := make(map[string]interface{})

	if req.Name != nil {
		updateField[order.NameField()] = *req.Name
	}
	if req.GroupID != nil {
		updateField[order.GroupIDField()] = *req.GroupID
	}
	if req.ProcessID != nil {
		updateField[order.ProcessIDField()] = *req.ProcessID
	}
	if req.OrderType != nil {
		updateField[order.OrderTypeField()] = *req.OrderType
	}
	if len(req.NodeType) != 0 {
		updateField[order.NodeTypeField()] = strings.Join(req.NodeType, ",")
	}
	if req.Label != nil {
		//checkOrder, err := s.orderRepo.GetByLabel(ctx, *req.Label)
		//if err != nil && !errors.Is(err, model.ErrRecordNotFound) {
		//	return errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		//}
		//if err == nil {
		//	if checkOrder.ID != id && len(*req.Label) != 0 {
		//		return errors.WithCodeError(ce.ErrorOrderDuplicateLabel.Code(), errors.New("duplicate order label error"))
		//	}
		//}

		updateField[order.LabelField()] = *req.Label
	}
	if req.Layout != nil {
		updateField[order.LayoutField()] = *req.Layout
	}
	if req.TaskUrl != nil {
		updateField[order.TaskUrlField()] = *req.TaskUrl
	}
	if req.TaskMethod != nil {
		updateField[order.TaskMethodField()] = *req.TaskMethod
	}
	if req.Status != nil {
		updateField[order.StatusField()] = *req.Status
	}
	if req.Sort != nil {
		updateField[order.SortField()] = *req.Sort
	}
	err = s.orderRepo.Update(ctx, id, updateField)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

func (s *orderService) Delete(ctx context.Context, id int64) error {
	_, err := s.orderRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	err = s.orderRepo.DeleteByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

func (s *orderService) Copy(ctx context.Context, id int64) error {
	order, err := s.orderRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	newOrder := model.NewOrder()
	newOrder.Name = order.Name + "@copy" + util.RandString(3)
	newOrder.GroupID = order.GroupID
	newOrder.OrderType = order.OrderType
	newOrder.ProcessID = order.ProcessID
	newOrder.NodeType = order.NodeType
	newOrder.Label = ""
	newOrder.Layout = order.Layout
	newOrder.TaskUrl = order.TaskUrl
	newOrder.TaskMethod = order.TaskMethod
	newOrder.Sort = order.Sort
	newOrder.Status = order.Status

	tx := s.db.WithContext(ctx).Begin()
	defer model.RecoverRollback(tx)

	orderTx := s.orderRepo.WithTx(s.db)
	orderFieldTx := s.orderFieldRepo.WithTx(tx)
	err = orderTx.Create(ctx, newOrder)
	if err != nil {
		tx.Rollback()
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	newOrderID := newOrder.ID
	orderFields, err := orderFieldTx.GetByOrder(ctx, id)
	if err != nil {
		tx.Rollback()
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	orderFieldList := make([]*model.OrderField, 0, len(orderFields))
	for _, orderField := range orderFields {
		newOrderField := model.NewOrderField()
		newOrderField.OrderID = newOrderID
		newOrderField.Name = orderField.Name
		newOrderField.Key = orderField.Key
		newOrderField.Component = orderField.Component
		newOrderField.VerRule = orderField.VerRule
		newOrderField.Placeholder = orderField.Placeholder
		newOrderField.DefaultVal = orderField.DefaultVal
		newOrderField.IsRequired = orderField.IsRequired
		newOrderField.IsTitle = orderField.IsTitle
		newOrderField.IsEdit = orderField.IsEdit
		newOrderField.IsClear = orderField.IsClear
		newOrderField.DisplayField = orderField.DisplayField
		newOrderField.DisplayVal = orderField.DisplayVal
		newOrderField.Description = orderField.Description
		newOrderField.Enum = orderField.Enum
		newOrderField.GroupName = orderField.GroupName
		newOrderField.Sort = orderField.Sort
		newOrderField.Status = orderField.Status
		orderFieldList = append(orderFieldList, newOrderField)
	}
	err = orderFieldTx.CreateInBatches(ctx, orderFieldList)
	if err != nil {
		tx.Rollback()
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	tx.Commit()
	return nil
}
