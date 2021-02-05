package model

import (
	"crypto/x509/pkix"
	"encoding/base32"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	Base    `valid:"required"`
	Name    string    `json:"name" valid:"notnull"`
	Email   string    `json:"email" valid:"notnull"`
	PixKeys []*PixKey `valid:"-"`
}

func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}
	return nil
}

func NewUser(bank *Bank, name string, email string) (*User, error) {
	user := User{
		Name:  name,
		Email: email,
	}
	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()

	err := User.isValid()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
