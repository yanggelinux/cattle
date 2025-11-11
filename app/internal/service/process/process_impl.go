package process

import (
	"context"
	"encoding/json"
	"github.com/samber/lo"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	processrepo "github.com/yanggelinux/cattle/internal/repository/process"
	"github.com/yanggelinux/cattle/internal/store/model"
	"github.com/yanggelinux/cattle/pkg/util"
	"gorm.io/datatypes"
	"strings"
)

type Properties struct {
	Width    float64   `json:"width"`
	Height   float64   `json:"height"`
	NodeInfo *NodeInfo `json:"nodeInfo"`
}

type NodeInfo struct {
	Name         string `json:"name"`
	Role         string `json:"role"`
	RoleName     string `json:"roleName"`
	ApprovalType int8   `json:"approvalType"`
	ApprovalEdit int8   `json:"approvalEdit"`
}

type Text struct {
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Value string  `json:"value"`
}
type Node struct {
	X          float64    `json:"x"`
	Y          float64    `json:"y"`
	ZIndex     float64    `json:"zIndex"`
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Text       Text       `json:"text"`
	Properties Properties `json:"properties"`
}

type Edge struct {
	SourceNodeId   string `json:"sourceNodeId"`
	TargetNodeId   string `json:"targetNodeId"`
	SourceAnchorId string `json:"sourceAnchorId"`
	TargetAnchorId string `json:"targetAnchorId"`
	X              int64  `json:"x"`
	Y              int64  `json:"y"`
	ZIndex         int64  `json:"zIndex"`
	ID             string `json:"id"`
	Type           string `json:"type"`
	Text           Text   `json:"text"`
}

type processService struct {
	processRepo processrepo.ProcessRepo
}

func NewProcessService() ProcessService {
	return &processService{processRepo: processrepo.NewProcessRepo()}
}

func (s *processService) GetList(ctx context.Context, req *request.GetProcessReq) (*result.ProcessResult, error) {

	//进行一个copy
	filter := &processrepo.ProcessFilter{
		Name:     req.Name,
		Status:   req.Status,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	records, total, err := s.processRepo.GetList(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	retList := make([]*result.ProcessRet, 0, len(records))
	for _, record := range records {
		ret := &result.ProcessRet{}
		ret.ID = record.ID
		ret.Name = record.Name
		ret.ProcInfo = record.ProcInfo
		ret.Status = record.Status
		ret.UpdatedTime = util.FormatTimeToString(record.UpdatedTime)
		ret.CreatedTime = util.FormatTimeToString(record.CreatedTime)
		retList = append(retList, ret)
	}
	resultData := &result.ProcessResult{}
	resultData.Total = total
	resultData.RetList = retList
	return resultData, nil
}

func (s *processService) GetDetail(ctx context.Context, id int64) (*result.ProcessDetailRet, error) {
	record, err := s.processRepo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	ret := &result.ProcessDetailRet{}
	ret.ID = record.ID
	ret.Name = record.Name
	ret.ProcInfo = record.ProcInfo
	ret.NodeData = record.NodeData
	ret.EdgeData = record.EdgeData
	ret.Status = record.Status
	ret.UpdatedTime = util.FormatTimeToString(record.UpdatedTime)
	ret.CreatedTime = util.FormatTimeToString(record.CreatedTime)

	return ret, nil
}

func (s *processService) Create(ctx context.Context, req *request.CreateProcessReq) (*result.ProcessOptResult, error) {
	process := model.NewProcess()
	process.Name = *req.Name
	if req.NodeData != nil {
		process.NodeData = req.NodeData
	}
	if req.EdgeData != nil {
		process.EdgeData = req.EdgeData
	}
	if req.Status == nil {
		process.Status = int8(1)
	} else {
		process.Status = *req.Status
	}
	checked := s.processRepo.CheckDuplicateEntry(ctx, *req.Name)
	if checked {
		return nil, errors.WithCodeError(ce.ErrorDBDuplicateEntry.Code(), errors.New("Duplicate Entry error"))
	}
	err := s.processRepo.Create(ctx, process)
	if err != nil {
		return nil, errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	resultData := &result.ProcessOptResult{}
	resultData.ID = process.ID
	return resultData, err
}
func (s *processService) genProcInfo(nodeData, edgeData datatypes.JSON) (datatypes.JSON, error) {

	nodes := make([]*Node, 0, 0)
	err := json.Unmarshal(nodeData, &nodes)
	if err != nil {
		return nil, err
	}
	edges := make([]*Edge, 0, 0)
	err = json.Unmarshal(edgeData, &edges)
	if err != nil {
		return nil, err
	}
	// 判断是否是开始节点和结束节点结尾
	// 判断是一条线，节点不出现分叉
	// 从开始到结束 生成proc info列表
	// 判断节点信息是否存在
	procInfoBytes, err := s.genCheckProcInfo(nodes, edges)
	if err != nil {
		return nil, err
	}
	return procInfoBytes, nil
}

func (s *processService) genCheckProcInfo(nodes []*Node, edges []*Edge) (datatypes.JSON, error) {
	startNodes := make([]*Node, 0, 0)
	endNodes := make([]*Node, 0, 0)
	nodeMap := make(map[string]*Node)
	for _, node := range nodes {
		nodeMap[node.ID] = node
	}

	// 记录所有 source 和 target
	sources := make(map[string]bool)
	targets := make(map[string]bool)

	for _, edge := range edges {
		sources[edge.SourceNodeId] = true
		targets[edge.TargetNodeId] = true
	}
	// 起始节点: 出现在 sources 中，但不在 targets 中
	for id := range sources {
		if !targets[id] {
			if node, ok := nodeMap[id]; ok {
				startNodes = append(startNodes, node)
			}
		}
	}
	// 结束节点: 出现在 targets 中，但不在 sources 中
	for id := range targets {
		if !sources[id] {
			if node, ok := nodeMap[id]; ok {
				endNodes = append(endNodes, node)
			}
		}
	}
	// 判断起始节点列表和结束节点列表的长度 不等于1 保存
	if len(startNodes) != 1 || len(endNodes) != 1 {
		return nil, errors.New("流程的开始或者结束节点有多个")
	}
	startNode := startNodes[0]
	endNode := endNodes[0]
	if startNode.Type != "procStart" {
		return nil, errors.New("流程不是以开始节点开始")
	}
	if endNode.Type != "procEnd" {
		return nil, errors.New("流程不是以结束节点结束")
	}
	if startNode.Properties.NodeInfo == nil {
		return nil, errors.New("开始节点没有输入节点信息")
	}
	if endNode.Properties.NodeInfo == nil {
		return nil, errors.New("结束节点没有编辑节点信息")
	}
	procInfo := make([]*model.PorcessNode, 0, len(nodes))

	startProcNode, err := s.genProcNodeByIDs([]string{startNode.ID}, nodeMap)
	if err != nil {
		return nil, err
	}
	procInfo = append(procInfo, startProcNode)
	nextID := startNode.ID
	has := make([]string, 0, 0)
	for i := 0; i <= len(nodes); i++ {
		nextIDs := s.getNextNodeIDs(nextID, edges)
		//if len(nextIDs) > 1 {
		//	// next节点大于1的情况
		//	return nil, errors.New("流程不是一条节点链路")
		//}
		if len(nextIDs) == 0 {
			// 结束节点
			if lo.Contains(has, nextID) {
				continue
			}
			nextNode := nodeMap[nextID]
			nextNodeInfo := nextNode.Properties.NodeInfo
			if nextNodeInfo == nil {
				return nil, errors.New("节点没有输入节点信息")
			}
			nextProcNode, err := s.genProcNodeByIDs([]string{nextID}, nodeMap)
			if err != nil {
				return nil, err
			}
			procInfo = append(procInfo, nextProcNode)
			has = append(has, nextID)
			break
		}

		// next节点等于1的情况
		nextID = nextIDs[0]
		if lo.Contains(has, nextID) {
			continue
		}
		nextNode := nodeMap[nextID]
		nextNodeInfo := nextNode.Properties.NodeInfo
		if nextNodeInfo == nil {
			return nil, errors.New("审批节点没有输入节点信息")
		}
		nextProcNode, err := s.genProcNodeByIDs(nextIDs, nodeMap)
		if err != nil {
			return nil, err
		}
		procInfo = append(procInfo, nextProcNode)
		has = append(has, nextIDs...)

	}
	procInfoBytes, err := json.Marshal(procInfo)
	if err != nil {
		return nil, errors.Wrap(err, "proc info json marshal failed")
	}
	return procInfoBytes, nil
}

func (s *processService) genProcNodeByIDs(nextIDs []string, nodeMap map[string]*Node) (*model.PorcessNode, error) {

	procNode := &model.PorcessNode{}
	var procNodeNames []string
	var procNodeType string
	var approvalTypes []int8
	var approvalEdits []int8
	approvalInfo := make([]*model.ApprovalInfo, 0, len(nextIDs))
	for _, id := range nextIDs {
		node, ok := nodeMap[id]
		if !ok {
			continue
		}
		nodeType := node.Type
		if nodeType != "procApproval" && len(nextIDs) > 1 {
			return nil, errors.New("并行节点必须是审批节点")
		}
		nodeInfo := node.Properties.NodeInfo
		if nodeInfo == nil {
			return nil, errors.New("审批节点没有输入节点信息")
		}
		procNodeName := nodeInfo.Name
		procNodeType = nodeType
		approvalType := nodeInfo.ApprovalType
		approvalEdit := nodeInfo.ApprovalEdit
		approvalTypes = append(approvalTypes, approvalType)
		arlInfo := &model.ApprovalInfo{}
		arlInfo.Role = nodeInfo.Role
		arlInfo.RoleName = nodeInfo.RoleName
		approvalInfo = append(approvalInfo, arlInfo)
		approvalEdits = append(approvalEdits, approvalEdit)
		procNodeNames = append(procNodeNames, procNodeName)
	}
	procNode.Name = strings.Join(procNodeNames, "\n")
	procNode.Type = procNodeType
	procNode.ApprovalType = lo.Max(approvalTypes)
	procNode.ApprovalEdit = lo.Max(approvalEdits)
	procNode.ApprovalInfo = approvalInfo
	return procNode, nil
}

func (s *processService) getNextNodeIDs(nodeID string, edges []*Edge) []string {
	nextNodeIDs := make([]string, 0, 0)
	for _, edge := range edges {
		if edge.SourceNodeId == nodeID {
			nextNodeIDs = append(nextNodeIDs, edge.TargetNodeId)
		}
	}
	return nextNodeIDs
}

func (s *processService) Update(ctx context.Context, req *request.UpdateProcessReq) error {
	id := *req.ID
	process, err := s.processRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	updateField := make(map[string]interface{})
	// 根据nodedata 和 edgedata 生成 proc_info ,并判断节点是否满足需求
	if req.NodeData != nil && req.EdgeData != nil {
		procInfo, err := s.genProcInfo(req.NodeData, req.EdgeData)
		if err != nil {
			return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
		}
		if procInfo != nil {
			updateField[process.ProcInfoField()] = procInfo
		}
	}
	if req.Name != nil {
		updateField[process.NameField()] = *req.Name
	}
	if req.NodeData != nil {
		updateField[process.NodeDataField()] = req.NodeData
	}
	if req.EdgeData != nil {
		updateField[process.EdgeDataField()] = req.EdgeData
	}
	if req.Status != nil {
		updateField[process.StatusField()] = *req.Status
	}
	err = s.processRepo.Update(ctx, id, updateField)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

func (s *processService) Delete(ctx context.Context, id int64) error {
	_, err := s.processRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	err = s.processRepo.DeleteByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

func (s *processService) Copy(ctx context.Context, id int64) error {
	process, err := s.processRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	newProcess := model.NewProcess()
	newProcess.Name = process.Name + "@copy" + util.RandString(3)
	newProcess.ProcInfo = process.ProcInfo
	newProcess.NodeData = process.NodeData
	newProcess.EdgeData = process.EdgeData
	newProcess.Status = process.Status
	err = s.processRepo.Create(ctx, newProcess)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}
