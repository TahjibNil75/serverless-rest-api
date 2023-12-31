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
	GetAllAuthorWithRetry() ([]models.Author, error)
	FindAuthor(username string) ([]models.Author, error)

	SaveArticle(article *models.Article) (*models.Article, error)
}

type Repository struct {
	db *gorm.DB
}

func ConnectToDB() (DatabaseIFace, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	// In ConnectToDB function
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	// Automatically migrate the schema
	err = db.AutoMigrate(&models.Author{})
	if err != nil {
		return nil, fmt.Errorf("error migrating schema: %v", err)
	}

	// return &Repository{db: db}, nil
	return &Repository{db: db}, nil

}
