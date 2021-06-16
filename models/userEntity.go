package models

import "gorm.io/gorm"
import "github.com/go-playground/validator/v10"

type User struct {
	gorm.Model
	ID       int
	Name     string `json:"name" validate:"min=1,required"`
	LastName string `json:"lastName" validate:"min=1,required"`
	Loans    []Loan `json:"-"`
}

func (u *User) Validate() error {
	return validator.New().Struct(u)
}
