package entity

import (
	"time"

	"gorm.io/gorm"
)
type Employee struct{
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Age uint `json:"age"`
	Birthday time.Time `json:"birthday"`
	Profile string `json:"profile"`
	PhoneNumber string `json:"phone_number"`
	Address string `json:"address"`
	Position string `json:"position"`

}