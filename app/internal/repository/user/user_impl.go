package user

import (
	"context"
	"fmt"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/store/model"
	"time"
)

type UserFilter struct {
	UserName *string
	Email    *string
	Page     *int
	PageSize *int
}

type userRepo struct {
	db *model.MDB
}

func NewUserRepo() UserRepo {
	return &userRepo{db: model.GetDB()}
}

func (r *userRepo) WithTx(tx *model.MDB) UserRepo {
	return &userRepo{db: tx}
}

func (r *userRepo) GetList(ctx context.Context, filter *UserFilter) ([]*model.User, int64, error) {
	var total int64
	records := make([]*model.User, 0, 0)
	tx := r.db.WithContext(ctx).Model(model.NewUser())
	tx.Where("is_deleted = ?", 0)
	if filter.UserName != nil && len(*filter.UserName) > 0 {
		tx.Where("user_name like ?", fmt.Sprintf("%%%s%%", *filter.UserName))
	}
	if filter.Email != nil && len(*filter.Email) > 0 {
		tx.Where("email like ?", fmt.Sprintf("%%%s%%", *filter.Email))
	}
	err := tx.Count(&total).Error
	if err != nil {
		return records, 0, err
	}
	if total == 0 {
		return records, 0, nil
	}
	tx.Order("last_login_time desc")
	tx.Select(
		"id,user_name,password,email,display_name,dept_name,origin,last_login_time," +
			"updated_time,created_time")
	if filter.Page != nil && filter.PageSize != nil {
		tx.Offset((*filter.Page - 1) * *filter.PageSize).Limit(*filter.PageSize)
	}
	err = tx.Find(&records).Error
	if err != nil {
		return records, 0, err
	}
	return records, total, nil
}

func (r *userRepo) GetUsers(ctx context.Context) ([]*model.User, error) {
	records := make([]*model.User, 0, 0)
	tx := r.db.WithContext(ctx).Model(model.NewUser())
	tx.Where("is_deleted = ?", 0)
	err := tx.Select(
		"id,user_name,password,email,display_name,dept_name,origin,last_login_time," +
			"updated_time,created_time").Find(&records).Error
	if err != nil {
		return records, err
	}
	return records, nil
}

func (r *userRepo) GetUsersByRole(ctx context.Context, roleName string) ([]*model.User, error) {
	records := make([]*model.User, 0, 0)
	tx := r.db.WithContext(ctx).Model(model.NewUser())
	err := tx.Select(
		"user.id,user.user_name,user.password,user.email,user.display_name,user.dept_name,user.origin,user.last_login_time,"+
			"user.updated_time,user.created_time").
		Joins("left join user_role_rel on user.id = user_role_rel.user_id").
		Joins("left join role on user_role_rel.role_id = role.id").
		Where("role.role_name = ? and user.is_deleted = 0", roleName).
		Find(&records).Error
	if err != nil {
		return records, err
	}
	return records, nil
}
func (r *userRepo) GetUsersByRoles(ctx context.Context, roleNames []string) ([]*model.User, error) {
	records := make([]*model.User, 0, 0)
	tx := r.db.WithContext(ctx).Model(model.NewUser())
	err := tx.Select(
		"user.id,user.user_name,user.password,user.email,user.display_name,user.dept_name,user.origin,user.last_login_time,"+
			"user.updated_time,user.created_time").
		Joins("left join user_role_rel on user.id = user_role_rel.user_id").
		Joins("left join role on user_role_rel.role_id = role.id").
		Where("role.role_name in ? and user.is_deleted = 0", roleNames).
		Find(&records).Error
	if err != nil {
		return records, err
	}
	return records, nil
}

func (r *userRepo) GetRoleRels(ctx context.Context, userID int64) ([]*model.UserRoleRel, error) {
	records := make([]*model.UserRoleRel, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewUserRoleRel()).Where("user_id = ?", userID).Find(&records).Error
	return records, err
}

func (r *userRepo) GetUserRoleRels(ctx context.Context) ([]*result.UserRoleRet, error) {
	records := make([]*result.UserRoleRet, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewRole()).
		Select("role.id as role_id,role.role_name,role.display_name,role.is_super,user_role_rel.user_id as user_id").
		Joins("inner join user_role_rel on role.id = user_role_rel.role_id").
		Where("role.is_deleted = 0").Find(&records).Error
	return records, err
}

func (r *userRepo) GetByID(ctx context.Context, id int64) (*model.User, error) {
	user := model.NewUser()
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *userRepo) GetByName(ctx context.Context, userName string) (*model.User, error) {
	user := model.NewUser()
	err := r.db.WithContext(ctx).Where("user_name = ? and is_deleted = 0", userName).First(&user).Error
	return user, err
}
func (r *userRepo) GetBySysName(ctx context.Context, userName string) (*model.User, error) {
	user := model.NewUser()
	err := r.db.WithContext(ctx).Where("user_name = ? and is_deleted = 0 and origin = 1", userName).First(&user).Error
	return user, err
}

func (r *userRepo) GetByLdapName(ctx context.Context, userName string) (*model.User, error) {
	user := model.NewUser()
	err := r.db.WithContext(ctx).Where("user_name = ? and is_deleted = 0 and origin = 2", userName).First(&user).Error
	return user, err
}

func (r *userRepo) CheckDuplicateEntry(ctx context.Context, userName string) bool {
	user := model.NewUser()
	err := r.db.WithContext(ctx).Where("user_name = ? and is_deleted = 0", userName).First(&user).Error
	if err == nil {
		return true
	}
	return false
}

// 创建
func (r *userRepo) Create(ctx context.Context, user *model.User) error {
	err := r.db.WithContext(ctx).Create(user).Error
	return err
}

// 更新, 可以更新多列指定字段
func (r *userRepo) Update(ctx context.Context, id int64, updateField map[string]interface{}) error {
	err := r.db.WithContext(ctx).Model(model.NewUser()).Where("id = ?", id).Updates(updateField).Error
	return err
}

// 更新 保存 支持非 0字段 主键存在更新所有字段，主键不存在插入
func (r *userRepo) CreateOrUpdate(ctx context.Context, user *model.User) error {
	err := r.db.WithContext(ctx).Save(user).Error
	return err
}

// 逻辑删除
func (r *userRepo) DeleteByID(ctx context.Context, id int64) error {
	nowTime := time.Now()
	user := model.NewUser()
	user.IsDeleted = 1
	user.DeletedTime = nowTime
	err := r.db.WithContext(ctx).Where("id = ?", id).Updates(user).Error
	return err
}

// 关联表相关的操作
func (r *userRepo) GetRoleRelByID(ctx context.Context, userID, roleID int64) (*model.UserRoleRel, error) {
	userRoleRel := model.NewUserRoleRel()
	err := r.db.WithContext(ctx).Where("user_id = ? and role_id = ?", userID, roleID).First(&userRoleRel).Error
	return userRoleRel, err
}

func (r *userRepo) CreateRoleRel(ctx context.Context, userRoleRel *model.UserRoleRel) error {
	err := r.db.WithContext(ctx).Create(userRoleRel).Error
	return err
}

// 创建
func (r *userRepo) CreateRoleRelInBatches(ctx context.Context, userRoleRels []*model.UserRoleRel) error {
	err := r.db.WithContext(ctx).CreateInBatches(userRoleRels, 100).Error
	return err
}

// 物理删除
func (r *userRepo) DeleteRoleRelByIDs(ctx context.Context, userID int64, deleteRoleIDs []int64) error {
	err := r.db.WithContext(ctx).Where("user_id = ? and role_id in ?", userID, deleteRoleIDs).Delete(model.NewUserRoleRel()).Error
	return err
}
