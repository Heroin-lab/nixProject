package services

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

type JwtCustomClaims struct {
	ID       int    `json:"id"`
	UserRole string `json:"userRole"`
	jwt.StandardClaims
}

func GenerateToken(userID int, userRole string, lifetimeMinutes int, secret string) (string, error) {
	claims := &JwtCustomClaims{
		userID,
		userRole,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(lifetimeMinutes)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenString string, secret string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("failed to parse token claims")
	}

	return claims, nil
}

func GetTokenFromBearerString(bearerString string) (string, error) {
	if bearerString == "" {
		return "", errors.New("Empty string!")
	}

	parts := strings.Split(bearerString, "Bearer")
	if len(parts) != 2 {
		return "", errors.New("To short token string!")
	}

	token := strings.TrimSpace(parts[1])
	if len(token) < 1 {
		return "", errors.New("To short token string!")
	}

	return token, nil
}
