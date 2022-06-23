package dto

type NewsUpdateDTO struct {
	ID        uint64 `json:"id" form:"id" binding:"required"`
	TagsID    uint64 `json:"news_id,omitempty"  form:"news_id,omitempty"`
	Name      string `json:"name" form:"name" binding:"required"`
	CreatedBy string `gorm:"created_by,omitempty" json:"created_by"`
	UpdatedBy string `gorm:"updated_by,omitempty" json:"updated_by"`
	DeletedBy string `gorm:"deleted_by,omitempty" json:"deleted_by"`
}

type NewsCreateDTO struct {
	TagsID      uint64 `gorm:"tags_id" json:"tags_id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Name        string `json:"name" form:"name" binding:"required"`
	CreatedBy   string `gorm:"created_by,omitempty" json:"created_by"`
	UpdatedBy   string `gorm:"updated_by,omitempty" json:"updated_by"`
	DeletedBy   string `gorm:"deleted_by,omitempty" json:"deleted_by"`
}
