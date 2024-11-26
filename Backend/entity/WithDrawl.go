package entity
import "gorm.io/gorm"
type WithDrawal struct {
	gorm.Model
	EmpFirstName string
	EmpLastName string 
    LaundryProduct string 

	EmployeeID uint
	Employee Employees `gorm:"foreignKey:EmployeeID"`

	StoreID uint 
	Store Store `gorm:"foreignKey:StoreID"`        
}