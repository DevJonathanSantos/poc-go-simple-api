package use_cases

import (
	"github.com/DevJonathanSantos/poc-go-simple-api/internal/entities"
	"github.com/DevJonathanSantos/poc-go-simple-api/internal/repositories"
	"github.com/google/uuid"
)

type getCategoryByIdUseCase struct {
	repository repositories.ICategoryRepository
}

func NewGetCategoryByIdUseCase(repository repositories.ICategoryRepository) *getCategoryByIdUseCase {
	return &getCategoryByIdUseCase{repository}
}

func (u *getCategoryByIdUseCase) Execute(id uuid.UUID) (*entities.Category, error) {

	category, _ := u.repository.GetById(id)

	return category, nil
}
