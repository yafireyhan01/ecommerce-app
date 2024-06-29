package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       int       `json:"id" gorm:"primaryKey"`
	Guid     uuid.UUID `json:"guid" gorm:"type:uuid;not null"`
	Name     string    `json:"name"`
	Email    string    `json:"email" gorm:"not null;unique"`
	Password string    `json:"-" gorm:"not null"`
	Role     string    `json:"role" gorm:"default:CUSTOMER"`
	gorm.Model
}

//func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
//	user.Guid = generateGUID() // Implement a function to generate a unique Guid
//	return
//}
