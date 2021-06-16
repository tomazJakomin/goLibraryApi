package repositories

import (
	"fmt"
	"github.com/tomazJakomin/go-base-app/models"
	"gorm.io/gorm"
)

type LoanRepository struct {
	db *gorm.DB
}

const decrease_book_quantity_query string = "UPDATE books SET quantity = quantity + 1 WHERE id = ?"
const deactivate_loan_query string = "UPDATE loans SET active = ? WHERE id = (SELECT id FROM loans WHERE user_id = ? AND book_id = ? and active = ? LIMIT 1)"

func NewLoanRepository(db *gorm.DB) *LoanRepository {
	return &LoanRepository{
		db: db,
	}
}

func (repo *LoanRepository) GetActiveLoansForUser(user models.User) (*[]models.Loan, error) {
	var loan []models.Loan
	fmt.Println(user.ID)
	result := repo.db.Where("user_id = ? AND active = ?", user.ID, true).Find(&loan)

	if result.Error != nil {
		return nil, result.Error
	}

	return &loan, nil
}

func (repo *LoanRepository) DeactivateActiveLoanForUserAndBookIncreaseBookQuantity(userId int, bookId int) error {
	fmt.Println(userId)

	tx := repo.db.Begin()

	if err := tx.Exec(decrease_book_quantity_query, bookId).Error; err != nil {
		tx.Rollback()

		return err
	}

	if err := tx.Exec(deactivate_loan_query, false, userId, bookId, true).Error; err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit().Error
}

func (repo LoanRepository) CreateBookLoanForUser(bookId int, userId int) *gorm.DB {

	loan := models.Loan{
		BookID: bookId,
		UserID: userId,
		Active: true,
	}

	return repo.db.Create(&loan)
}
