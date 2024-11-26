package entity

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	
	UserID uint 
	User User `gorm:"foreignKey:UserID"`
	
	MachineID uint 
	Machine Machine `gorm:"foreignKey:MachineID"`

	ReportDetails []ReportDetail `gorm:"foreignKey:ReportID"`
	RPResualts []RPResualt `gorm:"foreignKey:ReportID"`
} 