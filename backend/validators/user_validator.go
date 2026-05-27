package validators

import (
	"anime-streaming-platform/models"
	"errors"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)

// ValidateRegister validates registration data
func ValidateRegister(req *models.RegisterRequest) error {
	if strings.TrimSpace(req.Username) == "" {
		return errors.New("username is required")
	}
	if len(req.Username) < 3 || len(req.Username) > 30 {
		return errors.New("username must be between 3 and 30 characters")
	}
	if !usernameRegex.MatchString(req.Username) {
		return errors.New("username can only contain letters, numbers, and underscores")
	}
	if strings.TrimSpace(req.Email) == "" {
		return errors.New("email is required")
	}
	if !emailRegex.MatchString(req.Email) {
		return errors.New("invalid email format")
	}
	if len(req.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	if len(req.Password) > 100 {
		return errors.New("password must be less than 100 characters")
	}
	return nil
}

// ValidateLogin validates login data
func ValidateLogin(req *models.LoginRequest) error {
	if strings.TrimSpace(req.Email) == "" {
		return errors.New("email is required")
	}
	if strings.TrimSpace(req.Password) == "" {
		return errors.New("password is required")
	}
	return nil
}
