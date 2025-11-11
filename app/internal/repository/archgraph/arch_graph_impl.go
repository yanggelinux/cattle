package archgraph

import (
	"context"
	"fmt"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/store/model"
	"time"
)

type ArchGraphFilter struct {
	GroupID   *int64
	GraphName *string
	Status    *int8
	Page      *int
	PageSize  *int
}

type archGraphRepo struct {
	db *model.MDB
}

func NewArchGraphRepo() ArchGraphRepo {
	return &archGraphRepo{db: model.GetDB()}
}

func (r *archGraphRepo) WithTx(tx *model.MDB) ArchGraphRepo {
	return &archGraphRepo{db: tx}
}

func (r *archGraphRepo) GetList(ctx context.Context, filter *ArchGraphFilter) ([]*result.ArchGraphRecord, int64, error) {
	var total int64
	records := make([]*result.ArchGraphRecord, 0, 0)
	tx := r.db.WithContext(ctx).Model(model.NewArchGraph())
	tx.Where("arch_graph.is_deleted = ?", 0)
	if filter.GraphName != nil && len(*filter.GraphName) > 0 {
		tx.Where("arch_graph.graph_name like ?", fmt.Sprintf("%%%s%%", *filter.GraphName))
	}
	if filter.Status != nil {
		tx.Where("arch_graph.status = ?", *filter.Status)
	}
	if filter.GroupID != nil {
		tx.Where("arch_graph.group_id = ?", *filter.GroupID)
	}
	err := tx.Count(&total).Error
	if err != nil {
		return records, 0, err
	}
	if total == 0 {
		return records, 0, nil
	}
	tx.Order("arch_graph.created_time desc")
	tx.Select(
		"arch_graph.id,arch_graph.group_id,arch_graph.graph_name,arch_graph.graph_key," +
			"arch_graph.image_data,arch_graph.owner,arch_graph.status,arch_graph.is_shared," +
			"arch_graph.graph_label,arch_graph.updated_time,arch_graph.created_time,arch_group.group_name").
		Joins("left join arch_group on arch_graph.group_id = arch_group.id")
	if filter.Page != nil && filter.PageSize != nil {
		tx.Offset((*filter.Page - 1) * (*filter.PageSize)).Limit(*filter.PageSize)
	}
	err = tx.Find(&records).Error
	if err != nil {
		return nil, 0, err
	}
	return records, total, nil
}

func (r *archGraphRepo) GetCount(ctx context.Context, filter *ArchGraphFilter) (int64, error) {
	var total int64
	tx := r.db.WithContext(ctx).Model(model.NewArchGraph())
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

func (r *archGraphRepo) GetByID(ctx context.Context, id int64) (*model.ArchGraph, error) {
	archGraph := model.NewArchGraph()
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&archGraph).Error
	return archGraph, err
}

func (r *archGraphRepo) GetByName(ctx context.Context, graphName string) (*model.ArchGraph, error) {
	archGraph := model.NewArchGraph()
	err := r.db.WithContext(ctx).Where("graph_name = ? and is_deleted = 0", graphName).First(&archGraph).Error
	return archGraph, err
}
func (r *archGraphRepo) GetByLabel(ctx context.Context, graphLabel string) (*model.ArchGraph, error) {
	archGraph := model.NewArchGraph()
	err := r.db.WithContext(ctx).Where("graph_label = ? and is_deleted = 0", graphLabel).First(&archGraph).Error
	return archGraph, err
}

func (r *archGraphRepo) GetAll(ctx context.Context) ([]*model.ArchGraph, error) {
	archGraphs := make([]*model.ArchGraph, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewArchGraph()).
		Select("id,group_id,graph_name,node_data,edge_data").Where("is_deleted = 0 and status = 2").
		Order("created_time desc").Find(&archGraphs).Error
	return archGraphs, err
}

func (r *archGraphRepo) GetSomeAll(ctx context.Context, selectField string) ([]*model.ArchGraph, error) {
	archGraphs := make([]*model.ArchGraph, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewArchGraph()).
		Select(selectField).Where("is_deleted = 0 and status != 4").Find(&archGraphs).Error
	return archGraphs, err
}

func (r *archGraphRepo) GetCountByGroup(ctx context.Context, groupID int64) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(model.NewArchGraph()).Where("group_id = ? and status != 4 and is_deleted = 0", groupID).Count(&count).Error
	return count, err
}

func (r *archGraphRepo) GetEnabledList(ctx context.Context, graphKey string) ([]*model.ArchGraph, error) {
	archGraphs := make([]*model.ArchGraph, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewArchGraph()).Where("graph_key = ? and status = 2 and is_deleted = 0", graphKey).
		Order("created_time desc").Find(&archGraphs).Error
	return archGraphs, err
}

func (r *archGraphRepo) CheckDuplicateEntry(ctx context.Context, groupID int64, graphLabel string) bool {
	archGraph := model.NewArchGraph()
	err := r.db.WithContext(ctx).Where("group_id = ? and graph_label = ? and is_deleted = 0", groupID, graphLabel).First(&archGraph).Error
	if err == nil {
		return true
	}
	return false
}

// 校验组下是否有图
func (r *archGraphRepo) CheckHas(ctx context.Context, groupID int64) bool {
	archGraph := model.NewArchGraph()
	err := r.db.WithContext(ctx).Where("group_id = ? and is_deleted = 0", groupID).First(&archGraph).Error
	if err == nil {
		return true
	}
	return false
}

// 创建
func (r *archGraphRepo) Create(ctx context.Context, archGraph *model.ArchGraph) error {
	err := r.db.WithContext(ctx).Create(archGraph).Error
	return err
}

// 更新, 可以更新多列指定字段
func (r *archGraphRepo) Update(ctx context.Context, id int64, updateField map[string]interface{}) error {
	err := r.db.WithContext(ctx).Model(model.NewArchGraph()).Where("id = ?", id).Updates(updateField).Error
	return err
}

func (r *archGraphRepo) UpdateByIDs(ctx context.Context, ids []int64, updateField map[string]interface{}) error {
	err := r.db.WithContext(ctx).Model(model.NewArchGraph()).Where("id in ?", ids).Updates(updateField).Error
	return err
}

// 更新 保存 支持非 0字段 主键存在更新所有字段，主键不存在插入
func (r *archGraphRepo) CreateOrUpdate(ctx context.Context, archGraph *model.ArchGraph) error {
	err := r.db.WithContext(ctx).Save(archGraph).Error
	return err
}

// 逻辑删除
func (r *archGraphRepo) DeleteByID(ctx context.Context, id int64) error {
	nowTime := time.Now()
	archGraph := model.NewArchGraph()
	archGraph.IsDeleted = 1
	archGraph.DeletedTime = nowTime
	err := r.db.WithContext(ctx).Where("id = ?", id).Updates(archGraph).Error
	return err
}
