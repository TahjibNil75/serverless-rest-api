package common

import (
	"errors"

	"github.com/tahjib75/models"
	"gorm.io/gorm"
)

func (r *Repository) CreateAuthor(author *models.Author) (*models.Author, error) {
	var Author models.Author
	result := r.db.Where("email = ?", author.Email).First(&Author)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("author already exists")
	}
	err := r.db.Create(Author).Error
	return author, err
}

func (r Repository) GetAllAuthor() ([]models.Author, error) {
	var authors []models.Author
	err := r.db.Find(&authors).Error
	return authors, err
}
