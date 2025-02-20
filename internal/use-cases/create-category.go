package use_cases

import (
	"log"

	"github.com/DevJonathanSantos/poc-go-simple-api/internal/entities"
	"github.com/DevJonathanSantos/poc-go-simple-api/internal/repositories"
)

type createCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewCreateCategoryUseCase(repository repositories.ICategoryRepository) *createCategoryUseCase {
	return &createCategoryUseCase{repository}
}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	log.Println(category)
	u.repository.Save(category)

	return nil
}
