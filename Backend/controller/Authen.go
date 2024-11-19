package controller

import (
	
	"four/config"
	"four/entity"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"

	"golang.org/x/crypto/bcrypt"
)
func Login(c *fiber.Ctx) error {
	user := entity.Member{}
	inputadmin := entity.Employee{}
	inputexit := entity.Member{}
	inputadminexit := entity.Employee{}
	db := config.DB()

	err := c.BodyParser(&user)
	if err != nil {
   return	c.Status(500).JSON(fiber.Map{
	"status": "error", 
	"message": "Review your input", 
	"data": err})
	   
	}
	err = c.BodyParser(&inputadmin)
	if err != nil {
   return	c.Status(500).JSON(fiber.Map{
	"status": "error", 
	"message": "Review your input", 
	"data": err})
	   
	}
	err = db.Where("email = ?", user.Email).First(&inputexit).Error
	
	if err != nil {
	   err2 := db.Where("email = ?", inputadmin.Email).First(&inputadminexit).Error
	   if err2 != nil {
		   return c.SendStatus(http.StatusUnauthorized)
	   }
	   err2 = bcrypt.CompareHashAndPassword([]byte(inputadminexit.Password), []byte(inputadmin.Password))
	   if err2 != nil {
		   return c.SendStatus(http.StatusUnauthorized)
	   }
	    

	   claims := jwt.MapClaims{
		   "name":  inputadminexit.Firstname,
		   "id": inputadminexit.ID,
		   "role":  inputadminexit.Position,
		   "exp":   time.Now().Add(time.Minute * 1).Unix(),
	   }
   
	   // Create token
	   token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
   
	   //use load env file
	   LoadENV()
	  t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	   if err != nil {
		   return c.SendStatus(http.StatusUnauthorized)
	   }
	   c.Cookie(&fiber.Cookie{
		   Name:     "jwt",
		   Value:    t,
		   Expires:  time.Now().Add(time.Minute * 1),
		   // HTTPOnly: true,
	   })
   
	   return c.JSON(fiber.Map{
		"status": "success",
		"message": "login success Employee",
	    "position": "admin",
	     "token": t,})
	   
	}

	//compare password
   err = bcrypt.CompareHashAndPassword([]byte(inputexit.Password), []byte(user.Password))
   if err != nil {
   return	c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid email or password", "data": err})
	   
   }else {
	   claims := jwt.MapClaims{
		   "name":  inputexit.Username,
		   "id": inputexit.ID,
		   "role" :"user",
		   "exp":   time.Now().Add(time.Minute * 1).Unix(),
	   }
   
	   // Create token
	   token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
   
	   // Generate encoded token and send it as response.
	   t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	   if err != nil {
		   return c.SendStatus(http.StatusUnauthorized)
	   }
	   c.Cookie(&fiber.Cookie{
		   Name:     "jwt",
		   Value:    t,
		   Expires:  time.Now().Add(time.Minute * 1),
		   // HTTPOnly: true,
	   })
   

   return	c.Status(200).JSON(fiber.Map{"status": "success", "message": "success hi user", "token": t,"position": "user"})
	   
   }
}
func LoadENV() {
    env_err := godotenv.Load()
	   if env_err != nil {
	   log.Fatal("Error loading .env file")
	   }
}
func AuthRequired(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
    LoadENV()
	token, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
      return []byte(os.Getenv("SECRET_KEY")), nil
  })

  if err != nil || !token.Valid {
      return c.SendStatus(fiber.StatusUnauthorized)
  }
  
  return c.Next()
}
func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Minute),
		HTTPOnly: true,
	})
	return c.JSON(fiber.Map{"status": "success", "message": "Logged out successfully"})
}