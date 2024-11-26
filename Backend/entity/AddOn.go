package entity

import "gorm.io/gorm"

type AddOn struct {
	gorm.Model
	AddOnName string 
	Price float32 
	Description string 

	PackageID uint 
	Package Package `gorm:"foreignKey:PackageID"`

	AddOnDetails []AddOnDetail `gorm:"foreignKey:AddOnID"`
}