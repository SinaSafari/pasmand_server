package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Title       string `json:"title"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	Description string `json:"description"`
	Plaque      string `json:"plaque"`
	Unit        string `json:"unit"`
	User        User
	UserId      int
}
