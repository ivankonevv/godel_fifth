package models

import "github.com/google/uuid"

type Image struct {
	ID       uint      `json:"-" gorm:"primaryKey"`
	ImageURL string    `json:"image_url"`
	PostID   uuid.UUID `json:"-"`
}
