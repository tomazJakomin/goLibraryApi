package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID          int
	Title       string
	Description string
	Quantity    int
	Loans       []Loan `json:"-"`
}
