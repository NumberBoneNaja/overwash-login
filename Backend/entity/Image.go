package entity

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	ImagePath string

	OrderID uint 
	Order Order `gorm:"foreignKey:OrderID"`
}