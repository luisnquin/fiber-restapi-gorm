package models

import (
	"github.com/jinzhu/gorm"
)

type Employee struct {
	gorm.Model
	Fullname string
	Age      int
	Position string
	Company  string
}


type Company struct {
	gorm.Model
	Name          string `gorm:"unique;not_null"`
	FundationYear int
	Employees     []Employee `gorm:"foreignKey:Company;references:Name"`
}
