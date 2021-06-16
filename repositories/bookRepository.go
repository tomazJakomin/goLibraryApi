package repositories

import (
	"github.com/tomazJakomin/go-base-app/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	db gorm.DB
}

func NewBookRepository(db gorm.DB) BookRepository {
	return BookRepository{
		db: db,
	}
}

func (repo BookRepository) GetBook(id int) (*models.Book, error) {
	book := models.Book{}
	result := repo.db.First(&book, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

func (repo BookRepository) GetAvailableBooks() (*[]models.Book, error) {
	books := []models.Book{}
	result := repo.db.Where("quantity > 0").Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return &books, nil
}

func (repo BookRepository) UpdateBook(selectedBook models.Book) (*models.Book, error) {
	result := repo.db.Save(&selectedBook)

	if result.Error != nil {
		return nil, result.Error
	}

	return &selectedBook, nil
}
