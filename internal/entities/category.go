package entities

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewCategory(name string) (*Category, error) {
	category := &Category{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := category.IsValid()

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (c *Category) IsValid() error {

	if len(c.Name) < 5 {
		return fmt.Errorf("name must be greater than 5. got %d", len(c.Name))
	}

	return nil
}
