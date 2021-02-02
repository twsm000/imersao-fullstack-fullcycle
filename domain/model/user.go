package model

import (
	"time"

	validator "github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

// NewUser ...
func NewUser(name, email string) (*User, error) {
	u := &User{
		Name:  name,
		Email: email,
	}

	u.ID = uuid.NewV4().String()
	u.CreatedAt = time.Now()

	if err := u.isValid(); err != nil {
		return nil, err
	}
	return u, nil
}

// User ...
type User struct {
	Base  `valid:"required"`
	Name  string `json:"name" valid:"notnull"`
	Email string `json:"email" valid:"notnull"`
}

func (u *User) isValid() error {
	_, err := validator.ValidateStruct(u)
	if err != nil {
		return err
	}
	return nil
}
