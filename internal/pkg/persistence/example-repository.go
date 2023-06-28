package persistence

import (
	"fmt"

	"gitlab.com/JonasEtzold/go-service-template/internal/pkg/db"
	models "gitlab.com/JonasEtzold/go-service-template/internal/pkg/models/example"
	"gorm.io/gorm"
)

type ExampleRepository struct {
	db *gorm.DB
}

var exampleRepository *ExampleRepository

func GetExampleRepository() *ExampleRepository {
	if exampleRepository == nil {
		exampleRepository = &ExampleRepository{
			db: db.Get(),
		}
	}

	return exampleRepository
}

func (r *ExampleRepository) Get(id string) (*models.Example, error) {
	var example models.Example
	r.db.First(&example, id)

	if len(example.Name) < 1 {
		return nil, fmt.Errorf("Example not found")
	}
	return &example, nil
}

func (r *ExampleRepository) All() (*[]models.Example, error) {
	var examples []models.Example
	result := db.Get().Order("id asc").Find(&examples)
	return &examples, result.Error
}

func (r *ExampleRepository) Query(q *models.Example) (*[]models.Example, error) {
	var examples []models.Example
	result := db.Get().Session(&gorm.Session{QueryFields: true}).Order("id asc").Find(&examples, q)
	return &examples, result.Error
}

func (r *ExampleRepository) Add(example *models.Example) error {
	return r.db.Create(&example).Error
}

func (r *ExampleRepository) Update(example *models.Example) error {
	return db.Get().Save(&example).Error
}

func (r *ExampleRepository) Delete(example *models.Example) error {
	return db.Get().Unscoped().Delete(&example).Error
}
