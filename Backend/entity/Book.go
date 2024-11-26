package entity

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	BookStart time.Time 
	BookEnd time.Time  
	Price uint 

    UserID uint 
	User User `gorm:"foreignKey:UserID"`

	BookDetails []BookDetail `gorm:"foreignKey:BookID"`
	Payments []Payment `gorm:"foreignKey:BookID"`

} 