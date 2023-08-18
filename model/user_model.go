package model

type User struct {
	Id           uint
	Name         string
	Email        string
	Phone        string
	Password     string
	WalletAmount uint
	ApiKey       string
}

type UserCreateRequest struct {
	Name     string `validate:"required,min=3,max=16" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Phone    string `validate:"required,e164" json:"phone"`
	Password string `validate:"required,min=8" json:"password"`
}

type UserUpdateRequest struct {
	Id       uint   `validate:"required,number,gte=0" json:"id,omitempty"`
	Name     string `validate:"required,min=3,max=16" json:"name,omitempty"`
	Email    string `validate:"required,email" json:"email,omitempty"`
	Phone    string `validate:"required,e164" json:"phone,omitempty"`
	Password string `validate:"required,min=8" json:"password,omitempty"`
}

type UserDeleteRequest struct {
	Id uint `validate:"required,number,gte=0"`
}

type UserResponse struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	WalletAmount uint   `json:"wallet_amount"`
	ApiKey       string `json:"api_key"`
}
