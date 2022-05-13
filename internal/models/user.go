package models

import (
	"fmt"
	"net/http"
	"net/mail"
	"unicode"

	"github.com/FoxFurry/PBL-2022-GO/internal/httperr"
)

type User struct {
	ID       uint64 `json:"ID,omitempty"`
	UUID     string `json:"UUID"`
	Mail     string `json:"mail"`
	Password string `json:"password,omitempty"`
}

func (u *User) ValidateAll() error {
	if err := u.ValidateMail(); err != nil {
		return err
	} else if err := ValidatePassword(u.Password); err != nil {
		return err
	}

	return nil
}

func (u *User) ValidateMail() error {
	if _, err := mail.ParseAddress(u.Mail); err != nil {
		return httperr.New(err.Error(), http.StatusBadRequest)
	}
	return nil
}

/*
ValidatePassword
 * Password rules:
 * at least 8 characters
 * at least 1 number
 * at least 1 upper case
 * at least 1 special character
*/
func ValidatePassword(pass string) error {
	if len(pass) < 8 {
		return httperr.ValidationError("password", "should be at least 8 characters long")
	}

	var (
		numberPresent  bool
		upperPresent   bool
		specialPresent bool
	)
	for _, c := range pass {
		switch {
		case unicode.IsNumber(c):
			numberPresent = true
		case unicode.IsUpper(c):
			upperPresent = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			specialPresent = true
		case unicode.IsLetter(c):
			continue
		default:
			return httperr.ValidationError("password", fmt.Sprintf("unsupported character: %c", c))
		}
	}

	if !numberPresent {
		return httperr.ValidationError("password", "should contain at least one number")
	} else if !upperPresent {
		return httperr.ValidationError("password", "should contain at least one uppercase character")
	} else if !specialPresent {
		return httperr.ValidationError("password", "should contain at least one special character")
	}

	return nil
}
