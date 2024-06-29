package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID   int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Guid uuid.UUID `json:"guid" gorm:"type:uuid;not null"`
	Name string    `json:"name" gorm:"not null"`
	gorm.Model
}
