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
	app.Use(middleware)
    app.Post("/login", controller.Login)
	app.Post("/logout",controller.Logout)
	
	rounter := app.Group("/")
	{
		rounter.Use(controller.AuthRequired)
	    rounter.Get("/h", helloHandler)
		
	}
	
	
	log.Fatal(app.Listen(":8000"))
}
func helloHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Hello, World!"})
  }
func middleware(c *fiber.Ctx) error {
       corsMiddleware := cors.New(cors.Config{
        AllowOrigins: "http://localhost:5173", 
        AllowMethods: "GET,POST,PUT,DELETE",   
        AllowHeaders: "Origin, Content-Type, Accept",
        AllowCredentials: true, // อนุญาตให้ส่ง Cookie
    })
    return corsMiddleware(c)
}


