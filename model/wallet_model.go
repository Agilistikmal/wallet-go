package model

type Wallet struct {
	UserId uint `json:"user_id,omitempty" validate:"required,number,gte=0"`
	Amount uint `json:"amount,omitempty" validate:"required,number,gte=0"`
}
