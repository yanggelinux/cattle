package permission

import (
	"context"
	"fmt"
	"github.com/yanggelinux/cattle/internal/store/model"
	"time"
)

type PermissionFilter struct {
	Name     *string
	Code     *string
	Project  *string
	Page     *int
	PageSize *int
}
type permissionRepo struct {
	db *model.MDB
}

func NewPermissionRepo() PermissionRepo {
	return &permissionRepo{db: model.GetDB()}
}

func (r *permissionRepo) WithTx(tx *model.MDB) PermissionRepo {
	return &permissionRepo{db: tx}
}

func (r *permissionRepo) GetList(ctx context.Context, filter *PermissionFilter) ([]*model.Permission, int64, error) {
	var total int64
	records := make([]*model.Permission, 0, 0)
	tx := r.db.WithContext(ctx).Model(model.NewPermission())
	tx.Where("is_deleted = ?", 0)
	if filter.Name != nil {
		tx.Where("name like ?", fmt.Sprintf("%%%s%%", *filter.Name))
	}
	if filter.Code != nil {
		tx.Where("code like ?", fmt.Sprintf("%%%s%%", *filter.Code))
	}
	if filter.Project != nil {
		tx.Where("project = ?", *filter.Project)
	}
	err := tx.Count(&total).Error
	if err != nil {
		return records, 0, err
	}
	if total == 0 {
		return records, 0, nil
	}
	tx.Order("sort asc")
	tx.Select(
		"id,parent_id,name,code,uri,method,project,perm_type,is_enabled,sort,updated_time,created_time")
	if filter.Page != nil && filter.PageSize != nil {
		tx.Offset((*filter.Page - 1) * *filter.PageSize).Limit(*filter.PageSize)
	}
	err = tx.Find(&records).Error
	if err != nil {
		return records, 0, err
	}
	return records, total, nil
}

func (r *permissionRepo) GetByID(ctx context.Context, id int64) (*model.Permission, error) {
	permission := model.NewPermission()
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&permission).Error
	return permission, err
}

func (r *permissionRepo) GetByCode(ctx context.Context, code string) (*model.Permission, error) {
	permission := model.NewPermission()
	err := r.db.WithContext(ctx).Where("code = ? and is_deleted = 0", code).First(&permission).Error
	return permission, err
}

func (r *permissionRepo) CheckDuplicateEntry(ctx context.Context, code string, project string) bool {
	permission := model.NewPermission()
	err := r.db.WithContext(ctx).Where("code = ? and project = ? and is_deleted = 0", code, project).First(&permission).Error
	if err == nil {
		return true
	}
	return false
}

// 创建
func (r *permissionRepo) Create(ctx context.Context, permission *model.Permission) error {
	err := r.db.WithContext(ctx).Create(permission).Error
	return err
}

// 更新, 可以更新多列指定字段
func (r *permissionRepo) Update(ctx context.Context, id int64, updateField map[string]interface{}) error {
	err := r.db.WithContext(ctx).Model(model.NewPermission()).Where("id = ?", id).Updates(updateField).Error
	return err
}

// 更新 保存 支持非 0字段 主键存在更新所有字段，主键不存在插入
func (r *permissionRepo) CreateOrUpdate(ctx context.Context, permission *model.Permission) error {
	err := r.db.WithContext(ctx).Save(permission).Error
	return err
}

// 逻辑删除
func (r *permissionRepo) DeleteByID(ctx context.Context, id int64) error {
	nowTime := time.Now()
	permission := model.NewPermission()
	permission.IsDeleted = 1
	permission.DeletedTime = nowTime
	err := r.db.WithContext(ctx).Where("id = ?", id).Updates(permission).Error
	return err
}

func (r *permissionRepo) GetPermsByRole(ctx context.Context, roleID int64) ([]*model.Permission, error) {
	records := make([]*model.Permission, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewRolePermRel()).
		Select("permission.id,permission.parent_id,permission.name,permission.code,permission.uri,"+
			"permission.method,permission.project,permission.perm_type").
		Joins("left join permission on permission.id = role_perm_rel.perm_id").
		Where("role_perm_rel.role_id = ? and permission.is_deleted = 0 and permission.is_enabled = 1", roleID).Find(&records).Error

	if err != nil {
		return records, err
	}
	return records, nil
}

// 获取权限相关信息

func (r *permissionRepo) GetPermsByProject(ctx context.Context, project string) ([]*model.Permission, error) {
	records := make([]*model.Permission, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewPermission()).
		Select("permission.id,permission.parent_id,permission.name,permission.code,permission.uri,permission.method,permission.project,permission.perm_type").
		Where("permission.is_deleted = 0 and permission.is_enabled = 1 and project = ?", project).Find(&records).Error

	if err != nil {
		return records, err
	}
	return records, nil
}
