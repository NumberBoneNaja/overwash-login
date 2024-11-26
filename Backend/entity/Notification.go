package entity

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	DateTime time.Time 
	Explain string 

	UserID uint 
	User User `gorm:"foreignKey:UserID"`

	OrderID uint 
	Order Order `gorm:"foreignKey:OrderID"`
}