package controller

import (
	"restapi/config"
	"restapi/models"

	"github.com/gin-gonic/gin"
)

// @Summary Get all books
// @Description Get a list of all books with associated categories
// @Produce json
// @Success 200 {array} models.Book
// @Router /book [get]
func GetBooks(c *gin.Context) {
	books := []models.Book{}
	config.DB.Preload("Category").Find(&books)
	c.JSON(200, &books)
}

// @Summary Add a new book
// @Description Add a new book
// @Accept json
// @Produce json
// @Param book body models.Book true "Book object to be added"
// @Success 200 {object} models.Book
// @Router /book [post]
func AddBook(c *gin.Context) {
	var book models.Book
	c.BindJSON(&book)
	book.ID = 0
	config.DB.Create(&book)
	c.JSON(200, &book)
}

// @Summary Delete a book by ID
// @Description Delete a book by ID
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {string} string "Book has been deleted successfully."
// @Router /book/{id} [delete]
func DeleteBook(c *gin.Context) {
	var book models.Book
	config.DB.Where("id = ?", c.Param("id")).Delete(&book)
	c.JSON(200, "Book has been deleted successfully.")
}

// @Summary Update a book by ID
// @Description Update a book by ID
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body models.Book true "Updated book object"
// @Success 200 {object} models.Book
// @Router /book/{id} [put]
func UpdateBook(c *gin.Context) {
	var book models.Book
	config.DB.Where("id=?", c.Param("id")).First(&book)
	c.BindJSON(&book)
	config.DB.Save(&book)
	c.JSON(200, &book)
}
