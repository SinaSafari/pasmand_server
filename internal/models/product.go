package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title       string
	Description string
	Content     string
	ImageUrl    string
	Price       int16
	Category    Category
	CategoryID  int
}
