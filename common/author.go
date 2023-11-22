package common

import (
	"errors"
	"time"

	"github.com/tahjib75/models"
	"gorm.io/gorm"
)

func (r *Repository) CreateAuthor(author *models.Author) (*models.Author, error) {
	var existingAuthor models.Author
	result := r.db.Where("email = ?", author.Email).First(&existingAuthor)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// The record does not exist, so create it
		err := r.db.Create(author).Error

		// Introduce a delay before returning
		time.Sleep(500 * time.Millisecond)

		return author, err
	}

	// The record already exists
	return nil, errors.New("author already exists")
}

// func (r Repository) GetAllAuthor() ([]models.Author, error) {
// 	var authors []models.Author
// 	err := r.db.Find(&authors).Error
// 	return authors, err
// }

const maxRetries = 3
const retryDelay = 500 * time.Millisecond

func (r *Repository) GetAllAuthorWithRetry() ([]models.Author, error) {
	var authors []models.Author
	var err error
	for attempt := 1; attempt <= maxRetries; attempt++ {
		err = r.db.Find(&authors).Error
		if err == nil {
			return authors, nil
		}
		// Introduce a delay before retrying
		time.Sleep(retryDelay)
	}

	return nil, err
}

func (r *Repository) FindAuthor(email string) ([]models.Author, error) {
	var authors []models.Author
	err := r.db.Where("email = ?", email).Find(&authors).Error
	return authors, err
}
