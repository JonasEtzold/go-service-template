package example

import (
	"time"

	"gorm.io/gorm"
)

type Example struct {
	gorm.Model
	Name string `gorm:"column:name;not null;" json:"name" form:"name" binding:"required"`
	Text string `gorm:"column:text;not null;" json:"text" form:"text" binding:"required"`
}

func (m *Example) BeforeCreate(tx *gorm.DB) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Example) BeforeUpdate(tx *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
