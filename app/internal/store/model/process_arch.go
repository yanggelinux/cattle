package model

import (
	"time"
)

func NewProcessArch() *ProcessArch {
	return &ProcessArch{}
}

type ProcessArch struct {
	ID               int64     `json:"id" gorm:"column:id"`
	ImageHash        string    `json:"imageHash" gorm:"column:image_hash"`
	EnabledImageHash string    `json:"enabledImageHash" gorm:"column:enabled_image_hash"`
	ImageData        string    `json:"imageData" gorm:"column:image_data"`
	EnabledImageData string    `json:"enabledImageData" gorm:"column:enabled_image_data"`
	CreatedTime      time.Time `json:"createdTime"  gorm:"column:created_time"`
}

func (m *ProcessArch) TableName() string {
	return "process_arch"
}

func (m *ProcessArch) IDField() string {
	return "id"
}
func (m *ProcessArch) ImageHashField() string {
	return "image_hash"
}
func (m *ProcessArch) EnabledImageHashField() string {
	return "enabled_image_hash"
}
func (m *ProcessArch) ImageDataField() string {
	return "image_data"
}
func (m *ProcessArch) EnabledImageDataField() string {
	return "enabled_image_data"
}
func (m *ProcessArch) CreatedTimeField() string {
	return "created_time"
}
