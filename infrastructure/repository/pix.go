package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/twsm000/imersao-fullstack-fullcycle/domain/model"
)

// PixKeyRepositoryDB ...
type PixKeyRepositoryDB struct {
	Db *gorm.DB
}

// AddBank ...
func (r *PixKeyRepositoryDB) AddBank(bank *model.Bank) error {
	if err := r.Db.Create(bank).Error; err != nil {
		return err
	}

	return nil
}

// AddAccount ...
func (r *PixKeyRepositoryDB) AddAccount(account *model.Account) error {
	if err := r.Db.Create(account).Error; err != nil {
		return err
	}

	return nil
}

// RegisterKey ...
func (r *PixKeyRepositoryDB) RegisterKey(pixKey *model.PixKey) (*model.PixKey, error) {
	if err := r.Db.Create(pixKey).Error; err != nil {
		return nil, err
	}

	return pixKey, nil
}

// FindKeyByKind ...
func (r *PixKeyRepositoryDB) FindKeyByKind(key, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey

	r.Db.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)

	if pixKey.ID == "" {
		return nil, errors.New("no key was found")
	}

	return &pixKey, nil
}

// FindAccount ...
func (r *PixKeyRepositoryDB) FindAccount(id string) (*model.Account, error) {
	var account model.Account

	r.Db.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, errors.New("no key was found")
	}

	return &account, nil
}

// FindBank ...
func (r *PixKeyRepositoryDB) FindBank(id string) (*model.Bank, error) {
	var bank model.Bank

	r.Db.Preload("Bank").First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, errors.New("no key was found")
	}

	return &bank, nil
}
