package order

import (
	"context"
	"fmt"
	"github.com/yanggelinux/cattle/internal/store/model"
	"time"
)

type OrderFieldFilter struct {
	Name     *string
	Status   *int8
	OrderID  *int64
	Page     *int
	PageSize *int
}
type orderFieldRepo struct {
	db *model.MDB
}

func NewOrderFieldRepo() OrderFieldRepo {
	return &orderFieldRepo{db: model.GetDB()}
}

func (r *orderFieldRepo) WithTx(tx *model.MDB) OrderFieldRepo {
	return &orderFieldRepo{db: tx}
}

func (r *orderFieldRepo) GetList(ctx context.Context, filter *OrderFieldFilter) ([]*model.OrderField, int64, error) {
	var total int64
	records := make([]*model.OrderField, 0, 0)
	tx := r.db.WithContext(ctx).Model(model.NewOrderField())
	tx.Where("is_deleted = ?", 0)
	if filter.OrderID != nil {
		tx.Where("order_id = ?", *filter.OrderID)
	}
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
	tx.Order("sort,group_name asc")
	tx.Select(
		"id,order_id,name,`key`,component,placeholder,ver_rule,default_val," +
			"is_required,is_title,is_edit,is_clear,display_field," +
			"display_val,description,enum,group_name,sort," +
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

func (r *orderFieldRepo) GetByID(ctx context.Context, id int64) (*model.OrderField, error) {
	orderField := model.NewOrderField()
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&orderField).Error
	return orderField, err
}

func (r *orderFieldRepo) GetByOrder(ctx context.Context, orderID int64) ([]*model.OrderField, error) {
	orderFields := make([]*model.OrderField, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewOrderField()).
		Where("order_id = ? and is_deleted = 0", orderID).Find(&orderFields).Error
	return orderFields, err
}

func (r *orderFieldRepo) GetByOrderIDAndKey(ctx context.Context, orderID int64, key string) (*model.OrderField, error) {
	orderField := model.NewOrderField()
	err := r.db.WithContext(ctx).Where("order_id = ? and `key` = ? and is_deleted = 0", orderID, key).First(&orderField).Error
	return orderField, err
}

func (r *orderFieldRepo) CheckDuplicateEntry(ctx context.Context, key string, orderID int64) bool {
	orderField := model.NewOrderField()
	err := r.db.WithContext(ctx).Where("key = ? and order_id = ? and is_deleted = 0", key, orderID).First(&orderField).Error
	if err == nil {
		return true
	}
	return false
}

// 创建
func (r *orderFieldRepo) Create(ctx context.Context, orderField *model.OrderField) error {
	err := r.db.WithContext(ctx).Create(orderField).Error
	return err
}

func (r *orderFieldRepo) CreateInBatches(ctx context.Context, orderFields []*model.OrderField) error {
	err := r.db.WithContext(ctx).CreateInBatches(orderFields, 100).Error
	return err
}

// 更新, 可以更新多列指定字段
func (r *orderFieldRepo) Update(ctx context.Context, id int64, updateField map[string]interface{}) error {
	err := r.db.WithContext(ctx).Model(model.NewOrderField()).Where("id = ?", id).Updates(updateField).Error
	return err
}

// 更新 保存 支持非 0字段 主键存在更新所有字段，主键不存在插入
func (r *orderFieldRepo) CreateOrUpdate(ctx context.Context, orderField *model.OrderField) error {
	err := r.db.WithContext(ctx).Save(orderField).Error
	return err
}

// 逻辑删除
func (r *orderFieldRepo) DeleteByID(ctx context.Context, id int64) error {
	nowTime := time.Now()
	orderField := model.NewOrderField()
	orderField.IsDeleted = 1
	orderField.DeletedTime = nowTime
	err := r.db.WithContext(ctx).Where("id = ?", id).Updates(orderField).Error
	return err
}
