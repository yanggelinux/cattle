package team

import (
	"context"
	"fmt"
	"github.com/yanggelinux/cattle/internal/store/model"
	"time"
)

type TeamFilter struct {
	Name     *string
	Page     *int
	PageSize *int
}

type teamRepo struct {
	db *model.MDB
}

func NewTeamRepo() TeamRepo {
	return &teamRepo{
		db: model.GetDB(),
	}
}

func (r *teamRepo) WithTx(tx *model.MDB) TeamRepo {
	return &teamRepo{db: tx}
}

func (r *teamRepo) GetList(ctx context.Context, filter *TeamFilter) ([]*model.Team, int64, error) {
	var total int64
	records := make([]*model.Team, 0, 0)
	tx := r.db.WithContext(ctx).Model(model.NewTeam())
	tx.Where("is_deleted = ?", 0)
	if filter.Name != nil && len(*filter.Name) > 0 {
		tx.Where("name like ?", fmt.Sprintf("%%%s%%", *filter.Name))
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
		"id,name,leader,director,updated_time,created_time")
	if filter.Page != nil && filter.PageSize != nil {
		tx.Offset((*filter.Page - 1) * (*filter.PageSize)).Limit(*filter.PageSize)
	}
	err = tx.Find(&records).Error
	if err != nil {
		return records, 0, err
	}
	return records, total, nil
}

func (r *teamRepo) GetByID(ctx context.Context, id int64) (*model.Team, error) {
	team := model.NewTeam()
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&team).Error
	return team, err
}

func (r *teamRepo) GetByName(ctx context.Context, name string) (*model.Team, error) {
	team := model.NewTeam()
	err := r.db.WithContext(ctx).Where("name = ? and is_deleted = 0", name).First(&team).Error
	return team, err
}

func (r *teamRepo) GetAll(ctx context.Context) ([]*model.Team, error) {
	teams := make([]*model.Team, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewTeam()).Where("is_deleted = 0").Find(&teams).Error
	return teams, err
}

func (r *teamRepo) CheckDuplicateEntry(ctx context.Context, name string) bool {
	team := model.NewTeam()
	err := r.db.WithContext(ctx).Where("name = ? and is_deleted = 0", name).First(&team).Error
	if err == nil {
		return true
	}
	return false
}

// 创建
func (r *teamRepo) Create(ctx context.Context, team *model.Team) error {
	err := r.db.WithContext(ctx).Create(team).Error
	return err
}

// 更新, 可以更新多列指定字段
func (r *teamRepo) Update(ctx context.Context, id int64, updateField map[string]interface{}) error {
	err := r.db.WithContext(ctx).Model(model.NewTeam()).Where("id = ?", id).Updates(updateField).Error
	return err
}

// 更新 保存 支持非 0字段 主键存在更新所有字段，主键不存在插入
func (r *teamRepo) CreateOrUpdate(ctx context.Context, team *model.Team) error {
	err := r.db.WithContext(ctx).Save(team).Error
	return err
}

// 逻辑删除
func (r *teamRepo) DeleteByID(ctx context.Context, id int64) error {
	nowTime := time.Now()
	team := model.NewTeam()
	team.IsDeleted = 1
	team.DeletedTime = nowTime
	err := r.db.WithContext(ctx).Where("id = ?", id).Updates(team).Error
	return err
}
