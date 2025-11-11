package model

func NewRole() *Role {
	return &Role{Model: &Model{}}
}

type Role struct {
	ID          int64  `json:"id" gorm:"column:id"`
	RoleName    string `json:"roleName" gorm:"column:role_name"`
	DisplayName string `json:"displayName" gorm:"column:display_name"`
	IsSuper     int8   `json:"isSuper" gorm:"column:is_super"`
	*Model
}

func (m *Role) TableName() string {
	return "role"
}

func (m *Role) IDField() string {
	return "id"
}
func (m *Role) RoleNameField() string {
	return "role_name"
}
func (m *Role) DisplayNameField() string {
	return "display_name"
}
func (m *Role) IsSuperField() string {
	return "is_super"
}
