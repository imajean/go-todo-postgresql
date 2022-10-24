package models

import (
	"time"

	"gorm.io/gorm"
)

// used instead of gorm.Model to assign custom json key to these attributes
type GormModel struct {
	ID        uint           `gorm:"primaryKey" json:"id:"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Todo struct {
	GormModel

	Title string `json:"title"`
	Done  bool   `json:"done"`
}
