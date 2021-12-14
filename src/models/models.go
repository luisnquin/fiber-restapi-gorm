package models

import (
	"github.com/jinzhu/gorm"
)

type Company struct {
	gorm.Model
	Name          string `gorm:"unique;not_null"`
	Code          uint   `gorm:"primaryKey"`
	FundationYear int
}

type Employee struct {
	gorm.Model
	Fullname    string
	Age         int
	Position    string
	CompanyCode uint    `gorm:"primaryKey"`
}
