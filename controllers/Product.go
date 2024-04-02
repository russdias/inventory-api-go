package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/russelldias98/go-api/initializers"
	"github.com/russelldias98/go-api/models"
	"github.com/russelldias98/go-api/utils"
)

func GetProducts(c *gin.Context) {
	var products []models.Product

	results := initializers.DB.Find(&products)

	if results.Error != nil {
		utils.FormatErrorAndSend(c, "No products found!")
		return
	}

	utils.FormatJSONAndSend(c, products)
}

func GetProductByID(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	results := initializers.DB.Where("id = ?", id).First(&product)

	if results.Error != nil {
		utils.FormatErrorAndSend(c, "No product found!")
		return
	}

	utils.FormatJSONAndSend(c, product)
}

type ProductDeleteBody struct {
	SupplierID uint `json:"supplier_id"`
	CategoryID uint `json:"category_id"`
}

func DeleteProductByID(c *gin.Context) {
	id := c.Param("id")

	var body ProductDeleteBody

	if err := c.BindJSON(&body); err != nil {
		utils.FormatErrorAndSend(c, "Request body is invalid!")
		return
	}

	results := initializers.DB.Where("id = ? AND supplier_id = ? AND category_id = ?", id, body.SupplierID, body.CategoryID).Delete(&models.Product{})

	if results.Error != nil {
		utils.FormatErrorAndSend(c, "No product found!")
		return
	}

	utils.FormatJSONAndSend(c, "Product deleted successfully!")
}

type ProductBody struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	SupplierID  uint    `json:"supplier_id"`
	CategoryID  uint    `json:"category_id"`
}

func CreateProduct(c *gin.Context) {
	var body ProductBody

	if err := c.BindJSON(&body); err != nil {
		utils.FormatErrorAndSend(c, "Request body is invalid!")
		return
	}

	results := initializers.DB.Create(&models.Product{
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		Quantity:    body.Quantity,
		SupplierID:  body.SupplierID,
		CategoryID:  body.CategoryID,
	})

	if results.Error != nil {
		utils.FormatErrorAndSend(c, "Error creating product!")
		return
	}

	utils.FormatJSONAndSend(c, "Product created successfully!")
}

func UpdateProductByID(c *gin.Context) {
	id := c.Param("id")
	var body ProductBody

	if err := c.BindJSON(&body); err != nil {
		utils.FormatErrorAndSend(c, "Request body is invalid!")
		return
	}

	results := initializers.DB.Model(&models.Product{}).Where("id = ? AND supplier_id", id, body.SupplierID).Updates(&models.Product{
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		Quantity:    body.Quantity,
		SupplierID:  body.SupplierID,
		CategoryID:  body.CategoryID,
	})

	if results.Error != nil {
		utils.FormatErrorAndSend(c, "Error updating product!")
		return
	}

	utils.FormatJSONAndSend(c, "Product updated successfully!")
}
