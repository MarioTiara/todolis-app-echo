package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgress struct {
	db *gorm.DB
}

func NewPostGressDB() *postgress {
	dsn := "host=localhost user=root password=secret dbname=todolistwebapi port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &postgress{db: db}
}
