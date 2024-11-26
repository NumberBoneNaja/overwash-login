package entity

import (
	"time"

	"gorm.io/gorm"
)

type EmployeeRole struct {
	gorm.Model
	Role string 
	Salary float32 
	DateStart time.Time 

	Employees []Employees `gorm:"foreignKey:EmployeeRoleID"`
}