package form

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
)

//ValidatePassword : will validate the user entered password
func ValidatePassword(password string) string {

	if len(password) < 5 {
		return "password must be longer than 5 characters"
	}
	if (symbolexists(password)) != "" {
		return symbolexists(password)
	}

	return ""
}

func symbolexists(password string) string {
	f := func(r rune) bool {
		return r < 'A' || r > 'z'
	}
	g := func(r rune) bool {
		return r >= 'A' && r <= 'Z'
	}
	h := func(r rune) bool {
		return r >= 'a' && r <= 'z'
	}
	if strings.IndexFunc(password, g) == -1 {
		return "password must contain uppercase letters"
	}
	if strings.IndexFunc(password, h) == -1 {
		return "password must contain lowercase letters"
	}
	if strings.IndexFunc(password, f) == -1 {
		return "password must contain special characters"
	}
	return ""
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CompareHashToPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
