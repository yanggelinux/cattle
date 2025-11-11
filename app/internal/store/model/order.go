package model

func NewOrder() *Order {
	return &Order{Model: &Model{}}
}

type Order struct {
	ID         int64  `json:"id" gorm:"column:id"`
	Name       string `json:"name" gorm:"column:name"`
	GroupID    int64  `json:"groupID" gorm:"column:group_id"`
	ProcessID  int64  `json:"processID" gorm:"column:process_id"`
	OrderType  int8   `json:"orderType" gorm:"column:order_type"`
	NodeType   string `json:"nodeType" gorm:"column:node_type"`
	Label      string `json:"label" gorm:"column:label"`
	Layout     int8   `json:"layout" gorm:"column:layout"`
	TaskUrl    string `json:"taskUrl" gorm:"column:task_url"`
	TaskMethod string `json:"taskMethod" gorm:"column:task_method"`
	Sort       int64  `json:"sort" gorm:"column:sort"`
	Status     int8   `json:"status" gorm:"column:status"`
	*Model
}

func (m *Order) TableName() string {
	return "order"
}

func (m *Order) IDField() string {
	return "id"
}
func (m *Order) NameField() string {
	return "name"
}

func (m *Order) GroupIDField() string {
	return "group_id"
}
func (m *Order) ProcessIDField() string {
	return "process_id"
}
func (m *Order) OrderTypeField() string {
	return "order_type"
}
func (m *Order) NodeTypeField() string {
	return "node_type"
}
func (m *Order) LabelField() string {
	return "label"
}
func (m *Order) LayoutField() string {
	return "layout"
}
func (m *Order) TaskUrlField() string {
	return "task_url"
}
func (m *Order) TaskMethodField() string {
	return "task_method"
}
func (m *Order) SortField() string {
	return "sort"
}
func (m *Order) StatusField() string {
	return "status"
}
