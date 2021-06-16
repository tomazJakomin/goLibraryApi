package bookExchange

import (
	"github.com/tomazJakomin/go-base-app/repositories"
	"gorm.io/gorm"
)

type ExchangeService struct {
	db            *gorm.DB
	repository    *repositories.LoanRepository
	LoanService   loanService
	ReturnService returnService
}

func NewExchangeService(db *gorm.DB) *ExchangeService {
	loanRepo := repositories.NewLoanRepository(db)
	return &ExchangeService{
		db:            db,
		repository:    loanRepo,
		LoanService:   loanService{db: *db, loanRepo: *loanRepo, bookRepo: repositories.NewBookRepository(*db)},
		ReturnService: returnService{db: *db, loanRepo: *loanRepo},
	}
}
