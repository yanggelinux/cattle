package model

func NewArchGroup() *ArchGroup {
	return &ArchGroup{Model: &Model{}}
}

type ArchGroup struct {
	ID        int64  `json:"id" gorm:"column:id"`
	ParentID  int64  `json:"parentID" gorm:"column:parent_id"`
	GroupName string `json:"groupName" gorm:"column:group_name"`
	*Model
}

func (m *ArchGroup) TableName() string {
	return "arch_group"
}

func (m *ArchGroup) IDField() string {
	return "id"
}
func (m *ArchGroup) ParentIDField() string {
	return "parent_id"
}
func (m *ArchGroup) GroupNameField() string {
	return "group_name"
}
