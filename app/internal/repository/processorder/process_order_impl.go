package processorder

import (
	"context"
	"fmt"
	"github.com/yanggelinux/cattle/internal/store/model"
	"time"
)

// 架构图审批节点

type ProcessOrderFilter struct {
	GraphName  *string
	DemandName *string
	Owner      *string
	Approver   *string
	StartTime  *string
	EndTime    *string
	OrderType  *int8
	Status     *int8
	Title      *string
	Page       *int
	PageSize   *int
}

type processOrderRepo struct {
	db *model.MDB
}

func NewProcessOrderRepo() ProcessOrderRepo {
	return &processOrderRepo{db: model.GetDB()}
}
func (r *processOrderRepo) WithTx(tx *model.MDB) ProcessOrderRepo {
	return &processOrderRepo{db: tx}
}

func (r *processOrderRepo) GetList(ctx context.Context, filter *ProcessOrderFilter) ([]*model.ProcessOrder, int64, error) {
	var total int64
	records := make([]*model.ProcessOrder, 0, 0)
	tx := r.db.WithContext(ctx).Model(model.NewProcessOrder())
	tx.Where("is_deleted = ?", 0)
	if filter.GraphName != nil && len(*filter.GraphName) > 0 {
		tx.Where("graph_name like ?", fmt.Sprintf("%%%s%%", *filter.GraphName))
	}
	if filter.DemandName != nil && len(*filter.DemandName) > 0 {
		tx.Where("demand_name like ?", fmt.Sprintf("%%%s%%", *filter.DemandName))
	}
	if filter.Title != nil && len(*filter.Title) > 0 {
		tx.Where("title like ?", fmt.Sprintf("%%%s%%", *filter.Title))
	}
	if filter.Owner != nil && len(*filter.Owner) > 0 {
		tx.Where("owner = ?", *filter.Owner)
	}
	if filter.OrderType != nil {
		tx.Where("order_type = ?", *filter.OrderType)
	}
	if filter.Status != nil {
		if *filter.Status == 12 {
			tx.Where("status in (1,2)")
		} else if *filter.Status == 13 {
			tx.Where("status in (1,3)")
		} else if *filter.Status == 14 {
			tx.Where("status in (1,4)")
		} else if *filter.Status == 123 {
			tx.Where("status in (1,2,3)")
		} else if *filter.Status == 124 {
			tx.Where("status in (1,2,4)")
		} else if *filter.Status == 23 {
			tx.Where("status in (2,3)")
		} else if *filter.Status == 24 {
			tx.Where("status in (2,4)")
		} else if *filter.Status == -1 {
			tx.Where("status >0")
		} else {
			tx.Where("status = ?", *filter.Status)
		}
	} else {
		tx.Where("status > 0")
	}
	if filter.StartTime != nil && len(*filter.StartTime) > 0 {
		tx.Where("created_time >= ?", *filter.StartTime)
	}
	if filter.EndTime != nil && len(*filter.EndTime) > 0 {
		tx.Where("created_time <= ?", *filter.EndTime)
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
		"id,order_id,graph_id,title,env,graph_name,order_name,order_process,order_type,demand_name,owner,task_status,status," +
			"updated_time,created_time")
	if filter.Page != nil && filter.PageSize != nil {
		tx.Offset((*filter.Page - 1) * *filter.PageSize).Limit(*filter.PageSize)
	}
	err = tx.Find(&records).Error
	if err != nil {
		return records, 0, err
	}
	return records, total, nil
}

func (r *processOrderRepo) GetApprovalList(ctx context.Context, filter *ProcessOrderFilter) ([]*model.ProcessOrder, int64, error) {
	var total int64
	records := make([]*model.ProcessOrder, 0, 0)
	tx := r.db.WithContext(ctx).Model(model.NewProcessOrder())
	tx.Where("process_order.is_deleted = ?", 0)
	if filter.GraphName != nil && len(*filter.GraphName) > 0 {
		tx.Where("process_order.graph_name like ?", fmt.Sprintf("%%%s%%", *filter.GraphName))
	}
	if filter.DemandName != nil && len(*filter.DemandName) > 0 {
		tx.Where("process_order.demand_name like ?", fmt.Sprintf("%%%s%%", *filter.DemandName))
	}
	if filter.Owner != nil && len(*filter.Owner) > 0 {
		tx.Where("process_order.owner = ?", *filter.Owner)
	}
	if filter.Title != nil && len(*filter.Title) > 0 {
		tx.Where("process_order.title = ?", *filter.Title)
	}
	if filter.OrderType != nil {
		tx.Where("process_order.order_type = ?", *filter.OrderType)
	}
	if filter.Status != nil {
		tx.Where("process_order.status = ?", *filter.Status)
	} else {
		// 不展示未审批的数据
		tx.Where("process_order.status != ?", 0)
	}
	if filter.StartTime != nil && len(*filter.StartTime) > 0 {
		tx.Where("process_order.created_time >= ?", *filter.StartTime)
	}
	if filter.EndTime != nil && len(*filter.EndTime) > 0 {
		tx.Where("process_order.created_time <= ?", *filter.EndTime)
	}
	if filter.Approver != nil && len(*filter.Approver) > 0 {
		tx.Where("process_approval.approver = ?", *filter.Approver)
	}
	tx.Joins("left join process_approval on process_order.id = process_approval.order_id")
	err := tx.Count(&total).Error
	if err != nil {
		return records, 0, err
	}
	if total == 0 {
		return records, 0, nil
	}
	tx.Order("created_time desc")
	tx.Select(
		"process_order.id,process_order.order_id,process_order.graph_id,process_order.title,process_order.env," +
			"process_order.graph_name,process_order.order_name,process_order.order_process,process_order.order_type," +
			"process_order.demand_name,process_order.owner,process_order.status," +
			"process_order.updated_time,process_order.created_time")
	if filter.Page != nil && filter.PageSize != nil {
		tx.Offset((*filter.Page - 1) * *filter.PageSize).Limit(*filter.PageSize)
	}
	err = tx.Find(&records).Error
	if err != nil {
		return records, 0, err
	}
	return records, total, nil
}

func (r *processOrderRepo) GetCount(ctx context.Context, filter *ProcessOrderFilter) (int64, error) {
	var total int64
	tx := r.db.WithContext(ctx).Model(model.NewProcessOrder())
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

func (r *processOrderRepo) GetCountByTime(ctx context.Context, orderType int8) ([]*model.ProcessOrderDist, error) {
	records := make([]*model.ProcessOrderDist, 0, 0)
	tx := r.db.WithContext(ctx).Model(model.NewProcessOrder())
	err := tx.Select("DATE(created_time) AS dtt,COUNT(*) AS count").
		Where("created_time >= CURDATE() - INTERVAL 30 DAY and order_type = ? and is_deleted = 0", orderType).
		Group("dtt").Find(&records).Order("dtt").Error
	if err != nil {
		return records, err
	}
	return records, nil
}

func (r *processOrderRepo) GetByID(ctx context.Context, id int64) (*model.ProcessOrder, error) {
	processOrder := model.NewProcessOrder()
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&processOrder).Error
	return processOrder, err
}

func (r *processOrderRepo) GetByIDStatus(ctx context.Context, graphID int64, status int8) ([]*model.ProcessOrder, error) {
	processOrders := make([]*model.ProcessOrder, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewProcessOrder()).
		Where("graph_id = ? and status = ? and is_deleted = 0", graphID, status).
		Order("created_time desc").Find(&processOrders).Error
	return processOrders, err
}

func (r *processOrderRepo) GetByParentIDStatus(ctx context.Context, parentID int64, status int8) ([]*model.ProcessOrder, error) {
	processOrders := make([]*model.ProcessOrder, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewProcessOrder()).
		Where("parent_id = ? and status = ? and is_deleted = 0", parentID, status).
		Order("created_time desc").Find(&processOrders).Error
	return processOrders, err
}

func (r *processOrderRepo) GetByParentID(ctx context.Context, parentID int64) ([]*model.ProcessOrder, error) {
	processOrders := make([]*model.ProcessOrder, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewProcessOrder()).
		Where("parent_id = ? and is_deleted = 0", parentID).
		Order("created_time desc").Find(&processOrders).Error
	return processOrders, err
}

func (r *processOrderRepo) GetTaskOrder(ctx context.Context, orderIDs []int64) ([]*model.ProcessOrder, error) {
	processOrders := make([]*model.ProcessOrder, 0, 0)
	err := r.db.WithContext(ctx).Model(model.NewProcessOrder()).Select("id,order_id,status,order_info,created_time").
		Where("status = 2 and task_status = 0 and order_id in ? and is_deleted = 0 and created_time >= CURDATE() - INTERVAL 7 DAY", orderIDs).
		Order("created_time desc").Find(&processOrders).Error
	return processOrders, err
}

// 创建
func (r *processOrderRepo) Create(ctx context.Context, processOrder *model.ProcessOrder) error {
	err := r.db.WithContext(ctx).Create(processOrder).Error
	return err
}

// 更新, 可以更新多列指定字段
func (r *processOrderRepo) Update(ctx context.Context, id int64, updateField map[string]interface{}) error {
	err := r.db.WithContext(ctx).Model(model.NewProcessOrder()).Where("id = ?", id).Updates(updateField).Error
	return err
}

func (r *processOrderRepo) Updates(ctx context.Context, ids []int64, updateField map[string]interface{}) error {
	err := r.db.WithContext(ctx).Model(model.NewProcessOrder()).Where("id in ?", ids).Updates(updateField).Error
	return err
}

// 更新 保存 支持非 0字段 主键存在更新所有字段，主键不存在插入
func (r *processOrderRepo) CreateOrUpdate(ctx context.Context, processOrder *model.ProcessOrder) error {
	err := r.db.WithContext(ctx).Save(processOrder).Error
	return err
}

func (r *processOrderRepo) DeleteByID(ctx context.Context, id int64) error {
	nowTime := time.Now()
	processOrder := model.NewProcessOrder()
	processOrder.IsDeleted = 1
	processOrder.DeletedTime = nowTime
	err := r.db.WithContext(ctx).Where("id = ?", id).Updates(processOrder).Error
	return err
}

// 物理删除
func (r *processOrderRepo) DeleteByID2(ctx context.Context, id int64) error {
	processOrder := model.NewProcessOrder()
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(processOrder).Error
	return err
}
