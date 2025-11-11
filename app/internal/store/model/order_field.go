package model

func NewOrderField() *OrderField {
	return &OrderField{Model: &Model{}}
}

type OrderField struct {
	ID           int64  `json:"id" gorm:"column:id"`
	OrderID      int64  `json:"orderID" gorm:"column:order_id"`
	Name         string `json:"name" gorm:"column:name"`
	Key          string `json:"key" gorm:"column:key"`
	Component    string `json:"component" gorm:"column:component"`
	Placeholder  string `json:"placeholder" gorm:"column:placeholder"`
	VerRule      int8   `json:"verRule" gorm:"column:ver_rule"`
	DefaultVal   string `json:"defaultVal" gorm:"column:default_val"`
	IsRequired   int8   `json:"isRequired" gorm:"column:is_required"`
	IsTitle      int8   `json:"isTitle" gorm:"column:is_title"`
	IsEdit       int8   `json:"isEdit" gorm:"column:is_edit"`
	IsClear      int8   `json:"isClear" gorm:"column:is_clear"`
	DisplayField string `json:"displayField" gorm:"column:display_field"`
	DisplayVal   string `json:"displayVal" gorm:"column:display_val"`
	Description  string `json:"description" gorm:"column:description"`
	Enum         string `json:"enum" gorm:"column:enum"`
	GroupName    string `json:"groupName" gorm:"column:group_name"`
	Sort         int64  `json:"sort" gorm:"column:sort"`
	Status       int8   `json:"status" gorm:"column:status"`
	*Model
}

func (m *OrderField) TableName() string {
	return "order_field"
}

func (m *OrderField) IDField() string {
	return "id"
}
func (m *OrderField) OrderIDField() string {
	return "order_id"
}
func (m *OrderField) NameField() string {
	return "name"
}
func (m *OrderField) KeyField() string {
	return "key"
}
func (m *OrderField) ComponentField() string {
	return "component"
}
func (m *OrderField) PlaceholderField() string {
	return "placeholder"
}
func (m *OrderField) VerRuleField() string {
	return "ver_rule"
}
func (m *OrderField) DefaultValField() string {
	return "default_val"
}
func (m *OrderField) IsRequiredField() string {
	return "is_required"
}
func (m *OrderField) IsTitleField() string {
	return "is_title"
}
func (m *OrderField) IsEditField() string {
	return "is_edit"
}
func (m *OrderField) IsClearField() string {
	return "is_clear"
}
func (m *OrderField) DisplayFieldField() string {
	return "display_field"
}
func (m *OrderField) DisplayValField() string {
	return "display_val"
}
func (m *OrderField) DescriptionField() string {
	return "description"
}
func (m *OrderField) EnumField() string {
	return "enum"
}
func (m *OrderField) GroupNameField() string {
	return "group_name"
}
func (m *OrderField) SortField() string {
	return "sort"
}
func (m *OrderField) StatusField() string {
	return "status"
}
