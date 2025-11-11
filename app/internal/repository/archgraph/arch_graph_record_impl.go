package archgraph

import (
	"context"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/store/model"
)

type ArchGraphRecordFilter struct {
	GraphID *int64
}

type archGraphRecordRepo struct {
	db *model.MDB
}

func NewArchGraphRecordRepo() ArchGraphRecordRepo {
	return &archGraphRecordRepo{db: model.GetDB()}
}
func (r *archGraphRecordRepo) WithTx(tx *model.MDB) ArchGraphRecordRepo {
	return &archGraphRecordRepo{db: tx}
}

func (r *archGraphRecordRepo) GetList(ctx context.Context, filter *ArchGraphRecordFilter) ([]*model.ArchGraphRecord, error) {

	records := make([]*model.ArchGraphRecord, 0, 0)
	tx := r.db.WithContext(ctx).Model(model.NewArchGraphRecord())
	tx.Where("graph_id = ?", *filter.GraphID)
	tx.Order("created_time desc")
	err := tx.Select(
		"id,graph_id,image_data,created_time").Limit(10).Find(&records).Error
	if err != nil {
		return records, err
	}
	return records, nil
}

func (r *archGraphRecordRepo) GetSnapshotRecord(ctx context.Context, graphID int64) (*model.ArchGraphRecord, error) {

	records := make([]*model.ArchGraphRecord, 0, 0)

	tx := r.db.WithContext(ctx)
	err := tx.Model(model.NewArchGraphRecord()).Where("graph_id = ? and record_type = 1", graphID).Order("created_time asc").Find(&records).Error
	if err != nil {
		return nil, err
	}
	// 不超过15条不操作
	if len(records) <= 15 {
		return nil, errors.New("do not get last record,return")
	}
	// 按照时间倒序排序 获取第一条更新
	return records[0], nil
}

func (r *archGraphRecordRepo) GetByID(ctx context.Context, id int64) (*model.ArchGraphRecord, error) {
	archGraphRecord := model.NewArchGraphRecord()
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&archGraphRecord).Error
	return archGraphRecord, err
}

func (r *archGraphRecordRepo) GetSyncRecord(ctx context.Context, graphID int64) (*model.ArchGraphRecord, error) {
	archGraphRecord := model.NewArchGraphRecord()
	err := r.db.WithContext(ctx).Where("graph_id = ? and record_type = 2", graphID).Order("created_time desc").First(&archGraphRecord).Error
	return archGraphRecord, err
}

// 获取审批生效的数据
func (r *archGraphRecordRepo) GetEnabledRecords(ctx context.Context, graphID int64) ([]*model.ArchGraphRecord, error) {
	records := make([]*model.ArchGraphRecord, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewArchGraphRecord()).Where("graph_id = ? and record_type = 3", graphID).Order("created_time desc").Find(&records).Error
	return records, err
}

// 创建
func (r *archGraphRecordRepo) Create(ctx context.Context, archGraphRecord *model.ArchGraphRecord) error {
	err := r.db.WithContext(ctx).Create(archGraphRecord).Error
	return err
}

// 更新, 可以更新多列指定字段
func (r *archGraphRecordRepo) Update(ctx context.Context, id int64, updateField map[string]interface{}) error {
	err := r.db.WithContext(ctx).Model(model.NewArchGraphRecord()).Where("id = ?", id).Updates(updateField).Error
	return err
}

// 更新 保存 支持非 0字段 主键存在更新所有字段，主键不存在插入
func (r *archGraphRecordRepo) CreateOrUpdate(ctx context.Context, archGraphRecord *model.ArchGraphRecord) error {
	err := r.db.WithContext(ctx).Save(archGraphRecord).Error
	return err
}

// 物理删除
func (r *archGraphRecordRepo) DeleteByID(ctx context.Context, id int64) error {
	archGraphRecord := model.NewArchGraphRecord()
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(archGraphRecord).Error
	return err
}
