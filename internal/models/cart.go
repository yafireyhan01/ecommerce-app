package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	ID          uint      `gorm:"primaryKey"`
	Guid        uuid.UUID `gorm:"type:uuid;not null;unique"`
	UserGuid    uuid.UUID `gorm:"type:uuid;not null"`
	ProductGuid uuid.UUID `gorm:"type:uuid;not null"`
	Qty         int       `gorm:"not null"`
	gorm.Model
}
