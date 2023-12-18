package repository

import (
	"context"
	"database/sql"
	"rental_mobile_fiber/internal/entity"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) Create(ctx context.Context, user entity.User, tx *sql.DB) error {
	sqlCommand := "insert into users (name,email,password,no_hp) values (?,?,?,?)"
	_, err := tx.ExecContext(ctx, sqlCommand, user.Name, user.Email, user.Password, user.NoHp)
	if err != nil {
		panic(err)
	}

	return nil
}
