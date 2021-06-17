package bookExchange

import (
	"errors"
	"github.com/tomazJakomin/go-base-app/repositories"
	"gorm.io/gorm"
)

type returnService struct {
	db       gorm.DB
	loanRepo repositories.LoanRepository
}

func (service returnService) ReturnBookForUser(bookId int, userId int) error {
	loans, err := service.loanRepo.GetActiveLoansForUserIdAndBookId(userId, bookId)

	if err != nil || len(*loans) < 1 {
		return errors.New("No book loan registered")
	}

	err1 := service.loanRepo.DeactivateActiveLoanForUserAndBookIncreaseBookQuantity(userId, bookId)
	if err1 != nil {
		return err1
	}

	return err1
}
