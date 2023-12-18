package repository

import (
	"context"
	"database/sql"
	"rental_mobile_fiber/internal/entity"
	"rental_mobile_fiber/internal/model"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) Create(ctx context.Context, user entity.User, db *sql.DB) error {
	sqlCommand := "insert into users (name,email,password,no_hp) values (?,?,?,?)"
	_, err := db.ExecContext(ctx, sqlCommand, user.Name, user.Email, user.Password, user.NoHp)
	if err != nil {
		panic(err)
	}

	return nil
}

func (u *UserRepository) Get(ctx context.Context, request *model.LoginRequest, db *sql.DB) (entity.User, error) {
	sqlCommand := "select id,name,password,email,no_hp from users where email = ? limit 1"
	row, err := db.QueryContext(ctx, sqlCommand, request.Email)
	if err != nil {
		panic(err)
	}

	var user entity.User
	if row.Next() {
		err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.NoHp)
		if err != nil {
			return user, err
		}
	}

	return user, nil
}
