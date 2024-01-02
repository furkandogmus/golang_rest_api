package routes

import (
	"restapi/controller"

	"github.com/gin-gonic/gin"
)

func BookRoute(router *gin.Engine) {
	router.GET("/book/", controller.GetBooks)
	router.POST("/book/", controller.AddBook)
	router.DELETE("/book/:id", controller.DeleteBook)
	router.PUT("/book/:id", controller.UpdateBook)
}
