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
	Create(ctx context.Context, request model.UserModelCreateRequest) model.UserModelResponse
	Update(ctx context.Context, request model.UserModelUpdateRequest) model.UserModelResponse
	Delete(ctx context.Context, userId uint)
	FindById(ctx context.Context, userId uint) model.UserModelResponse
	FindAll(ctx context.Context) []model.UserModelResponse
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

func (service *UserServiceImpl) Create(ctx context.Context, request model.UserModelCreateRequest) model.UserModelResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		panic(err)
	}

	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.TxCommitOrRollback(tx)

	user := model.UserModel{
		Name:     request.Name,
		Email:    request.Email,
		Phone:    request.Phone,
		Password: request.Password,
	}

	user = service.UserRepository.Create(ctx, tx, user)

	return helper.UserModelToUserModelResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request model.UserModelUpdateRequest) model.UserModelResponse {
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
	user = model.UserModel{
		Id:           user.Id,
		Name:         request.Name,
		Email:        request.Email,
		Phone:        request.Phone,
		Password:     request.Password,
		WalletAmount: request.WalletAmount,
	}

	user = service.UserRepository.Update(ctx, tx, user)

	return helper.UserModelToUserModelResponse(user)
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

func (service *UserServiceImpl) FindById(ctx context.Context, userId uint) model.UserModelResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.TxCommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(handler.NewNotFoundError(err.Error()))
	}
	return helper.UserModelToUserModelResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []model.UserModelResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.TxCommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)

	var userResponses []model.UserModelResponse
	for _, user := range users {
		userResponses = append(userResponses, helper.UserModelToUserModelResponse(user))
	}
	return userResponses
}
