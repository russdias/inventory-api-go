package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/russelldias98/go-api/initializers"
	"github.com/russelldias98/go-api/models"
	"github.com/russelldias98/go-api/utils"
)

func CheckStockForProduct(c *gin.Context) {
	var product models.Product

	supplierId, productId := c.Param("supplierId"), c.Param("productId")

	result := initializers.DB.Joins("Supplier").Where("suppliers.id = ? and products.id = ?", supplierId, productId).First(&product)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(200, gin.H{"stock": product.Quantity})
}

func ListOutOfStockProducts(c *gin.Context) {
	var products []models.Product

	result := initializers.DB.Where("quantity = ?", 0).Find(&products)

	if result.Error != nil {
		utils.FormatErrorAndSend(c, "No products found!")
		return
	}

	utils.FormatJSONAndSend(c, products)
}

func ListLowStockProducts(c *gin.Context) {
	t := c.Query("threshold")

	if t == "" {
		utils.FormatErrorAndSend(c, "Threshold not provided!")
		return
	}

	threshold, _ := strconv.Atoi(t)

	var products []models.Product

	result := initializers.DB.Where("quantity < ?", threshold).Find(&products)

	if result.Error != nil {
		utils.FormatErrorAndSend(c, "No products found!")
		return
	}

	utils.FormatJSONAndSend(c, products)
}
