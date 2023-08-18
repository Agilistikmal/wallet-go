package helper

import "github.com/agilistikmal/wallet-go/model"

func UserToUserResponse(user model.User) model.UserResponse {
	return model.UserResponse{
		Id:           user.Id,
		Name:         user.Name,
		Email:        user.Email,
		Phone:        user.Phone,
		WalletAmount: user.WalletAmount,
		ApiKey:       user.ApiKey,
	}
}

func UserToWallet(user model.User) model.Wallet {
	return model.Wallet{
		UserId: user.Id,
		Amount: user.WalletAmount,
	}
}
