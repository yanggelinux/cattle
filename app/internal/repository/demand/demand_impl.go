package demand

import (
	"context"
	"fmt"
	"github.com/yanggelinux/cattle/internal/store/model"
	"time"
)

type DemandFilter struct {
	Status   *int8
	Name     *string
	Page     *int
	PageSize *int
}

type demandRepo struct {
	db *model.MDB
}

func NewDemandRepo() DemandRepo {
	return &demandRepo{db: model.GetDB()}
}

func (r *demandRepo) WithTx(tx *model.MDB) DemandRepo {
	return &demandRepo{db: tx}
}

func (r *demandRepo) GetList(ctx context.Context, filter *DemandFilter) ([]*model.Demand, int64, error) {
	var total int64
	records := make([]*model.Demand, 0, 0)
	tx := r.db.WithContext(ctx).Model(model.NewDemand())
	tx.Where("is_deleted = ?", 0)
	if filter.Name != nil && len(*filter.Name) > 0 {
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

	tx.Order("created_time desc")
	tx.Select(
		"id,name,demand_type,order_no,biz,owner,description,opinion,review_process,evaluation,is_evaluate," +
			"status,updated_time,created_time")
	if filter.Page != nil && filter.PageSize != nil {
		tx.Offset((*filter.Page - 1) * (*filter.PageSize)).Limit(*filter.PageSize)
	}
	err = tx.Find(&records).Error
	if err != nil {
		return records, total, err
	}
	return records, total, nil
}

func (r *demandRepo) GetCount(ctx context.Context, filter *DemandFilter) (int64, error) {
	var total int64
	tx := r.db.WithContext(ctx).Model(model.NewDemand())
	tx.Where("is_deleted = ?", 0)
	if filter.Status != nil {
		tx.Where("status = ?", *filter.Status)
	}
	err := tx.Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *demandRepo) GetByID(ctx context.Context, id int64) (*model.Demand, error) {
	demand := model.NewDemand()
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&demand).Error
	return demand, err
}

func (r *demandRepo) GetByName(ctx context.Context, name string) (*model.Demand, error) {
	demand := model.NewDemand()
	err := r.db.WithContext(ctx).Where("name = ? and is_deleted = 0", name).First(&demand).Error
	return demand, err
}

func (r *demandRepo) GetByStatus(ctx context.Context, status int8) ([]*model.Demand, error) {
	demands := make([]*model.Demand, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewDemand()).Where("status = ? and is_deleted = 0 and is_evaluate = 0", status).Find(&demands).Error
	return demands, err
}

func (r *demandRepo) CheckDuplicateEntry(ctx context.Context, name string) bool {
	demand := model.NewDemand()
	err := r.db.WithContext(ctx).Where("name = ? and is_deleted = 0", name).First(&demand).Error
	if err == nil {
		return true
	}
	return false
}

// 创建
func (r *demandRepo) Create(ctx context.Context, demand *model.Demand) error {
	err := r.db.WithContext(ctx).Create(demand).Error
	return err
}

// 更新, 可以更新多列指定字段
func (r *demandRepo) Update(ctx context.Context, id int64, updateField map[string]interface{}) error {
	err := r.db.WithContext(ctx).Model(model.NewDemand()).Where("id = ?", id).Updates(updateField).Error
	return err
}

// 更新 保存 支持 0字段 主键存在更新所有字段，主键不存在插入
func (r *demandRepo) CreateOrUpdate(ctx context.Context, demand *model.Demand) error {
	err := r.db.WithContext(ctx).Save(demand).Error
	return err
}

// 逻辑删除
func (r *demandRepo) DeleteByID(ctx context.Context, id int64) error {
	nowTime := time.Now()
	demand := model.NewDemand()
	demand.IsDeleted = 1
	demand.DeletedTime = nowTime
	err := r.db.WithContext(ctx).Where("id = ?", id).Updates(demand).Error
	return err
}
