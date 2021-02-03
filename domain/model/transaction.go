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
		AccountFromID: accountFrom.ID,
		AccountFrom:   accountFrom,
		Amount:        amount,
		PixKeyIDTo:    pixKey.ID,
		PixKeyTo:      pixKey,
		Status:        TransactionPending,
		Description:   description,
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
	Base              `valid:"required"`
	AccountFrom       *Account          `valid:"-"`
	AccountFromID     string            `gorm:"column:account_from_id;type:uuid" valid:"notnull"`
	Amount            float64           `json:"amount" valid:"notnull" gorm:"type:float"`
	PixKeyTo          *PixKey           `valid:"-"`
	PixKeyIDTo        string            `gorm:"column:pix_key_id_to;type:uuid;" valid:"notnull"`
	Status            TransactionStatus `json:"status" valid:"notnull" gorm:"type:varchar(20)"`
	Description       string            `json:"description" valid:"notnull" gorm:"type:varchar(255)"`
	CancelDescription string            `json:"cancel_description" valid:"-" gorm:"type:varchar(255)"`
}

func (t *Transaction) isValid() error {
	if t.Amount <= 0 {
		return errors.New("the amount must be greater then 0")
	}

	if !(transactionStatusMap[t.Status]) {
		return errors.New("invalid status for the transaction")
	}

	if t.PixKeyTo.AccountID == t.AccountFromID {
		return errors.New("the source and the destination account cannot be the same")
	}

	if _, err := validator.ValidateStruct(t); err != nil {
		return err
	}
	return nil
}

// Cancel ...
func (t *Transaction) Cancel(description string) error {
	t.CancelDescription = description
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
