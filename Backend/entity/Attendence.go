package entity

import (
	"time"

	"gorm.io/gorm"
)

type Attendance struct {
	gorm.Model
	Date time.Time 
	CheckIn time.Time 
	CheckOut time.Time 
	Status string 

	EmployeeID uint 
	Employee Employees `gorm:"foreignKey:EmployeeID"`  


	
}