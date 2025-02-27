package use_cases

import (
	"github.com/DevJonathanSantos/poc-go-simple-api/internal/entities"
	"github.com/DevJonathanSantos/poc-go-simple-api/internal/repositories"
)

type listCategoriesUseCase struct {
	repository repositories.ICategoryRepository
}

func NewListCategoriesUseCase(repository repositories.ICategoryRepository) *listCategoriesUseCase {
	return &listCategoriesUseCase{repository}
}

func (u *listCategoriesUseCase) Execute() ([]*entities.Category, error) {

	categories, _ := u.repository.List()

	return categories, nil
}
