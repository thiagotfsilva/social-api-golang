package models

import (
	handlehash "api-devbook/src/utils/handleHash"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (u *User) validate(step string) error {
	if u.Name == "" {
		return errors.New("name is required")
	}

	if u.Nick == "" {
		return errors.New("nick is required")
	}

	if u.Email == "" {
		return errors.New("email is required")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("email is invalid")
	}

	if step == "register" && u.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func (u *User) format(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)

	if step == "register" {
		hashPassword, err := handlehash.Hash(u.Password)
		if err != nil {
			return err
		}

		u.Password = string(hashPassword)
	}

	return nil
}

// Prepare chama os metodos para validar e formatar o usuario
func (u *User) Prepare(step string) error {
	if err := u.validate(step); err != nil {
		return err
	}

	if err := u.format(step); err != nil {
		return err
	}

	return nil
}
