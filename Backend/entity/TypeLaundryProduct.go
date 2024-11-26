package entity

import "gorm.io/gorm"

type TypeLaundryProduct struct {
	gorm.Model
	TypeLundry string 

	Stores []Store `gorm:"foreignKey:TypeLaundryProductID"`
}