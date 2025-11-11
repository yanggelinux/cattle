package model

func NewPermission() *Permission {
	return &Permission{Model: &Model{}}
}

type Permission struct {
	ID        int64  `json:"id" gorm:"column:id"`
	ParentID  int64  `json:"parentID" gorm:"column:parent_id"`
	Name      string `json:"name" gorm:"column:name" `
	Code      string `json:"code" gorm:"column:code" `
	Uri       string `json:"uri" gorm:"column:uri"`
	Method    string `json:"method" gorm:"column:method"`
	Project   string `json:"project" gorm:"column:project"`
	PermType  int8   `json:"permType" gorm:"column:perm_type"`
	IsEnabled int8   `json:"isEnabled" gorm:"column:is_enabled"`
	Sort      int64  `json:"sort" gorm:"column:sort"`

	*Model
}

func (m *Permission) TableName() string {
	return "permission"
}

func (m *Permission) IDField() string {
	return "id"
}
func (m *Permission) ParentIDField() string {
	return "parent_id"
}
func (m *Permission) NameField() string {
	return "name"
}
func (m *Permission) CodeField() string {
	return "code"
}
func (m *Permission) UriField() string {
	return "uri"
}
func (m *Permission) MethodField() string {
	return "method"
}
func (m *Permission) ProjectField() string {
	return "project"
}
func (m *Permission) PermTypeField() string {
	return "perm_type"
}
func (m *Permission) IsEnabledField() string {
	return "is_enabled"
}
func (m *Permission) SortField() string {
	return "sort"
}
