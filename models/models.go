package models

import (
	"gorm.io/gorm"
)

type Supplier struct {
	gorm.Model
	Name        string `json:"name"`
	ContactInfo string `json:"contact_info"`
	Products    []Product
}

type Category struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Products    []Product `json:"products"`
}

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	SupplierID  uint    `json:"supplier_id"`
	CategoryID  uint    `json:"category_id"`
	Supplier    Supplier
	Category    Category
}
