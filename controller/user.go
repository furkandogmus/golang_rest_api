package controller

import (
	"restapi/config"
	"restapi/models"

	"github.com/gin-gonic/gin"
)

// @Summary Get all users
// @Description Get a list of all users with associated books and categories
// @Produce json
// @Success 200 {array} models.User
// @Router /user [get]
func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Preload("Books.Category").Find(&users)
	c.JSON(200, &users)
}

// @Summary Create a new user
// @Description Create a new user
// @Accept json
// @Produce json
// @Param user body models.User true "User object to be created"
// @Success 200 {object} models.User
// @Router /user [post]
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	user.ID = 0
	config.DB.Create(&user)
	c.JSON(200, &user)
}

// @Summary Delete a user by ID
// @Description Delete a user by ID
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {string} string "User has been deleted successfully."
// @Router /user/{id} [delete]
func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, "User has been deleted successfully.")
}

// @Summary Update a user by ID
// @Description Update a user by ID
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "Updated user object"
// @Success 200 {object} models.User
// @Router /user/{id} [put]
func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id=?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}
