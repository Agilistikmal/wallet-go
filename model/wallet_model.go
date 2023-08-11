package model

type WalletUpdateRequest struct {
	UserId uint `json:"user_id,omitempty" validate:"required,number,gte=0"`
	Amount uint `json:"amount,omitempty" validate:"required,number,gte=0"`
}

type WalletResponse struct {
	UserId uint `json:"user_id,omitempty" validate:"required,number,gte=0"`
	Amount uint `json:"amount,omitempty" validate:"required,number,gte=0"`
}
