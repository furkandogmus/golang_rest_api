package controller

import (
	"restapi/config"
	"restapi/models"

	"github.com/gin-gonic/gin"
)

// @Summary Get all categories
// @Description Get a list of all categories
// @Produce json
// @Success 200 {array} models.Category
// @Router /category [get]
func GetCategories(c *gin.Context) {
	categories := []models.Category{}
	config.DB.Find(&categories)
	c.JSON(200, &categories)
}

// @Summary Create a new category
// @Description Create a new category
// @Accept json
// @Produce json
// @Param category body models.Category true "Category object to be created"
// @Success 200 {object} models.Category
// @Router /category [post]
func CreateCategory(c *gin.Context) {
	var category models.Category
	c.BindJSON(&category)
	category.ID = 0
	config.DB.Create(&category)
	c.JSON(200, &category)
}

// @Summary Delete a category by ID
// @Description Delete a category by ID
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {string} string "User has been deleted successfully."
// @Router /category/{id} [delete]
func DeleteCategory(c *gin.Context) {
	var category models.Category
	config.DB.Where("id = ?", c.Param("id")).Delete(&category)
	c.JSON(200, "Category has been deleted successfully.")
}

// @Summary Update a category by ID
// @Description Update a category by ID
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body models.Category true "Updated category object"
// @Success 200 {object} models.Category
// @Router /category/{id} [put]
func UpdateCategory(c *gin.Context) {
	var category models.Category
	config.DB.Where("id=?", c.Param("id")).First(&category)
	c.BindJSON(&category)
	config.DB.Save(&category)
	c.JSON(200, &category)
}
