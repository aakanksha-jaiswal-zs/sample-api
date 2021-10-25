package models

import "time"

type Employee struct {
	ID          int
	Name        string
	Designation string
	CreatedAt   time.Time
}
