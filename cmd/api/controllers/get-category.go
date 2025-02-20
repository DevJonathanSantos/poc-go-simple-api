package controllers

import (
	"net/http"

	"github.com/DevJonathanSantos/poc-go-simple-api/internal/repositories"
	use_cases "github.com/DevJonathanSantos/poc-go-simple-api/internal/use-cases"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetCategoryById(ctx *gin.Context, repository repositories.ICategoryRepository) {
	idParam := ctx.Param("id") // Get ID from URL param

	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid UUID format",
		})
		return
	}

	useCase := use_cases.NewGetCategoryByIdUseCase(repository)

	category, err := useCase.Execute(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success":  true,
		"category": category,
	})
}
