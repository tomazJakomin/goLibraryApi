package bookExchange

import (
	"errors"
	"github.com/tomazJakomin/go-base-app/models"
	"github.com/tomazJakomin/go-base-app/repositories"
	"gorm.io/gorm"
)

const MAX_BOOK_PER_USER = 5

type loanService struct {
	db       gorm.DB
	loanRepo repositories.LoanRepository
	bookRepo repositories.BookRepository
}

func (service *loanService) ExchangeBookForUser(bookId int, userId int) error {
	limitationError := service.validateUserExchangeLimit(userId)

	if limitationError != nil {
		return limitationError
	}

	loadedBook, errBook := service.bookRepo.GetBook(bookId)

	if errBook != nil {
		return errors.New("Can not find book")
	}

	if loadedBook.Quantity == 0 {
		return errors.New("No more books left!")
	}

	return service.makeExchange(loadedBook, userId)
}

func (service *loanService) validateUserExchangeLimit(userId int) error {
	user := models.User{ID: userId}
	loans, err := service.loanRepo.GetActiveLoansForUser(user)
	if err != nil {
		return err
	}

	numberOfBooksLent := len(*loans)

	if numberOfBooksLent == MAX_BOOK_PER_USER {
		return errors.New("Max exchange limit reached!")
	}

	return nil
}

func (service *loanService) makeExchange(exchangingBook *models.Book, userId int) error {
	tx := service.loanRepo.CreateBookLoanForUser(exchangingBook.ID, userId)

	if tx.Error != nil {
		return errors.New(tx.Error.Error())
	}

	exchangingBook.Quantity--
	_, err := service.bookRepo.UpdateBook(*exchangingBook)

	if err != nil {
		tx.Rollback()
		return errors.New(err.Error())
	}

	return nil
}
