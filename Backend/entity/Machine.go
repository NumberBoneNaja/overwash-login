package entity

import "gorm.io/gorm"

type Machine  struct {
	gorm.Model
	MachineName string 

	MachineTypeID string 
	MachineType MachineType `gorm:"foreignKey:MachineTypeID"`

	MachineStatusID string 
	MachineStatus MachineStatus `gorm:"foreignKey:MachineStatusID"`

	BookDetails []BookDetail `gorm:"foreignKey:MachineID"`
	Reports []Report `gorm:"foreignKey:MachineID"` 
}