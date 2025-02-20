package main

import (
	"github.com/DevJonathanSantos/poc-go-simple-api/cmd/api/controllers"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")

	categoryRoutes.POST("/", controllers.CreateCategory)
}
