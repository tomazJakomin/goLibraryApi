package database

import (
	"fmt"
	"github.com/tomazJakomin/go-base-app/internal/config"
	"github.com/tomazJakomin/go-base-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb(config config.Config) *gorm.DB {
	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Ljubljana",
		config.Database.Host,
		config.Database.User,
		config.Database.Pass,
		config.Database.DbName,
		config.Database.Port)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Migrator().DropTable("loans", "books", "users")
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.Loan{})

	initialDataSetter(db)

	return db
}
