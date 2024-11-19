package config

import (
	"fmt"
	"four/entity"
	"time"

	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)


var db *gorm.DB


func DB() *gorm.DB {

   return db

}


func ConnectionDB() {

   database, err := gorm.Open(sqlite.Open("Overwash.db?cache=shared"), &gorm.Config{})

   if err != nil {

       panic("failed to connect database")

   }

   fmt.Println("connected database")

   db = database

}

func SetupDatabase()  {
	db.AutoMigrate(&entity.Member{}, &entity.Employee{})
    hashedpassword,_ := HashPassword("1234")
	BirthDay ,_ := time.Parse("2006-01-02", "1988-11-12") 
	db.Create(&entity.Member{
		Username: "Member", 
		Email: "Member@gmail.com", 
		Password: hashedpassword, 
		Age: 20,
		Birthday: BirthDay,
		 Profile: "admin", 
		 PhoneNumber: "123456789", 
		 Address: "123 Main St"})


	db.Create(&entity.Employee{
		Firstname: "admin", 
		Lastname: "admin", 
		Email: "admin@gmail.com", 
		Password: hashedpassword, 
		Age: 20,
		Birthday: BirthDay,
		 Profile: "admin", 
		 PhoneNumber: "123456789", 
		 Address: "123 Main St",
		 Position: "admin",
	})
	println("create attribute successfully")

}

// Use format time

