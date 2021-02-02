package model

import (
	"errors"
	"time"

	validator "github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

var transactionStatusMap map[TransactionStatus]bool

func init() {
	transactionStatusMap = make(map[TransactionStatus]bool)
	transactionStatusMap[TransactionPending] = true
	transactionStatusMap[TransactionCompleted] = true
	transactionStatusMap[TransactionCanceled] = true
	transactionStatusMap[TransactionConfirmed] = true
}

// TransactionStatus ...
type TransactionStatus string

const (
	// TransactionPending ...
	TransactionPending TransactionStatus = "pending"

	// TransactionCompleted ...
	TransactionCompleted TransactionStatus = "completed"

	// TransactionCanceled ...
	TransactionCanceled TransactionStatus = "canceled"

	// TransactionConfirmed ...
	TransactionConfirmed TransactionStatus = "confirmed"
)

// NewTransaction ...
func NewTransaction(accountFrom *Account, amount float64, pixKey *PixKey, description string) (*Transaction, error) {
	t := &Transaction{
		AccountFrom: accountFrom,
		Amount:      amount,
		PixKeyTo:    pixKey,
		Status:      TransactionPending,
		Description: description,
	}

	t.ID = uuid.NewV4().String()
	t.CreatedAt = time.Now()

	if err := t.isValid(); err != nil {
		return nil, err
	}
	return t, nil
}

// TransactionRepositoryInterface ...
type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

// Transaction ...
type Transaction struct {
	Base              `valid:"requried"`
	AccountFrom       *Account          `valid:"-"`
	Amount            float64           `json:"amount" valid:"notnull"`
	PixKeyTo          *PixKey           `valid:"-"`
	Status            TransactionStatus `json:"status" valid:"notnull"`
	Description       string            `json:"description" valid:"notnull"`
	CancelDescription string            `json:"cancel_description" valid:"-"`
}

func (t *Transaction) isValid() error {
	if t.Amount <= 0 {
		return errors.New("the amount must be greater then 0")
	}

	if !(transactionStatusMap[t.Status]) {
		return errors.New("invalid status for the transaction")
	}

	if t.PixKeyTo.AccountID == t.AccountFrom.ID {
		return errors.New("the source and the destination account cannot be the same")
	}

	if _, err := validator.ValidateStruct(t); err != nil {
		return err
	}
	return nil
}

// Cancel ...
func (t *Transaction) Cancel(description string) error {
	t.Description = description
	return t.updateStatus(TransactionCanceled)
}

// Confirm ...
func (t *Transaction) Confirm() error {
	return t.updateStatus(TransactionConfirmed)
}

// Complete ...
func (t *Transaction) Complete() error {
	return t.updateStatus(TransactionCompleted)
}

func (t *Transaction) updateStatus(newStatus TransactionStatus) error {
	t.Status = newStatus
	t.UpdateAt = time.Now()
	return t.isValid()
}

// Transactions ...
type Transactions struct {
	List []*Transaction
}
