package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type Payment struct {
	ID           int       `gorm:"primaryKey"`
	Guid         uuid.UUID `gorm:"type:uuid;not null"`
	CheckoutGuid uuid.UUID `gorm:"type:uuid;not null"`
	PaymentDate  time.Time `gorm:"default:current_timestamp"`
}
