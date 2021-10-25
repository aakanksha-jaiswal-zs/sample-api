package student

import (
	"strings"

	"example.com/sample-api/errors"
	"example.com/sample-api/models"
	"example.com/sample-api/services"
	"example.com/sample-api/stores"
)

type service struct {
	// inject the store layer dependency
	store stores.Student
}

func New(s stores.Student) services.Student {
	return service{store: s}
}

func (s service) GetAll() ([]models.Student, error) {
	return s.store.GetAll()
}

func (s service) Get(id int64) (models.Student, error) {
	if id < 0 {
		return models.Student{}, errors.InvalidParam{Param: []string{"id"}}
	}

	return s.store.Get(id)
}

//todo service logic for create
func (s service) Create(student models.Student) (models.Student, error) {
	err := checkMandatoryFields(student)
	if err != nil {
		return models.Student{}, err
	}

	id, err := s.store.Create(student)
	if err != nil {
		return models.Student{}, err
	}

	return s.Get(id)
}

func checkMandatoryFields(student models.Student) error {
	params := make([]string, 0, 2)

	if strings.TrimSpace(student.Name) == "" {
		// missing param name
		params = append(params, "name")
	}

	if strings.TrimSpace(student.Major) == "" {
		// missing param major
		params = append(params, "major")
	}

	if len(params) > 0 {
		return errors.MissingParam{Param: params}
	}

	return nil
}

func (s service) Update(student models.Student) (models.Student, error) {
	err := validate(student)
	if err != nil {
		return models.Student{}, err
	}

	_, err = s.Get(student.ID)
	if err != nil {
		return models.Student{}, err
	}

	err = s.store.Update(student)
	if err != nil {
		return models.Student{}, err
	}

	return s.Get(student.ID)
}

func (s service) Delete(id int64) error {
	_, err := s.Get(id)
	if err != nil {
		return err
	}

	return s.store.Delete(id)
}

func validate(student models.Student) error {
	// add validation for empty update
	if strings.TrimSpace(student.Name) == "" && strings.TrimSpace(student.Major) == "" {
		return errors.MissingParam{Param: []string{"name", "major"}}
	}

	if len(student.Major) > 3 {
		return errors.InvalidParam{Param: []string{"major"}}
	}

	return nil
}
