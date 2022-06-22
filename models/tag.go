package models

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	ID          uint64         `gorm:"primary_key:auto_increment" json:"id"`
	Title       string         `gorm:"type:varchar(255)" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	News        *[]News        `gorm:"foreignKey:TagsId" json:"news"`
	CreatedBy   string         `gorm:"created_by,omitempty" json:"created_by"`
	UpdatedBy   string         `gorm:"updated_by,omitempty" json:"updated_by"`
	DeletedBy   string         `gorm:"deleted_by,omitempty" json:"deleted_by"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
