package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type Checkout struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Guid       uuid.UUID `json:"guid" gorm:"type:uuid;not null"`
	UserGuid   uuid.UUID `json:"user_guid" gorm:"type:uuid;not null"`
	CartGuid   uuid.UUID `json:"cart_guid" gorm:"type:uuid;not null"`
	OrderDate  time.Time `json:"order_date"`
	TotalPrice float64   `json:"total_price" gorm:"type:decimal(10,2);not null"`
	Status     string    `json:"status" gorm:"type:enum('PENDING', 'PAID');not null"`
	gorm.Model
}

//type Cart struct {
//	ID           int       `json:"id" gorm:"primaryKey;autoIncrement"`
//	Guid         uuid.UUID `json:"guid" gorm:"type:uuid;not null"`
//	UserGuid     uuid.UUID `json:"user_guid" gorm:"type:uuid;not null"`
//	ProductGuid  uuid.UUID `json:"product_guid" gorm:"type:uuid;not null"`
//	Qty          int       `json:"qty" gorm:"not null"`
//	Product      Product   `gorm:"foreignKey:ProductGuid"`
//}
//
//type Product struct {
//	ID           int       `json:"id" gorm:"primaryKey;autoIncrement"`
//	Guid         uuid.UUID `json:"guid" gorm:"type:uuid;not null"`
//	CategoryGuid uuid.UUID `json:"category_guid" gorm:"type:uuid;not null"`
//	Name         string    `json:"name" gorm:"not null"`
//	Description  string    `json:"description"`
//	Price        float64   `json:"price" gorm:"type:decimal(10,2);not null"`
//	StockQty     int       `json:"stock_qty" gorm:"not null"`
//}
