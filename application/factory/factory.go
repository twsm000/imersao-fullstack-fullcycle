package factory

import (
	"github.com/jinzhu/gorm"
	"github.com/twsm000/imersao-fullstack-fullcycle/application/usecase"
	"github.com/twsm000/imersao-fullstack-fullcycle/infrastructure/repository"
)

// TransactionUseCaseFactory ...
func TransactionUseCaseFactory(db *gorm.DB) *usecase.TransactionUseCase {
	pixRepo := &repository.PixKeyRepositoryDB{Db: db}
	transactionRepo := &repository.TransactionRepositoryDB{Db: db}
	uc := &usecase.TransactionUseCase{
		TransactionRepository: transactionRepo,
		PixRepository:         pixRepo,
	}

	return uc
}
