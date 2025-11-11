package model

import "gorm.io/datatypes"

func NewTest() *Test {
	return &Test{Model: &Model{}}
}

type Test struct {
	ID   int64          `json:"id" gorm:"column:id"`
	Name string         `json:"name" gorm:"column:name"`
	Data datatypes.JSON `json:"data" gorm:"column:data" type:"json"`
	*Model
}

func (m *Test) TableName() string {
	return "test"
}
