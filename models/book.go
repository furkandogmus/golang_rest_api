package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Author     string   `json:"author"`
	Title      string   `json:"title"`
	ISBN       string   `json:"isbn"`
	CategoryID int      `json:"category_id" `
	Category   Category `json:"category"`
	UserID     uint     `json:"user_id"`
}
