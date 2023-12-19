package usecase

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"rental_mobile_fiber/internal/entity"
	"rental_mobile_fiber/internal/model"
	"rental_mobile_fiber/internal/model/convert"
	"rental_mobile_fiber/internal/repository"
)

type EventUseCase struct {
	DB          *sql.DB
	Validate    *validator.Validate
	EventRepo   *repository.EventRepository
	viperConfig *viper.Viper
}

func NewEventUseCase(DB *sql.DB, validate *validator.Validate, eventRepo *repository.EventRepository, viperConfig *viper.Viper) *EventUseCase {
	return &EventUseCase{DB: DB, Validate: validate, EventRepo: eventRepo, viperConfig: viperConfig}
}

func (e *EventUseCase) Create(ctx context.Context, request *model.EventCreate) (*model.EventResponse, error) {
	err := e.Validate.Struct(request)
	if err != nil {
		return nil, fiber.ErrBadRequest
	}

	event := entity.Event{
		Thumbnail:   request.Thumbnail,
		Title:       request.Title,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
		Description: request.Description,
		MaxJoin:     request.MaxJoin,
		Price:       request.Price,
		IsPublish:   request.IsPublish,
		Location:    request.Location,
	}

	eventId, err := e.EventRepo.Create(ctx, event, e.DB)

	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	event.ID = eventId

	return convert.EventToResponse(&event), nil
}

func (e *EventUseCase) Update(ctx context.Context, request *model.EventUpdate, eventId int) (*model.EventResponse, error) {
	err := e.Validate.Struct(request)
	if err != nil {
		return nil, err
	}

	event := entity.Event{
		ID:          eventId,
		Thumbnail:   request.Thumbnail,
		Title:       request.Title,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
		Description: request.Description,
		MaxJoin:     request.MaxJoin,
		Price:       request.Price,
		IsPublish:   request.IsPublish,
		Location:    request.Location,
	}

	err = e.EventRepo.Update(ctx, event, e.DB)

	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return convert.EventToResponse(&event), nil
}

func (e EventUseCase) Destroy(ctx context.Context, eventId int) error {

	err := e.EventRepo.Delete(ctx, eventId, e.DB)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return nil
}
