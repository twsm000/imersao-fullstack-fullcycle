package usecase

import (
	"github.com/twsm000/imersao-fullstack-fullcycle/domain/model"
)

// PixUseCase ...
type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

// RegisterKey ...
func (p *PixUseCase) RegisterKey(key, kind, accountID string) (*model.PixKey, error) {
	account, err := p.PixKeyRepository.FindAccount(accountID)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(account, kind, key)
	if err != nil {
		return nil, err
	}

	pixKey, err = p.PixKeyRepository.RegisterKey(pixKey)
	if err != nil {
		return nil, err
	}

	return pixKey, nil
}

// FindKeyByKind ...
func (p *PixUseCase) FindKeyByKind(key, kind string) (*model.PixKey, error) {
	return p.PixKeyRepository.FindKeyByKind(key, kind)
}
