package main

import (
	"restapi/config"
	_ "restapi/docs"
	"restapi/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.New()
	url := ginSwagger.URL("localhost/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	config.Connect()
	routes.UserRoute(router)
	routes.BookRoute(router)
	routes.CategoryRoute(router)
	router.Run(":8080")

}
