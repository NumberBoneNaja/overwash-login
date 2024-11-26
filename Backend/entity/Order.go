package entity

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Price float32 
    Status string
	UserID uint
	User User `gorm:"foreignKey:UserID"`
	
	PackageID uint 
	Package Package `gorm:"foreignKey:PackageID"`

	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID"`

	Notifications []Notification `gorm:"foreignKey:OrderID"`

	Images []Image `gorm:"foreignKey:OrderID"`

	Payments []Payment `gorm:"foreignKey:OrderID"`

	AddOnDetails []AddOnDetail `gorm:"foreignKey:OrderID"` 
}