package model

import (
	"gorm.io/datatypes"
)

func NewArchGraph() *ArchGraph {
	return &ArchGraph{Model: &Model{}}
}

type ArchGraph struct {
	ID         int64          `json:"id" gorm:"column:id"`
	GroupID    int64          `json:"groupID" gorm:"column:group_id"`
	GraphName  string         `json:"graphName" gorm:"column:graph_name"`
	GraphKey   string         `json:"graphKey" gorm:"column:graph_key"`
	GraphLabel string         `json:"graphLabel" gorm:"column:graph_label"`
	NodeData   datatypes.JSON `json:"nodeData" gorm:"column:node_data" type:"json"`
	EdgeData   datatypes.JSON `json:"edgeData" gorm:"column:edge_data" type:"json"`
	ImageData  string         `json:"imageData" gorm:"column:image_data"`
	Owner      string         `json:"owner" gorm:"column:owner"`
	Status     int8           `json:"status" gorm:"column:status"`
	IsShared   int8           `json:"isShared" gorm:"column:is_shared"`
	*Model
}

func (m *ArchGraph) TableName() string {
	return "arch_graph"
}

func (m *ArchGraph) IDField() string {
	return "id"
}
func (m *ArchGraph) GroupIDField() string {
	return "group_id"
}
func (m *ArchGraph) GraphNameField() string {
	return "graph_name"
}
func (m *ArchGraph) GraphKeyField() string {
	return "graph_key"
}
func (m *ArchGraph) GraphLabelField() string {
	return "graph_label"
}
func (m *ArchGraph) NodeDataField() string {
	return "node_data"
}
func (m *ArchGraph) EdgeDataField() string {
	return "edge_data"
}
func (m *ArchGraph) ImageDataField() string {
	return "image_data"
}
func (m *ArchGraph) OwnerField() string {
	return "owner"
}
func (m *ArchGraph) StatusField() string {
	return "status"
}
func (m *ArchGraph) IsSharedField() string {
	return "is_shared"
}
