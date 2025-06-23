package models

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost     = 12
	minNameLen     = 2
	minPasswordLen = 7
)

// server-side parameters
type CreateUserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// TODO: Return APIValidateUserError and then adjust the returned struct
func (params CreateUserParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.Name) < minNameLen {
		errors["name"] = fmt.Sprintf("name length should be at least %d characters", minNameLen)
	}
	if len(params.Password) < minPasswordLen {
		errors["password"] = fmt.Sprintf("password length should be at least %d characters", minPasswordLen)
	}
	if !isEmailValid(params.Email) {
		errors["email"] = "email is invalid"
	}
	return errors
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

// client-side parameters
type User struct {
	ID                int32  `json:"id,omitempty"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	EncryptedPassword string `json:"-"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encpwd, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		Name:              params.Name,
		Email:             params.Email,
		EncryptedPassword: string(encpwd),
	}, nil
}
