package entity

import "gorm.io/gorm"

type PackageMemberShipType struct {
	gorm.Model
	Name string 
	Price float32 
	HowLong uint 
	Explain string 

	MemberShips []MemberShip `gorm:"foreignKey:PackageMemberShipTypeID"`
} 