package stores

import "example.com/sample-api/models"

type Student interface {
	GetAll() ([]models.Student, error)
	Get(id int64) (models.Student, error)
	Create(student models.Student) (int64, error)
	Update(student models.Student) error
	Delete(id int64) error
}
