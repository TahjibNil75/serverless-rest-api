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
	refreshTokenExpiration := time.Time{}

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

func UpdateAllTokens(accessToken string, refreshToken bool) (string, error) {
	// Initialize a claims struct to store token information
	claims := &JwtClaim{}

	// Parse the existing token without verifying the signature
	_, _, err := new(jwt.Parser).ParseUnverified(accessToken, claims)
	if err != nil {
		return "", err
	}

	// Determine the expiration time based on whether it's a refresh token or not
	var expirationTime time.Duration
	if refreshToken {
		expirationTime = 0 // Set to 0 for no expiration (refresh token)
	} else {
		expirationTime = time.Hour * 6 // Set expiration to 6 hours for access token
	}

	// Set the expiration time in the claims
	claims.ExpiresAt = time.Now().Add(expirationTime).Unix()

	// Create a new token with the updated expiration time
	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedNewAccessToken, err := newAccessToken.SignedString([]byte(JwtToken))
	if err != nil {
		return "", err
	}

	return signedNewAccessToken, nil
}
