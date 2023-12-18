package usecase

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"rental_mobile_fiber/internal/entity"
	"rental_mobile_fiber/internal/model"
	"rental_mobile_fiber/internal/model/convert"
	"rental_mobile_fiber/internal/repository"
)

type UserUseCase struct {
	DB       *sql.DB
	Validate *validator.Validate
	userRepo *repository.UserRepository
}

func NewUserUseCase(DB *sql.DB, validate *validator.Validate, userRepo *repository.UserRepository) *UserUseCase {
	return &UserUseCase{DB: DB, Validate: validate, userRepo: userRepo}
}

func (u *UserUseCase) Create(ctx context.Context, request *model.RegisterUser) (*model.UserResponse, error) {

	err := u.Validate.Struct(request)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	user := entity.User{
		Password: string(password),
		Name:     request.Name,
		Email:    request.Email,
		NoHp:     request.NoHp,
	}

	if err := u.userRepo.Create(ctx, user, u.DB); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return convert.UserToResponse(&user), nil

}
