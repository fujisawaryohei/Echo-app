package entities

import "time"

type User struct {
	id         string
	name       string
	email      string
	created_at time.Time
	updated_at time.Time
}
