package model

import (
	"time"

	validator "github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

// NewAccount ...
func NewAccount(bank *Bank, owner, number string) (*Account, error) {
	account := &Account{
		BankID: bank.ID,
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
	BankID    string    `gorm:"column:bank_id;type:uuid;not null" valid:"-"`
	OwnerName string    `json:"owner_name" valid:"notnull" gorm:"column:owner_name;type:varchar(255);not null"`
	Number    string    `json:"number" valid:"notnull" gorm:"type:varchar(20)"`
	PixKeys   []*PixKey `valid:"-" gorm:"ForeignKey:AccountID"`
}

func (a *Account) isValid() error {
	_, err := validator.ValidateStruct(a)
	if err != nil {
		return err
	}
	return nil
}
