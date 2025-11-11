package archgraph

import (
	"context"
	"github.com/yanggelinux/cattle/internal/store/model"
)

type ArchGraphReviewFilter struct {
	GraphID  *int64
	GraphKey *string
	Page     *int
	PageSize *int
}

type archGraphReviewRepo struct {
	db *model.MDB
}

func NewArchGraphReviewRepo() ArchGraphReviewRepo {
	return &archGraphReviewRepo{db: model.GetDB()}
}
func (r *archGraphReviewRepo) WithTx(tx *model.MDB) ArchGraphReviewRepo {
	return &archGraphReviewRepo{db: tx}
}

func (r *archGraphReviewRepo) GetList(ctx context.Context, filter *ArchGraphReviewFilter) ([]*model.ArchGraphReview, int64, error) {
	var total int64
	records := make([]*model.ArchGraphReview, 0, 0)
	tx := r.db.WithContext(ctx).Model(model.NewArchGraphReview())
	if filter.GraphID != nil {
		tx = tx.Where("graph_id=?", *filter.GraphID)
	}
	if filter.GraphKey != nil && len(*filter.GraphKey) > 0 {
		tx = tx.Where("graph_key=?", *filter.GraphKey)
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
		"id,graph_id,graph_key,content,reviewer,notify_party,created_time")
	if filter.Page != nil && filter.PageSize != nil {
		tx.Offset((*filter.Page - 1) * (*filter.PageSize)).Limit(*filter.PageSize)
	}
	err = tx.Find(&records).Error
	if err != nil {
		return records, total, err
	}
	return records, total, nil
}

func (r *archGraphReviewRepo) GetByID(ctx context.Context, id int64) (*model.ArchGraphReview, error) {
	archGraphReview := model.NewArchGraphReview()
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&archGraphReview).Error
	return archGraphReview, err
}

// 创建
func (r *archGraphReviewRepo) Create(ctx context.Context, archGraphReview *model.ArchGraphReview) error {
	err := r.db.WithContext(ctx).Create(archGraphReview).Error
	return err
}

// 更新, 可以更新多列指定字段
func (r *archGraphReviewRepo) Update(ctx context.Context, id int64, updateField map[string]interface{}) error {
	err := r.db.WithContext(ctx).Model(model.NewArchGraphReview()).Where("id = ?", id).Updates(updateField).Error
	return err
}

// 更新 保存 支持非 0字段 主键存在更新所有字段，主键不存在插入
func (r *archGraphReviewRepo) CreateOrUpdate(ctx context.Context, archGraphReview *model.ArchGraphReview) error {
	err := r.db.WithContext(ctx).Save(archGraphReview).Error
	return err
}

// 物理删除
func (r *archGraphReviewRepo) DeleteByID(ctx context.Context, id int64) error {
	archGraphReview := model.NewArchGraphReview()
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(archGraphReview).Error
	return err
}
