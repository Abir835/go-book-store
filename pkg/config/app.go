package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"os"
)

var (
	db *gorm.DB
)

func Connect() {
	dbUser := os.Getenv("USERNAME")
	dbPass := os.Getenv("PASSWORD")
	dbName := os.Getenv("DATABASE")

	dsn := fmt.Sprintf("%s:%s/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser, dbPass, dbName)

	d, err := gorm.Open("mssql", dsn)
	if err != nil {
		panic("failed to connect database")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
