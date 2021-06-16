package models

import (
	"gorm.io/gorm"
)

type Loan struct {
	gorm.Model
	ID     int
	UserID int
	BookID int
	Active bool
}
