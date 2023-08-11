package service

import (
	"context"
	"database/sql"
	"github.com/agilistikmal/wallet-go/handler"
	"github.com/agilistikmal/wallet-go/helper"
	"github.com/agilistikmal/wallet-go/model"
	"github.com/agilistikmal/wallet-go/repository"
	"github.com/go-playground/validator"
)

type WalletService interface {
	Update(ctx context.Context, request model.Wallet) model.Wallet
}

type WalletServiceImpl struct {
	WalletRepository repository.WalletRepository
	UserRepository   repository.UserRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewWalletService(walletRepository repository.WalletRepository, userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) WalletService {
	return &WalletServiceImpl{WalletRepository: walletRepository, UserRepository: userRepository, DB: DB, Validate: validate}
}

func (service *WalletServiceImpl) Update(ctx context.Context, request model.Wallet) model.Wallet {
	err := service.Validate.Struct(request)
	if err != nil {
		panic(err)
	}

	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.TxCommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.UserId)
	if err != nil {
		panic(handler.NewNotFoundError(err.Error()))
	}
	user.WalletAmount = request.Amount

	wallet := service.WalletRepository.Update(ctx, tx, user)

	return wallet
}
