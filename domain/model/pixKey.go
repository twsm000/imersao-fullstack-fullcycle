package model

import (
	"errors"
	"time"

	validator "github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

// NewPixKey ...
func NewPixKey(account *Account, kind, key string) (*PixKey, error) {
	pk := &PixKey{
		Account: account,
		Kind:    kind,
		Key:     key,
		Status:  "active",
	}

	pk.ID = uuid.NewV4().String()
	pk.CreatedAt = time.Now()

	if err := pk.isValid(); err != nil {
		return nil, err
	}
	return pk, nil
}

// PixKeyRepositoryInterface ...
type PixKeyRepositoryInterface interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
}

// PixKey ...
type PixKey struct {
	Base      `valid:"required"`
	AccountID string   `json:"account_id" valid:"notnull"`
	Account   *Account `valid:"-"`
	Kind      string   `json:"kind" valid:"notnull"`
	Key       string   `json:"key" valid:"notnull"`
	Status    string   `json:"status" valid:"notnull"`
}

func (pk *PixKey) isValid() error {	
	if pk.Kind != "email" && pk.Kind != "cpf" {
		return errors.New("invalid type of key")
	}
	
	if pk.Status != "active" && pk.Status != "inactive" {
		return errors.New("invalid status")
	}
	
	_, err := validator.ValidateStruct(pk)
	if err != nil {
		return err
	}
	return nil
}
