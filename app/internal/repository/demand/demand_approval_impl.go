package demand

import (
	"context"
	"github.com/yanggelinux/cattle/internal/store/model"
)

type demandApprovalRepo struct {
	db *model.MDB
}

func NewDemandApprovalRepo() DemandApprovalRepo {
	return &demandApprovalRepo{db: model.GetDB()}
}
func (r *demandApprovalRepo) WithTx(tx *model.MDB) DemandApprovalRepo {
	return &demandApprovalRepo{db: tx}
}

func (r *demandApprovalRepo) GetByID(ctx context.Context, id int64) (*model.DemandApproval, error) {
	demandApproval := model.NewDemandApproval()
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&demandApproval).Error
	return demandApproval, err
}

func (r *demandApprovalRepo) GetByDemand(ctx context.Context, demandID int64) ([]*model.DemandApproval, error) {
	demandApprovals := make([]*model.DemandApproval, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewDemandApproval()).Where("demand_id = ?", demandID).Find(&demandApprovals).Error
	return demandApprovals, err
}

// 创建
func (r *demandApprovalRepo) Create(ctx context.Context, demandApproval *model.DemandApproval) error {
	err := r.db.WithContext(ctx).Create(demandApproval).Error
	return err
}

// 更新, 可以更新多列指定字段
func (r *demandApprovalRepo) Update(ctx context.Context, id int64, updateField map[string]interface{}) error {
	err := r.db.WithContext(ctx).Model(model.NewDemandApproval()).Where("id = ?", id).Updates(updateField).Error
	return err
}

// 更新 保存 支持非 0字段 主键存在更新所有字段，主键不存在插入
func (r *demandApprovalRepo) CreateOrUpdate(ctx context.Context, demandApproval *model.DemandApproval) error {
	err := r.db.WithContext(ctx).Save(demandApproval).Error
	return err
}

// 物理删除
func (r *demandApprovalRepo) DeleteByID(ctx context.Context, id int64) error {
	demandApproval := model.NewDemandApproval()
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(demandApproval).Error
	return err
}
