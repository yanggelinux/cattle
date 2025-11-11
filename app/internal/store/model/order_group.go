package model

func NewOrderGroup() *OrderGroup {
	return &OrderGroup{Model: &Model{}}
}

type OrderGroup struct {
	ID     int64  `json:"id" gorm:"column:id"`
	Name   string `json:"name" gorm:"column:name"`
	Sort   int64  `json:"sort" gorm:"column:sort"`
	Status int8   `json:"status" gorm:"column:status"`
	*Model
}

func (m *OrderGroup) TableName() string {
	return "order_group"
}

func (m *OrderGroup) IDField() string {
	return "id"
}
func (m *OrderGroup) NameField() string {
	return "name"
}
func (m *OrderGroup) SortField() string {
	return "sort"
}
func (m *OrderGroup) StatusField() string {
	return "status"
}
