package main

import (
	"github.com/DevJonathanSantos/poc-go-simple-api/cmd/api/controllers"
	"github.com/DevJonathanSantos/poc-go-simple-api/internal/repositories"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")
	categoryRepository := repositories.NewCategoryRepository()

	categoryRoutes.POST("/", func(ctx *gin.Context) {

		controllers.CreateCategory(ctx, categoryRepository)
	})

	categoryRoutes.GET("/", func(ctx *gin.Context) {

		controllers.ListCategories(ctx, categoryRepository)
	})

	categoryRoutes.GET("/:id", func(ctx *gin.Context) {
		controllers.GetCategoryById(ctx, categoryRepository)
	})
}
