package common

import (
	"github.com/go-playground/validator/v10"
)

type AuthorSignupValidator struct {
	UserName        string `json:"username" binding:"required,min=4,max=255" validator:"required,min=4,max=255"`
	Email           string `json:"email" binding:"required,email" validator:"required,email" unique:"true"`
	Password        string `json:"password" binding:"required,min=4,max=255" validator:"required,min=4,max=255"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

func ValidateAuthorSignup(authorSignupValidator AuthorSignupValidator) error {
	validate := validator.New()
	return validate.Struct(authorSignupValidator)
}

type LoginValidator struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=255"`
}

func ValidateAuthorLogin(authorLoginValidator LoginValidator) error {
	validate := validator.New()
	return validate.Struct(authorLoginValidator)
}

type ArticleValidator struct {
	Name    string `json:"name" binding:"required,min=3,max=255"`
	Content string `json:"content" binding:"required,min=300,max=2000"`
}

// ValidateArticle validates the ArticleValidator struct.
func ValidateArticle(articleValidator ArticleValidator) error {
	validate := validator.New()
	return validate.Struct(articleValidator)
}
