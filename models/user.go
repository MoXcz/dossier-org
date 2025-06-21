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

type CreateUserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (params CreateUserParams) Validate() []string {
	errors := []string{}
	if len(params.Name) < minNameLen {
		errors = append(errors, fmt.Sprintf("name lenght should be at least %d characters", minNameLen))
	}
	if len(params.Password) < minPasswordLen {
		errors = append(errors, fmt.Sprintf("password lenght should be at least %d characters", minPasswordLen))
	}
	if !isEmailValid(params.Email) {
		errors = append(errors, "email is invalid")
	}
	return errors
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

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
