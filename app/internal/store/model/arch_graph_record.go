package model

import (
	"gorm.io/datatypes"
	"time"
)

func NewArchGraphRecord() *ArchGraphRecord {
	return &ArchGraphRecord{}
}

type ArchGraphRecord struct {
	ID          int64          `json:"id" gorm:"column:id"`
	GraphID     int64          `json:"graphID" gorm:"column:graph_id"`
	NodeData    datatypes.JSON `json:"nodeData" gorm:"column:node_data" type:"json"`
	EdgeData    datatypes.JSON `json:"edgeData" gorm:"column:edge_data" type:"json"`
	ImageData   string         `json:"imageData" gorm:"column:image_data"`
	RecordType  int8           `json:"recordType" gorm:"column:record_type"`
	CreatedTime time.Time      `json:"createdTime"  gorm:"column:created_time"`
}

func (m *ArchGraphRecord) TableName() string {
	return "arch_graph_record"
}

func (m *ArchGraphRecord) IDField() string {
	return "id"
}
func (m *ArchGraphRecord) GraphIDField() string {
	return "graph_id"
}
func (m *ArchGraphRecord) NodeDataField() string {
	return "node_data"
}
func (m *ArchGraphRecord) EdgeDataField() string {
	return "edge_data"
}
func (m *ArchGraphRecord) ImageDataField() string {
	return "image_data"
}
func (m *ArchGraphRecord) RecordTypeField() string {
	return "record_type"
}
func (m *ArchGraphRecord) CreatedTimeField() string {
	return "created_time"
}
