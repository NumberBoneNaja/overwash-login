package entity

import "gorm.io/gorm"

type ReportDetail struct {
	gorm.Model
	Massage string 
	ImagePath string `json:"image_path"`
	
	ReportID uint 
	Report Report `gorm:"foreignKey:ReportID"`
}