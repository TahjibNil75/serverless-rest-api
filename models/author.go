package models

import "time"

type Author struct {
	Uid             uint      `json:"uid" gorm:"primaryKey:unique"`
	UserName        string    `json:"userName" gorm:"not null" validate:"required, min=3,max=50"`
	Email           string    `json:"email" gorm:"not null:unique" validate:"email, required"`
	Password        string    `json:"password" gorm:"not null" validate:"required"`
	PasswordConfirm string    `json:"passwordConfirm" bindings:"required"`
	CreatedAt       time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt       time.Time `json:"updatedAt" gorm:"not null"`
	Token           *string   `json:"token"`
	RefreshToken    *string   `json:"refreshToken"`
}

type GetAuthors struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
}
