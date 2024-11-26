package entity

import "gorm.io/gorm"

type Package struct {
    gorm.Model
    PackageName string 
    Explain     string 

    Orders     []Order     `gorm:"foreignKey:PackageID"`   
    AddOns     []AddOn     `gorm:"foreignKey:PackageID"`   
    ClothTypes []ClothType `gorm:"foreignKey:PackageID"`   
}