package repository

import (
	"context"
	"database/sql"
	"github.com/agilistikmal/wallet-go/helper"
	"github.com/agilistikmal/wallet-go/model"
)

type WalletRepository interface {
	Update(ctx context.Context, tx *sql.Tx, user model.User) model.Wallet
}

type WalletRepositoryImpl struct {
}

func NewWalletRepository() WalletRepository {
	return &WalletRepositoryImpl{}
}

func (repository *WalletRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user model.User) model.Wallet {
	SQL := "UPDATE user SET wallet_amount = ? WHERE id = ?"
	_, err := tx.Exec(SQL, user.WalletAmount, user.Id)
	if err != nil {
		panic(err)
	}
	return helper.UserToWallet(user)
}
