package model

import (
	"encoding/json"
	"fmt"

	validator "github.com/asaskevich/govalidator"
)

// NewTransaction ...
func NewTransaction() *Transaction {
	return new(Transaction)
}

// Transaction ...
type Transaction struct {
	ID           string  `json:"id" validate:"required,uuid"`
	AccountID    string  `json:"account_id" validate:"required,uuid"`
	Amount       float64 `json:"amount" validate:"required,numeric"`
	PixKeyTo     string  `json:"pixKeyTo" validate:"required"`
	PixKeyKindTo string  `json:"pix_key_kind_to" validate:"required"`
	Description  string  `json:"descrition" validate:"required"`
	Status       string  `json:"status" validate:"required"`
	Error        string  `json:"error" validate:"required"`
}

func (t *Transaction) isValid() error {
	if _, err := validator.ValidateStruct(t); err != nil {
		return fmt.Errorf("Error during Transaction validation: %s", err.Error())
	}
	return nil
}

// ParseJSON ...
func (t *Transaction) ParseJSON(data []byte) error {
	if err := json.Unmarshal(data, t); err != nil {
		return err
	}

	if err := t.isValid(); err != nil {
		return err
	}

	return nil
}


// ToJSON
func (t *Transaction) ToJSON() ([]byte, error) {
	if err := t.isValid(); err != nil {
		return nil, err
	}

	data, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return data, nil
}