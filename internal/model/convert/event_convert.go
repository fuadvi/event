package convert

import (
	"rental_mobile_fiber/internal/entity"
	"rental_mobile_fiber/internal/model"
)

func EventToResponse(event *entity.Event) *model.EventResponse {
	return &model.EventResponse{
		ID:          event.ID,
		Thumbnail:   event.Thumbnail,
		Title:       event.Title,
		StartDate:   event.StartDate,
		EndDate:     event.EndDate,
		Description: event.Description,
		MaxJoin:     event.MaxJoin,
		Price:       event.Price,
		IsPublish:   event.IsPublish,
		Location:    event.Location,
	}
}
