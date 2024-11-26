package controller

import (
	"fmt"
	"four/config"
	"four/entity"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateOrder(c *fiber.Ctx) error {
    
	var Orders entity.Order
    db := config.DB()

	if err := c.BodyParser(&Orders); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error", "message": "Review your input", "data":  err.Error()})
	}
    var packages entity.Package
	db.Where("id = ?", Orders.PackageID).First(&packages)
	if packages.ID == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}
	var users entity.User
	db.Where("id = ?", Orders.UserID).First(&users)
	if users.ID == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}
	err := db.Create(&Orders).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error", "message": "Review your input", "data":  err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":   Orders,
		
	})
}

func CreateOrderDetail(c *fiber.Ctx) error {
	var OrderDetail entity.OrderDetail
	db := config.DB()
	if err := c.BodyParser(&OrderDetail); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error", "message": "Review your input", "data":  err.Error()})
	}
    var orders entity.Order
	db.Where("id = ?", OrderDetail.OrderID).First(&orders)
	if orders.ID == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}
	var cloths entity.ClothType
	db.Where("id = ?", OrderDetail.ClothTypeID).First(&cloths)
	if cloths.ID == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}


	err := db.Create(&OrderDetail).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error", "message": "Review your input", "data":  err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",	
		"data":   OrderDetail,
	})
}
func CreateAddOnDetail(c *fiber.Ctx) error {
	var AddOnDetail entity.AddOnDetail
    db := config.DB()
	 if err := c.BodyParser(&AddOnDetail); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	 }
	 err := db.Create(&AddOnDetail).Error
	 if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	 }
	 return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",	
		"data":   AddOnDetail,
	})

}

func GetAllPackage(c *fiber.Ctx) error {

	db := config.DB()
	var Packages []entity.Package
	err := db.Find(&Packages).Error
	if err != nil{
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(Packages)

}

func GetAddOnByPackageID(c *fiber.Ctx) error {
	var AddOn []entity.AddOn
	
	db := config.DB()
	if err := db.Where("package_id = ?", c.Params("id")).Find(&AddOn).Error; err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(AddOn)
}
func GetClothByPackageID(c *fiber.Ctx) error {
	id,err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}


	var Cloth []entity.ClothType
	db := config.DB()
	if err := db.Where("package_id = ?", uint(id)).Find(&Cloth).Error; err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(Cloth)
}



 func UpdateStatusOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	db := config.DB()
	var order entity.Order
	check :=db.Where("id = ?", id).First(&order).Error
    if check != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

    if err := c.BodyParser(&order); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	

	// db.Save(&Order)
	db.Model(&order).Select("status").Updates(entity.Order{Status: order.Status})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Updatesuccess",

		"data" : order,
	})
 }

 func DeleteOrderByID(c *fiber.Ctx)  error{
	db := config.DB()
	id := c.Params("id")

	err:=db.Where("id = ?", id).Delete(&entity.Order{}).RowsAffected
	if err == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error", "message": "not found",
		})
	}
	return c.SendStatus(fiber.StatusOK)
	
 }
 func UploadImage(c *fiber.Ctx ) error {
	file,err := c.FormFile("file")
	db := config.DB()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	uploadDir := "./images/ProfOrder"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
   
	err = c.SaveFile(file,"./images/ProfOrder/"+file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "upload file failed",
			"error": err.Error(),
		})
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	Image := entity.Image{
		ImagePath: file.Filename,
		OrderID: uint(id),
	}
	result := db.Create(&Image)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "upload file failed",
			
		})
	}
	return c.JSON(fiber.Map{
        "message": "Image uploaded successfully",
        "image":   Image,
    })
 }

 func GetImagesByOrderID(c *fiber.Ctx) error {
	var images []entity.Image
	db := config.DB()
	result := db.Where("order_id = ?", c.Params("id")).Find(&images)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	for i := range images {
		images[i].ImagePath = fmt.Sprintf("http://localhost:8000/images/ProfOrder/%s", images[i].ImagePath)
	}

	return c.JSON(images)
}
func DeleteImageByID(c *fiber.Ctx) error {
    // extract image name from params
    id := c.Params("id")

	db := config.DB()
	var image entity.Image
	check :=db.Where("id = ?", id).First(&image).Error
	if check != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

    // delete image from ./images
    err := os.Remove(fmt.Sprintf("./images/ProfOrder/%s", image.ImagePath))
    if err != nil {
        log.Println(err)
        return c.JSON(fiber.Map{"status": 500, "message": "Server Error", "data": nil})
    }

	// delete image from database
	if err := db.Where("id = ?", id).Delete(&entity.Image{}); err.Error != nil {
	
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "delete image failed",
			// "error":   err.Error(),
		})
	}

    return c.JSON(fiber.Map{"status": 201, "message": "Image deleted successfully", "data": nil})
}
func UploadMultipleImages(c *fiber.Ctx) error {
	db := config.DB()
	forms, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ 
			"error" : err.Error()})
	}
	files := forms.File["images"]
    if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "No files uploaded",
		})
	}
	uploadDir := "./images/ProfOrder"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	for _, file := range files {
	err = c.SaveFile(file,"./images/ProfOrder/"+file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "upload file failed",
			"error": err.Error(),
		})
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	Image := entity.Image{
		ImagePath: file.Filename,
		OrderID: uint(id),
	}
	result := db.Create(&Image)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "upload file failed",
			
		})
	}
		
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Images uploaded successfully"})
}