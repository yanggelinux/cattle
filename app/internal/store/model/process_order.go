package model

import (
	"gorm.io/datatypes"
	"time"
)

func NewProcessOrder() *ProcessOrder {
	return &ProcessOrder{Model: &Model{}}
}

type ProcessOrder struct {
	ID               int64          `json:"id" gorm:"column:id"`
	GraphID          int64          `json:"graphID" gorm:"column:graph_id"`
	OrderID          int64          `json:"orderID" gorm:"column:order_id"`
	Title            string         `json:"title" gorm:"column:title"`
	Env              string         `json:"env" gorm:"column:env"`
	GraphName        string         `json:"graphName" gorm:"column:graph_name"`
	OrderName        string         `json:"orderName" gorm:"column:order_name"`
	DemandName       string         `json:"demandName" gorm:"column:demand_name"`
	Owner            string         `json:"owner" gorm:"column:owner"`
	OrderField       datatypes.JSON `json:"orderField" gorm:"column:order_field" type:"json"`
	OrderInfo        datatypes.JSON `json:"orderInfo" gorm:"column:order_info" type:"json"`
	OrderProcess     datatypes.JSON `json:"orderProcess" gorm:"column:order_process" type:"json"`
	ImageHash        string         `json:"imageHash" gorm:"column:image_hash"`
	EnabledImageHash string         `json:"enabledImageHash" gorm:"column:enabled_image_hash"`
	OrderType        int8           `json:"orderType" gorm:"column:order_type"`
	Description      string         `json:"description" gorm:"column:description"`
	TaskStatus       int8           `json:"taskStatus" gorm:"column:task_status"`
	TaskResult       datatypes.JSON `json:"taskResult" gorm:"column:task_result" type:"json"`
	Status           int8           `json:"status" gorm:"column:status"`
	*Model
}

func (m *ProcessOrder) TableName() string {
	return "process_order"
}

func (m *ProcessOrder) IDField() string {
	return "id"
}
func (m *ProcessOrder) GraphIDField() string {
	return "graph_id"
}
func (m *ProcessOrder) OrderIDField() string {
	return "order_id"
}
func (m *ProcessOrder) TitleField() string {
	return "title"
}
func (m *ProcessOrder) EnvField() string {
	return "env"
}
func (m *ProcessOrder) GraphNameField() string {
	return "graph_name"
}
func (m *ProcessOrder) OrderNameField() string {
	return "order_name"
}
func (m *ProcessOrder) DemandNameField() string {
	return "demand_name"
}
func (m *ProcessOrder) OwnerField() string {
	return "owner"
}
func (m *ProcessOrder) OrderFieldField() string {
	return "order_field"
}
func (m *ProcessOrder) OrderInfoField() string {
	return "order_info"
}
func (m *ProcessOrder) OrderProcessField() string {
	return "order_process"
}
func (m *ProcessOrder) ImageHashField() string {
	return "image_hash"
}
func (m *ProcessOrder) EnabledImageHashField() string {
	return "enabled_image_hash"
}
func (m *ProcessOrder) OrderTypeField() string {
	return "order_type"
}
func (m *ProcessOrder) DescriptionField() string {
	return "description"
}
func (m *ProcessOrder) TaskStatusField() string {
	return "task_status"
}
func (m *ProcessOrder) TaskResultField() string {
	return "task_result"
}
func (m *ProcessOrder) StatusField() string {
	return "status"
}

type ProcessOrderDist struct {
	Count int64     `json:"count"`
	DTT   time.Time `json:"dtt"`
}
