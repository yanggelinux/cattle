package archgroup

import (
	"context"
	"fmt"
	"github.com/yanggelinux/cattle/internal/store/model"
	"time"
)

type ArchGroupFilter struct {
	ParentID  *int64
	GroupName *string
}

type archGroupRepo struct {
	db *model.MDB
}

func NewArchGroupRepo() ArchGroupRepo {
	return &archGroupRepo{db: model.GetDB()}
}

func (r *archGroupRepo) WithTx(tx *model.MDB) ArchGroupRepo {
	return &archGroupRepo{db: tx}
}

func (r *archGroupRepo) GetList(ctx context.Context, filter *ArchGroupFilter) ([]*model.ArchGroup, error) {

	records := make([]*model.ArchGroup, 0, 0)
	tx := r.db.WithContext(ctx).Model(model.NewArchGroup())
	tx.Where("is_deleted = ?", 0)
	if filter.GroupName != nil && len(*filter.GroupName) > 0 {
		tx.Where("group_name like ?", fmt.Sprintf("%%%s%%", *filter.GroupName))
	}
	if filter.ParentID != nil {
		tx.Where("parent_id = ?", *filter.ParentID)
	}
	tx.Order("created_time desc")
	err := tx.Select(
		"id,parent_id,group_name,updated_time,created_time").Find(&records).Error
	if err != nil {
		return records, err
	}
	return records, nil
}

func (r *archGroupRepo) GetByID(ctx context.Context, id int64) (*model.ArchGroup, error) {
	archGroup := model.NewArchGroup()
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&archGroup).Error
	return archGroup, err
}

func (r *archGroupRepo) GetByParent(ctx context.Context, parentID int64) ([]*model.ArchGroup, error) {
	records := make([]*model.ArchGroup, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewArchGroup()).
		Where("parent_id = ? and is_deleted =0", parentID).Find(&records).Error
	return records, err
}

func (r *archGroupRepo) GetAll(ctx context.Context) ([]*model.ArchGroup, error) {
	records := make([]*model.ArchGroup, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewArchGroup()).
		Where("is_deleted = 0").Find(&records).Error
	return records, err
}

func (r *archGroupRepo) GetCountByParent(ctx context.Context, parentID int64) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(model.NewArchGroup()).Where("parent_id = ? and is_deleted =0", parentID).Count(&count).Error
	return count, err
}

func (r *archGroupRepo) GetByName(ctx context.Context, groupName string) (*model.ArchGroup, error) {
	archGroup := model.NewArchGroup()
	err := r.db.WithContext(ctx).Where("group_name = ? and is_deleted = 0", groupName).First(&archGroup).Error
	return archGroup, err
}

func (r *archGroupRepo) CheckDuplicateEntry(ctx context.Context, parentID int64, groupName string) bool {
	archGroup := model.NewArchGroup()
	err := r.db.WithContext(ctx).Where("group_name = ? and parent_id = ? and is_deleted = 0", groupName, parentID).First(&archGroup).Error
	if err == nil {
		return true
	}
	return false
}

func (r *archGroupRepo) CheckHas(ctx context.Context, parentID int64) bool {
	archGroup := model.NewArchGroup()
	err := r.db.WithContext(ctx).Where("parent_id = ? and is_deleted = 0", parentID).First(&archGroup).Error
	if err == nil {
		return true
	}
	return false
}

// 创建
func (r *archGroupRepo) Create(ctx context.Context, archGroup *model.ArchGroup) error {
	err := r.db.WithContext(ctx).Create(archGroup).Error
	return err
}

// 更新, 可以更新多列指定字段
func (r *archGroupRepo) Update(ctx context.Context, id int64, updateField map[string]interface{}) error {
	err := r.db.WithContext(ctx).Model(model.NewArchGroup()).Where("id = ?", id).Updates(updateField).Error
	return err
}

// 更新 保存 支持 0字段 主键存在更新所有字段，主键不存在插入
func (r *archGroupRepo) CreateOrUpdate(ctx context.Context, archGroup *model.ArchGroup) error {
	err := r.db.WithContext(ctx).Save(archGroup).Error
	return err
}

// 逻辑删除
func (r *archGroupRepo) DeleteByID(ctx context.Context, id int64) error {
	nowTime := time.Now()
	archGroup := model.NewArchGroup()
	archGroup.IsDeleted = 1
	archGroup.DeletedTime = nowTime
	err := r.db.WithContext(ctx).Where("id = ?", id).Updates(archGroup).Error
	return err
}
