package archgraph

import (
	"context"
	"encoding/json"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	graphrepo "github.com/yanggelinux/cattle/internal/repository/archgraph"
	userrepo "github.com/yanggelinux/cattle/internal/repository/user"
	"github.com/yanggelinux/cattle/internal/store/model"
	"github.com/yanggelinux/cattle/pkg/log"
	"github.com/yanggelinux/cattle/pkg/util"
	"go.uber.org/zap"
	"strings"
	"time"
)

type archGraphService struct {
	archGraphRepo       graphrepo.ArchGraphRepo
	archGraphRecordRepo graphrepo.ArchGraphRecordRepo
	archGraphReviewRepo graphrepo.ArchGraphReviewRepo
	userRepo            userrepo.UserRepo
}

func NewArchGraphService() ArchGraphService {
	return &archGraphService{
		archGraphRepo:       graphrepo.NewArchGraphRepo(),
		archGraphRecordRepo: graphrepo.NewArchGraphRecordRepo(),
		archGraphReviewRepo: graphrepo.NewArchGraphReviewRepo(),
		userRepo:            userrepo.NewUserRepo(),
	}
}

func (s *archGraphService) GetList(ctx context.Context, req *request.GetArchGraphReq) (*result.ArchGraphResult, error) {

	filter := &graphrepo.ArchGraphFilter{
		GroupID:   req.GroupID,
		GraphName: req.GraphName,
		Status:    req.Status,
		Page:      req.Page,
		PageSize:  req.PageSize,
	}
	records, total, err := s.archGraphRepo.GetList(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	retList := make([]*result.ArchGraphRet, 0, len(records))
	userDept := s.getUserDept(ctx)
	userName := util.GetUserName(ctx)
	uGroup := userDept[userName]
	for _, record := range records {
		// 从这里判断是否展示
		isShared := record.IsShared
		status := record.Status
		owner := record.Owner
		oGroup := userDept[owner]
		// 如果没有组什么都看不到
		if len(oGroup) == 0 || len(uGroup) == 0 || status == 4 {
			continue
		}
		isAppend := false
		//审批通过的情况
		//if status == 2 {
		//	isAppend = true
		//}
		//处理逻辑为 对 uGroup 为devops 用户 进行特殊处理，这个组能看到自己组下的所有数据，以及其它组 status 为2  以及 isShared为1的数据
		//对uGroup 不为 devops的用户，只能看到 userName == owner 以及 status 为2  以及 isShared为1 的数据 但是看不到 oGroup 为devops 的数据
		if uGroup == "devops" {
			// devops组的逻辑
			if oGroup == uGroup || status == 2 || isShared == 1 {
				isAppend = true
			}
		} else {
			// 非devops组逻辑
			// if (owner == userName || status == 2 || isShared == 1) && oGroup != "devops" {
			if owner == userName || status == 2 || isShared == 1 {
				isAppend = true
			}
		}
		isSuper := util.GetSuper(ctx)
		if isSuper == "1" {
			isAppend = true
		}
		if !isAppend {
			continue
		}
		ret := &result.ArchGraphRet{}
		ret.ID = record.ID
		ret.GroupID = record.GroupID
		ret.GroupName = record.GroupName
		ret.GraphName = record.GraphName
		ret.GraphKey = record.GraphKey
		ret.GraphLabel = record.GraphLabel
		ret.ImageData = record.ImageData
		ret.Owner = owner
		ret.Status = status
		ret.IsShared = isShared
		ret.UpdatedTime = util.FormatTimeToString(record.UpdatedTime)
		ret.CreatedTime = util.FormatTimeToString(record.CreatedTime)
		retList = append(retList, ret)
	}
	resultData := &result.ArchGraphResult{}
	resultData.Total = total
	resultData.RetList = retList
	return resultData, nil
}

func (s *archGraphService) getUserDept(ctx context.Context) map[string]string {
	userDept := make(map[string]string)
	userRecords, err := s.userRepo.GetUsers(ctx)
	if err != nil {
		return userDept
	}
	for _, userRecord := range userRecords {
		userDept[userRecord.UserName] = userRecord.DeptName
	}
	return userDept
}

func (s *archGraphService) GetByID(ctx context.Context, id int64) (*result.ArchGraphDetailResult, error) {
	archGraph, err := s.archGraphRepo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	nodeData := archGraph.NodeData
	edgeData := archGraph.EdgeData
	if archGraph.NodeData == nil {
		nodeData = []byte("[]")
	}
	if archGraph.EdgeData == nil {
		edgeData = []byte("[]")
	}
	resultData := &result.ArchGraphDetailResult{}
	resultData.ID = id
	resultData.GroupID = archGraph.GroupID
	resultData.GraphName = archGraph.GraphName
	resultData.GraphKey = archGraph.GraphKey
	resultData.GraphLabel = archGraph.GraphLabel
	resultData.NodeData = nodeData
	resultData.EdgeData = edgeData
	resultData.ImageData = archGraph.ImageData
	resultData.ImageHash = util.HashBase64(archGraph.ImageData)
	resultData.Owner = archGraph.Owner
	resultData.Status = archGraph.Status
	resultData.IsShared = archGraph.IsShared
	resultData.UpdatedTime = util.FormatTimeToString(archGraph.UpdatedTime)
	resultData.CreatedTime = util.FormatTimeToString(archGraph.CreatedTime)
	return resultData, nil
}

func (s *archGraphService) Create(ctx context.Context, req *request.CreateArchGraphReq) (*result.ArchGraphOptResult, error) {
	archGraph := model.NewArchGraph()
	archGraph.GroupID = *req.GroupID
	archGraph.GraphName = *req.GraphName
	archGraph.Owner = *req.Owner
	graphKey := util.GenUUIDv4()
	archGraph.GraphKey = graphKey
	if req.GraphLabel != nil {
		archGraph.GraphLabel = strings.Trim(*req.GraphLabel, " ")
	}
	if req.Status != nil {
		archGraph.Status = *req.Status
	}
	// 对 graph label 做重复检查
	if req.GraphLabel != nil && len(*req.GraphLabel) > 0 {
		checked := s.archGraphRepo.CheckDuplicateEntry(ctx, *req.GroupID, *req.GraphLabel)
		if checked {
			return nil, errors.WithCodeError(ce.ErrorDBDuplicateEntry.Code(), errors.New("Duplicate Entry error"))
		}
	}

	err := s.archGraphRepo.Create(ctx, archGraph)
	if err != nil {
		return nil, errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	resultData := &result.ArchGraphOptResult{}
	resultData.ID = archGraph.ID
	return resultData, err
}

func (s *archGraphService) Save(ctx context.Context, req *request.UpdateArchGraphReq) error {
	id := *req.ID
	archGraph, err := s.archGraphRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	status := *req.Status
	// 针对状态是审批通过的图保存，生成一个新的图
	if status == 2 {
		userName := util.GetUserName(ctx)
		newArchGraph := model.NewArchGraph()
		newArchGraph.GroupID = archGraph.GroupID
		newArchGraph.GraphName = archGraph.GraphName + "@new" + util.RandString(3)
		newArchGraph.GraphKey = archGraph.GraphKey
		newArchGraph.GraphLabel = archGraph.GraphLabel
		newArchGraph.Owner = userName
		newArchGraph.Status = 0
		newArchGraph.IsShared = 0
		newArchGraph.ImageData = *req.ImageData
		newArchGraph.NodeData = req.NodeData
		newArchGraph.EdgeData = req.EdgeData
		err = s.archGraphRepo.Create(ctx, newArchGraph)
		if err != nil {
			return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
		}
		//注意这里只是创建图 并不对已经生效的图修改保存
		return nil
	}
	updateField := make(map[string]interface{})
	updateField[archGraph.ImageDataField()] = *req.ImageData
	updateField[archGraph.NodeDataField()] = req.NodeData
	updateField[archGraph.EdgeDataField()] = req.EdgeData
	err = s.archGraphRepo.Update(ctx, id, updateField)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	// 异步创建快照
	go func() {
		lastRecord, err := s.archGraphRecordRepo.GetSnapshotRecord(ctx, id)
		// 创建
		if err != nil {
			archGraphRecord := model.NewArchGraphRecord()
			archGraphRecord.GraphID = id
			archGraphRecord.NodeData = req.NodeData
			archGraphRecord.EdgeData = req.EdgeData
			archGraphRecord.ImageData = *req.ImageData
			archGraphRecord.RecordType = 1
			err = s.archGraphRecordRepo.Create(ctx, archGraphRecord)
			if err != nil {
				log.Logger.Error("create process graph record error", zap.Error(err))
			}
		} else {
			lastRecord.GraphID = id
			lastRecord.NodeData = req.NodeData
			lastRecord.EdgeData = req.EdgeData
			lastRecord.ImageData = *req.ImageData
			lastRecord.RecordType = 1
			lastRecord.CreatedTime = time.Now()
			err = s.archGraphRecordRepo.CreateOrUpdate(ctx, lastRecord)
			if err != nil {
				log.Logger.Error("create process graph record error", zap.Error(err))
			}
		}
	}()
	return nil
}

// 只是更新操作
func (s *archGraphService) Update(ctx context.Context, req *request.UpdateArchGraphReq) error {
	id := *req.ID
	archGraph, err := s.archGraphRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	// 只是更新
	updateField := make(map[string]interface{})
	if req.GraphName != nil {
		updateField[archGraph.GraphNameField()] = *req.GraphName
	}
	if req.GroupID != nil {
		updateField[archGraph.GroupIDField()] = *req.GroupID
	}
	if req.Status != nil {
		updateField[archGraph.StatusField()] = *req.Status
	}
	if req.ImageData != nil {
		updateField[archGraph.ImageDataField()] = *req.ImageData
	}
	if req.NodeData != nil {
		updateField[archGraph.NodeDataField()] = req.NodeData
	}
	if req.EdgeData != nil {
		updateField[archGraph.EdgeDataField()] = req.EdgeData
	}
	if req.Owner != nil {
		updateField[archGraph.OwnerField()] = *req.Owner
	}
	if req.IsShared != nil {
		updateField[archGraph.IsSharedField()] = *req.IsShared
	}
	if req.GraphLabel != nil {
		if len(*req.GraphLabel) > 0 && *req.GraphLabel != archGraph.GraphLabel {
			checked := s.archGraphRepo.CheckDuplicateEntry(ctx, *req.GroupID, *req.GraphLabel)
			if checked {
				return errors.WithCodeError(ce.ErrorDBDuplicateEntry.Code(), errors.New("Duplicate Entry error"))
			}
		}
		updateField[archGraph.GraphLabelField()] = strings.Trim(*req.GraphLabel, " ")
	}
	// 编辑时候如果没有key，自动生成一个key
	if len(archGraph.GraphKey) == 0 {
		updateField[archGraph.GraphKeyField()] = util.GenUUIDv4()
	}
	err = s.archGraphRepo.Update(ctx, id, updateField)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

func (s *archGraphService) Delete(ctx context.Context, id int64) error {
	_, err := s.archGraphRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	err = s.archGraphRepo.DeleteByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

func (s *archGraphService) Copy(ctx context.Context, id int64) error {
	archGraph, err := s.archGraphRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	userName := util.GetUserName(ctx)
	newArchGraph := model.NewArchGraph()
	newArchGraph.GraphName = archGraph.GraphName + "@copy" + util.RandString(3)
	newArchGraph.GroupID = archGraph.GroupID
	newArchGraph.GraphLabel = ""
	newArchGraph.GraphKey = util.GenUUIDv4()
	newArchGraph.NodeData = archGraph.NodeData
	newArchGraph.EdgeData = archGraph.EdgeData
	newArchGraph.ImageData = archGraph.ImageData
	newArchGraph.Owner = userName
	err = s.archGraphRepo.Create(ctx, newArchGraph)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

// 默认取最新10条记录
func (s *archGraphService) GetRecordList(ctx context.Context, req *request.GetArchGraphRecordReq) (*result.ArchGraphRecordResult, error) {

	filter := &graphrepo.ArchGraphRecordFilter{
		GraphID: req.GraphID,
	}
	records, err := s.archGraphRecordRepo.GetList(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	retList := make([]*result.ArchGraphRecordRet, 0, len(records))
	for _, record := range records {
		ret := &result.ArchGraphRecordRet{}
		ret.ID = record.ID
		ret.GraphID = record.GraphID
		ret.ImageData = record.ImageData
		ret.ImageHash = util.HashBase64(record.ImageData)
		ret.CreatedTime = util.FormatTimeToString(record.CreatedTime)
		retList = append(retList, ret)
	}
	resultData := &result.ArchGraphRecordResult{}
	resultData.RetList = retList
	return resultData, nil
}

func (s *archGraphService) GetEnabledRecord(ctx context.Context, graphID int64) (*result.ArchGraphRecordRet, error) {
	archGraph, err := s.archGraphRepo.GetByID(ctx, graphID)
	if err != nil {
		return nil, errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	records, err := s.archGraphRepo.GetEnabledList(ctx, archGraph.GraphKey)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	resultData := &result.ArchGraphRecordRet{}
	if len(records) == 0 {
		return resultData, nil
	}
	record := records[0]
	resultData.GraphID = graphID
	resultData.ImageData = record.ImageData
	resultData.CreatedTime = util.FormatTimeToString(record.CreatedTime)
	return resultData, nil
}

func (s *archGraphService) SelectRecord(ctx context.Context, req *request.SelectArchGraphReq) error {
	id := *req.ID
	graphID := *req.GraphID
	archGraphRecord, err := s.archGraphRecordRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	archGraph, err := s.archGraphRepo.GetByID(ctx, graphID)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}

	updateField := make(map[string]interface{})
	updateField[archGraph.NodeDataField()] = archGraphRecord.NodeData
	updateField[archGraph.EdgeDataField()] = archGraphRecord.EdgeData
	updateField[archGraph.ImageDataField()] = archGraphRecord.ImageData
	err = s.archGraphRepo.Update(ctx, graphID, updateField)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

func (s *archGraphService) GetReviewList(ctx context.Context, req *request.GetArchGraphReviewReq) (*result.ArchGraphReviewResult, error) {

	filter := &graphrepo.ArchGraphReviewFilter{
		GraphID:  req.GraphID,
		GraphKey: req.GraphKey,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	records, total, err := s.archGraphReviewRepo.GetList(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	retList := make([]*result.ArchGraphReviewRet, 0, len(records))
	for _, record := range records {
		ret := &result.ArchGraphReviewRet{}
		ret.ID = record.ID
		ret.GraphID = record.GraphID
		ret.GraphKey = record.GraphKey
		ret.Content = record.Content
		ret.Reviewer = record.Reviewer
		ret.NotifyParty = strings.Split(record.NotifyParty, ",")
		ret.CreatedTime = util.FormatTimeToString(record.CreatedTime)
		retList = append(retList, ret)
	}
	resultData := &result.ArchGraphReviewResult{}
	resultData.Total = total
	resultData.RetList = retList
	return resultData, nil
}

func (s *archGraphService) CreateReview(ctx context.Context, req *request.CreateArchGraphReviewReq) error {

	archGraphReview := &model.ArchGraphReview{}
	userName := util.GetUserName(ctx)
	archGraphReview.GraphID = *req.GraphID
	archGraphReview.GraphKey = *req.GraphKey
	archGraphReview.Content = *req.Content
	archGraphReview.Reviewer = userName
	if len(req.NotifyParty) > 0 {
		archGraphReview.NotifyParty = strings.Join(req.NotifyParty, ",")
	}
	err := s.archGraphReviewRepo.Create(ctx, archGraphReview)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

func (s *archGraphService) DeleteReview(ctx context.Context, id int64) error {
	_, err := s.archGraphReviewRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	err = s.archGraphReviewRepo.DeleteByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

// open api
type ArchGraphNode struct {
	//架构节点ID
	ID string
	// 前一节点ID列表
	NextNodeIDs []string
	// 后一节点ID列表
	PrevNodeIDs []string
	// 前一边ID列表
	NextEdgeIDs []string
	// 后一边ID列表
	PrevEdgeIDs []string
	// 关联图ID
	GraphID int64
}

func (s *archGraphService) GetDataByLabel(ctx context.Context, req *request.GetArchGraphDataReq) (*result.ArchGraphData, error) {

	graphLabel := *req.GraphLabel
	record, err := s.archGraphRepo.GetByLabel(ctx, graphLabel)
	if err != nil {
		return nil, errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	nodeData := record.NodeData
	edgeData := record.EdgeData
	nodes := make([]*Node, 0, 0)
	err = json.Unmarshal(nodeData, &nodes)
	if err != nil {
		return nil, errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
	}
	edges := make([]*Edge, 0, 0)
	err = json.Unmarshal(edgeData, &edges)
	if err != nil {
		return nil, errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
	}
	nodeInfos := make([]*result.NodeInfo, 0, len(nodes))
	archGraphNodes := make([]*ArchGraphNode, 0, 0)
	for _, node := range nodes {
		// 在这里判断 节点只有 text 而 没有 selected
		nodeInfo := &result.NodeInfo{}
		nodeInfo.ID = node.ID
		nodeInfo.Type = node.Type
		nodeInfo.Name = node.Text.Value
		// 架构图的节点不放到节点列表
		if node.Type == "archGraph" {
			if node.Properties.LinkGraph == nil {
				continue
			}
			graphID := node.Properties.LinkGraph.ID
			archGraphNode := &ArchGraphNode{}
			archGraphNode.ID = node.ID
			archGraphNode.GraphID = graphID
			archGraphNodes = append(archGraphNodes, archGraphNode)
			continue
		}
		nodeInfos = append(nodeInfos, nodeInfo)
	}
	edgeInfos := make([]*result.EdgeInfo, 0, len(edges))
	for _, edge := range edges {
		edgeInfo := &result.EdgeInfo{}
		edgeInfo.ID = edge.ID
		edgeInfo.SourceID = edge.SourceNodeId
		edgeInfo.TargetID = edge.TargetNodeId
		// 在这里不全架构图节点的信息，并记录和架构图节点相连的边的信息
		s.genArchGraphNodes(edge, archGraphNodes)
		edgeInfos = append(edgeInfos, edgeInfo)
	}

	// 在这里遍历 架构图节点信息
	archNodeInfos := make([]*result.NodeInfo, 0, 0)
	arcEdgeInfos := make([]*result.EdgeInfo, 0, 0)
	for _, archGraphNode := range archGraphNodes {
		graphID := archGraphNode.GraphID
		archNodeInfo, archEdgeInfo, err := s.genArchGraphInfo(ctx, graphID)
		if err != nil {
			return nil, errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		}
		startNodes, endNodes := s.findStartAndEndNodes(archNodeInfo, archEdgeInfo)
		if len(startNodes) != 1 || len(endNodes) != 1 {
			return nil, errors.WithCodeError(ce.ErrorRelGraphNode.Code(), err)
		}
		archNodeInfos = append(archNodeInfos, archNodeInfo...)
		arcEdgeInfos = append(arcEdgeInfos, archEdgeInfo...)
		startNodeID := startNodes[0].ID
		endNodeID := endNodes[0].ID
		// 把链接架构图节点的边的 sourceID 换成 endNodeID， targetID 换成 startNodeID
		nextEdgeIDs := archGraphNode.NextEdgeIDs
		prevEdgeIDs := archGraphNode.PrevEdgeIDs
		// 处理边信息,进行source 或者target 节点替换
		s.dealEdges(edgeInfos, prevEdgeIDs, startNodeID, "target")
		s.dealEdges(edgeInfos, nextEdgeIDs, endNodeID, "source")
	}
	nodeInfos = append(nodeInfos, archNodeInfos...)
	edgeInfos = append(edgeInfos, arcEdgeInfos...)
	resultData := &result.ArchGraphData{}
	resultData.NodeInfos = nodeInfos
	resultData.EdgeInfos = edgeInfos
	return resultData, nil
}

func (s *archGraphService) dealEdges(edges []*result.EdgeInfo, edgeIDs []string, nodeID string, typ string) {
	for _, edgeID := range edgeIDs {
		for _, edge := range edges {
			if edge.ID == edgeID {
				if typ == "source" {
					edge.SourceID = nodeID
				}
				if typ == "target" {
					edge.TargetID = nodeID
				}
			}
		}
	}
}
func (s *archGraphService) genArchGraphInfo(ctx context.Context, graphID int64) ([]*result.NodeInfo, []*result.EdgeInfo, error) {
	record, err := s.archGraphRepo.GetByID(ctx, graphID)
	if err != nil {
		return nil, nil, err
	}
	nodeData := record.NodeData
	edgeData := record.EdgeData
	nodes := make([]*Node, 0, 0)
	err = json.Unmarshal(nodeData, &nodes)
	if err != nil {
		return nil, nil, err
	}
	edges := make([]*Edge, 0, 0)
	err = json.Unmarshal(edgeData, &edges)
	if err != nil {
		return nil, nil, err
	}
	nodeInfos := make([]*result.NodeInfo, 0, len(nodes))
	for _, node := range nodes {
		// 在这里判断 节点只有 text 而 没有 selected
		nodeInfo := &result.NodeInfo{}
		nodeInfo.ID = node.ID
		nodeInfo.Type = node.Type
		nodeInfo.Name = node.Text.Value
		nodeInfos = append(nodeInfos, nodeInfo)
	}
	edgeInfos := make([]*result.EdgeInfo, 0, len(edges))
	for _, edge := range edges {
		edgeInfo := &result.EdgeInfo{}
		edgeInfo.ID = edge.ID
		edgeInfo.SourceID = edge.SourceNodeId
		edgeInfo.TargetID = edge.TargetNodeId
		edgeInfos = append(edgeInfos, edgeInfo)
	}
	return nodeInfos, edgeInfos, nil
}

func (s *archGraphService) genArchGraphNodes(edge *Edge, archGraphNodes []*ArchGraphNode) {
	edgeID := edge.ID
	sourceID := edge.SourceNodeId
	targetID := edge.TargetNodeId
	for _, archGraphNode := range archGraphNodes {
		if archGraphNode.ID == sourceID {
			nextNodeIDs := archGraphNode.NextNodeIDs
			nextNodeIDs = append(nextNodeIDs, targetID)
			nextEdgeIDs := archGraphNode.NextEdgeIDs
			nextEdgeIDs = append(nextEdgeIDs, edgeID)
			archGraphNode.NextNodeIDs = nextNodeIDs
			archGraphNode.NextEdgeIDs = nextEdgeIDs
		}
		if archGraphNode.ID == targetID {
			prevNodeIDs := archGraphNode.PrevNodeIDs
			prevNodeIDs = append(prevNodeIDs, sourceID)
			prevEdgeIDs := archGraphNode.PrevEdgeIDs
			prevEdgeIDs = append(prevEdgeIDs, edgeID)
			archGraphNode.PrevNodeIDs = prevNodeIDs
			archGraphNode.PrevEdgeIDs = prevEdgeIDs
		}
	}
}

func (s *archGraphService) findStartAndEndNodes(nodes []*result.NodeInfo, edges []*result.EdgeInfo) ([]*result.NodeInfo, []*result.NodeInfo) {
	startNodes := make([]*result.NodeInfo, 0, 0)
	endNodes := make([]*result.NodeInfo, 0, 0)
	nodeMap := make(map[string]*result.NodeInfo)
	for _, node := range nodes {
		nodeMap[node.ID] = node
	}

	// 记录所有 source 和 target
	sources := make(map[string]bool)
	targets := make(map[string]bool)

	for _, edge := range edges {
		sources[edge.SourceID] = true
		targets[edge.TargetID] = true
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
	return startNodes, endNodes
}
