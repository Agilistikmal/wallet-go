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

type UserService interface {
	Create(ctx context.Context, request model.UserCreateRequest) model.UserResponse
	Update(ctx context.Context, request model.UserUpdateRequest) model.UserResponse
	Delete(ctx context.Context, userId uint)
	FindById(ctx context.Context, userId uint) model.UserResponse
	FindAll(ctx context.Context) []model.UserResponse
	UpdateWallet(ctx context.Context, request model.WalletUpdateRequest) model.WalletResponse
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request model.UserCreateRequest) model.UserResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		panic(err)
	}

	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.TxCommitOrRollback(tx)

	user := model.User{
		Name:     request.Name,
		Email:    request.Email,
		Phone:    request.Phone,
		Password: request.Password,
	}

	user = service.UserRepository.Create(ctx, tx, user)

	return helper.UserToUserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request model.UserUpdateRequest) model.UserResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		panic(err)
	}

	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.TxCommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(handler.NewNotFoundError(err.Error()))
	}
	user = model.User{
		Id:       user.Id,
		Name:     request.Name,
		Email:    request.Email,
		Phone:    request.Phone,
		Password: request.Password,
	}

	user = service.UserRepository.Update(ctx, tx, user)

	return helper.UserToUserResponse(user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId uint) {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.TxCommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(handler.NewNotFoundError(err.Error()))
	}
	service.UserRepository.Delete(ctx, tx, user)
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId uint) model.UserResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.TxCommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(handler.NewNotFoundError(err.Error()))
	}
	return helper.UserToUserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []model.UserResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.TxCommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)

	var userResponses []model.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, helper.UserToUserResponse(user))
	}
	return userResponses
}

func (service *UserServiceImpl) UpdateWallet(ctx context.Context, request model.WalletUpdateRequest) model.WalletResponse {
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

	user = service.UserRepository.UpdateWallet(ctx, tx, user)

	return helper.WalletToWalletResponse(user)
}
