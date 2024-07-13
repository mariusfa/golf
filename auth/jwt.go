package auth

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// TODO: implement jwt.go

// use jwt for authentication

type UserClaims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

func ParseToken(token string, secret string) (string, error) {
	var userClaims UserClaims
	_, err := jwt.ParseWithClaims(token, &userClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", fmt.Errorf("Error parsing token: %w", err)
	}

	return userClaims.UserId, nil
}

func CreateToken(userId string, secret string, expires *time.Time) (string, error) {
	now := jwt.NewNumericDate(time.Now())
	var expiresAt *jwt.NumericDate

	if expires == nil {
		defaultExpires := time.Now().Add(24 * time.Hour)
		expiresAt = jwt.NewNumericDate(defaultExpires)
	} else {
		expiresAt = jwt.NewNumericDate(*expires)
	}

	claims := UserClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  now,
			NotBefore: now,
		},
	}

	signedToken, err := createSignedTokenWithClaims(secret, claims)
	if err != nil {
		return "", fmt.Errorf("Error creating signed token: %w", err)
	}
	return signedToken, nil
}

func createSignedTokenWithClaims(secret string, userClaims UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
