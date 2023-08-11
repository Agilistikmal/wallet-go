package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/agilistikmal/wallet-go/model"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user model.User) model.User
	Update(ctx context.Context, tx *sql.Tx, user model.User) model.User
	Delete(ctx context.Context, tx *sql.Tx, user model.User)
	FindById(ctx context.Context, tx *sql.Tx, userId uint) (model.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []model.User
	UpdateWallet(ctx context.Context, tx *sql.Tx, user model.User) model.User
}

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user model.User) model.User {
	SQL := "INSERT INTO user(name, email, phone, password) VALUES (?, ? ,? ,?)"
	result, err := tx.Exec(SQL, user.Name, user.Email, user.Phone, user.Password)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	user.Id = uint(id)
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user model.User) model.User {
	SQL := "UPDATE user SET name = ?, email = ?, phone = ?, password = ? WHERE id = ?"
	_, err := tx.Exec(SQL, user.Name, user.Email, user.Phone, user.Password, user.Id)
	if err != nil {
		panic(err)
	}
	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user model.User) {
	SQL := "DELETE FROM user WHERE id = ?"
	_, err := tx.Exec(SQL, user.Id)
	if err != nil {
		panic(err)
	}
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId uint) (model.User, error) {
	SQL := "SELECT id, name, email, phone, wallet_amount FROM user WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	user := model.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.WalletAmount)
		if err != nil {
			panic(err)
		}
		return user, nil
	} else {
		return user, errors.New("User not found.")
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []model.User {
	SQL := "SELECT id, name, email, phone, wallet_amount FROM user"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.WalletAmount)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return users
}

func (repository *UserRepositoryImpl) UpdateWallet(ctx context.Context, tx *sql.Tx, user model.User) model.User {
	SQL := "UPDATE user SET wallet_amount = ? WHERE id = ?"
	_, err := tx.Exec(SQL, user.WalletAmount, user.Id)
	if err != nil {
		panic(err)
	}
	return user
}
