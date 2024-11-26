package entity

import "gorm.io/gorm"

type BookDetail struct {
	gorm.Model
	Price uint 
    
	MachineID uint 
	Machine Machine `gorm:"foreignKey:MachineID"`

	BookID uint `json:"book_id"`
	Book Book `gorm:"foreignKey:BookID"`
}