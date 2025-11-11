package archgroup

import (
	"context"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	"github.com/yanggelinux/cattle/internal/repository/archgraph"
	grouprepo "github.com/yanggelinux/cattle/internal/repository/archgroup"
	"github.com/yanggelinux/cattle/internal/store/model"
	"github.com/yanggelinux/cattle/pkg/util"
)

type archGroupService struct {
	archGroupRepo grouprepo.ArchGroupRepo
	archGraphRepo archgraph.ArchGraphRepo
}

func NewArchGroupService() ArchGroupService {
	return &archGroupService{
		archGroupRepo: grouprepo.NewArchGroupRepo(),
		archGraphRepo: archgraph.NewArchGraphRepo(),
	}
}

func (s *archGroupService) GetList(ctx context.Context, req *request.GetArchGroupReq) (*result.ArchGroupResult, error) {

	filter := &grouprepo.ArchGroupFilter{
		ParentID:  req.ParentID,
		GroupName: req.GroupName,
	}
	records, err := s.archGroupRepo.GetList(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	allRecords, err := s.archGroupRepo.GetAll(ctx)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	selectFiled := "`id`,`group_id`"
	graphs, err := s.archGraphRepo.GetSomeAll(ctx, selectFiled)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	idMap := s.buildDescendantMap(allRecords)
	retList := make([]*result.ArchGroupRet, 0, len(records))
	for _, record := range records {
		ret := &result.ArchGroupRet{}
		itemCount := s.getGraphCount(record.ID, idMap, graphs)
		ret.ID = record.ID
		ret.ParentID = record.ParentID
		ret.GroupName = record.GroupName
		ret.ItemCount = itemCount
		ret.UpdatedTime = util.FormatTimeToString(record.UpdatedTime)
		ret.CreatedTime = util.FormatTimeToString(record.CreatedTime)
		retList = append(retList, ret)
	}
	resultData := &result.ArchGroupResult{}
	resultData.RetList = retList
	return resultData, nil
}

func (s *archGroupService) getGraphCount(id int64, idMap map[int64][]int64, graphs []*model.ArchGraph) int64 {
	var count int64
	count = s.getCountByGroup(id, graphs)
	ids, ok := idMap[id]
	if !ok {
		return count
	}
	for _, _id := range ids {
		cnt := s.getCountByGroup(_id, graphs)
		count = count + cnt
	}
	return count
}
func (s *archGroupService) getCountByGroup(groupID int64, graphs []*model.ArchGraph) int64 {
	var count int64
	for _, graph := range graphs {
		if graph.GroupID == groupID {
			count++
		}
	}
	return count
}

func (s *archGroupService) buildDescendantMap(flatList []*model.ArchGroup) map[int64][]int64 {
	idMap := make(map[int64][]int64)    // 最终结果
	childMap := make(map[int64][]int64) // 父 -> 直接子节点

	// 构建父子映射
	for _, node := range flatList {
		if node.ParentID != 0 {
			childMap[node.ParentID] = append(childMap[node.ParentID], node.ID)
		}
	}

	// 递归获取所有子节点
	var getAllChildren func(parentID int64) []int64
	getAllChildren = func(parentID int64) []int64 {
		children := childMap[parentID]
		if len(children) == 0 {
			return nil
		}

		all := make([]int64, 0, len(children))
		for _, c := range children {
			all = append(all, c)
			all = append(all, getAllChildren(c)...) // 递归收集
		}

		// 只有有后代时才加到结果
		if len(all) > 0 {
			idMap[parentID] = all
		}

		return all
	}
	// 遍历所有节点
	for _, node := range flatList {
		getAllChildren(node.ID)
	}
	return idMap
}

func (s *archGroupService) Create(ctx context.Context, req *request.CreateArchGroupReq) error {
	archGroup := model.NewArchGroup()
	archGroup.GroupName = *req.GroupName
	archGroup.ParentID = *req.ParentID
	checked := s.archGroupRepo.CheckDuplicateEntry(ctx, *req.ParentID, *req.GroupName)
	if checked {
		return errors.WithCodeError(ce.ErrorDBDuplicateEntry.Code(), errors.New("Duplicate Entry error"))
	}
	err := s.archGroupRepo.Create(ctx, archGroup)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return err
}

func (s *archGroupService) Update(ctx context.Context, req *request.UpdateArchGroupReq) error {
	id := *req.ID
	archGroup, err := s.archGroupRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	updateField := make(map[string]interface{})
	if req.GroupName != nil {
		updateField[archGroup.GroupNameField()] = *req.GroupName
	}
	if req.ParentID != nil {
		updateField[archGroup.ParentIDField()] = *req.ParentID
	}
	err = s.archGroupRepo.Update(ctx, id, updateField)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

func (s *archGroupService) Delete(ctx context.Context, id int64) error {
	// 删除组前先检查是否有图
	checked := s.archGraphRepo.CheckHas(ctx, id)
	if checked {
		return errors.WithCodeError(ce.ErrorGroupHasGraphRecord.Code(), errors.New("group has graph record error"))
	}
	checked2 := s.archGroupRepo.CheckHas(ctx, id)
	if checked2 {
		return errors.WithCodeError(ce.ErrorGroupHasRecord.Code(), errors.New("group has group record error"))
	}
	_, err := s.archGroupRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	err = s.archGroupRepo.DeleteByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}
