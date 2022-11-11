package config

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "sqlserver://login:password@localhost:1433?database=Test"
	d, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
