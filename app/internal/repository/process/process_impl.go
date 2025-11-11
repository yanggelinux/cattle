package process

import (
	"context"
	"fmt"
	"github.com/yanggelinux/cattle/internal/store/model"
	"time"
)

type ProcessFilter struct {
	Name     *string
	Status   *int8
	Page     *int
	PageSize *int
}
type processRepo struct {
	db *model.MDB
}

func NewProcessRepo() ProcessRepo {
	return &processRepo{db: model.GetDB()}
}

func (r *processRepo) WithTx(tx *model.MDB) ProcessRepo {
	return &processRepo{db: tx}
}

func (r *processRepo) GetList(ctx context.Context, filter *ProcessFilter) ([]*model.Process, int64, error) {
	var total int64
	records := make([]*model.Process, 0, 0)
	tx := r.db.WithContext(ctx).Model(model.NewProcess())
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
	tx.Order("created_time desc")
	tx.Select(
		"id,name,proc_info,node_data,edge_data,status,updated_time,created_time")
	if filter.Page != nil && filter.PageSize != nil {
		tx.Offset((*filter.Page - 1) * *filter.PageSize).Limit(*filter.PageSize)
	}
	err = tx.Find(&records).Error
	if err != nil {
		return records, 0, err
	}
	return records, total, nil
}

func (r *processRepo) GetByID(ctx context.Context, id int64) (*model.Process, error) {
	process := model.NewProcess()
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&process).Error
	return process, err
}

func (r *processRepo) GetAll(ctx context.Context) ([]*model.Process, error) {
	processes := make([]*model.Process, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewProcess()).Where("is_deleted = 0").Find(&processes).Error
	return processes, err
}

func (r *processRepo) CheckDuplicateEntry(ctx context.Context, name string) bool {
	process := model.NewProcess()
	err := r.db.WithContext(ctx).Where("name = ? and is_deleted = 0", name).First(&process).Error
	if err == nil {
		return true
	}
	return false
}

// 创建
func (r *processRepo) Create(ctx context.Context, process *model.Process) error {
	err := r.db.WithContext(ctx).Create(process).Error
	return err
}

// 更新, 可以更新多列指定字段
func (r *processRepo) Update(ctx context.Context, id int64, updateField map[string]interface{}) error {
	err := r.db.WithContext(ctx).Model(model.NewProcess()).Where("id = ?", id).Updates(updateField).Error
	return err
}

// 更新 保存 支持非 0字段 主键存在更新所有字段，主键不存在插入
func (r *processRepo) CreateOrUpdate(ctx context.Context, process *model.Process) error {
	err := r.db.WithContext(ctx).Save(process).Error
	return err
}

// 逻辑删除
func (r *processRepo) DeleteByID(ctx context.Context, id int64) error {
	nowTime := time.Now()
	process := model.NewProcess()
	process.IsDeleted = 1
	process.DeletedTime = nowTime
	err := r.db.WithContext(ctx).Where("id = ?", id).Updates(process).Error
	return err
}
