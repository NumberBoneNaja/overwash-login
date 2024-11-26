//ตัวนี้ใช้จริง
package entity

import (
	"gorm.io/gorm"
)

type Employees struct {
	gorm.Model
    Firstname string 
	Lastname string 
	Password string 
	Email string `gorm:"unique"`
	Address string 
	PhoneNumber string 
	
	// EmployeeRoleID uint
	// EmployeeRole EmployeeRole `gorm:"foreignKey:EmployeeRoleID"`

	

	// RPResualts []RPResualt `gorm:"foreignKey:EmployeeID"`
	// WithDrawals []WithDrawal `gorm:"foreignKey:EmployeeID"`
	// Attendances []Attendance `gorm:"foreignKey:EmployeeID"`
}