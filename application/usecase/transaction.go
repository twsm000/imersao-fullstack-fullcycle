package usecase

import "github.com/twsm000/imersao-fullstack-fullcycle/domain/model"

// TransactionUseCase ...
type TransactionUseCase struct {
	TransactionRepository model.TransactionRepositoryInterface
	PixRepository model.PixKeyRepositoryInterface
}


// Register ...
func (t *TransactionUseCase) Register(accountID string, amount float64, pixKeyTo string, pixKeyKindTo string, description string) (*model.Transaction, error) {
	account, err := t.PixRepository.FindAccount(accountID)
	if err != nil {
		return nil, err
	}

	pixKey, err := t.PixRepository.FindKeyByKind(pixKeyTo, pixKeyKindTo)
	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(account, amount, pixKey, description)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

// Confirm ...
func (t *TransactionUseCase) Confirm(transactionID string) (*model.Transaction, error) {
	transaction, err := t.TransactionRepository.Find(transactionID)
	if err != nil {
		return nil, err
	}

	transaction.Status = model.TransactionConfirmed
	err = t.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

// Complete ...
func (t *TransactionUseCase) Complete(transactionID string) (*model.Transaction, error) {
	transaction, err := t.TransactionRepository.Find(transactionID)
	if err != nil {
		return nil, err
	}

	transaction.Status = model.TransactionCompleted
	err = t.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

// Cancel ...
func (t *TransactionUseCase) Cancel(transactionID, reason string) (*model.Transaction, error) {
	transaction, err := t.TransactionRepository.Find(transactionID)
	if err != nil {
		return nil, err
	}

	transaction.Status = model.TransactionCanceled
	transaction.CancelDescription = reason
	err = t.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}