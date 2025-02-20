package main

import "github.com/gin-gonic/gin"

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")

	categoryRoutes.POST("/")
}
