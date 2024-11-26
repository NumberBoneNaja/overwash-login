package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string 
	Password string 
	Email string `gorm:"unique"`
	Age uint 
	Birthday time.Time 
	ProfilePath string  //เอาไว้เก็บรูป profile
	PhoneNumber string 
	Address string 
	
	GenderID uint // fk กับตาราง gender 1 to 1
	Gender Gender `gorm:"foreignKey:GenderID"`
    
	// MemberShips []MemberShip `gorm:"foreignKey:UserID"` //ส่ง UserID เป็น FK ในตารางmembership

	Notifications []Notification `gorm:"foreignKey:UserID"` //ส่ง UserID เป็น FK ในตารางnotification
    
	// HistoryRewards []HistoryReward `gorm:"foreignKey:UserID"` // UserID เป็น FK ในตาราง HistoryReward

	// Reports []Report `gorm:"foreignKey:UserID"` // ส่ง  UserID เป็น FK ในตาราง Report

	// Books []Book `gorm:"foreignKey:UserID"` // ส่ง  UserID เป็น FK ในตาราง Book

	Payments []Payment `gorm:"foreignKey:UserID"` // ส่ง  UserID เป็น FK ในตาราง Payment

	Orders []Order `gorm:"foreignKey:UserID"` // ส่ง  UserID เป็น FK ในตาราง Order



	 


	
}