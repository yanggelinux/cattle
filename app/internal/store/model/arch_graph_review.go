package model

import (
	"time"
)

func NewArchGraphReview() *ArchGraphReview {
	return &ArchGraphReview{}
}

type ArchGraphReview struct {
	ID          int64     `json:"id" gorm:"column:id"`
	GraphID     int64     `json:"graphID" gorm:"column:graph_id"`
	GraphKey    string    `json:"graphKey" gorm:"column:graph_key"`
	Content     string    `json:"content" gorm:"column:content"`
	Reviewer    string    `json:"reviewer" gorm:"column:reviewer"`
	NotifyParty string    `json:"notifyParty" gorm:"column:notify_party"`
	CreatedTime time.Time `json:"createdTime"  gorm:"column:created_time"`
}

func (m *ArchGraphReview) TableName() string {
	return "arch_graph_review"
}

func (m *ArchGraphReview) IDField() string {
	return "id"
}
func (m *ArchGraphReview) GraphIDField() string {
	return "graph_id"
}
func (m *ArchGraphReview) GraphKeyField() string {
	return "graph_key"
}
func (m *ArchGraphReview) ContentField() string {
	return "content"
}
func (m *ArchGraphReview) ReviewerField() string {
	return "reviewer"
}
func (m *ArchGraphReview) NotifyPartyField() string {
	return "notify_party"
}
func (m *ArchGraphReview) CreatedTimeField() string {
	return "created_time"
}
