package processorder

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/dto/biz"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	graphrepo "github.com/yanggelinux/cattle/internal/repository/archgraph"
	orderrepo "github.com/yanggelinux/cattle/internal/repository/order"
	processrepo "github.com/yanggelinux/cattle/internal/repository/process"
	porderrepo "github.com/yanggelinux/cattle/internal/repository/processorder"
	rolerepo "github.com/yanggelinux/cattle/internal/repository/role"
	teamrepo "github.com/yanggelinux/cattle/internal/repository/team"
	userrepo "github.com/yanggelinux/cattle/internal/repository/user"
	"github.com/yanggelinux/cattle/internal/store/model"
	"github.com/yanggelinux/cattle/pkg/log"
	"github.com/yanggelinux/cattle/pkg/util"
	"go.uber.org/zap"
	"gorm.io/datatypes"
	"strings"
)

type processOrderService struct {
	db                  *model.MDB
	processOrderRepo    porderrepo.ProcessOrderRepo
	processApprovalRepo porderrepo.ProcessApprovalRepo
	processArchRepo     porderrepo.ProcessArchRepo
	processRepo         processrepo.ProcessRepo
	archGraphRepo       graphrepo.ArchGraphRepo
	archGraphRecordRepo graphrepo.ArchGraphRecordRepo
	orderRepo           orderrepo.OrderRepo
	orderFieldRepo      orderrepo.OrderFieldRepo
	userRepo            userrepo.UserRepo
	roleRepo            rolerepo.RoleRepo
	teamRepo            teamrepo.TeamRepo
}

func NewProcessOrderService() ProcessOrderService {
	return &processOrderService{
		db:                  model.GetDB(),
		processOrderRepo:    porderrepo.NewProcessOrderRepo(),
		processApprovalRepo: porderrepo.NewProcessApprovalRepo(),
		processArchRepo:     porderrepo.NewProcessArchRepo(),
		processRepo:         processrepo.NewProcessRepo(),
		archGraphRepo:       graphrepo.NewArchGraphRepo(),
		archGraphRecordRepo: graphrepo.NewArchGraphRecordRepo(),
		orderRepo:           orderrepo.NewOrderRepo(),
		orderFieldRepo:      orderrepo.NewOrderFieldRepo(),
		userRepo:            userrepo.NewUserRepo(),
		roleRepo:            rolerepo.NewRoleRepo(),
		teamRepo:            teamrepo.NewTeamRepo(),
	}
}

func (s *processOrderService) genActiveIndex(orderProcess []*model.PorcessNode) int8 {
	i := 0
	for _, v := range orderProcess {
		if v.Status == 0 {
			break
		}
		i += 1
	}
	return int8(i)
}

func (s *processOrderService) GetList(ctx context.Context, req *request.GetProcessOrderReq) (*result.ProcessOrderResult, error) {
	var err error
	filter := &porderrepo.ProcessOrderFilter{
		Title:      req.Title,
		GraphName:  req.GraphName,
		OrderType:  req.OrderType,
		DemandName: req.DemandName,
		Owner:      req.Owner,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		Status:     req.Status,
		Page:       req.Page,
		PageSize:   req.PageSize,
	}
	// 对我的待办特殊处理
	if req.Label != nil && *req.Label == "todo" {
		// 13 status 为 1 或 3
		status := int8(13)
		filter.Page = nil
		filter.PageSize = nil
		filter.Status = &status
	}
	records, total, err := s.processOrderRepo.GetList(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}

	roleNames, err := s.getRolesByUser(ctx)
	if err != nil {
		return nil, errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
	}
	teamNameInfo, _ := s.genTeamInfo(ctx)

	retList := make([]*result.ProcessOrderRet, 0, len(records))
	userName := util.GetUserName(ctx)
	for _, record := range records {
		var orderProcess []*model.PorcessNode
		err := json.Unmarshal(record.OrderProcess, &orderProcess)
		if err != nil {
			continue
		}
		activeIndex := s.genActiveIndex(orderProcess)
		curNodeIndex := activeIndex
		if activeIndex == int8(len(orderProcess)) {
			curNodeIndex = activeIndex - 1
		}
		curOrderNode := orderProcess[curNodeIndex]
		// 获取我的待办
		hasApproval := s.isApprovalOrder(userName, roleNames, curOrderNode, teamNameInfo)
		if req.Label != nil && *req.Label == "todo" {
			// 有审批权限 或者 是自己工单且状态是审批不通过 的非
			if !(hasApproval == int8(1) || (record.Status == 3 && record.Owner == userName && record.OrderType > 2)) {
				continue
			}
		}
		ret := &result.ProcessOrderRet{}
		ret.ID = record.ID
		ret.OrderID = record.OrderID
		ret.GraphID = record.GraphID
		ret.Title = record.Title
		ret.Env = record.Env
		ret.GraphName = record.GraphName
		ret.OrderName = record.OrderName
		ret.OrderProcess = orderProcess
		ret.OrderType = record.OrderType
		ret.DemandName = record.DemandName
		ret.Owner = record.Owner
		ret.TaskStatus = record.TaskStatus
		ret.Status = record.Status
		ret.ActiveIndex = activeIndex
		ret.CurOrderNode = curOrderNode
		ret.HasApproval = hasApproval
		ret.UpdatedTime = util.FormatTimeToString(record.UpdatedTime)
		ret.CreatedTime = util.FormatTimeToString(record.CreatedTime)
		retList = append(retList, ret)
	}
	resultData := &result.ProcessOrderResult{}
	resultData.Total = total
	resultData.RetList = retList
	return resultData, nil
}

func (s *processOrderService) GetApprovalList(ctx context.Context, req *request.GetProcessOrderReq) (*result.ProcessOrderResult, error) {
	filter := &porderrepo.ProcessOrderFilter{
		GraphName:  req.GraphName,
		OrderType:  req.OrderType,
		DemandName: req.DemandName,
		Owner:      req.Owner,
		Approver:   req.Approver,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		Status:     req.Status,
		Page:       req.Page,
		PageSize:   req.PageSize,
	}
	records, total, err := s.processOrderRepo.GetApprovalList(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	retList := make([]*result.ProcessOrderRet, 0, len(records))
	for _, record := range records {
		var orderProcess []*model.PorcessNode
		err := json.Unmarshal(record.OrderProcess, &orderProcess)
		if err != nil {
			continue
		}
		activeIndex := s.genActiveIndex(orderProcess)
		curNodeIndex := activeIndex
		if activeIndex == int8(len(orderProcess)) {
			curNodeIndex = activeIndex - 1
		}
		curOrderNode := orderProcess[curNodeIndex]
		ret := &result.ProcessOrderRet{}
		ret.ID = record.ID
		ret.OrderID = record.OrderID
		ret.GraphID = record.GraphID
		ret.Title = record.Title
		ret.Env = record.Env
		ret.GraphName = record.GraphName
		ret.OrderName = record.OrderName
		ret.OrderProcess = orderProcess
		ret.OrderType = record.OrderType
		ret.DemandName = record.DemandName
		ret.Owner = record.Owner
		ret.Status = record.Status
		ret.ActiveIndex = activeIndex
		ret.CurOrderNode = curOrderNode
		ret.UpdatedTime = util.FormatTimeToString(record.UpdatedTime)
		ret.CreatedTime = util.FormatTimeToString(record.CreatedTime)
		retList = append(retList, ret)
	}
	resultData := &result.ProcessOrderResult{}
	resultData.Total = total
	resultData.RetList = retList
	return resultData, nil
}

func (s *processOrderService) getRolesByUser(ctx context.Context) ([]string, error) {
	roleNames := make([]string, 0, 0)
	userName := util.GetUserName(ctx)
	user, err := s.userRepo.GetByName(ctx, userName)
	if err != nil {
		return roleNames, errors.Wrap(err, "获取角色信息时，查询用户信息失败")
	}
	userID := user.ID
	roles, err := s.roleRepo.GetRoles(ctx, userID)
	if err != nil {
		return roleNames, errors.Wrap(err, "获取角色信息时，查询角色信息失败")
	}
	for _, role := range roles {
		roleNames = append(roleNames, role.RoleName)
	}
	return roleNames, nil
}

// 当前审批节点的工单 对当前用户 是否有审批权限，还需要处理不同部门 组长 和 总监 的情况
func (s *processOrderService) isApprovalOrder(approver string, roleNames []string, curOrderNode *model.PorcessNode, teamNameInfo map[string]*biz.TeamInfo) int8 {
	if len(curOrderNode.Approver) > 0 && curOrderNode.Approver == approver {
		return 1
	}
	var leaders, directors []string
	approvalInfo := curOrderNode.ApprovalInfo
	deptName := curOrderNode.DeptName
	teamInfo, ok := teamNameInfo[deptName]
	if ok {
		leaders = strings.Split(teamInfo.Leader, ",")
		directors = strings.Split(teamInfo.Director, ",")
	}
	leaderRoles := []string{
		"leader",
		"teamLeader",
	}
	directorRoles := []string{
		"director",
	}
	for _, aprInfo := range approvalInfo {
		// 在这里判断角色类型
		// 审批的角色是组长或者总监
		if lo.Contains(leaderRoles, aprInfo.Role) {
			if len(leaders) > 0 {
				if lo.Contains(leaders, approver) {
					return 1
				} else {
					return 0
				}
			}
		}
		if lo.Contains(directorRoles, aprInfo.Role) {
			if len(directors) > 0 {
				if lo.Contains(directors, approver) {
					return 1
				} else {
					return 0
				}
			}
		}
		if aprInfo.Status == 0 && lo.Contains(roleNames, aprInfo.Role) {
			return 1
		}
	}
	return 0
}

func (s *processOrderService) genTeamInfo(ctx context.Context) (map[string]*biz.TeamInfo, error) {
	teamNameInfo := make(map[string]*biz.TeamInfo)
	teams, err := s.teamRepo.GetAll(ctx)
	if err != nil {
		return teamNameInfo, errors.Wrap(err, "get team info error")
	}
	for _, team := range teams {
		teamNameInfo[team.Name] = &biz.TeamInfo{
			Name:     team.Name,
			Leader:   team.Leader,
			Director: team.Director,
		}
	}
	return teamNameInfo, nil
}

func (s *processOrderService) GetByID(ctx context.Context, id int64) (*result.ProcessOrderDetailResult, error) {

	record, err := s.processOrderRepo.GetByID(ctx, id)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}

	// 获取工单关联的订单信息
	orderLayout := int8(2)
	orderLabel := ""
	// order type 大于2 的才 获取工单信息
	if record.OrderType > 2 {
		orderRecord, err := s.orderRepo.GetByID(ctx, record.OrderID)
		if err != nil {
			err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
			return nil, err
		}
		orderLayout = orderRecord.Layout
		orderLabel = orderRecord.Label
	}
	var orderProcess []*model.PorcessNode
	err = json.Unmarshal(record.OrderProcess, &orderProcess)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	activeIndex := s.genActiveIndex(orderProcess)
	curNodeIndex := activeIndex
	if activeIndex == int8(len(orderProcess)) {
		curNodeIndex = activeIndex - 1
	}
	curOrderNode := orderProcess[curNodeIndex]
	// 如果有子工单获取子工单
	imageHash := record.ImageHash
	enabledImageHash := record.EnabledImageHash
	var imageData, enabledImageData string
	processArch1, err := s.processArchRepo.GetImageHash(ctx, imageHash)
	if err == nil {
		imageData = processArch1.ImageData
	}
	processArch2, err := s.processArchRepo.GetEnabledImageHash(ctx, enabledImageHash)
	if err == nil {
		enabledImageData = processArch2.EnabledImageData
	}
	taskResult := record.TaskResult
	if taskResult == nil || len(taskResult) == 0 {
		taskResult = []byte("{}")
	}
	orderField := record.OrderField
	if orderField == nil || len(orderField) == 0 {
		orderField = []byte("[]")
	}
	resultData := &result.ProcessOrderDetailResult{}
	resultData.ID = record.ID
	resultData.OrderID = record.OrderID
	resultData.GraphID = record.GraphID
	resultData.Title = record.Title
	resultData.Env = record.Env
	resultData.GraphName = record.GraphName
	resultData.OrderName = record.OrderName
	resultData.OrderLayout = orderLayout
	resultData.OrderLabel = orderLabel
	resultData.DemandName = record.DemandName
	resultData.OrderInfo = record.OrderInfo
	resultData.OrderField = orderField
	resultData.OrderProcess = orderProcess
	resultData.ActiveIndex = activeIndex
	resultData.CurOrderNode = curOrderNode
	resultData.ImageHash = imageHash
	resultData.EnabledImageHash = enabledImageHash
	resultData.ImageData = imageData
	resultData.EnabledImageData = enabledImageData
	resultData.OrderType = record.OrderType
	resultData.Description = record.Description
	resultData.TaskStatus = record.TaskStatus
	resultData.TaskResult = taskResult
	resultData.Status = record.Status
	resultData.Opinion = s.getApprovalOpinion(ctx, record.ID)
	resultData.UpdatedTime = util.FormatTimeToString(record.UpdatedTime)
	resultData.CreatedTime = util.FormatTimeToString(record.CreatedTime)
	return resultData, nil
}

func (s *processOrderService) GetUnapprovedList(ctx context.Context, req *request.GetUnapprovedOrderReq) (*result.ProcessOrderUnapprovedResult, error) {

	graphID := *req.GraphID
	status := int8(0)
	records, err := s.processOrderRepo.GetByIDStatus(ctx, graphID, status)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	retList := make([]*result.ProcessOrderUnapprovedRet, 0, len(records))
	for _, record := range records {
		ret := &result.ProcessOrderUnapprovedRet{}
		// 根据title 查询
		if req.Title != nil && len(*req.Title) > 0 {
			if strings.Contains(record.Title, *req.Title) {
				continue
			}
		}
		ret.ID = record.ID
		ret.Title = record.Title
		ret.Env = record.Env
		ret.GraphID = record.GraphID
		ret.OrderID = record.OrderID
		ret.GraphName = record.GraphName
		ret.OrderType = record.OrderType
		ret.OrderName = record.OrderName
		ret.DemandName = record.DemandName
		ret.Status = record.Status
		ret.OrderInfo = record.OrderInfo
		ret.UpdatedTime = util.FormatTimeToString(record.UpdatedTime)
		ret.CreatedTime = util.FormatTimeToString(record.CreatedTime)
		retList = append(retList, ret)
	}
	resultData := &result.ProcessOrderUnapprovedResult{}
	resultData.RetList = retList
	return resultData, nil
}

func (s *processOrderService) getApprovalOpinion(ctx context.Context, id int64) string {

	records, err := s.processApprovalRepo.GetByOrder(ctx, id)
	if err != nil {
		return ""
	}
	opinions := make([]string, 0, len(records))
	for _, record := range records {
		approver := record.Approver
		opinion := record.Opinion
		action := record.Action
		createTime := util.FormatTimeToString(record.CreatedTime)
		txt := fmt.Sprintf("审批时间：%s，审批人：%s，审批操作：%s，审批意见：%s", createTime, approver, action, opinion)
		opinions = append(opinions, txt)
	}
	return strings.Join(opinions, "@")
}

//分配审批人

func (s *processOrderService) AssignApprover(ctx context.Context, req *request.AssignApproverReq) error {
	id := *req.ID
	approver := *req.Approver
	processOrder, err := s.processOrderRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
	}

	orderProcess := processOrder.OrderProcess
	var orderProcessData []*model.PorcessNode
	err = json.Unmarshal(orderProcess, &orderProcessData)
	if err != nil {
		return errors.WithCodeError(ce.Error.Code(), err)
	}

	activeIndex := s.genActiveIndex(orderProcessData)
	curNodeIndex := activeIndex
	if int(activeIndex) == len(orderProcessData)-1 {
		curNodeIndex = activeIndex - 1
	}

	// 获取当前正在审批的节点
	curNode := orderProcessData[curNodeIndex]
	curNode.Approver = approver
	orderProcessJson, err := json.Marshal(orderProcessData)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	// 更新状态为0
	updateField := make(map[string]interface{})
	updateField[processOrder.OrderProcessField()] = orderProcessJson
	err = s.processOrderRepo.Update(ctx, id, updateField)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

// 工单类型 1 架构图申请工单 2 架构图变更工单 3 请求类资源工单 4 请求类非资源工单 5 非请求工单
// 创建工单
func (s *processOrderService) Create(ctx context.Context, req *request.CreateProcessOrderReq) error {

	orderType := *req.OrderType
	userName := util.GetUserName(ctx)
	deptName := util.GetDeptName(ctx)
	processOrder := &model.ProcessOrder{}
	if req.GraphID != nil {
		processOrder.GraphID = *req.GraphID
	}
	if req.OrderID != nil {
		processOrder.OrderID = *req.OrderID
	}
	if req.Title != nil {
		processOrder.Title = *req.Title
	}
	if req.Env != nil {
		processOrder.Env = *req.Env
	}
	if req.GraphName != nil {
		processOrder.GraphName = *req.GraphName
	}
	if req.OrderName != nil {
		processOrder.OrderName = *req.OrderName
	}
	if req.OrderType != nil {
		processOrder.OrderType = *req.OrderType
	}
	if req.DemandName != nil {
		processOrder.DemandName = *req.DemandName
	}
	if req.Owner != nil {
		processOrder.Owner = *req.Owner
	} else {
		processOrder.Owner = userName
	}
	// 生成imageHash
	if req.ImageData != nil && len(*req.ImageData) != 0 {
		imageHash := util.HashBase64(*req.ImageData)
		processOrder.ImageHash = imageHash
	}
	if req.EnabledImageData != nil && len(*req.EnabledImageData) != 0 {
		enabledImageHash := util.HashBase64(*req.EnabledImageData)
		processOrder.EnabledImageHash = enabledImageHash
	}
	if req.Description != nil {
		processOrder.Description = *req.Description
	}
	if req.OrderInfo != nil {
		// 在这里写一个 校验orderInfo的方法
		verMsg, err := s.validateOrderInfo(ctx, *req.OrderID, req.OrderInfo)
		if err != nil && len(verMsg) > 0 {
			log.Logger.Error(verMsg, zap.Error(err))
			return errors.WithCodeError(ce.ErrorValidateForm.Code(), err)
		}
		processOrder.OrderInfo = req.OrderInfo
	}
	processOrder.Status = 1
	// 请求资源工单创建初始状态为未审批，提交审批的操作
	if orderType == 3 {
		processOrder.Status = 0
	}
	// orderType 1 2 是架构图图审批
	if orderType <= 2 {
		if len(GraphOrderPorcess) > 1 {
			if len(GraphOrderPorcess[0].ApprovalInfo) > 0 {
				GraphOrderPorcess[0].ApprovalInfo[0].ApproverName = userName
				GraphOrderPorcess[0].ApprovalInfo[0].Approver = userName
				GraphOrderPorcess[0].ApprovalInfo[0].Status = 1
				GraphOrderPorcess[0].Status = 1
			}
		}
		for _, node := range GraphOrderPorcess {
			node.DeptName = deptName
		}
		orderProcess, err := json.Marshal(GraphOrderPorcess)
		if err != nil {
			return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
		}
		processOrder.OrderProcess = orderProcess
	} else {
		// 非架构图图的工单，通过order_id 获取流程信息
		if req.OrderID == nil {
			return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), errors.New("工单ID为空"))
		}
		orderID := *req.OrderID
		processData, err := s.getOrderProcess(ctx, orderID)
		if err != nil {
			return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
		}
		if len(processData) > 1 {
			if len(processData[0].ApprovalInfo) > 0 {
				processData[0].ApprovalInfo[0].ApproverName = userName
				processData[0].ApprovalInfo[0].Approver = userName
				processData[0].ApprovalInfo[0].Status = 1
				processData[0].Status = 1
			}
		}
		for _, node := range processData {
			node.DeptName = deptName
		}
		orderProcess, err := json.Marshal(processData)
		if err != nil {
			return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
		}
		processOrder.OrderProcess = orderProcess

	}
	tx := s.db.WithContext(ctx).Begin()
	defer model.RecoverRollback(tx)
	processOrderTx := s.processOrderRepo.WithTx(tx)
	archGraphTx := s.archGraphRepo.WithTx(tx)
	err := processOrderTx.Create(ctx, processOrder)
	if err != nil {
		tx.Rollback()
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	// 架构图相关的工单才更新架构图的状态
	if orderType <= 2 {
		// 创建成功后修改流程图的状态为 1 审批中
		if req.GraphID == nil {
			return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), errors.New("没有传入graphID"))
		}
		graphID := *req.GraphID
		updateField := make(map[string]interface{})
		updateField["status"] = 1
		err = archGraphTx.Update(ctx, graphID, updateField)
		if err != nil {
			tx.Rollback()
			return errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		}
		// 生成流程架构信息，只有架构图工单才会生成架构图流程信息
		err = s.genProcessOrderArch(ctx, req.ImageData, req.EnabledImageData)
		if err != nil {
			tx.Rollback()
			return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
		}
	}
	tx.Commit()
	return nil
}

func (s *processOrderService) genProcessOrderArch(ctx context.Context, imageData, enabledImageData *string) error {
	processArch := &model.ProcessArch{}
	isCreate := false
	if imageData != nil && len(*imageData) != 0 {
		imageHash := util.HashBase64(*imageData)
		_, err := s.processArchRepo.GetImageHash(ctx, imageHash)
		if err != nil {
			if !errors.Is(err, model.ErrRecordNotFound) {
				return errors.Wrap(err, "查询流程工单架构图信息失败")
			}
			// 创建一条记录
			processArch.ImageData = *imageData
			processArch.ImageHash = imageHash
			isCreate = true
		}
	}
	if enabledImageData != nil && len(*enabledImageData) != 0 {
		enabledImageHash := util.HashBase64(*enabledImageData)
		_, err := s.processArchRepo.GetEnabledImageHash(ctx, enabledImageHash)
		if err != nil {
			if !errors.Is(err, model.ErrRecordNotFound) {
				return errors.Wrap(err, "查询流程工单架构图信息失败")
			}
			// 创建一条记录
			processArch.EnabledImageData = *enabledImageData
			processArch.EnabledImageHash = enabledImageHash
			isCreate = true
		}
	}
	if isCreate {
		err := s.processArchRepo.Create(ctx, processArch)
		if err != nil {
			return errors.Wrap(err, "创建流程工单架构图信息失败")
		}
	}
	return nil
}

func (s *processOrderService) getOrderProcess(ctx context.Context, orderID int64) ([]*model.PorcessNode, error) {

	order, err := s.orderRepo.GetByID(ctx, orderID)
	if err != nil {
		return nil, errors.Wrap(err, "获取工单信息错误")
	}
	process, err := s.processRepo.GetByID(ctx, order.ProcessID)
	if err != nil {
		return nil, errors.Wrap(err, "获取工单流程信息失败")
	}

	procData := make([]*model.PorcessNode, 0, 0)
	procInfo := process.ProcInfo
	err = json.Unmarshal(procInfo, &procData)
	if err != nil {
		return nil, errors.Wrap(err, "解析流程信息失败")
	}
	return procData, nil
}

// 对创建没有提交申请的工单 提交申请
func (s *processOrderService) Apply(ctx context.Context, req *request.ApplyProcessOrderReq) error {
	graphID := *req.GraphID
	orderType := *req.OrderType
	processOrder := &model.ProcessOrder{}
	// orderType 3 请求资源类工单
	if orderType != 3 {
		return errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), errors.New("order type is not 3"))
	}
	tx := s.db.WithContext(ctx).Begin()
	defer model.RecoverRollback(tx)
	processOrderTx := s.processOrderRepo.WithTx(tx)
	archGraphTx := s.archGraphRepo.WithTx(tx)
	ids := req.IDs
	archGraph, err := archGraphTx.GetByID(ctx, graphID)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
	}
	imageHash := util.HashBase64(archGraph.ImageData)
	updateField := make(map[string]interface{})
	updateField[processOrder.StatusField()] = 1
	updateField[processOrder.ImageHashField()] = imageHash
	// 根據ids 批量更新
	err = processOrderTx.Updates(ctx, ids, updateField)
	if err != nil {
		tx.Rollback()
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	// 生成一条流程工单的记录,防止直接修改状态的情况产生
	err = s.genProcessOrderArch(ctx, &archGraph.ImageData, nil)
	if err != nil {
		tx.Rollback()
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	tx.Commit()
	return nil
}

// 这个方法只针对 order type = 3 请求资源工单
func (s *processOrderService) ReApply(ctx context.Context, req *request.ReApplyProcessOrderReq) error {
	id := *req.ID
	processOrder, err := s.processOrderRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
	}
	if processOrder.OrderType != 3 {
		return errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), errors.New("order type is not 3"))
	}

	// 更新工单流程状态
	userName := util.GetUserName(ctx)
	orderProcess := processOrder.OrderProcess
	var orderProcessData []*model.PorcessNode
	err = json.Unmarshal(orderProcess, &orderProcessData)
	if err != nil {
		return errors.WithCodeError(ce.Error.Code(), err)
	}
	if len(orderProcessData) > 1 {
		if len(orderProcessData[0].ApprovalInfo) > 0 {
			orderProcessData[0].ApprovalInfo[0].ApproverName = userName
			orderProcessData[0].ApprovalInfo[0].Approver = userName
			orderProcessData[0].ApprovalInfo[0].Status = 1
			orderProcessData[0].Status = 1
		}
	}
	orderProcessJson, err := json.Marshal(orderProcessData)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	// 更新状态为0
	updateField := make(map[string]interface{})
	updateField[processOrder.OrderProcessField()] = orderProcessJson

	updateField[processOrder.StatusField()] = 0
	err = s.processOrderRepo.Update(ctx, id, updateField)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

// 重新编辑工单时更新，在审批时候编辑工单时更新
func (s *processOrderService) Update(ctx context.Context, req *request.UpdateProcessOrderReq) error {
	id := *req.ID
	orderType := *req.OrderType
	processOrder, err := s.processOrderRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	updateField := make(map[string]interface{})
	if req.DemandName != nil {
		updateField[processOrder.DemandNameField()] = *req.DemandName
	}
	if req.Env != nil {
		processOrder.Env = *req.Env
	}
	if req.OrderInfo != nil {
		// 在这里写一个 校验orderInfo的方法
		verMsg, err := s.validateOrderInfo(ctx, processOrder.OrderID, req.OrderInfo)
		if err != nil && len(verMsg) > 0 {
			log.Logger.Error(verMsg, zap.Error(err))
			return errors.WithCodeError(ce.ErrorValidateForm.Code(), err)
		}
		updateField[processOrder.OrderInfoField()] = req.OrderInfo
	}
	// 特别注意这里的逻辑
	// order type 为 4 5 请求非资源工单 非请求工单 更新的时候一定要更新状态，只有在审批失败时才重新提交，这里注意在审批时候也可能修改工单
	// isApproval 有值的时候是在审批时候更新，这时候不更新流程的状态
	if (orderType == 4 || orderType == 5) && (req.IsApproval == nil) {
		userName := util.GetUserName(ctx)
		updateField[processOrder.StatusField()] = 1
		orderProcess := processOrder.OrderProcess
		var orderProcessData []*model.PorcessNode
		err = json.Unmarshal(orderProcess, &orderProcessData)
		if err != nil {
			return errors.WithCodeError(ce.Error.Code(), err)
		}
		if len(orderProcessData) > 1 {
			if len(orderProcessData[0].ApprovalInfo) > 0 {
				orderProcessData[0].ApprovalInfo[0].ApproverName = userName
				orderProcessData[0].ApprovalInfo[0].Approver = userName
				orderProcessData[0].ApprovalInfo[0].Status = 1
				orderProcessData[0].Status = 1
			}
		}
		orderProcessJson, err := json.Marshal(orderProcessData)
		if err != nil {
			return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
		}
		updateField[processOrder.OrderProcessField()] = orderProcessJson
	}
	err = s.processOrderRepo.Update(ctx, id, updateField)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

func (s *processOrderService) Delete(ctx context.Context, id int64) error {
	_, err := s.processOrderRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	err = s.processOrderRepo.DeleteByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

// 审批架构图的工单 orderTyoe 1 2
func (s *processOrderService) Approve(ctx context.Context, req *request.ApproveProcessOrderReq) error {
	id := *req.ID
	action := *req.Action
	// 审批人
	approver := *req.Approver
	approverName := *req.ApproverName
	graphID := *req.GraphID
	processOrder, err := s.processOrderRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	// 通过approver 获取roles
	roleNames, err := s.getRolesByUser(ctx)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	// 图状态 0 未审批 1 审批中 2 审批成功 3审批失败 4 审批成功失效
	orderProcess := processOrder.OrderProcess
	var orderProcessData []*model.PorcessNode
	err = json.Unmarshal(orderProcess, &orderProcessData)
	if err != nil {
		return errors.WithCodeError(ce.Error.Code(), err)
	}
	// 处理工单流程
	var procNodeName string
	if req.ProcNodeName != nil {
		procNodeName = *req.ProcNodeName
	}
	status, orderProcessData := s.dealOrderProcess(action, approver, approverName, procNodeName, roleNames, orderProcessData)

	orderProcessJSON, err := json.Marshal(orderProcessData)
	if err != nil {
		return errors.WithCodeError(ce.Error.Code(), err)
	}
	// 事务相关
	tx := s.db.WithContext(ctx).Begin()
	defer model.RecoverRollback(tx)

	processOrderTx := s.processOrderRepo.WithTx(tx)
	processApprovalTx := s.processApprovalRepo.WithTx(tx)
	actionText := "通过"
	if status == 3 {
		actionText = "不通过"
	}
	processApproval := &model.ProcessApproval{}
	processApproval.Approver = approverName
	processApproval.OrderID = id
	processApproval.Opinion = *req.Opinion
	processApproval.Status = status
	processApproval.Action = actionText
	// 创建一条审批记录
	err = processApprovalTx.Create(ctx, processApproval)
	if err != nil {
		tx.Rollback()
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	// 更新订单，状态，节点信息
	updateOrder := make(map[string]interface{})
	updateOrder["status"] = status
	updateOrder["order_process"] = orderProcessJSON

	//审批不通过 清空一些字段的情况
	if processOrder.OrderType > 2 && status == 3 {
		orderInfo, err := s.clearOrderInfo(ctx, processOrder.OrderID, processOrder.OrderInfo)
		if err != nil {
			log.Logger.Error("clear order info err:", zap.Error(err))
		} else {
			updateOrder["order_info"] = orderInfo
		}
	}

	err = processOrderTx.Update(ctx, id, updateOrder)
	if err != nil {
		tx.Rollback()
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	// 非架构图工单的情况
	if processOrder.OrderType > 2 {
		tx.Commit()
		return nil
	}
	// 如果状态是大于2的，属于资源申请相关，就不往下走了，下面是架构图审批的操作
	// 如果审批生效 生成一条最新的审批记录，数据从process_record表里面查
	archGraphTx := s.archGraphRepo.WithTx(tx)
	archGraph, err := archGraphTx.GetByID(ctx, graphID)
	if err != nil {
		tx.Rollback()
		return errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
	}
	// 审批动作是success  status = 2 的情况下 graphKey 一致的 archGraph 中 修改老的图的状态为4
	if status == 2 {
		graphKey := archGraph.GraphKey
		enabledArchGraphs, err := archGraphTx.GetEnabledList(ctx, graphKey)
		if err != nil {
			tx.Rollback()
			return errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		}
		ids := make([]int64, 0, 0)
		for _, enabledArchGraph := range enabledArchGraphs {
			ids = append(ids, enabledArchGraph.ID)
		}
		updateEnabledGraph := make(map[string]interface{})
		updateEnabledGraph["status"] = 4
		err = archGraphTx.UpdateByIDs(ctx, ids, updateEnabledGraph)
		if err != nil {
			tx.Rollback()
			return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
		}
	}
	//修改archGraph状态
	updateGraph := make(map[string]interface{})
	updateGraph["status"] = status
	err = archGraphTx.Update(ctx, graphID, updateGraph)
	if err != nil {
		tx.Rollback()
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	// 重新生成一条记录
	archGraphRecord := model.NewArchGraphRecord()
	archGraphRecord.GraphID = graphID
	archGraphRecord.NodeData = archGraph.NodeData
	archGraphRecord.EdgeData = archGraph.EdgeData
	archGraphRecord.ImageData = archGraph.ImageData
	// 3 是审批通过的架构图
	archGraphRecord.RecordType = 3
	err = s.archGraphRecordRepo.Create(ctx, archGraphRecord)
	if err != nil {
		tx.Rollback()
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	tx.Commit()
	return nil
}
func (s *processOrderService) dealOrderProcess(action, approver, approverName, procNodeName string, roleNames []string, orderProcessData []*model.PorcessNode) (int8, []*model.PorcessNode) {
	status := int8(1)
	activeIndex := s.genActiveIndex(orderProcessData)
	curNodeIndex := activeIndex
	if int(activeIndex) == len(orderProcessData)-1 {
		curNodeIndex = activeIndex - 1
	}
	//审批通过
	if action == "pass" {
		// 获取当前正在审批的节点
		curNode := orderProcessData[curNodeIndex]
		status = 1
		approvalType := curNode.ApprovalType
		//
		for i := range curNode.ApprovalInfo {
			// 并行的审批节点状态是0 并且当前用户有审批的权限
			if curNode.ApprovalInfo[i].Status == 0 && (lo.Contains(roleNames, curNode.ApprovalInfo[i].Role) || approver == curNode.Approver) {
				curNode.ApprovalInfo[i].Approver = approver
				curNode.ApprovalInfo[i].ApproverName = approverName
				curNode.ApprovalInfo[i].Status = 1
			}
		}
		// approvalType 0 一个子节点通过就全通过，1 所有的通过才通过
		if approvalType == 0 {
			curNode.Status = 1
		} else {
			allPass := true
			// 这个时候判断当前节点的ApprovalInfo 中所有元素的 status 的状态
			for _, aprInfo := range curNode.ApprovalInfo {
				// 并行的审批节点状态是0 并且当前用户有审批的权限
				if aprInfo.Status == 0 {
					allPass = false
				}
			}
			// ApprovalInfo 中所有元素的 status 的状态 都为1，当前节点状态才为1
			if allPass {
				curNode.Status = 1
			}
		}
		// 更新最后一个节点
		if int(activeIndex) == len(orderProcessData)-2 {
			status = 2
			// 审批成功 最后一个节点变成成功
			lastNodeIndex := len(orderProcessData) - 1
			orderProcessData[lastNodeIndex].Status = 1
			orderProcessData[lastNodeIndex].ApprovalInfo[0].Status = 1
		}

	} else if action == "success" {
		status = 2
		// 修改当前节点
		// 修改当前之后节点
		for _, procNode := range orderProcessData[curNodeIndex:] {
			//if procNode.Type == "procStart" {
			//	continue
			//}
			procNode.Status = 1
			for i := range procNode.ApprovalInfo {
				// 并行的审批节点状态是0 并且当前用户有审批的权限
				procNode.ApprovalInfo[i].Approver = approver
				procNode.ApprovalInfo[i].ApproverName = approverName
				procNode.ApprovalInfo[i].Status = 1
			}
		}
	} else {
		// 审批失败 status = 3 , 除了procStart 节点 说有节点的状态都变为0
		status = 3
		if len(procNodeName) == 0 {
			for _, procNode := range orderProcessData {
				procNode.Status = 0
				for i := range procNode.ApprovalInfo {
					// 并行的审批节点状态是0 并且当前用户有审批的权限
					procNode.ApprovalInfo[i].Status = 0
				}
			}
		} else {
			for i := len(orderProcessData) - 1; i >= 0; i-- {
				procNode := orderProcessData[i]
				procNode.Status = 0
				for i := range procNode.ApprovalInfo {
					// 并行的审批节点状态是0 并且当前用户有审批的权限
					procNode.ApprovalInfo[i].Status = 0
				}
				if procNode.Name == procNodeName {
					break
				}

			}
		}

	}
	return status, orderProcessData
}

// 校验orderInfo方法
func (s *processOrderService) validateOrderInfo(ctx context.Context, orderID int64, orderInfo datatypes.JSON) (string, error) {
	orderFields, err := s.orderFieldRepo.GetByOrder(ctx, orderID)
	if err != nil {
		return "", errors.Wrap(err, "校验工单信息，获取工单字段错误")
	}
	orderKeyField := make(map[string]*model.OrderField)
	var ignoreComponents []string = []string{"select", "multipleSelect", "dateTimePicker", "datePicker"}
	orderKeyName := make(map[string]string)
	for _, orderField := range orderFields {
		component := orderField.Component
		verRule := orderField.VerRule
		// 没有校验规则的不校验
		if verRule == 1 {
			continue
		}
		// 需要选择的字段不校验
		if lo.Contains(ignoreComponents, component) {
			continue
		}
		orderKeyField[orderField.Key] = orderField
		orderKeyName[orderField.Key] = orderField.Name
	}
	//
	orderInfoData := &biz.OrderInfo{}
	err = json.Unmarshal(orderInfo, &orderInfoData)
	if err != nil {
		return "", err
	}
	formData := orderInfoData.FormData
	groupFormDataInfo := orderInfoData.GroupFormDataInfo

	for key, val := range formData {
		field, ok := orderKeyField[key]
		if !ok {
			continue
		}
		verRule := field.VerRule
		if !util.CheckValue(verRule, val) {
			ruleTxt, ok := util.VerRuleMapping[verRule]
			if !ok {
				ruleTxt = "校验失败"
			}
			name, ok := orderKeyName[key]
			if !ok {
				name = key
			}
			msg := fmt.Sprintf("字段:%s,校验失败:字段格式:%s", name, ruleTxt)
			return msg, errors.New(msg)
		}
	}
	for group, formDatas := range groupFormDataInfo {
		i := 1
		for _, form := range formDatas {
			for key, val := range form {
				field, ok := orderKeyField[key]
				if !ok {
					continue
				}
				verRule := field.VerRule
				if !util.CheckValue(verRule, val) {
					ruleTxt, ok := util.VerRuleMapping[verRule]
					if !ok {
						ruleTxt = "校验失败"
					}
					name, ok := orderKeyName[key]
					if !ok {
						name = key
					}
					msg := fmt.Sprintf("分组:%s,第%d个表单,字段:%s,校验失败:字段格式:%s", group, i, name, ruleTxt)
					return msg, errors.New(msg)
				}
			}
			i++
		}
	}
	return "", nil
}

func (s *processOrderService) clearOrderInfo(ctx context.Context, orderID int64, orderInfo datatypes.JSON) (datatypes.JSON, error) {
	orderFields, err := s.orderFieldRepo.GetByOrder(ctx, orderID)
	if err != nil {
		return []byte("{}"), errors.Wrap(err, "校验工单信息，获取工单字段错误")
	}
	orderClearKey := make(map[string]struct{})
	for _, orderField := range orderFields {
		isClear := orderField.IsClear
		// 不需要清空的字段
		if isClear == 0 {
			continue
		}
		orderClearKey[orderField.Key] = struct{}{}
	}
	//
	orderInfoData := &biz.OrderInfo{}
	err = json.Unmarshal(orderInfo, &orderInfoData)
	if err != nil {
		return []byte("{}"), err
	}
	formData := orderInfoData.FormData
	groupFormDataInfo := orderInfoData.GroupFormDataInfo

	for key, _ := range orderClearKey {
		formData[key] = ""
	}
	for _, formDatas := range groupFormDataInfo {
		for _, form := range formDatas {
			for key, _ := range orderClearKey {
				form[key] = ""
			}
		}
	}
	orderInfoJson, err := json.Marshal(orderInfoData)
	if err != nil {
		return []byte("{}"), err
	}
	return orderInfoJson, nil
}
