package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"link-shortener/configs"
)

type DB struct {
	*gorm.DB
}

func NewDB(conf *configs.Config) *DB {
	db, err := gorm.Open(postgres.Open(conf.Db.Dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return &DB{db}
}
