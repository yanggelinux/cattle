package model

import (
	"time"
)

func NewRolePermRel() *RolePermRel {
	return &RolePermRel{}
}

type RolePermRel struct {
	ID          int64     `json:"id" gorm:"column:id"`
	RoleID      int64     `json:"roleID" gorm:"column:role_id"`
	PermID      int64     `json:"permID" gorm:"column:perm_id"`
	CreatedTime time.Time `json:"createdTime"  gorm:"column:created_time"`
}

func (m *RolePermRel) TableName() string {
	return "role_perm_rel"
}

func (m *RolePermRel) IDField() string {
	return "id"
}

func (m *RolePermRel) RoleIDField() string {
	return "role_id"
}
func (m *RolePermRel) PermIDField() string {
	return "perm_id"
}
func (m *RolePermRel) CreatedTimeField() string {
	return "created_time"
}
