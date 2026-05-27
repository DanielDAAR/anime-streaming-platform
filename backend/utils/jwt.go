package utils

import (
	"errors"
	"os"
	"strconv"
	"time"

	"anime-streaming-platform/models"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(getJWTSecret())

func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "default-secret-key-change-in-production"
	}
	return secret
}

// GenerateToken creates a new JWT token for a user
func GenerateToken(user models.User) (string, error) {
	expirationHours := 24
	if envHours := os.Getenv("JWT_EXPIRATION_HOURS"); envHours != "" {
		if h, err := strconv.Atoi(envHours); err == nil {
			expirationHours = h
		}
	}

	claims := jwt.MapClaims{
		"userId":   user.ID.Hex(),
		"username": user.Username,
		"email":    user.Email,
		"role":     string(user.Role),
		"exp":      time.Now().Add(time.Hour * time.Duration(expirationHours)).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken validates and parses a JWT token
func ValidateToken(tokenString string) (*models.JWTClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &models.JWTClaims{
			UserID:   getStringClaim(claims, "userId"),
			Username: getStringClaim(claims, "username"),
			Email:    getStringClaim(claims, "email"),
			Role:     models.UserRole(getStringClaim(claims, "role")),
		}, nil
	}

	return nil, errors.New("invalid token claims")
}

func getStringClaim(claims jwt.MapClaims, key string) string {
	if val, ok := claims[key].(string); ok {
		return val
	}
	return ""
}
