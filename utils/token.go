package utils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

var JwtToken = []byte("JWT_SECRET")

type JwtClaim struct {
	Email string `json:"email"`
	Uid   uint   `json:"uid"`
	jwt.StandardClaims
}

func CreateToken(email string, uid uint) (accessToken string, refreshToken string, err error) {
	accessTokenExpiration := time.Now().Add(6 * time.Hour)
	refreshTokenExpiration := time.Now().Add(168 * time.Hour)

	claims := &JwtClaim{
		Email: email,
		Uid:   uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessTokenExpiration.Unix(),
		},
	}
	refreshClaims := &JwtClaim{
		Email: email,
		Uid:   uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshTokenExpiration.Unix(),
		},
	}
	// Create a new JWT token with the specified claims and signing method (HS256)
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(JwtToken))
	if err != nil {
		log.Panic(err)
		return "", "", err
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(JwtToken))
	if err != nil {
		log.Panic(err)
		//return "", err 			// not enough return values; have (string, error);want (string, string, error)
		return "", "", err
	}

	return accessToken, refreshToken, nil

}
