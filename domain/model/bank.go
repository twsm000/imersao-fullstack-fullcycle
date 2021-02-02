package model

import (
	"time"

	validator "github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	validator.SetFieldsRequiredByDefault(true)
}

// NewBank ...
func NewBank(code, name string) (*Bank, error) {
	bank := &Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	if err := bank.isValid(); err != nil {
		return nil, err
	}

	return bank, nil
}

// Bank ...
type Bank struct {
	Base    `valid:"required"`
	Code    string     `json:"code" valid:"notnull"`
	Name    string     `json:"name" valid:"notnull"`
	Account []*Account `valid:"-"`
}

func (b *Bank) isValid() error {
	_, err := validator.ValidateStruct(b)
	if err != nil {
		return err
	}
	return nil
}
