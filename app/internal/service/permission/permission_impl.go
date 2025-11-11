package permission

import (
	"context"
	"github.com/samber/lo"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	permrepo "github.com/yanggelinux/cattle/internal/repository/permission"
	rolerepo "github.com/yanggelinux/cattle/internal/repository/role"
	"github.com/yanggelinux/cattle/internal/store/model"
	"github.com/yanggelinux/cattle/pkg/util"
)

type permissionService struct {
	db             *model.MDB
	permissionRepo permrepo.PermissionRepo
	roleRepo       rolerepo.RoleRepo
}

func NewPermissionService() PermissionService {
	return &permissionService{db: model.GetDB(), permissionRepo: permrepo.NewPermissionRepo(), roleRepo: rolerepo.NewRoleRepo()}
}

func (s *permissionService) GetList(ctx context.Context, req *request.GetPermissionReq) (*result.PermissionResult, error) {

	filter := &permrepo.PermissionFilter{
		Name:     req.Name,
		Code:     req.Code,
		Project:  req.Project,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	records, total, err := s.permissionRepo.GetList(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	retList := make([]*result.PermissionRet, 0, len(records))
	for _, record := range records {
		ret := &result.PermissionRet{}
		ret.ID = record.ID
		ret.ParentID = record.ParentID
		ret.Name = record.Name
		ret.Code = record.Code
		ret.Uri = record.Uri
		ret.Method = record.Method
		ret.Project = record.Project
		ret.PermType = record.PermType
		ret.IsEnabled = record.IsEnabled
		ret.Sort = record.Sort
		ret.UpdatedTime = util.FormatTimeToString(record.UpdatedTime)
		ret.CreatedTime = util.FormatTimeToString(record.CreatedTime)
		retList = append(retList, ret)
	}
	resultData := &result.PermissionResult{}
	resultData.Total = total
	resultData.RetList = s.genTree2(retList)
	return resultData, nil
}

func (s *permissionService) Create(ctx context.Context, req *request.CreatePermissionReq) error {
	permission := model.NewPermission()
	permission.ParentID = *req.ParentID
	permission.Name = *req.Name
	permission.Code = *req.Code
	permission.Uri = *req.Uri
	permission.Method = *req.Method
	permission.Project = *req.Project
	permission.PermType = *req.PermType
	permission.IsEnabled = *req.IsEnabled
	permission.Sort = *req.Sort
	checked := s.permissionRepo.CheckDuplicateEntry(ctx, *req.Code, *req.Project)
	if checked {
		return errors.WithCodeError(ce.ErrorDBDuplicateEntry.Code(), errors.New("Duplicate Entry error"))
	}
	err := s.permissionRepo.Create(ctx, permission)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return err
}

func (s *permissionService) Update(ctx context.Context, req *request.UpdatePermissionReq) error {
	id := *req.ID
	permission, err := s.permissionRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	updateField := make(map[string]interface{})
	if req.ParentID != nil {
		updateField[permission.ParentIDField()] = *req.ParentID
	}
	if req.Name != nil {
		updateField[permission.NameField()] = *req.Name
	}
	if req.Code != nil {
		updateField[permission.CodeField()] = *req.Code
	}
	if req.Uri != nil {
		updateField[permission.UriField()] = *req.Uri
	}
	if req.Method != nil {
		updateField[permission.MethodField()] = *req.Method
	}
	if req.Project != nil {
		updateField[permission.ProjectField()] = *req.Project
	}
	if req.PermType != nil {
		updateField[permission.PermTypeField()] = *req.PermType
	}
	if req.IsEnabled != nil {
		updateField[permission.IsEnabledField()] = *req.IsEnabled
	}
	if req.Sort != nil {
		updateField[permission.SortField()] = *req.Sort
	}
	err = s.permissionRepo.Update(ctx, id, updateField)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

func (s *permissionService) Delete(ctx context.Context, id int64) error {
	_, err := s.permissionRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	err = s.permissionRepo.DeleteByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

func (s *permissionService) GetRolePermList(ctx context.Context, req *request.GetRolePermReq) (*result.PermTreeData, error) {
	roleID := *req.RoleID
	project := *req.Project
	rolePermRecords, err := s.permissionRepo.GetPermsByRole(ctx, roleID)
	if err != nil {
		return nil, err
	}

	allPermRecords, err := s.permissionRepo.GetPermsByProject(ctx, project)
	if err != nil {
		return nil, err
	}

	//如果是超级管理员给所有权限
	if *req.IsSuper == 1 {
		rolePermRecords = allPermRecords
	}

	rolePermIDList := make([]int64, 0, len(rolePermRecords))
	for _, rolePermRecord := range rolePermRecords {
		if rolePermRecord.Project != project {
			continue
		}
		rolePermIDList = append(rolePermIDList, rolePermRecord.ID)
	}

	allPermTreeData := s.genTree(allPermRecords)
	permTreeData := &result.PermTreeData{}
	permTreeData.RolePermIDList = rolePermIDList
	permTreeData.AllPermTreeData = allPermTreeData
	return permTreeData, nil
}

func (s *permissionService) UpdateRolePerm(ctx context.Context, req *request.UpdateRolePermReq) error {
	roleID := *req.RoleID
	permIDList := req.PermIDList
	project := *req.Project
	rolePermRecords, err := s.permissionRepo.GetPermsByRole(ctx, roleID)
	if err != nil {
		return err
	}
	rolePermIDList := make([]int64, 0, len(rolePermRecords))
	for _, rolePermRecord := range rolePermRecords {
		if rolePermRecord.Project != project {
			continue
		}
		rolePermIDList = append(rolePermIDList, rolePermRecord.ID)
	}
	createList := make([]*model.RolePermRel, 0, 0)
	deleteIDList := make([]int64, 0, 0)
	// 进行id的diff操作
	for _, permID := range permIDList {
		// 新的id 在旧的列表不存在，创建这个权限
		if !lo.Contains(rolePermIDList, permID) {
			createList = append(createList, &model.RolePermRel{
				RoleID: roleID,
				PermID: permID,
			})
		}
	}
	for _, permID := range rolePermIDList {
		// 旧的id 在新的列表不存在，删除这个权限
		if !lo.Contains(permIDList, permID) {
			deleteIDList = append(deleteIDList, permID)
		}
	}
	// 进行批量操作,事务操作
	tx := s.db.WithContext(ctx).Begin()
	defer model.RecoverRollback(tx)
	roleTx := s.roleRepo.WithTx(tx)
	// 批量创建角色权限
	if len(createList) > 0 {
		err = roleTx.CreatePermRelInBatches(ctx, createList)
		if err != nil {
			tx.Rollback()
			return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
		}
	}
	// 批量删除角色权限
	if len(deleteIDList) > 0 {
		err = roleTx.DeletePermRelByIDs(ctx, roleID, deleteIDList)
		if err != nil {
			tx.Rollback()
			return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
		}
	}
	//提交
	tx.Commit()

	return nil
}

// todo 使用泛型优化合并俩代码
func (s *permissionService) genTree(records []*model.Permission) []*result.PermTreeNode {
	idMap := make(map[int64]*result.PermTreeNode, len(records)+1)
	topNode := &result.PermTreeNode{Permission: &model.Permission{ID: 0, ParentID: 0, Name: "root"}}
	idMap[topNode.ID] = topNode

	for _, record := range records {
		// 先判断 当前ID 对应的节点是否存在
		nNode, ok := idMap[record.ID]
		if ok {
			// 存在就补录节点数据
			nNode.Permission = record
		} else {
			// 不存在就新建一个节点，并放入map
			nNode = &result.PermTreeNode{Permission: record}
			idMap[nNode.ID] = nNode
		}
		// 根据parentId 获取 parentNode
		pNode, ok := idMap[nNode.ParentID]
		if !ok { // 获取不到，说明父节点还没有，先进行添加
			pNode = &result.PermTreeNode{}
			idMap[nNode.ParentID] = pNode
		}

		// 将当前节点添加到父节点的 array 中
		pNode.Children = append(pNode.Children, nNode)
	}
	return topNode.Children
}
func (s *permissionService) genTree2(records []*result.PermissionRet) []*result.ResTreeNode {
	idMap := make(map[int64]*result.ResTreeNode, len(records)+1)
	topNode := &result.ResTreeNode{PermissionRet: &result.PermissionRet{ID: 0, ParentID: 0, Name: "root"}}
	idMap[topNode.ID] = topNode

	for _, record := range records {
		// 先判断 当前ID 对应的节点是否存在
		nNode, ok := idMap[record.ID]
		if ok {
			// 存在就补录节点数据
			nNode.PermissionRet = record
		} else {
			// 不存在就新建一个节点，并放入map
			nNode = &result.ResTreeNode{PermissionRet: record}
			idMap[nNode.ID] = nNode
		}
		// 根据parentId 获取 parentNode
		pNode, ok := idMap[nNode.ParentID]
		if !ok { // 获取不到，说明父节点还没有，先进行添加
			pNode = &result.ResTreeNode{}
			idMap[nNode.ParentID] = pNode
		}

		// 将当前节点添加到父节点的 array 中
		pNode.Children = append(pNode.Children, nNode)
	}
	return topNode.Children
}
