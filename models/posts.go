package models

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	PostID      uuid.UUID `json:"post_id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"notnull" validate:"required,max=200"`
	Description string    `json:"description" validate:"max=1000"`
	Price       float32   `json:"price"`
	Images      []Image   `json:"images" gorm:"foreignKey:PostID" validate:"max=3"`
	CreatedAt   time.Time `json:"created_at"`
}
