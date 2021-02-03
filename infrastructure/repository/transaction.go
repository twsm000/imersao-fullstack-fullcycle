package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/twsm000/imersao-fullstack-fullcycle/domain/model"
)

// TransactionRepositoryDB ...
type TransactionRepositoryDB struct {
	Db *gorm.DB
}

// Register ...
func (r *TransactionRepositoryDB) Register(transaction *model.Transaction) error {
	if err := r.Db.Create(transaction).Error; err != nil {
		return err
	}

	return nil
}

// Save ...
func (r *TransactionRepositoryDB) Save(transaction *model.Transaction) error {
	if err := r.Db.Save(transaction).Error; err != nil {
		return err
	}

	return nil
}

// Find ...
func (r *TransactionRepositoryDB) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	r.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, errors.New("no key was found")
	}

	return &transaction, nil
}
