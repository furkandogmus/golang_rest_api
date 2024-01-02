package routes

import (
	"restapi/controller"

	"github.com/gin-gonic/gin"
)

func CategoryRoute(router *gin.Engine) {
	router.GET("/category/", controller.GetCategories)
	router.POST("/category/", controller.CreateCategory)
	router.DELETE("/category/:id", controller.DeleteCategory)
	router.PUT("/category/:id", controller.UpdateCategory)
}
