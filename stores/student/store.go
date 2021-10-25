package student

import (
	"database/sql"
	"log"
	"strings"
	"time"

	"example.com/sample-api/errors"
	"example.com/sample-api/models"
	"example.com/sample-api/stores"
)

type store struct {
	*sql.DB // database dependency
}

func New(db *sql.DB) stores.Student {
	return store{DB: db}
}

func (s store) GetAll() ([]models.Student, error) {
	rows, err := s.DB.Query(getAllStudents)
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	var (
		student   models.Student
		createdAt string
	)

	students := make([]models.Student, 0)

	for rows.Next() {
		err := rows.Scan(&student.ID, &student.Name, &student.Major, &createdAt)
		if err != nil {
			return nil, errors.DB{Err: err}
		}

		student.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
		if err != nil {
			return nil, errors.DB{Err: err}
		}

		students = append(students, student)
	}

	return students, nil
}

func (s store) Get(id int64) (models.Student, error) {
	row := s.DB.QueryRow(getStudentByID, id)

	var createdAt string

	student := models.Student{ID: id}

	err := row.Scan(&student.Name, &student.Major, &createdAt)
	if err != nil {
		return models.Student{}, errors.DB{Err: err}
	}

	student.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
	if err != nil {
		return models.Student{}, errors.DB{Err: err}
	}

	return student, nil
}

func (s store) Create(student models.Student) (int64, error) {
	res, err := s.DB.Exec(insertStudent, student.Name, student.Major, time.Now()) // .Format("2006-01-02 15:04:05")
	if err != nil {
		return 0, errors.DB{Err: err}
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, errors.DB{Err: err}
	}

	log.Printf("student id %d has been created", id)
	return id, nil
}

func (s store) Update(student models.Student) error {
	query, args := getUpdate(student)
	if query == "" {
		return nil
	}

	_, err := s.DB.Exec(query, args...)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}

func getUpdate(student models.Student) (query string, args []interface{}) {
	fields := make([]string, 0)

	if student.Name != "" {
		fields = append(fields, "name = ?")
		args = append(args, student.Name)
	}

	if student.Major != "" {
		fields = append(fields, "major = ?")
		args = append(args, student.Major)
	}

	if len(args) > 0 {
		query = "update student set " + strings.Join(fields, ", ") + " where id = ?"
		args = append(args, student.ID)
	}

	return
}

func (s store) Delete(id int64) error {
	_, err := s.DB.Exec(deleteStudentByID, id)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}
