package entity

import "gorm.io/gorm"


type OrderDetail struct {
	gorm.Model
	Quantity    uint    
	Price       float32  

	OrderID     uint     
	Order       Order     `gorm:"foreignKey:OrderID"`
	
	ClothTypeID uint     
	ClothType   ClothType `gorm:"foreignKey:ClothTypeID"`
}