package common

import (
	"fmt"
	"os"

	"github.com/tahjib75/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseIFace interface {
	CreateAuthor(author *models.Author) (*models.Author, error)
	GetAllAuthor() ([]models.Author, error)
}

type Repository struct {
	db *gorm.DB
}

func ConnectToDB() (DatabaseIFace, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	// Automatically migrate the schema
	err = db.AutoMigrate(&models.Author{})
	if err != nil {
		return nil, err
	}
	return &Repository{db: db}, nil

}
