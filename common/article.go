package common

import (
	"log"

	"github.com/tahjib75/models"
)

// func (r *Repository) SaveArticle(article *models.Article) (*models.Article, error) {
// 	err := r.db.Create(article).Error
// 	return article, err
// }

func (r *Repository) SaveArticle(article *models.Article) (*models.Article, error) {
	err := r.db.Create(article).Error
	if err != nil {
		log.Printf("Error saving article: %v", err)
	} else {
		log.Printf("Article saved successfully: %v", article)
	}
	return article, err
}
