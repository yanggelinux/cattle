package role

import (
	"context"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	rolerepo "github.com/yanggelinux/cattle/internal/repository/role"
	"github.com/yanggelinux/cattle/internal/store/model"
	"github.com/yanggelinux/cattle/pkg/util"
)

type roleService struct {
	roleRepo rolerepo.RoleRepo
}

func NewRoleService() RoleService {
	return &roleService{roleRepo: rolerepo.NewRoleRepo()}
}

func (s *roleService) GetList(ctx context.Context, req *request.GetRoleReq) (*result.RoleResult, error) {

	//进行一个copy
	filter := &rolerepo.RoleFilter{
		RoleName: req.RoleName,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	records, total, err := s.roleRepo.GetList(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	retList := make([]*result.RoleRet, 0, len(records))
	for _, record := range records {
		ret := &result.RoleRet{}
		ret.ID = record.ID
		ret.RoleName = record.RoleName
		ret.DisplayName = record.DisplayName
		ret.IsSuper = record.IsSuper
		ret.UpdatedTime = util.FormatTimeToString(record.UpdatedTime)
		ret.CreatedTime = util.FormatTimeToString(record.CreatedTime)
		retList = append(retList, ret)
	}
	resultData := &result.RoleResult{}
	resultData.Total = total
	resultData.RetList = retList
	return resultData, nil
}

func (s *roleService) Create(ctx context.Context, req *request.CreateRoleReq) error {
	role := model.NewRole()
	role.RoleName = *req.RoleName
	role.DisplayName = *req.DisplayName
	role.IsSuper = *req.IsSuper
	checked := s.roleRepo.CheckDuplicateEntry(ctx, *req.RoleName)
	if checked {
		return errors.WithCodeError(ce.ErrorDBDuplicateEntry.Code(), errors.New("Duplicate Entry error"))
	}
	err := s.roleRepo.Create(ctx, role)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return err
}

func (s *roleService) Update(ctx context.Context, req *request.UpdateRoleReq) error {
	id := *req.ID
	role, err := s.roleRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	updateField := make(map[string]interface{})

	if req.RoleName != nil {
		updateField[role.RoleNameField()] = *req.RoleName
	}
	if req.DisplayName != nil {
		updateField[role.DisplayNameField()] = *req.DisplayName
	}
	if req.IsSuper != nil {
		updateField[role.IsSuperField()] = *req.IsSuper
	}
	err = s.roleRepo.Update(ctx, id, updateField)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

func (s *roleService) Delete(ctx context.Context, id int64) error {
	_, err := s.roleRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	err = s.roleRepo.DeleteByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}
