package user

import (
	"context"
	"github.com/samber/lo"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	rolerepo "github.com/yanggelinux/cattle/internal/repository/role"
	userrepo "github.com/yanggelinux/cattle/internal/repository/user"
	"github.com/yanggelinux/cattle/internal/store/model"
	"github.com/yanggelinux/cattle/pkg/util"
	"gorm.io/gorm"
	"strings"
	"time"
)

type userService struct {
	db       *model.MDB
	userRepo userrepo.UserRepo
	roleRepo rolerepo.RoleRepo
}

func NewUserService() UserService {
	return &userService{
		db:       model.GetDB(),
		userRepo: userrepo.NewUserRepo(),
		roleRepo: rolerepo.NewRoleRepo(),
	}
}

func (s *userService) GetList(ctx context.Context, req *request.GetUserReq) (*result.UserResult, error) {

	filter := &userrepo.UserFilter{
		UserName: req.UserName,
		Email:    req.Email,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	records, total, err := s.userRepo.GetList(ctx, filter)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	// 获取全部的映射关系
	roleNamesMaping, roleIDsMapping, err := s.getUserRoles(ctx)
	if err != nil {
		err = errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
		return nil, err
	}
	retList := make([]*result.UserRet, 0, len(records))
	for _, record := range records {
		ret := &result.UserRet{}
		roleNames := roleNamesMaping[record.ID]
		roleIDs := roleIDsMapping[record.ID]
		ret.ID = record.ID
		ret.UserName = record.UserName
		ret.Password = record.Password
		ret.Email = record.Email
		ret.DisplayName = record.DisplayName
		ret.DeptName = record.DeptName
		ret.RoleNames = strings.Join(roleNames, ",")
		ret.RoleIDs = roleIDs
		ret.Origin = record.Origin
		ret.LastLoginTime = util.FormatTimeToString(record.LastLoginTime)
		ret.UpdatedTime = util.FormatTimeToString(record.UpdatedTime)
		ret.CreatedTime = util.FormatTimeToString(record.CreatedTime)
		retList = append(retList, ret)
	}
	resultData := &result.UserResult{}
	resultData.Total = total
	resultData.RetList = retList
	return resultData, nil
}

func (s *userService) getUserRoles(ctx context.Context) (map[int64][]string, map[int64][]int64, error) {
	roleNamesMaping := make(map[int64][]string)
	roleIDsMapping := make(map[int64][]int64)
	records, err := s.userRepo.GetUserRoleRels(ctx)
	if err != nil {
		return roleNamesMaping, roleIDsMapping, err
	}
	for _, record := range records {
		userID := record.UserID
		if _, ok := roleNamesMaping[userID]; !ok {
			roleNamesMaping[userID] = []string{record.DisplayName}
			roleIDsMapping[userID] = []int64{record.RoleID}
		} else {
			roleNamesMaping[userID] = append(roleNamesMaping[userID], record.DisplayName)
			roleIDsMapping[userID] = append(roleIDsMapping[userID], record.RoleID)
		}
	}
	return roleNamesMaping, roleIDsMapping, nil
}

func (s *userService) Create(ctx context.Context, req *request.CreateUserReq) error {
	// 事务
	tx := s.db.WithContext(ctx).Begin()
	defer model.RecoverRollback(tx)

	userTx := s.userRepo.WithTx(tx)
	user := model.NewUser()
	user.UserName = *req.UserName
	user.Password = *req.Password
	user.DisplayName = *req.DisplayName
	user.Email = *req.Email
	user.DeptName = *req.DeptName
	user.Origin = 1
	user.LastLoginTime = time.Now()
	checked := userTx.CheckDuplicateEntry(ctx, *req.UserName)
	if checked {
		return errors.WithCodeError(ce.ErrorDBDuplicateEntry.Code(), errors.New("Duplicate Entry error"))
	}
	err := userTx.Create(ctx, user)
	if err != nil {
		tx.Rollback()
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}

	// 在这里创建用户角色关系
	roleIDs := req.RoleIDs
	// 不创建角色关联
	if len(roleIDs) == 0 {
		return nil
	}
	userRoleRels := make([]*model.UserRoleRel, 0, 0)
	for _, roleID := range roleIDs {
		userRoleRel := &model.UserRoleRel{
			RoleID: roleID,
			UserID: user.ID,
		}
		userRoleRels = append(userRoleRels, userRoleRel)
	}
	if len(userRoleRels) > 0 {
		err = s.userRepo.CreateRoleRelInBatches(ctx, userRoleRels)
		if err != nil {
			tx.Rollback()
			return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
		}

	}
	tx.Commit()
	return nil
}

func (s *userService) Update(ctx context.Context, req *request.UpdateUserReq) error {
	id := *req.ID
	tx := s.db.WithContext(ctx).Begin()
	defer model.RecoverRollback(tx)
	userTx := s.userRepo.WithTx(tx)

	user, err := userTx.GetByID(ctx, id)
	if err != nil {
		tx.Rollback()
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	updateField := make(map[string]interface{})
	if req.UserName != nil {
		updateField[user.UserNameField()] = *req.UserName
	}
	if *req.Origin == 1 && req.Password != nil {
		updateField[user.PasswordField()] = *req.Password
	}
	if req.DisplayName != nil {
		updateField[user.DisplayNameField()] = *req.DisplayName
	}
	if req.Email != nil {
		updateField[user.EmailField()] = *req.Email
	}
	if req.DeptName != nil {
		updateField[user.DeptNameField()] = *req.DeptName
	}
	if req.Origin != nil {
		updateField[user.OriginField()] = *req.Origin
	}
	err = userTx.Update(ctx, id, updateField)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	// 在这里更新用户角色关系
	records, err := userTx.GetRoleRels(ctx, user.ID)
	if err != nil {
		tx.Rollback()
		return errors.WithCodeError(ce.ErrorDBQueryFailed.Code(), err)
	}
	// 注意这里为空的情况,不往下走,至少要有一个角色
	if len(req.RoleIDs) == 0 {
		tx.Commit()
		return nil
	}
	roleIDs := req.RoleIDs
	oldRoleIDs := make([]int64, 0, len(records))
	for _, record := range records {
		oldRoleIDs = append(oldRoleIDs, record.RoleID)
	}
	createList := make([]*model.UserRoleRel, 0, 0)
	deleteIDList := make([]int64, 0, 0)

	for _, roleID := range roleIDs {
		if !lo.Contains(oldRoleIDs, roleID) {
			createList = append(createList, &model.UserRoleRel{
				RoleID: roleID,
				UserID: id,
			})
		}
	}
	for _, roleID := range oldRoleIDs {
		if !lo.Contains(roleIDs, roleID) {
			deleteIDList = append(deleteIDList, roleID)
		}
	}
	if len(createList) > 0 {
		err = userTx.CreateRoleRelInBatches(ctx, createList)
		if err != nil {
			tx.Rollback()
			return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
		}
	}
	// 批量删除用户角色
	if len(deleteIDList) > 0 {
		err = userTx.DeleteRoleRelByIDs(ctx, id, deleteIDList)
		if err != nil {
			tx.Rollback()
			return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
		}
	}
	tx.Commit()
	return nil
}

func (s *userService) Delete(ctx context.Context, id int64) error {
	_, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBQueryNotFound.Code(), err)
	}
	err = s.userRepo.DeleteByID(ctx, id)
	if err != nil {
		return errors.WithCodeError(ce.ErrorDBOperateFailed.Code(), err)
	}
	return nil
}

// 创建角色名为user的普通用户
func (s *userService) CreateUserRole(ctx context.Context, userID int64, roleName string) error {
	roleM, err := s.roleRepo.GetByName(ctx, roleName)
	if err != nil {
		return err
	}
	roleID := roleM.ID
	userRoleRel, err := s.userRepo.GetRoleRelByID(ctx, userID, roleID)
	// 不存在创建
	if errors.Is(err, gorm.ErrRecordNotFound) {
		userRoleRel.UserID = userID
		userRoleRel.RoleID = roleID
		err := s.userRepo.CreateRoleRel(ctx, userRoleRel)
		if err != nil {
			return err
		}
	}
	return nil
}
