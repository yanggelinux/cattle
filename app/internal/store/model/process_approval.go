package model

import (
	"time"
)

func NewProcessApproval() *ProcessApproval {
	return &ProcessApproval{}
}

type ProcessApproval struct {
	ID          int64     `json:"id" gorm:"column:id"`
	OrderID     int64     `json:"orderID" gorm:"column:order_id"`
	Approver    string    `json:"approver" gorm:"column:approver"`
	Action      string    `json:"action" gorm:"column:action"`
	Opinion     string    `json:"opinion" gorm:"column:opinion"`
	Status      int8      `json:"status" gorm:"column:status"`
	CreatedTime time.Time `json:"createdTime"  gorm:"column:created_time"`
}

func (m *ProcessApproval) TableName() string {
	return "process_approval"
}

func (m *ProcessApproval) IDField() string {
	return "id"
}
func (m *ProcessApproval) OrderIDField() string {
	return "order_id"
}
func (m *ProcessApproval) ApproverField() string {
	return "approver"
}
func (m *ProcessApproval) ActionField() string {
	return "action"
}
func (m *ProcessApproval) OpinionField() string {
	return "opinion"
}
func (m *ProcessApproval) StatusField() string {
	return "status"
}
func (m *ProcessApproval) CreatedTimeField() string {
	return "created_time"
}
