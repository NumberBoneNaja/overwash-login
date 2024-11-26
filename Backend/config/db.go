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
	// db.AutoMigrate(&entity.User{}, &entity.Employees{},&entity.AddOn{},&entity.AddOnDetail{},&entity.Attendance{},
	// entity.Book{},&entity.BookDetail{},&entity.Package{},entity.ClothType{},entity.EmployeeRole{},&entity.Gender{},&entity.HistoryReward{},
    // &entity.Image{},&entity.Machine{},&entity.MachineStatus{},&entity.MachineType{},&entity.MemberShip{},
    // &entity.Notification{},&entity.Order{},&entity.OrderDetail{},&entity.Payment{},&entity.Report{},
    //  &entity.RPResualt{},&entity.Reward{},&entity.Store{},&entity.TypeLaundryProduct{},&entity.WithDrawal{},&entity.ReportDetail{})
	db.AutoMigrate(&entity.Gender{},&entity.User{}, &entity.Employees{},&entity.Package{},&entity.ClothType{},&entity.AddOn{},&entity.Order{},
		&entity.OrderDetail{},&entity.AddOnDetail{},&entity.Image{},&entity.Payment{},&entity.Notification{})
    hashedpassword,_ := HashPassword("1234")
	BirthDay ,_ := time.Parse("2006-01-02", "1988-11-12") 


	GenderMale := entity.Gender{Name: "Male"}
	GenderFemale := entity.Gender{Name: "Female"}

	db.FirstOrCreate(&GenderMale, &entity.Gender{Name: "Male"})
	db.FirstOrCreate(&GenderFemale, &entity.Gender{Name: "Female"})

     PackageFirst := entity.Package{PackageName: "ซักผับ", Explain: "เหมาะสำหรับเสื้อลำลองในชีวิตประจำวัน"}
	 PackageSecond := entity.Package{PackageName: "ซักรองเท้า", Explain: "ดูแลรองเท้าอย่างพิถีพิถัน ทุกประเภท"}
     PackageThird := entity.Package{PackageName: "ซักเครื่องนอนและอื่น", Explain: "ซักอบเครื่องนอน ผ้าที่ใช้ในบ้านและห้องน้ำ"}
	// Create packages
	db.FirstOrCreate(&PackageFirst,&entity.Package{PackageName: "ซักผับ", Explain: "เหมาะสำหรับเสื้อลำลองในชีวิตประจำวัน"})
	db.FirstOrCreate(&PackageSecond,&entity.Package{PackageName: "ซักรองเท้า", Explain: "ดูแลรองเท้าอย่างพิถีพิถัน ทุกประเภท"})
	db.FirstOrCreate(&PackageThird,&entity.Package{PackageName: "ซักเครื่องนอนและอื่น", Explain: "ซักอบเครื่องนอน ผ้าที่ใช้ในบ้านและห้องน้ำ"})

	ClothypeFirst := entity.ClothType{TypeName: "เสื้อสำลอง", Price: 10.00, PackageID: 1}
	ClothypeSecond := entity.ClothType{TypeName: "รองเท้า", Price: 485.00, PackageID: 2}
	ClothypeThird := entity.ClothType{TypeName: "เครื่องนอนไม่หนา", Price: 100.00, PackageID: 3}
	ClothypeFourth := entity.ClothType{TypeName: "เครื่องนอนหนา", Price: 200.00, PackageID: 3}
	// // Create cloth types
	db.FirstOrCreate(&ClothypeFirst,&entity.ClothType{TypeName: "เสื้อสำลอง", Price: 10.00, PackageID: 1})
	db.FirstOrCreate(&ClothypeSecond,&entity.ClothType{TypeName: "รองเท้า", Price: 485.00, PackageID: 2})
	db.FirstOrCreate(&ClothypeThird,&entity.ClothType{TypeName: "เครื่องนอนไม่หนา", Price: 100.00, PackageID: 3})
	db.FirstOrCreate(&ClothypeFourth,&entity.ClothType{TypeName: "เครื่องนอนหนา", Price: 200.00, PackageID: 3})

	AddOnFirst := entity.AddOn{AddOnName: "ขจัดคราบ", Price: 10.00, Description: "ขจัดคราบล้ำลึก", PackageID: 1}
	AddOnSecond := entity.AddOn{AddOnName: "น้ำหอมร้องเท้า", Price: 10.00, Description: "น้ำหอมร้องเท้า", PackageID: 2}
	AddOnThird := entity.AddOn{AddOnName: "ซักอุณภูมิสูง 60 องศา", Price: 20.00, Description: "ซักอุณภูมิสูง เพื่อขจัดแบคทีเรีย", PackageID: 3}
	AddOnFourth := entity.AddOn{AddOnName: "ปกป้องรองเท้า", Price: 20.00, Description: "เคลือบน้ำยากันน้ำ ความชื้น และคราบสกปรก", PackageID: 2}

	// // Create add-ons
	db.FirstOrCreate(&AddOnFirst,&entity.AddOn{AddOnName: "ขจัดคราบ", Price: 10.00, Description: "ขจัดคราบล้ำลึก", PackageID: 1})
	db.FirstOrCreate(&AddOnSecond,&entity.AddOn{AddOnName: "น้ำหอมร้องเท้า", Price: 10.00, Description: "น้ำหอมร้องเท้า", PackageID: 2})
	db.FirstOrCreate(&AddOnThird,&entity.AddOn{AddOnName: "ซักอุณภูมิสูง 60 องศา", Price: 20.00, Description: "ซักอุณภูมิสูง เพื่อขจัดแบคทีเรีย", PackageID: 3})
	db.FirstOrCreate(&AddOnFourth,&entity.AddOn{AddOnName: "ปกป้องรองเท้า", Price: 20.00, Description: "เคลือบน้ำยากันน้ำ ความชื้น และคราบสกปรก", PackageID: 2})

	// // Create orders (ensure UserID and PackageID exist)
	// db.FirstOrCreate(&entity.Order{Price: 100.00, UserID: 1, Status: "รอชําระเงิน", PackageID: 1})





	db.FirstOrCreate(&entity.User{
		Username: "Member", 
		Email: "Member@gmail.com", 
		Password: hashedpassword, 
		Age: 20,
		Birthday: BirthDay,
		ProfilePath: "Member", 
		PhoneNumber: "123456789", 
		Address: "123 Main St",
	    GenderID: 1,} )


	db.FirstOrCreate(&entity.Employees{
		Firstname: "admin", 
		Lastname: "admin", 
		Email: "admin@gmail.com", 
		Password: hashedpassword, 
		Address: "123 Main St",
		PhoneNumber: "123456789", 
		// EmployeeRoleID: 1,
		
	})
	println("create attribute successfully")

}

// Use format time

