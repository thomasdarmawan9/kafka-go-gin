package user_domain

import (
	"fmt"
	"users/db"
	"users/utils/error_formats"
	"users/utils/error_utils"
)

var UserRepo userDomain = &userRepo{}

const (
	queryGetUserByEmail = `SELECT id, email, password, role from users WHERE email = $1`
	queryCreateUser     = `INSERT INTO users (email, password, address, role) VALUES ($1, $2, $3, $4) RETURNING id, email, address, created_at`
)

type userDomain interface {
	GetUserByEmail(string) (*User, error_utils.MessageErr)
	CreateUser(*User) (*User, error_utils.MessageErr)
}

type userRepo struct{}

func (u *userRepo) GetUserByEmail(email string) (*User, error_utils.MessageErr) {
	db := db.GetDB()
	row := db.QueryRow(queryGetUserByEmail, email)

	var user User

	if err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Role); err != nil {
		return nil, error_formats.ParseError(err)
	}

	return &user, nil
}

func (u *userRepo) CreateUser(userReq *User) (*User, error_utils.MessageErr) {
	db := db.GetDB()
	row := db.QueryRow(queryCreateUser, userReq.Email, userReq.Password, userReq.Address, userReq.Role)

	var user User
	err := row.Scan(&user.Id, &user.Email, &user.Address, &user.CreatedAt)

	if err != nil {
		fmt.Println("err create user dao:", err)
		return nil, error_formats.ParseError(err)
	}
	return &user, nil
}
