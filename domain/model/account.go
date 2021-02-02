package model

import (
	"time"

	validator "github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

// NewAccount ...
func NewAccount(bank *Bank, owner, number string) (*Account, error) {
	account := &Account{
		Bank:      bank,
		OwnerName: owner,
		Number:    number,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	if err := account.isValid(); err != nil {
		return nil, err
	}

	return account, nil
}

// Account ...
type Account struct {
	Base      `valid:"required"`
	Bank      *Bank     `valid:"-"`
	OwnerName string    `json:"owner_name" valid:"notnull"`
	Number    string    `json:"number" valid:"notnull"`
	PixKeys   []*PixKey `valid:"-"`
}

func (a *Account) isValid() error {
	_, err := validator.ValidateStruct(a)
	if err != nil {
		return err
	}
	return nil
}
