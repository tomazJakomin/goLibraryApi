package bookExchange

import (
	"github.com/tomazJakomin/go-base-app/repositories"
	"gorm.io/gorm"
)

type returnService struct {
	db       gorm.DB
	loanRepo repositories.LoanRepository
}

func (service returnService) ReturnBookForUser(bookId int, userId int) error {
	err1 := service.loanRepo.DeactivateActiveLoanForUserAndBookIncreaseBookQuantity(userId, bookId)
	if err1 != nil {
		return err1
	}

	println(err1)

	return err1
}
