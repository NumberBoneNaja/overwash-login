package entity

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	PaymentStart time.Time 
	PaymentEnd time.Time  
	Price float32

	// BookID uint 
	// Book Book `gorm:"foreignKey:BookID"`

	OrderID uint 
	Order Order `gorm:"foreignKey:OrderID"`

	UserID uint 
	User User `gorm:"foreignKey:UserID"`
}