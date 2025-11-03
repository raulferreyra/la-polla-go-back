package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open(url string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(url), &gorm.Config{})
}
