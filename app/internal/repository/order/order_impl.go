package order

import (
	"context"
	"fmt"
	"github.com/yanggelinux/cattle/internal/store/model"
	"time"
)

type OrderFilter struct {
	Name      *string
	Status    *int8
	OrderType *int8
	Page      *int
	PageSize  *int
}
type orderRepo struct {
	db *model.MDB
}

func NewOrderRepo() OrderRepo {
	return &orderRepo{db: model.GetDB()}
}

func (r *orderRepo) WithTx(tx *model.MDB) OrderRepo {
	return &orderRepo{db: tx}
}

func (r *orderRepo) GetList(ctx context.Context, filter *OrderFilter) ([]*model.Order, int64, error) {
	var total int64
	records := make([]*model.Order, 0, 0)
	tx := r.db.WithContext(ctx).Model(model.NewOrder())
	tx.Where("is_deleted = ?", 0)
	if filter.Name != nil {
		tx.Where("name like ?", fmt.Sprintf("%%%s%%", *filter.Name))
	}
	if filter.Status != nil {
		tx.Where("status = ?", *filter.Status)
	}
	if filter.OrderType != nil {
		tx.Where("order_type = ?", *filter.OrderType)
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
		"id,name,group_id,process_id,order_type,node_type,label,layout,task_url,task_method,sort," +
			"status,updated_time,created_time")
	if filter.Page != nil && filter.PageSize != nil {
		tx.Offset((*filter.Page - 1) * *filter.PageSize).Limit(*filter.PageSize)
	}
	err = tx.Find(&records).Error
	if err != nil {
		return records, 0, err
	}
	return records, total, nil
}

func (r *orderRepo) GetByID(ctx context.Context, id int64) (*model.Order, error) {
	order := model.NewOrder()
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&order).Error
	return order, err
}

func (r *orderRepo) GetByName(ctx context.Context, name string) (*model.Order, error) {
	order := model.NewOrder()
	err := r.db.WithContext(ctx).Where("name = ? and is_deleted = 0", name).First(&order).Error
	return order, err
}

func (r *orderRepo) GetByLabel(ctx context.Context, label string) ([]*model.Order, error) {
	orders := make([]*model.Order, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewOrder()).Where("label = ? and is_deleted = 0", label).Find(&orders).Error
	return orders, err
}

func (r *orderRepo) GetByLabels(ctx context.Context, labels []string) ([]*model.Order, error) {
	orders := make([]*model.Order, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewOrder()).Where("label in ? and is_deleted = 0", labels).Find(&orders).Error
	return orders, err
}

func (r *orderRepo) GetAll(ctx context.Context) ([]*model.Order, error) {
	orders := make([]*model.Order, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewOrder()).Where("is_deleted = 0").Find(&orders).Error
	return orders, err
}

func (r *orderRepo) CheckDuplicateEntry(ctx context.Context, name string) bool {
	order := model.NewOrder()
	err := r.db.WithContext(ctx).Where("name = ? and is_deleted = 0", name).First(&order).Error
	if err == nil {
		return true
	}
	return false
}

func (r *orderRepo) CheckDuplicateLabel(ctx context.Context, label string) bool {
	// 不校验为空的情况
	if len(label) == 0 {
		return false
	}
	order := model.NewOrder()
	err := r.db.WithContext(ctx).Where("label = ? and is_deleted = 0", label).First(&order).Error
	if err == nil {
		return true
	}
	return false
}

// 创建
func (r *orderRepo) Create(ctx context.Context, order *model.Order) error {
	err := r.db.WithContext(ctx).Create(order).Error
	return err
}

// 更新, 可以更新多列指定字段
func (r *orderRepo) Update(ctx context.Context, id int64, updateField map[string]interface{}) error {
	err := r.db.WithContext(ctx).Model(model.NewOrder()).Where("id = ?", id).Updates(updateField).Error
	return err
}

// 更新 保存 支持非 0字段 主键存在更新所有字段，主键不存在插入
func (r *orderRepo) CreateOrUpdate(ctx context.Context, order *model.Order) error {
	err := r.db.WithContext(ctx).Save(order).Error
	return err
}

// 逻辑删除
func (r *orderRepo) DeleteByID(ctx context.Context, id int64) error {
	nowTime := time.Now()
	order := model.NewOrder()
	order.IsDeleted = 1
	order.DeletedTime = nowTime
	err := r.db.WithContext(ctx).Where("id = ?", id).Updates(order).Error
	return err
}
