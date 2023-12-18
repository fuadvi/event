package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"rental_mobile_fiber/internal/model"
	"rental_mobile_fiber/internal/usecase"
)

type UserController struct {
	UseCase *usecase.UserUseCase
}

func NewUserController(useCase *usecase.UserUseCase) *UserController {
	return &UserController{UseCase: useCase}
}

func (c UserController) Register(ctx *fiber.Ctx) error {
	var request *model.RegisterUser

	if err := ctx.BodyParser(&request); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

func (c UserController) Login(ctx *fiber.Ctx) error {
	var request *model.LoginRequest

	if err := ctx.BodyParser(&request); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}

	response, err := c.UseCase.Login(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}
