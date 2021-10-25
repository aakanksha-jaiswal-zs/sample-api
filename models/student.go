package models

import "time"

type Student struct {
	ID        int64
	Name      string
	Major     string
	CreatedAt time.Time
}

//todo: json tags
