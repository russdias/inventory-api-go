package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/russelldias98/go-api/initializers"
	"github.com/russelldias98/go-api/models"
	"github.com/russelldias98/go-api/utils"
)

func GetSuppliers(c *gin.Context) {
	var suppliers []models.Supplier

	results := initializers.DB.Find(&suppliers)

	if results.Error != nil {
		utils.FormatErrorAndSend(c, "No suppliers found!")
		return
	}

	utils.FormatJSONAndSend(c, suppliers)
}

func GetSupplierByID(c *gin.Context) {
	var supplier models.Supplier
	id := c.Param("id")

	results := initializers.DB.Where("id = ?", id).First(&supplier)

	if results.Error != nil {
		utils.FormatErrorAndSend(c, "No supplier found!")
		return
	}

	utils.FormatJSONAndSend(c, supplier)
}

func DeleteSupplierByID(c *gin.Context) {
	id := c.Param("id")

	results := initializers.DB.Where("id = ?", id).Delete(&models.Supplier{})

	if results.Error != nil {
		utils.FormatErrorAndSend(c, "No supplier found!")
		return
	}

	utils.FormatJSONAndSend(c, "Supplier deleted successfully!")
}

type SupplierBody struct {
	Name        string `json:"name"`
	ContactInfo string `json:"contact_info"`
}

func CreateSupplier(c *gin.Context) {
	var body SupplierBody

	if err := c.BindJSON(&body); err != nil {
		utils.FormatErrorAndSend(c, "Request body is invalid!")
		return
	}

	results := initializers.DB.Create(&models.Supplier{
		Name:        body.Name,
		ContactInfo: body.ContactInfo,
	})

	if results.Error != nil {
		utils.FormatErrorAndSend(c, "Error creating supplier!")
		return
	}

	utils.FormatJSONAndSend(c, "Supplier created successfully!")
}

func UpdateSupplierByID(c *gin.Context) {
	id := c.Param("id")
	var body SupplierBody

	if err := c.BindJSON(&body); err != nil {
		utils.FormatErrorAndSend(c, "Request body is invalid!")
		return
	}

	results := initializers.DB.Model(&models.Supplier{}).Where("id = ?", id).Updates(&models.Supplier{
		Name:        body.Name,
		ContactInfo: body.ContactInfo,
	})

	if results.Error != nil {
		utils.FormatErrorAndSend(c, "Error updating supplier!")
		return
	}

	utils.FormatJSONAndSend(c, "Supplier updated successfully!")
}

func GetProductsBySupplierID(c *gin.Context) {
	var products []models.Product
	id := c.Param("supplierId")

	results := initializers.DB.Where("supplier_id = ?", id).Find(&products)

	if results.Error != nil {
		utils.FormatErrorAndSend(c, "No products found!")
		return
	}

	utils.FormatJSONAndSend(c, products)
}
