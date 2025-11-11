package model

import (
	"time"
)

func NewDemandApproval() *DemandApproval {
	return &DemandApproval{}
}

type DemandApproval struct {
	ID          int64     `json:"id" gorm:"column:id"`
	DemandID    int64     `json:"demandID" gorm:"column:demand_id"`
	Approver    string    `json:"approver" gorm:"column:approver"`
	Action      string    `json:"action" gorm:"column:action"`
	Opinion     string    `json:"opinion" gorm:"column:opinion"`
	Status      int8      `json:"status" gorm:"column:status"`
	CreatedTime time.Time `json:"createdTime"  gorm:"column:created_time"`
}

func (m *DemandApproval) TableName() string {
	return "demand_approval"
}

func (m *DemandApproval) IDField() string {
	return "id"
}
func (m *DemandApproval) DemandIDField() string {
	return "demand_id"
}
func (m *DemandApproval) ApproverField() string {
	return "approver"
}
func (m *DemandApproval) ActionField() string {
	return "action"
}
func (m *DemandApproval) OpinionField() string {
	return "opinion"
}
func (m *DemandApproval) StatusField() string {
	return "status"
}
func (m *DemandApproval) CreatedTimeField() string {
	return "created_time"
}
