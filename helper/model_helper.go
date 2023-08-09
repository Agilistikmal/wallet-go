package helper

import "github.com/agilistikmal/wallet-go/model"

func UserModelToUserModelResponse(user model.UserModel) model.UserModelResponse {
	return model.UserModelResponse{
		Id:           user.Id,
		Name:         user.Name,
		Email:        user.Email,
		Phone:        user.Phone,
		WalletAmount: user.WalletAmount,
	}
}
