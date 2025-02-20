package controllers

import (
	"net/http"

	"github.com/DevJonathanSantos/poc-go-simple-api/internal/repositories"
	use_cases "github.com/DevJonathanSantos/poc-go-simple-api/internal/use-cases"
	"github.com/gin-gonic/gin"
)

func ListCategories(ctx *gin.Context, repository repositories.ICategoryRepository) {
	useCase := use_cases.NewListCategoriesUseCase(repository)

	categories, err := useCase.Execute()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success":    true,
		"categories": categories,
	})
}
