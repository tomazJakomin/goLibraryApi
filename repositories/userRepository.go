package repositories

import (
	"github.com/tomazJakomin/go-base-app/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db gorm.DB
}

func NewUserRepository(db gorm.DB) UserRepository {
	return UserRepository{db: db}
}

func (userRepo *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	result := userRepo.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repo UserRepository) GetUser(id string) (*models.User, error) {
	user := models.User{}
	result := repo.db.First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo UserRepository) GetAllUsers() (*[]models.User, error) {
	users := []models.User{}
	result := repo.db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return &users, nil
}
