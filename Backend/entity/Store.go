package entity
import "gorm.io/gorm"
type Store struct {
	gorm.Model
	Brand string 
	Stock uint 
	//TypeLundryProduct
	
	TypeLaundryProductID uint 
	TypeLaundryProduct TypeLaundryProduct `gorm:"foreignKey:TypeLaundryProductID"`
	
	
}