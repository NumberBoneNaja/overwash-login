package entity
import "gorm.io/gorm"
type RPResualt struct {
	gorm.Model
	Status string 

	ReportID uint 
	Report Report `gorm:"foreignKey:ReportID"`
 
    EmployeeID uint 
	Employees Employees `gorm:"foreignKey:EmployeeID"`
	
}