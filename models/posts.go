package models

import (
	"time"

	"gorm.io/gorm"
)


type BaseModel struct {
    ID        uint           `json:"id" gorm:"primaryKey"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" swaggerignore:"true"` 
}

type Post struct {
	BaseModel
	PicAddres   string        `json:"pic_address"`
	Title       string        `json:"title"`
	Caption     string        `json:"caption"`
	PostComment []PostComment `gorm:"constraint:OnDelete:CASCADE;"`
	UserID      uint
}

type PostRegister struct {
	PicAddres string `json:"pic_address"`
	Title     string `json:"title" validate:"max=250"`
	Caption   string `json:"caption" validate:"max=1000"`
}

type PostComment struct {
	BaseModel
	Text   string `json:"text"`
	UserID uint   `gorm:"not null"`
	PostID uint   `gorm:"not null"`
	
}

type PostCommentsRegister struct {
	Text string `json:"text" binding:"max=250,required"`
}
