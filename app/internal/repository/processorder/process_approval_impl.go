package processorder

import (
	"context"
	"github.com/yanggelinux/cattle/internal/store/model"
)

type processApprovalRepo struct {
	db *model.MDB
}

func NewProcessApprovalRepo() ProcessApprovalRepo {
	return &processApprovalRepo{db: model.GetDB()}
}
func (r *processApprovalRepo) WithTx(tx *model.MDB) ProcessApprovalRepo {
	return &processApprovalRepo{db: tx}
}

func (r *processApprovalRepo) GetByID(ctx context.Context, id int64) (*model.ProcessApproval, error) {
	processApproval := model.NewProcessApproval()
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&processApproval).Error
	return processApproval, err
}

func (r *processApprovalRepo) GetByOrder(ctx context.Context, orderID int64) ([]*model.ProcessApproval, error) {
	processApprovals := make([]*model.ProcessApproval, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewProcessApproval()).Where("order_id = ?", orderID).Find(&processApprovals).Error
	return processApprovals, err
}
func (r *processApprovalRepo) GetByApprover(ctx context.Context, approver string) ([]*model.ProcessApproval, error) {
	processApprovals := make([]*model.ProcessApproval, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewProcessApproval()).Where("approver = ?", approver).Find(&processApprovals).Error
	return processApprovals, err
}

// 创建
func (r *processApprovalRepo) Create(ctx context.Context, processApproval *model.ProcessApproval) error {
	err := r.db.WithContext(ctx).Create(processApproval).Error
	return err
}

// 更新, 可以更新多列指定字段
func (r *processApprovalRepo) Update(ctx context.Context, id int64, updateField map[string]interface{}) error {
	err := r.db.WithContext(ctx).Model(model.NewProcessApproval()).Where("id = ?", id).Updates(updateField).Error
	return err
}

// 更新 保存 支持非 0字段 主键存在更新所有字段，主键不存在插入
func (r *processApprovalRepo) CreateOrUpdate(ctx context.Context, processApproval *model.ProcessApproval) error {
	err := r.db.WithContext(ctx).Save(processApproval).Error
	return err
}

// 物理删除
func (r *processApprovalRepo) DeleteByID(ctx context.Context, id int64) error {
	processApproval := model.NewProcessApproval()
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(processApproval).Error
	return err
}
