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

func NewPostGressDB(dsn string) *postgress {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &postgress{db: db}
}

func (p *postgress) GetDB() *gorm.DB {
	return p.db
}
