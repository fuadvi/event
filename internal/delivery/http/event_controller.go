package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"rental_mobile_fiber/internal/model"
	"rental_mobile_fiber/internal/usecase"
	"strconv"
)

type EventController struct {
	EventCase *usecase.EventUseCase
}

func NewEventController(eventCase *usecase.EventUseCase) *EventController {
	return &EventController{EventCase: eventCase}
}

func (e *EventController) Create(ctx *fiber.Ctx) error {
	var request *model.EventCreate

	if err := ctx.BodyParser(&request); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	response, err := e.EventCase.Create(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse[*model.EventResponse]{Data: response})
}

func (e *EventController) Update(ctx *fiber.Ctx) error {

	var request *model.EventUpdate

	params := ctx.Params("id")

	eventId, err := strconv.Atoi(params)
	if err != nil {
		return ctx.JSON(model.WebResponse[interface{}]{Errors: err.Error()})
	}

	if err := ctx.BodyParser(&request); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return ctx.JSON(model.WebResponse[interface{}]{Errors: err.Error()})
	}

	response, err := e.EventCase.Update(ctx.UserContext(), request, eventId)
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse[*model.EventResponse]{Data: response})
}

func (e *EventController) Delete(ctx *fiber.Ctx) error {

	params := ctx.Params("id")
	eventId, err := strconv.Atoi(params)

	if err != nil {
		return ctx.JSON(model.WebResponse[interface{}]{Errors: err.Error()})
	}

	if err := e.EventCase.Destroy(ctx.UserContext(), eventId); err != nil {
		return ctx.JSON(model.WebResponse[interface{}]{Errors: err.Error()})
	}

	return ctx.JSON(model.WebResponse[string]{Data: "success delete event"})
}

func (e *EventController) Get(ctx *fiber.Ctx) error {
	params := ctx.Params("id")
	eventId, err := strconv.Atoi(params)

	if err != nil {
		return ctx.JSON(model.WebResponse[interface{}]{Errors: err.Error()})
	}

	response, err := e.EventCase.GetEvent(ctx.UserContext(), eventId)
	if err != nil {
		return ctx.JSON(model.WebResponse[interface{}]{Errors: err.Error()})
	}

	return ctx.JSON(model.WebResponse[*model.EventResponse]{Data: response})
}

func (e *EventController) List(ctx *fiber.Ctx) error {

	response, err := e.EventCase.ListEvent(ctx.UserContext())
	if err != nil {
		return ctx.JSON(model.WebResponse[interface{}]{Errors: err.Error()})
	}

	return ctx.JSON(model.WebResponse[[]*model.EventResponse]{Data: response})
}
