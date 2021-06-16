package database

import (
	"github.com/tomazJakomin/go-base-app/models"
	"gorm.io/gorm"
)

func initialDataSetter(db *gorm.DB) {
	books := []models.Book{
		{Title: "Pragmatic programmer", Quantity: 2, Description: "From journeyman to master"},
		{Title: "Design patterns", Quantity: 4, Description: "Elements of reusable object-oriented software"},
		{Title: "Clean code", Quantity: 1, Description: "A handbook of agile software craftmanship"},
	}

	db.Create(&books)
}
