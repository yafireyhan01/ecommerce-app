package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Guid         uuid.UUID `json:"guid" gorm:"type:uuid;not null"`
	CategoryGuid uuid.UUID `json:"category_guid" gorm:"type:uuid;not null"`
	Name         string    `json:"name" gorm:"not null"`
	Description  string    `json:"description"`
	Price        float64   `json:"price" gorm:"type:decimal(10,2);not null"`
	StockQty     int       `json:"stock_qty" gorm:"not null"`
	gorm.Model
}
