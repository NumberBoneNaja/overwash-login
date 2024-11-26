package entity

import (
	"time"

	"gorm.io/gorm"
)
type  HistoryReward struct {
	gorm.Model
	Expire time.Time 
    
    RewardID uint    // fk 
	Reward Reward `gorm:"foreignKey:RewardID"`

	UserID uint `json:"user_id"`  //fk
	User User `gorm:"foreignKey:UserID"`
}