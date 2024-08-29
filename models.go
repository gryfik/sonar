package main

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	CategoryID  uint
	Category    Category
}

type Category struct {
	gorm.Model
	Name     string
	Products []Product
}

type Cart struct {
	gorm.Model
	Products []Product `gorm:"many2many:cart_products;"`
}
