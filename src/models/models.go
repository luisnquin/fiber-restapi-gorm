package models

import (
	"github.com/jinzhu/gorm"
)

type Company struct {
	gorm.Model
	Name          string `gorm:"unique;not_null" json:"name"`
	FundationYear int    `gorm:"not_null" json:"fundation_year"`
}

type Employee struct {
	gorm.Model
	Fullname    string `gorm:"not_null" json:"fullname"`
	Age         int    `gorm:"not_null" json:"age"`
	Position    string `gorm:"not_null" json:"position"`
	CompanyCode int    `gorm:"not_null" json:"company_code"`
}

type EmployeesAndCompany struct {
	Fullname      string `json:"fullname"`
	Age           int    `json:"age"`
	Position      string `json:"position"`
	CompanyCode   int    `json:"company_code"`
	Name          string `json:"name"`
	FundationYear int    `json:"fundation_year"`
}
