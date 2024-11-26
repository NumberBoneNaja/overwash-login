package entity

import "gorm.io/gorm"

type Gender struct {
	gorm.Model
	Name string 

	// 1 gender เชื่อมได้หลาย User
	Users []User `gorm:"foreignKey:GenderID"` //ส่ง fk ตาราง user
	
}