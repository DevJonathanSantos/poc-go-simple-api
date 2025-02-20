package repositories

import (
	"github.com/DevJonathanSantos/poc-go-simple-api/internal/entities"
	"github.com/google/uuid"
)

type ICategoryRepository interface {
	Save(category *entities.Category) error
	List() ([]*entities.Category, error)
	GetById(id uuid.UUID) (*entities.Category, error)
}
