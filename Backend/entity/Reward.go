package entity
import (
	"gorm.io/gorm"
)
type Reward struct {
	gorm.Model
	Point uint
	Name string 
	Description float32 
	ImagePath string 

	HistoryRewards []HistoryReward `gorm:"foreignKey:RewardID"`

}