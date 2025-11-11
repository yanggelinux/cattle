package role

import (
	"context"
	"fmt"
	"github.com/yanggelinux/cattle/internal/store/model"
	"time"
)

type RoleFilter struct {
	RoleName *string
	Page     *int
	PageSize *int
}

type roleRepo struct {
	db *model.MDB
}

func NewRoleRepo() RoleRepo {
	return &roleRepo{
		db: model.GetDB(),
	}
}

func (r *roleRepo) WithTx(tx *model.MDB) RoleRepo {
	return &roleRepo{db: tx}
}

func (r *roleRepo) GetList(ctx context.Context, filter *RoleFilter) ([]*model.Role, int64, error) {
	var total int64
	records := make([]*model.Role, 0, 0)
	tx := r.db.WithContext(ctx).Model(model.NewRole())
	tx.Where("is_deleted = ?", 0)
	if filter.RoleName != nil && len(*filter.RoleName) > 0 {
		tx.Where("role_name like ?", fmt.Sprintf("%%%s%%", *filter.RoleName))
	}
	err := tx.Count(&total).Error
	if err != nil {
		return records, 0, err
	}
	if total == 0 {
		return records, 0, nil
	}
	tx.Order("created_time desc")
	tx.Select(
		"id,role_name,display_name,is_super,updated_time,created_time")
	if filter.Page != nil && filter.PageSize != nil {
		tx.Offset((*filter.Page - 1) * (*filter.PageSize)).Limit(*filter.PageSize)
	}
	err = tx.Find(&records).Error
	if err != nil {
		return records, 0, err
	}
	return records, total, nil
}

// 获取角色相关信息
func (r *roleRepo) GetRoles(ctx context.Context, userID int64) ([]*model.Role, error) {
	records := make([]*model.Role, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewRole()).
		Select("role.id,role.role_name,role.display_name,role.is_super").
		Joins("left join user_role_rel on role.id = user_role_rel.role_id").
		Where("user_role_rel.user_id = ? and role.is_deleted = 0", userID).Find(&records).Error
	if err != nil {
		return records, err
	}
	return records, nil
}

func (r *roleRepo) GetByID(ctx context.Context, id int64) (*model.Role, error) {
	role := model.NewRole()
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&role).Error
	return role, err
}

func (r *roleRepo) GetByName(ctx context.Context, roleName string) (*model.Role, error) {
	role := model.NewRole()
	err := r.db.WithContext(ctx).Where("role_name = ? and is_deleted = 0", roleName).First(&role).Error
	return role, err
}

func (r *roleRepo) CheckDuplicateEntry(ctx context.Context, roleName string) bool {
	role := model.NewRole()
	err := r.db.WithContext(ctx).Where("role_name = ? and is_deleted = 0", roleName).First(&role).Error
	if err == nil {
		return true
	}
	return false
}

// 创建
func (r *roleRepo) Create(ctx context.Context, role *model.Role) error {
	err := r.db.WithContext(ctx).Create(role).Error
	return err
}

// 更新, 可以更新多列指定字段
func (r *roleRepo) Update(ctx context.Context, id int64, updateField map[string]interface{}) error {
	err := r.db.WithContext(ctx).Model(model.NewRole()).Where("id = ?", id).Updates(updateField).Error
	return err
}

// 更新 保存 支持非 0字段 主键存在更新所有字段，主键不存在插入
func (r *roleRepo) CreateOrUpdate(ctx context.Context, role *model.Role) error {
	err := r.db.WithContext(ctx).Save(role).Error
	return err
}

// 逻辑删除
func (r *roleRepo) DeleteByID(ctx context.Context, id int64) error {
	nowTime := time.Now()
	role := model.NewRole()
	role.IsDeleted = 1
	role.DeletedTime = nowTime
	err := r.db.WithContext(ctx).Where("id = ?", id).Updates(role).Error
	return err
}

// 权限关联表相关

// 创建
func (r *roleRepo) CreatePermRelInBatches(ctx context.Context, rolePermRels []*model.RolePermRel) error {
	err := r.db.WithContext(ctx).CreateInBatches(rolePermRels, 100).Error
	return err
}

func (r *roleRepo) DeletePermRelByIDs(ctx context.Context, roleID int64, deletePermIDs []int64) error {
	err := r.db.WithContext(ctx).Where("role_id = ? and perm_id in ?", roleID, deletePermIDs).Delete(model.NewRolePermRel()).Error
	return err
}
