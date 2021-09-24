package entities

import "time"

type User struct {
	Name       string
	Email      string
	Created_at time.Time
	Updated_at time.Time
}
