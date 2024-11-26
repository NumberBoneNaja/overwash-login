package entity

import "gorm.io/gorm"

type MemberShip struct {
	gorm.Model
	Point uint 
	Status string 
	MemberShipStart uint 
	MemberShipEnd uint 

	PackageMemberShipTypeID uint 
	PackageMemberShipType PackageMemberShipType `gorm:"foreignKey:PackageMemberShipTypeID"`

	UserID uint
	User User `gorm:"foreignKey:UserID"`
}