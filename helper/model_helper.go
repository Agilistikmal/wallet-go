package helper

import "github.com/agilistikmal/wallet-go/model"

func UserToUserResponse(user model.User) model.UserResponse {
	return model.UserResponse{
		Id:           user.Id,
		Name:         user.Name,
		Email:        user.Email,
		Phone:        user.Phone,
		WalletAmount: user.WalletAmount,
	}
}

func WalletToWalletResponse(user model.User) model.WalletResponse {
	return model.WalletResponse{
		UserId: user.Id,
		Amount: user.WalletAmount,
	}
}
