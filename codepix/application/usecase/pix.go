package usecase

import (
	"errors"
	"github.com/lucasfernnades/imersao/codepix-go/domain/model"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (pix *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := pix.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)
	if err != nil {
		return nil, err
	}

	pix.PixKeyRepository.RegisterKey(pixKey)
	if pixKey.ID == "" {
		return nil, errors.New("unable to create new key at the moment")
	}

	return pixKey, nil
}

func (pix *PixUseCase) FindKeyByKind(key string, kind string) (*model.PixKey, error) {
	pixKey, err := pix.PixKeyRepository.FindKeyByKind(key, kind)
	if err != nil {
		return nil, err
	}

	return pixKey, nil
}
