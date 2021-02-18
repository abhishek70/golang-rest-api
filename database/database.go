package database

import (
	"fmt"
	"github.com/abhishek70/golang-rest-api/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var (
	DB *gorm.DB
)

func Initialize() {

	var DbUser 		=	os.Getenv("DB_USER")
	var DbPassword	=	os.Getenv("DB_PASS")
	var DbHost		=	os.Getenv("DB_HOST")
	var DbPort		= 	os.Getenv("DB_PORT")
	var DbName		=	os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Add table suffix when creating tables
	_ = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Product{})


	// Create
	db.Create(&model.Product{Sku: "D42", Name: "Product One"})

	DB = db
}