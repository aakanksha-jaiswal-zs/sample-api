package models

import "time"

type Student struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Major     string    `json:"major"`
	CreatedAt time.Time `json:"created_at"`
}

//todo: json tags
