package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/russelldias98/go-api/initializers"
	"github.com/russelldias98/go-api/models"
	"github.com/russelldias98/go-api/utils"
)

func GetCategories(c *gin.Context) {
	var categories []models.Category

	results := initializers.DB.Find(&categories)

	if results.Error != nil {
		utils.FormatErrorAndSend(c, "No categories found!")
		return
	}

	utils.FormatJSONAndSend(c, categories)
}

func GetCategoryByID(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	results := initializers.DB.Where("id = ?", id).First(&category)

	if results.Error != nil {
		utils.FormatErrorAndSend(c, "No category found!")
		return
	}

	utils.FormatJSONAndSend(c, category)
}

func DeleteCategoryByID(c *gin.Context) {
	id := c.Param("id")

	results := initializers.DB.Where("id = ?", id).Delete(&models.Category{})

	if results.Error != nil {
		utils.FormatErrorAndSend(c, "No category found!")
		return
	}

	utils.FormatJSONAndSend(c, "Category deleted successfully!")
}

type CategoryBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateCategory(c *gin.Context) {
	var body CategoryBody

	if err := c.BindJSON(&body); err != nil {
		utils.FormatErrorAndSend(c, "Request body is invalid!")
		return
	}

	results := initializers.DB.Create(&models.Category{
		Name:        body.Name,
		Description: body.Description,
	})

	if results.Error != nil {
		utils.FormatErrorAndSend(c, "Error creating category!")
		return
	}

	utils.FormatJSONAndSend(c, "Category created successfully!")
}

func UpdateCategoryByID(c *gin.Context) {
	id := c.Param("id")
	var body CategoryBody

	if err := c.BindJSON(&body); err != nil {
		utils.FormatErrorAndSend(c, "Request body is invalid!")
		return
	}

	results := initializers.DB.Model(&models.Category{}).Where("id = ?", id).Updates(&models.Category{
		Name:        body.Name,
		Description: body.Description,
	})

	if results.Error != nil {
		utils.FormatErrorAndSend(c, "Error updating category!")
		return
	}

	utils.FormatJSONAndSend(c, "Category updated successfully!")
}
