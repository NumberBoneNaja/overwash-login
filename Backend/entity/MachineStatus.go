package entity

import "gorm.io/gorm"

type MachineStatus struct {
	gorm.Model
	Status string

	Machines []Machine `gorm:"foreignKey:MachineStatusID"`
}