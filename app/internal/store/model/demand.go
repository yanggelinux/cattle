package model

import "gorm.io/datatypes"

func NewDemand() *Demand {
	return &Demand{Model: &Model{}}
}

type Demand struct {
	ID            int64          `json:"id" gorm:"column:id"`
	Name          string         `json:"name" gorm:"column:name"`
	DemandType    int8           `json:"demandType" gorm:"column:demand_type"`
	OrderNo       string         `json:"orderNo" gorm:"column:order_no"`
	Biz           string         `json:"biz" gorm:"column:biz"`
	Owner         string         `json:"owner" gorm:"column:owner"`
	Description   string         `json:"description" gorm:"column:description"`
	Opinion       string         `json:"opinion" gorm:"column:opinion"`
	ReviewProcess datatypes.JSON `json:"reviewProcess" gorm:"column:review_process"`
	Evaluation    datatypes.JSON `json:"evaluation" gorm:"column:evaluation"`
	IsEvaluate    int8           `json:"isEvaluate" gorm:"column:is_evaluate"`
	Status        int8           `json:"status" gorm:"column:status"`
	*Model
}

func (m *Demand) TableName() string {
	return "demand"
}
func (m *Demand) IDField() string {
	return "id"
}
func (m *Demand) NameField() string {
	return "name"
}
func (m *Demand) DemandTypeField() string {
	return "demand_type"
}
func (m *Demand) OrderNoField() string {
	return "order_no"
}
func (m *Demand) BizField() string {
	return "Biz"
}
func (m *Demand) OwnerField() string {
	return "owner"
}
func (m *Demand) DescriptionField() string {
	return "Description"
}
func (m *Demand) OpinionField() string {
	return "Opinion"
}
func (m *Demand) ReviewProcessField() string {
	return "review_process"
}
func (m *Demand) EvaluationField() string {
	return "evaluation"
}
func (m *Demand) IsEvaluateField() string {
	return "is_evaluate"
}
func (m *Demand) StatusField() string {
	return "status"
}
