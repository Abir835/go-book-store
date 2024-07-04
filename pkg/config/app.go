package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	db *gorm.DB
)

func Connect() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("PASSWORD")
	dbName := os.Getenv("DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser, dbPass, dbName)

	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
