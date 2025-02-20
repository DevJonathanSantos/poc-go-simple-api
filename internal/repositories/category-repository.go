package repositories

import (
	"errors"

	"github.com/DevJonathanSantos/poc-go-simple-api/internal/entities"
	"github.com/google/uuid"
)

type categoryRepository struct {
	db []*entities.Category
}

func NewCategoryRepository() *categoryRepository {
	return &categoryRepository{
		db: make([]*entities.Category, 0),
	}
}

func (r *categoryRepository) Save(category *entities.Category) error {
	r.db = append(r.db, category)
	return nil
}

func (r *categoryRepository) List() ([]*entities.Category, error) {
	return r.db, nil
}

func (r *categoryRepository) GetById(id uuid.UUID) (*entities.Category, error) {
	for _, category := range r.db {
		if category.ID == id {
			return category, nil
		}
	}
	return nil, errors.New("category not found")
}
