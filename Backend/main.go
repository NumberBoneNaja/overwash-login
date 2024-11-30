// เขียน hello world ด้วย go
package main

import (
	// "fmt"
	"four/config"
	"four/controller"
	"log"

	// "os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/joho/godotenv"
	
)

func main() {

	config.ConnectionDB()
    config.SetupDatabase()
    app := fiber.New()
	
	app.Use(middleware())
	app.Post("/mutipleupload/:id", controller.UploadMultipleImages)
    app.Post("/login", controller.Login)
	app.Post("/logout",controller.Logout)
	app.Static("/images/ProfOrder", "./images/ProfOrder")
	rounter := app.Group("/")
	{
		rounter.Use(controller.AuthRequired)
	    rounter.Get("/h", helloHandler)

		// Order
		rounter.Post("/CreateOrder",controller.CreateOrder)
		rounter.Post("/CreateOrderDetail",controller.CreateOrderDetail)
		rounter.Post("/CreateAddOnDetail",controller.CreateAddOnDetail)
		rounter.Patch("/UpdateOrder/:id",controller.UpdateStatusOrder)
        rounter.Delete("/DeleteOrderByID/:id",controller.DeleteOrderByID)
       // Package
		rounter.Get("/GetAllPackage",controller.GetAllPackage)
		rounter.Get("/GetClothByPackageID/:id",controller.GetClothByPackageID)
		rounter.Get("/GetAddOnByPackageID/:id",controller.GetAddOnByPackageID)

		//Image
		rounter.Post("/UploadImage/:id",controller.UploadImage)
		rounter.Get("/GetImagesByOrderID/:id",controller.GetImagesByOrderID)
		rounter.Delete("/DeleteImage/:id",controller.DeleteImageByID)
		rounter.Post("/mutipleupload/:id", controller.UploadMultipleImages)

		//Create package and add on
		rounter.Post("/CreatePackage",controller.CreatePackage)
		rounter.Post("/CreateAddOn",controller.CreateAddOn)
		rounter.Post("/CreateCloth",controller.CreateClothType)
		
	}
	
	
	log.Fatal(app.Listen(":8000"))
}
func helloHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Hello, World!"})
  }
  func middleware() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",  // ระบุ URL ที่อนุญาตให้เข้าถึง
		AllowMethods:     "GET,POST,PUT,DELETE",    // ระบุ HTTP methods ที่อนุญาต
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization", // เพิ่ม Authorization ที่นี่
		AllowCredentials: true,                      // อนุญาตให้ส่ง Cookie
	})
}
