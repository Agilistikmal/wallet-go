package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/agilistikmal/wallet-go/helper"
	"github.com/agilistikmal/wallet-go/model"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user model.User) model.User
	Update(ctx context.Context, tx *sql.Tx, user model.User) model.User
	Delete(ctx context.Context, tx *sql.Tx, user model.User)
	FindById(ctx context.Context, tx *sql.Tx, userId uint) (model.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []model.User
}

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user model.User) model.User {
	hashPassword, errHash := helper.HashPassword(user.Password)
	apiKey := helper.GenerateRandomString(30)
	if errHash != nil {
		panic(errHash)
	}
	SQL := "INSERT INTO user(name, email, phone, password, api_key) VALUES (?, ? ,? ,?, ?)"
	result, err := tx.Exec(SQL, user.Name, user.Email, user.Phone, hashPassword, apiKey)
	if err != nil {
		panic(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	user.Id = uint(id)
	user.ApiKey = apiKey
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user model.User) model.User {
	SQL := "UPDATE user SET name = ?, email = ?, phone = ?, password = ? WHERE id = ?"
	_, err := tx.Exec(SQL, user.Name, user.Email, user.Phone, user.Password, user.Id)
	if err != nil {
		panic(err.Error())
	}
	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user model.User) {
	SQL := "DELETE FROM user WHERE id = ?"
	_, err := tx.Exec(SQL, user.Id)
	if err != nil {
		panic(err.Error())
	}
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId uint) (model.User, error) {
	SQL := "SELECT id, name, email, phone, wallet_amount, api_key FROM user WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	user := model.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.WalletAmount, &user.ApiKey)
		if err != nil {
			panic(err.Error())
		}
		return user, nil
	} else {
		return user, errors.New("User not found.")
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []model.User {
	SQL := "SELECT id, name, email, phone, wallet_amount, api_key FROM user"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.WalletAmount, &user.ApiKey)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}
	return users
}
