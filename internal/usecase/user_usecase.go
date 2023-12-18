package usecase

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"log"
	"rental_mobile_fiber/internal/entity"
	"rental_mobile_fiber/internal/model"
	"rental_mobile_fiber/internal/model/convert"
	"rental_mobile_fiber/internal/repository"
	"time"
)

type UserUseCase struct {
	DB          *sql.DB
	Validate    *validator.Validate
	userRepo    *repository.UserRepository
	viperConfig *viper.Viper
}

func NewUserUseCase(DB *sql.DB, validate *validator.Validate, userRepo *repository.UserRepository, viperConfig *viper.Viper) *UserUseCase {
	return &UserUseCase{DB: DB, Validate: validate, userRepo: userRepo, viperConfig: viperConfig}
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

func (u *UserUseCase) Login(ctx context.Context, request *model.LoginRequest) (*model.UserResponse, error) {
	err := u.Validate.Struct(request)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	user, err := u.userRepo.Get(ctx, request, u.DB)

	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, fiber.ErrBadRequest
	}

	claims := jwt.MapClaims{
		"name": user.Name,
		"exp":  time.Now().Add(time.Second * 10).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(u.viperConfig.GetString("jwt.secret")))
	if err != nil {
		log.Printf("token.SignedString: %v", tokenString)
		return &model.UserResponse{}, fiber.ErrInternalServerError
	}

	log.Printf("token.SignedString: %v", tokenString)
	log.Printf("viper login: %v", u.viperConfig.GetString("jwt.secret"))

	userResponse := convert.UserToResponse(&user)
	userResponse.Token = tokenString

	return userResponse, nil
}
