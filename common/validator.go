package common

import (
	"github.com/go-playground/validator/v10"
)

type AuthorSignupValidator struct {
	UserName string `json:"username" binding:"required,min=4,max=255" validator:"required,min=4,max=255"`
	Email    string `json:"email" binding:"required,email" validator:"required,email" unique:"true"`
}

func ValidateAuthorSignup(authorSignupValidator AuthorSignupValidator) error {
	validate := validator.New()
	return validate.Struct(authorSignupValidator)
}
