package order

import (
	"context"
	"fmt"
	"github.com/yanggelinux/cattle/internal/store/model"
	"time"
)

type OrderGroupFilter struct {
	Name     *string
	Status   *int8
	Page     *int
	PageSize *int
}
type orderGroupRepo struct {
	db *model.MDB
}

func NewOrderGroupRepo() OrderGroupRepo {
	return &orderGroupRepo{db: model.GetDB()}
}

func (r *orderGroupRepo) WithTx(tx *model.MDB) OrderGroupRepo {
	return &orderGroupRepo{db: tx}
}

func (r *orderGroupRepo) GetList(ctx context.Context, filter *OrderGroupFilter) ([]*model.OrderGroup, int64, error) {
	var total int64
	records := make([]*model.OrderGroup, 0, 0)
	tx := r.db.WithContext(ctx).Model(model.NewOrderGroup())
	tx.Where("is_deleted = ?", 0)
	if filter.Name != nil {
		tx.Where("name like ?", fmt.Sprintf("%%%s%%", *filter.Name))
	}
	if filter.Status != nil {
		tx.Where("status = ?", *filter.Status)
	}
	err := tx.Count(&total).Error
	if err != nil {
		return records, 0, err
	}
	if total == 0 {
		return records, 0, nil
	}
	tx.Order("sort desc")
	tx.Select(
		"id,name,sort,status,updated_time,created_time")
	if filter.Page != nil && filter.PageSize != nil {
		tx.Offset((*filter.Page - 1) * *filter.PageSize).Limit(*filter.PageSize)
	}
	err = tx.Find(&records).Error
	if err != nil {
		return records, 0, err
	}
	return records, total, nil
}

func (r *orderGroupRepo) GetByID(ctx context.Context, id int64) (*model.OrderGroup, error) {
	orderGroup := model.NewOrderGroup()
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&orderGroup).Error
	return orderGroup, err
}

func (r *orderGroupRepo) GetAll(ctx context.Context) ([]*model.OrderGroup, error) {
	orderGroups := make([]*model.OrderGroup, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewOrderGroup()).Where("is_deleted = 0").Find(&orderGroups).Error
	return orderGroups, err
}

func (r *orderGroupRepo) CheckDuplicateEntry(ctx context.Context, name string) bool {
	orderGroup := model.NewOrderGroup()
	err := r.db.WithContext(ctx).Where("name = ? and is_deleted = 0", name).First(&orderGroup).Error
	if err == nil {
		return true
	}
	return false
}

// 创建
func (r *orderGroupRepo) Create(ctx context.Context, orderGroup *model.OrderGroup) error {
	err := r.db.WithContext(ctx).Create(orderGroup).Error
	return err
}

// 更新, 可以更新多列指定字段
func (r *orderGroupRepo) Update(ctx context.Context, id int64, updateField map[string]interface{}) error {
	err := r.db.WithContext(ctx).Model(model.NewOrderGroup()).Where("id = ?", id).Updates(updateField).Error
	return err
}

// 更新 保存 支持非 0字段 主键存在更新所有字段，主键不存在插入
func (r *orderGroupRepo) CreateOrUpdate(ctx context.Context, orderGroup *model.OrderGroup) error {
	err := r.db.WithContext(ctx).Save(orderGroup).Error
	return err
}

// 逻辑删除
func (r *orderGroupRepo) DeleteByID(ctx context.Context, id int64) error {
	nowTime := time.Now()
	orderGroup := model.NewOrderGroup()
	orderGroup.IsDeleted = 1
	orderGroup.DeletedTime = nowTime
	err := r.db.WithContext(ctx).Where("id = ?", id).Updates(orderGroup).Error
	return err
}
