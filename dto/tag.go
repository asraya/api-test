package dto

//TagUpdateDTO is a model that client use when updating a tag
type TagUpdateDTO struct {
	ID        uint64 `json:"id" form:"id" binding:"required"`
	Name      string `json:"name" form:"name" binding:"required"`
	CreatedBy string `gorm:"created_by,omitempty" json:"created_by"`
	UpdatedBy string `gorm:"updated_by,omitempty" json:"updated_by"`
	DeletedBy string `gorm:"deleted_by,omitempty" json:"deleted_by"`
}

//TagCreateDTO is is a model that clinet use when create a new tag
type TagCreateDTO struct {
	NewsId int    `json:"news_id" form:"news_id" binding:"required"`
	Name   string `json:"name" form:"name" binding:"required"`
}
