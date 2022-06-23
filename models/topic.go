package models

import (
	"time"

	"gorm.io/gorm"
)

//User represents users table in database
type Topic struct {
	ID        uint64         `gorm:"primary_key:auto_increment" json:"id"`
	Name      string         `json:"name" form:"name" binding:"required"`
	News      []News         `gorm:"foreignKey:TopicId"`
	CreatedBy string         `gorm:"created_by,omitempty" json:"created_by"`
	UpdatedBy string         `gorm:"updated_by,omitempty" json:"updated_by"`
	DeletedBy string         `gorm:"deleted_by,omitempty" json:"deleted_by"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
