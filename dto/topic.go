package dto

type TopicUpdateDTO struct {
	ID        uint64 `json:"id" form:"id" binding:"required"`
	NewsID    uint64 `gorm:"news_id" json:"news_id"`
	Name      string `json:"name" form:"name" binding:"required"`
	CreatedBy string `gorm:"created_by,omitempty" json:"created_by"`
	UpdatedBy string `gorm:"updated_by,omitempty" json:"updated_by"`
	DeletedBy string `gorm:"deleted_by,omitempty" json:"deleted_by"`
}

type TopicCreateDTO struct {
	Name      string `json:"name" form:"name" binding:"required"`
	NewsID    uint64 `gorm:"news_id" json:"news_id"`
	CreatedBy string `gorm:"created_by,omitempty" json:"created_by"`
	UpdatedBy string `gorm:"updated_by,omitempty" json:"updated_by"`
	DeletedBy string `gorm:"deleted_by,omitempty" json:"deleted_by"`
}
