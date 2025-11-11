package processorder

import (
	"context"
	"github.com/yanggelinux/cattle/internal/store/model"
)

type processArchRepo struct {
	db *model.MDB
}

func NewProcessArchRepo() ProcessArchRepo {
	return &processArchRepo{db: model.GetDB()}
}
func (r *processArchRepo) WithTx(tx *model.MDB) ProcessArchRepo {
	return &processArchRepo{db: tx}
}

func (r *processArchRepo) GetByID(ctx context.Context, id int64) (*model.ProcessArch, error) {
	processArch := model.NewProcessArch()
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&processArch).Error
	return processArch, err
}

func (r *processArchRepo) GetImageHash(ctx context.Context, imageHash string) (*model.ProcessArch, error) {
	processArch := model.NewProcessArch()
	err := r.db.WithContext(ctx).Where("image_hash = ?", imageHash).First(&processArch).Error
	return processArch, err
}

func (r *processArchRepo) GetEnabledImageHash(ctx context.Context, enabledImageHash string) (*model.ProcessArch, error) {
	processArch := model.NewProcessArch()
	err := r.db.WithContext(ctx).Where("enabled_image_hash = ?", enabledImageHash).First(&processArch).Error
	return processArch, err
}

// 创建
func (r *processArchRepo) Create(ctx context.Context, processArch *model.ProcessArch) error {
	err := r.db.WithContext(ctx).Create(processArch).Error
	return err
}

// 更新, 可以更新多列指定字段
func (r *processArchRepo) Update(ctx context.Context, id int64, updateField map[string]interface{}) error {
	err := r.db.WithContext(ctx).Model(model.NewProcessArch()).Where("id = ?", id).Updates(updateField).Error
	return err
}

// 更新 保存 支持非 0字段 主键存在更新所有字段，主键不存在插入
func (r *processArchRepo) CreateOrUpdate(ctx context.Context, processArch *model.ProcessArch) error {
	err := r.db.WithContext(ctx).Save(processArch).Error
	return err
}

// 物理删除
func (r *processArchRepo) DeleteByID(ctx context.Context, id int64) error {
	processArch := model.NewProcessArch()
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(processArch).Error
	return err
}
