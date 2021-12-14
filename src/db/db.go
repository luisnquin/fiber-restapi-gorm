package db

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/luisnquin/fiber-restapi-gorm/models"
)

var DB *gorm.DB
var err error

func Init() *gorm.DB {
	dsn := os.Getenv("DSN")
	DB, err = gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	// DB.DropTableIfExists(&models.Company{}, &models.Employee{})
	DB.AutoMigrate(&models.Company{}, &models.Employee{})

	return DB
}
