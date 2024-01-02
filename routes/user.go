package routes

import (
	"restapi/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/user/", controller.GetUsers)
	router.POST("/user/", controller.CreateUser)
	router.DELETE("/user/:id", controller.DeleteUser)
	router.PUT("/user/:id", controller.UpdateUser)
}
