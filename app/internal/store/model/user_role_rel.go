package model

import (
	"time"
)

func NewUserRoleRel() *UserRoleRel {
	return &UserRoleRel{}
}

type UserRoleRel struct {
	ID          int64     `json:"id" gorm:"column:id"`
	UserID      int64     `json:"userID" gorm:"column:user_id"`
	RoleID      int64     `json:"roleID" gorm:"column:role_id"`
	CreatedTime time.Time `json:"createdTime"  gorm:"column:created_time"`
}

func (m *UserRoleRel) TableName() string {
	return "user_role_rel"
}

func (m *UserRoleRel) IDField() string {
	return "id"
}
func (m *UserRoleRel) UserIDField() string {
	return "user_id"
}
func (m *UserRoleRel) RoleIDField() string {
	return "role_id"
}
func (m *UserRoleRel) CreatedTimeField() string {
	return "created_time"
}
