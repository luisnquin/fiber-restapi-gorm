package models

import (
	"github.com/jinzhu/gorm"
)

type Company struct {
	gorm.Model
	Name          string `gorm:"unique;not_null" json:"name"`
	FundationYear int    `json:"fundation_year"`
}

type Employee struct {
	gorm.Model
	Fullname    string `json:"fullname"`
	Age         int    `json:"age"`
	Position    string `json:"position"`
	CompanyCode int    `gorm:"primaryKey" json:"company_code"`
}
