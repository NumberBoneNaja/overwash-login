package entity
import "gorm.io/gorm"
type ClothType struct {
    gorm.Model
    TypeName     string     
    Price        float32  
     
    PackageID    uint       
    Package      Package    `gorm:"foreignKey:PackageID" `
	
    OrderDetails []OrderDetail `gorm:"foreignKey:ClothTypeID" `
}