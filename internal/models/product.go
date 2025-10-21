package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName  string `json:"productName"`
	Quantity string `json:"quantity"`
}