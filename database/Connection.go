package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres dbname=echoapp_development port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	return db, err
}
