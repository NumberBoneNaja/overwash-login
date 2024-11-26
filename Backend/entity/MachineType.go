package entity

import "gorm.io/gorm"

type MachineType struct {
	gorm.Model
	Name string
	Price float32 

	Machines []Machine `gorm:"foreignKey:MachineTypeID"`
} 