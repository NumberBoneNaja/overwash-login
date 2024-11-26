package entity

import "gorm.io/gorm"

type AddOnDetail struct {
	gorm.Model
	Price float32 

	AddOnID uint 
	AddOn AddOn `gorm:"foreignKey:AddOnID"`

	OrderID uint 
	Order Order `gorm:"foreignKey:OrderID"`
}